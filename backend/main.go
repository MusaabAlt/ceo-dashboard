package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type App struct {
	DB *pgxpool.Pool
}

func main() {
	// Load env vars
	_ = godotenv.Load()
	_ = godotenv.Load("../.env")

	port := getEnv("APP_PORT", "8080")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		getEnv("DB_USER", "ceo"),
		getEnv("DB_PASSWORD", "ceo123"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5433"),
		getEnv("DB_NAME", "ceo_dashboard"),
		getEnv("DB_SSLMODE", "disable"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal("DB connect error:", err)
	}
	if err := db.Ping(ctx); err != nil {
		log.Fatal("DB ping error:", err)
	}

	app := &App{DB: db}

	r := chi.NewRouter()
	r.Get("/health", app.health)
	r.Get("/db/ping", app.dbPing)

	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func (a *App) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (a *App) dbPing(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var x int
	if err := a.DB.QueryRow(ctx, "SELECT 1").Scan(&x); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"db": "connected",
	})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
