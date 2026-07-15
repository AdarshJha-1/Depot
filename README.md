# 📦 Depot

A lightweight image hosting service built with Go.

Upload an image through a simple web interface and receive a shareable URL to access it.

## Features

- Image upload via multipart form
- Static image serving
- UUID-based filenames
- MIME type validation
- 4 MB upload limit
- Rate limiting
- Security headers
- Minimal HTMX frontend
- Docker support

## Tech Stack

- Go
- HTMX
- Docker

## Project Structure

```text
.
├── cmd/
├── internal/
│   ├── handler/
│   ├── middleware/
│   ├── routes/
│   └── server/
├── web/
├── uploads/
└── Dockerfile
```

## Running Locally

```bash
git clone https://github.com/AdarshJha-1/Depot.git
cd Depot

go run ./cmd/depot
```

Open:

```
http://localhost:6969
```

## Docker

Build the image:

```bash
docker build -t depot .
```

Run the container:

```bash
docker run -p 6969:6969 depot
```

Then open:

```
http://localhost:6969
```

## API

### Upload an image

```http
POST /upload
```

Form field:

```
img
```

### Access an uploaded image

```http
GET /uploads/{filename}
```

Example:

```
/uploads/c470322a-cfdf-4011-b5ea-a2971105d810.png
```

## Preview

<img width="800" alt="Depot Upload Page" src="./uploads/c470322a-cfdf-4011-b5ea-a2971105d810.jpg">
