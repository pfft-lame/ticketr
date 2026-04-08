package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ticketr/internal/config"
	"ticketr/internal/db"
	repo "ticketr/internal/repository"
)

func main() {
	validate := newValidator()

	env, err := config.LoadEnv(validate)
	if err != nil {
		log.Fatalf("Error loading ENVs\n%s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbInstance, err := db.New(ctx, env.Dsn)
	if err != nil {
		log.Fatalf("DB connection failed: %s", err)
	}

	queries := repo.New(dbInstance)

	cfg := cfg{
		port: env.Port,
	}
	app := application{
		cfg:      cfg,
		db:       dbInstance,
		queries:  queries,
		validate: validate,
	}

	err = run(&app)
	if err != nil {
		log.Fatalln("Failed to Shutdown Gracefully: ", err)
	}

	log.Println("Graceful Shutdown complete.")
}

func run(app *application) error {
	svr := app.server(app.mount())

	go func() {
		log.Println("Server started in port:", app.cfg.port)
		if err := svr.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln("Failed to start the server: ", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		return fmt.Errorf("Failed to Shutdown server: %w", err)
	}
	log.Println("Closed Sever.")

	app.db.Close()
	log.Println("Closed DB.")

	return nil
}
