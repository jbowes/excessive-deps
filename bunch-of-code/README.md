# Sample Full-Stack Web App

This is a sample full-stack web application with a Go backend and a Node (Vite + React + TypeScript) frontend.

- Backend: `backend/` (Go, chi router, in-memory storage, services, middleware)
- Frontend: `frontend/` (Vite + React + TypeScript)
- IaC: `k8s/` (basic Kubernetes manifests)
- Docker: `backend/Dockerfile`, `frontend/Dockerfile`

## Getting Started

### Backend
1. Install Go 1.21+.
2. From `backend/`, run:
   - `go mod tidy`
   - `go run ./cmd/server`

The backend listens on `:8080` by default.

### Frontend
1. Install Node 18+ and pnpm or npm.
2. From `frontend/`, run:
   - `pnpm install` (or `npm install`)
   - `pnpm dev` (or `npm run dev`)

The frontend dev server listens on `:5173` by default and proxies API requests to the backend.

### API
- `GET /health` – health check.
- `GET /api/users` – list users.
- `POST /api/users` – create user.
- `GET /api/todos` – list todos.
- `POST /api/todos` – create todo.

### Docker
- Build backend: `docker build -t sample-backend ./backend`
- Build frontend: `docker build -t sample-frontend ./frontend`

### Kubernetes (sample only)
Apply manifests from `k8s/` after adjusting namespace and ingress domain.

## Project Structure

```
backend/
  cmd/server/main.go
  internal/
    config/config.go
    http/
      router.go
      handlers/
        health.go
        users.go
        todos.go
      middleware/
        auth.go
        logging.go
    models/
      user.go
      todo.go
      project.go
    services/
      user_service.go
      todo_service.go
    storage/
      memory/
        user_repo.go
        todo_repo.go
  pkg/
    logger/logger.go
    utils/password.go
    utils/validation.go
frontend/
  index.html
  package.json
  tsconfig.json
  vite.config.ts
  src/
    main.tsx
    App.tsx
    styles.css
    components/
      Header.tsx
      TodoList.tsx
      TodoItem.tsx
      UserList.tsx
    hooks/
      useFetch.ts
    services/
      api.ts
    types/
      index.ts
k8s/
  namespace.yaml
  backend-deployment.yaml
  service-backend.yaml
  frontend-deployment.yaml
  service-frontend.yaml
  ingress.yaml
```

## Notes
- Storage is in-memory for simplicity.
- Add persistence by implementing real repositories under `internal/storage`.
- This repository is designed for demonstration, with lots of source code across layers.
