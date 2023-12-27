package cmd

import (
	"fmt"
	"net/http"
	"time"

	"go-oauth2-server/config"
	"go-oauth2-server/services"

	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
	"gopkg.in/tylerb/graceful.v1"
)

// RunServer runs the app
func RunServer(configBackend config.ConfigBackendType, port int) error {
	cnf, db, err := initConfigDB(true, true, configBackend)
	if err != nil {
		return err
	}
	defer func() {
		d, err := db.DB()
		if err == nil {
			d.Close()
		}
	}()

	// start the services
	if err := services.Init(cnf, db); err != nil {
		return err
	}
	defer services.Close()

	// Start a classic negroni app
	app := negroni.New()
	app.Use(negroni.NewRecovery())
	app.Use(negroni.NewLogger())
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(negroni.NewStatic(http.Dir("public")))

	// Create a router instance
	router := mux.NewRouter()

	// Add routes
	services.HealthService.RegisterRoutes(router, "/v1")
	services.OauthService.RegisterRoutes(router, "/v1/oauth")
	services.WebService.RegisterRoutes(router, "/web")

	// Set the router
	app.UseHandler(router)

	graceful.Run(fmt.Sprintf(":%v", port), 5*time.Second, app)

	return nil
}
