package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/hnpatil/messages/entity"
	"github.com/hnpatil/messages/utils/config"
	_ "github.com/lib/pq"
)

func GetInstance(cfg *config.Config) *entity.Client {
	driver, err := sql.Open(dialect.Postgres, cfg.GetValue(config.DB_URL))
	if err != nil {
		panic(fmt.Sprintf("failed opening connection to postgres: %v", err))
	}

	options := []entity.Option{entity.Driver(driver), entity.Debug()}

	client := entity.NewClient(options...)

	if err := client.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	return client
}
