package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/nathancamolez-dev/go-orbit/internal/api"
)

func main() {
	gob.Register(uuid.UUID{})
	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(
		ctx,
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)

	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
}

func run(ctx context.Context) error {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := cfg.Build()
	if err != nil {
		return err
	}

	logger = logger.Named("go-orbit-api")
	defer logger.Sync()

	pool, err := pgxpool.New(
		ctx,
		fmt.Sprintf(
			"user = %s password = %s host = %s port = %s dbname = %s ",
			os.Getenv("GO_DATABASE_USER"),
			os.Getenv("GO_DATABASE_PASSWORD"),
			os.Getenv("GO_DATABASE_HOST"),
			os.Getenv("GO_DATABASE_PORT"),
			os.Getenv("GO_DATABASE_NAME"),
		),
	)

	if err != nil {
		return err
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return err
	}

	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 12 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	api := api.API{
		Router: chi.NewRouter(),
	}

	api.BindRoutes()

	logger.Info("Starting a server on port 8080")

	if err := http.ListenAndServe("localhost:8080", api.Router); err != nil {
		panic(err)
	}

	return nil
}
