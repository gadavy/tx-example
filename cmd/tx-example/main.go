package main

import (
	"log"
	"os"

	"github.com/gadavy/tx-example/config"
	auth "github.com/gadavy/tx-example/internal/app/api/users/v1"
	"github.com/gadavy/tx-example/internal/app/repositories"
	"github.com/gadavy/tx-example/internal/app/services"
	"github.com/gadavy/tx-example/internal/pkg/clock"
	"github.com/gadavy/tx-example/internal/pkg/crypto"
	"github.com/jmoiron/sqlx"
	"github.com/loghole/tron"
)

func main() {
	app, err := tron.New(config.TronOptions()...)
	if err != nil {
		log.Fatalf("can't create app: %v", err)
	}

	defer app.Close()

	database, err := sqlx.Connect("postgres", "some connection string")
	if err != nil {
		log.Fatalf("can't connect to postgres: %v", err)
	}

	// Init repositories.
	var (
		usersRepo        = repositories.NewUsersRepo(app.TraceLogger(), database)
		sessionsRepo     = repositories.NewSessionRepo(app.TraceLogger(), database)
		loginHistoryRepo = repositories.NewLoginHistoryRepo(app.TraceLogger(), database)
	)

	// Init service
	var (
		authService = services.NewAuthService(
			app.TraceLogger(),
			clock.NewClock(),
			usersRepo,
			sessionsRepo,
			loginHistoryRepo,
			crypto.NewTokenGenerator(os.Getenv("some secret")),
			crypto.NewBCrypt(crypto.DefaultCost),
		)
	)

	// Init implementations
	authV1Impl := auth.NewImplementation(app.TraceLogger(), authService)

	if err := app.Run(authV1Impl); err != nil {
		app.Logger().Fatalf("can't run app: %v", err)
	}

	if err := app.Wait(); err != nil {
		app.Logger().Errorf("wait: %v", err)
	}
}
