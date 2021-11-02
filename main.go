package main

import (
	"context"
	"fmt"
	"github.com/86soft/healthyro-recipes/adapters/dao"
	"github.com/86soft/healthyro-recipes/app"
	"github.com/86soft/healthyro-recipes/app/command"
	"github.com/86soft/healthyro-recipes/app/query"
	"github.com/86soft/healthyro-recipes/ports"
	"github.com/86soft/healthyro/recipe"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	//dsn := os.Getenv("db_conn")
	dsn := os.Getenv("db_conn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to db")
	}

	app := app.NewApplication(dao.NewRecipeRepository(db))
	srv := ports.NewGrpcServer(app)
	recipe.RegisterRecipeServiceServer()
}
