package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"log/slog"
	"os"
	"ozone/internal/pkg/config"
	"ozone/internal/pkg/profile/repo"
	"ozone/internal/pkg/profile/usecase"
	"ozone/internal/utils/logger"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {

	// STEUP CONFIG
	cfg := config.MustLoad()

	//::: SETUP LOGGER (*logFile)

	logFile, err := os.OpenFile(cfg.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening logFile: %v", err)
		return err

	}

	defer func() {
		err = errors.Join(err, logFile.Close())
	}()

	loger := logger.Set(cfg.Environmet, logFile)

	log.Println(cfg)
	loger.Info("Starting "+cfg.GRPC.AuthContainerIP,
		slog.String("GRPC adddr", fmt.Sprintf("%s:%d", cfg.GRPC.AuthContainerIP, cfg.GRPC.AuthPort)),
		slog.String("LOG file ", cfg.LogFilePath))

	loger.Debug("debug messages are enabled")

	// ::::::::::::DATABASE::::::::::::::::
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Database.DBuser, cfg.Database.DBpass, cfg.Database.DBhost, cfg.Database.DBport, cfg.Database.DBname)
	pool, err := pgxpool.Connect(context.Background(), connStr)

	if err != nil {
		loger.Error("Error connecting to database", "err", err)
		return err

	}
	err = pool.Ping(context.Background())
	if err != nil {
		loger.Error("Error pinging database", "err", err)
		return err
	}

	// ::::::::::::DATABASE::::::::::::::::

	//::: create repo,usecase,grpc

	profileRepo := repo.NewProfileRepo(pool)
	profileUseCase := usecase.NewProfileUseCase(profileRepo, cfg)
	return nil
}
