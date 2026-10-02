# PixelForge — Version 0

A Go-based image processing and image generation application designed to work with image files. PixelForge provides a backend for uploading, processing, transforming, storing, and retrieving images, with a React-based frontend for interacting with the system.

**Version 0** focuses on building the core image-processing pipeline, authentication, image storage, database management, and the initial application workflow. AI-powered image features are planned for future versions.

---

## Backend Setup

### 1. Install Docker

Install Docker and Docker Compose on your system.

Verify the installation:

```bash
docker --version
docker compose version
```

### 2. Clone the Repository

```bash
git clone https://github.com/MilanBist/image-generator.git<your-repository-url>
cd image-generator/backend
```

### 3. Start PostgreSQL

Start the PostgreSQL 16 container:

```bash
docker compose up -d
```

Check that the container is running:

```bash
docker ps
```

You should see a PostgreSQL container similar to:

```text
backend-db-1   postgres:16   ...   0.0.0.0:5435->5432/tcp
```


### 4. Connect to PostgreSQL

Open the PostgreSQL shell inside the Docker container:

```bash
docker exec -it backend-db-1 psql -U postgres -d pixelForge
```

### 5. Run the Backend

From the backend directory:

```bash
go run .
```

The backend should now be running and ready to receive requests from the frontend.

### 6. Stop PostgreSQL

When finished, stop the PostgreSQL container:

```bash
docker compose down
```

---

## Database Setup

PostgreSQL is configured using Docker Compose, while **Goose** is used to manage database migrations.

The database currently uses PostgreSQL 16.

Example Docker Compose configuration:

```yaml
services:

  db:

    image: postgres:16

    environment:

      POSTGRES_USER: postgres

      POSTGRES_PASSWORD: postgres

      POSTGRES_DB: pixelForge

    ports:

      - "5435:5432"

    volumes:

      - postgres_data:/var/lib/postgresql/data

volumes:

  postgres_data:
```

Start the database:

```bash
docker compose up -d
```

Run the database migrations:

```bash
goose -dir database/migration postgres "$DATABASE_URL" up
```

Goose keeps track of applied migrations so that database changes can be managed consistently across development environments.

---

## Version 0

PixelForge v0 establishes the foundation of the application and focuses on the core backend and frontend workflow.

### Core Features

* User registration and login
* JWT-based authentication
* Access and refresh token flow
* Authentication middleware
* Logging middleware
* RAW image file uploading
* RAW image parsing and image extraction
* Image metadata storage
* Generated image storage
* PostgreSQL database integration
* Goose database migrations
* Image retrieval and downloading
* Image history
* Image preview through the frontend
* Basic image transformations
* Image resizing
* Image processing pipeline
* Backend unit/handler testing
* React frontend integrated with the Go backend

The v0 implementation is primarily focused on understanding and building the complete application pipeline from **file upload → processing → database metadata → image storage → retrieval → frontend display**.

---

## Project Architecture

PixelForge currently follows a backend/frontend architecture:

```text
                 ┌─────────────────┐
                 │  React Frontend │
                 └────────┬────────┘
                          │ HTTP
                          ▼
                 ┌─────────────────┐
                 │   Go Backend    │
                 │                 │
                 │  REST API       │
                 │  Authentication │
                 │  Image Handler  │
                 │  Processing     │
                 └───────┬─────────┘
                         │
              ┌──────────┴──────────┐
              ▼                     ▼
       ┌─────────────┐       ┌──────────────┐
       │ PostgreSQL  │       │ File Storage │
       │             │       │              │
       │ Metadata    │       │ RAW / Images │
       └─────────────┘       └──────────────┘
```

The database stores information about users, uploaded files, generated images, and image-processing history, while the actual image files are stored separately.

---

## Testing

The project includes handler-level tests for the backend.

The tests use interfaces and fake implementations where appropriate so that handlers can be tested without depending on the actual database or file storage system.

The initial testing strategy focuses on verifying the behavior of individual handlers and their interactions with their dependencies.

Integration testing will be expanded in future versions.

---

## Future Improvements

PixelForge v0 provides the foundation for more advanced image-processing and AI functionality.

Planned improvements for future versions include:

### Backend Improvements

* Integration tests for the complete request flow
* Improved error handling
* Request validation improvements
* Rate limiting
* Caching
* Concurrency improvements
* Idempotency for image-processing requests
* Pagination for image history
* Better logging and observability
* Background processing for expensive image operations
* Improved storage management

### Image Processing

* Additional image transformations
* Advanced filters
* Brightness and contrast adjustments
* More resizing options
* Image format conversion
* Image optimization
* Batch image processing

### AI Features

* AI-powered image transformations
* Prompt-based image editing
* AI-assisted image generation
* Natural-language image operations
* WebLLM integration
* Integration with external AI APIs
* AI-powered image feature detection

### Infrastructure

* Object storage such as S3-compatible storage
* Redis-based caching
* Background job processing
* Containerized production deployment
* CI/CD pipeline
* Better production configuration and secrets management

### Frontend
* Improved image editor
* Advanced image preview/lightbox
* Better history organization
* Drag-and-drop uploads
* Progress indicators for long-running operations
* More interactive image-processing controls
* Implement the search bar with the webllm.

---

## Acknowledgements

PixelForge was built as a learning-focused project to explore backend engineering, image processing, database design, authentication, testing, and the integration of modern AI capabilities.

The project makes use of several open-source technologies and libraries, including:

* **Go** — Backend development
* **PostgreSQL** — Database
* **Docker** — Database/container environment
* **Goose** — Database migrations
* **React** — Frontend
* **JWT** — Authentication
* **pgx** — PostgreSQL driver and connection pooling

Special thanks to the open-source community and the documentation and learning resources that make it possible to build projects like PixelForge.

---

## Project Status

**Current Version: v0**

PixelForge v0 is focused on completing the core application workflow and establishing a solid foundation for future development.

The next versions will focus on deeper testing, backend scalability, advanced image-processing capabilities, and AI integration.
