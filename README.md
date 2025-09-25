# Salon Finder Backend

A backend service for Salon Finder.

## 🚀 How to Run

1. Clone the repo
   ```bash
   git clone https://github.com/aluyapeter/salon-finder-backend.git
   cd salon-finder-backend
   ```
   
2. Run the API
```bash
go run ./cmd/api
```

3. Visit health check
```bash
http://localhost:8080/health
```

# ⚙️ Environment Variables
PORT — API port (default 8080)

DATABASE_URL — Database connection string

# 🛠️ Development
make run → run locally

make build → build binary

make lint → run linter
