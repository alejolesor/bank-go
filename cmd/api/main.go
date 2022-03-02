package main

import (
	"VERITRAN/cmd/api/app"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := httplog.NewLogger("http-log", httplog.Options{
		Concise: true,
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))

	router(r)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func router(r *chi.Mux) {
	operations := app.DependencyInjection()

	r.Post("/deposit", operations.Deposite)
	r.Post("/create", operations.Create)
	r.Post("/withDrawal", operations.WithDrawal)
	r.Post("/transfer", operations.Transfer)

}
