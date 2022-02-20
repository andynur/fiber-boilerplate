package main

import (
	"flag"
	"log"

	"github.com/andynur/fiber-boilerplate/app"
	"github.com/andynur/fiber-boilerplate/database/migrations"
	"github.com/andynur/fiber-boilerplate/routes"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
)

func main() {
	configFile := flag.String("config", "config.yml", "User Config file from user")
	migrate := flag.Bool("migrate", false, "Update db structure")
	flag.Parse()
	app.Load(*configFile)
	if app.Http.Profiler.Enabled {
		_, _ = profiler.Start(profiler.Config{
			ApplicationName: app.Http.Server.Name,
			ServerAddress:   app.Http.Profiler.Server,
		})
	}

	app.Http.Server.Version = app.Version
	if *migrate {
		migrations.Migrate()
	} else {
		routes.LoadRoutes(app.Http.Server.App)
		app.Http.Route404()
		log.Fatal(app.Http.Server.ServeWithGraceFullShutdown())
	}

}
