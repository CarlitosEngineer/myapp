package main

import (
	"fmt"

	"myapp/configs"
	app "myapp/internal/app"
	dbpkg "myapp/internal/db"
	"myapp/internal/modules/countries"
)

func main() {
	cfg := configs.Load()
	isDev := true

	gdb, err := dbpkg.NewGorm(cfg, isDev)
	if err != nil {
		panic(err)
	}

	// AutoMigrate del módulo
	if err := gdb.AutoMigrate(&countries.Country{}); err != nil {
		panic(err)
	}

	srv := app.NewServer(gdb)
	srv.HealthRoutes()
	srv.RegisterRoutes()

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	_ = srv.Start(addr)
}
