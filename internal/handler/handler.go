package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	uploadDir = "./uploads"
)

func HandlerPong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

func HandlerUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.Body = http.MaxBytesReader(w, r.Body, 4<<20)
		err := r.ParseMultipartForm(4 << 20)
		if err != nil {
			http.Error(w, "failed to parse form", http.StatusBadRequest)
			return
		}

		// http.DetectContentType()

		file, handler, err := r.FormFile("img")
		if err != nil {
			http.Error(w, "error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		buffer := make([]byte, 512)

		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			http.Error(w, "failed to read file", http.StatusBadRequest)
			return
		}

		contentType := http.DetectContentType(buffer[:n])

		allowed := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/webp": true,
			"image/gif":  true,
		}

		if !allowed[contentType] {
			http.Error(w, "unsupported file type", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, "failed to rewind file", http.StatusInternalServerError)
			return
		}

		ext := filepath.Ext(handler.Filename)
		fileName := uuid.NewString() + ext

		err = os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			fmt.Printf("dir Error: %v\n", err)
			http.Error(w, "failed to initialize storage", http.StatusInternalServerError)
			return
		}

		dst, err := os.Create(filepath.Join(uploadDir, fileName))
		if err != nil {
			http.Error(w, "failed to save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "failed to write file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		url := "/uploads/" + fileName

		fmt.Fprintf(w, `
			<div class="result-card">
				<h3>✅ Upload Successful</h3>

				<p>
					<a href="%s" target="_blank">%s</a>
				</p>

				<img src="%s" alt="Uploaded Image">
			</div>
		`, url, url, url)
	}
}

func HandlerGetUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		img := filepath.Clean(r.PathValue("img"))
		path := filepath.Join(uploadDir, img)
		http.ServeFile(w, r, path)
	}
}
