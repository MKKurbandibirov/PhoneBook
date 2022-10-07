package main

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	pkgpostgres "phonebook/pkg/postgres"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	dsn := "postgres://magomed:1111@localhost:5442/phonebook" +
		"?sslmode=disable"

	pool, err := pkgpostgres.NewPool(dsn, logger)
	if err != nil {
		logger.Fatal(err)
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		logger.Fatal(fmt.Errorf("couldn't ping database: %w", err))
	}

	//repo := repository.NewRepository(pool, logger)
	//fmt.Println(repo.Numbers(context.Background()))
}
