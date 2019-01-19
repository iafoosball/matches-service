// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/iafoosball/matches-service/matches"
	"github.com/iafoosball/matches-service/restapi/operations"
	"log"
	"net/http"
)

//go:generate swagger generate server --target .. --name matches --spec ../swagger.yml

var ConfigurationFlags = struct {
	DatabaseHost     string `long:"dbhost" description:"The database host url" default:"arangodb" env:"dbhost"`
	DatabasePort     int    `long:"dbport" description:"The database port" default:"8529" env:"dbport"`
	DatabaseUser     string `long:"dbuser" description:"The database user" default:"root" env:"dbuser"`
	DatabasePassword string `long:"dbpassword" description:"The database password for the user" default:"matches-password" env:"dbpassword"`
}{}

func configureFlags(api *operations.MatchesAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "Additional configuration parameters",
			LongDescription:  "Additional configuration parameters, mainly used to configure the arangodb connection.",
			Options:          &ConfigurationFlags,
		},
	}
}

func configureAPI(api *operations.MatchesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	log.SetFlags(log.Ltime | log.Lshortfile)
	matches.InitDatabase(ConfigurationFlags.DatabaseHost, ConfigurationFlags.DatabasePort, ConfigurationFlags.DatabaseUser, ConfigurationFlags.DatabasePassword)

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	//[Start: Goals end points]
	api.GetGoalsHandler = operations.GetGoalsHandlerFunc(matches.PagedGoals())
	api.PostGoalsHandler = operations.PostGoalsHandlerFunc(matches.CreateGoal())
	api.PostGoalsBatchCreateHandler = operations.PostGoalsBatchCreateHandlerFunc(matches.BatchCreateGoals())
	//[End: Goals end points]

	//[Start: Matches end points]
	api.PostMatchesHandler = operations.PostMatchesHandlerFunc(matches.CreateMatch())
	api.GetMatchesHandler = operations.GetMatchesHandlerFunc(matches.PagedMatches())
	//[End: Matches end points]

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
	matches.Addr(addr)
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	// TODO: encapsulate this into separate function (for visitor pattern over handler i.e. return authenticate(handler)
	// TODO: figure out better pay of passing auth-service address
	// This middleware was commented out for project presentation.
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	v := &auth.JWTValidator {
	//		Protocol:"http",
	//		Hostname: "auth-service",
	//		Port:8070,
	//	}
	//
	//	authStr := r.Header.Get("Authorization")
	//	if ok, err := v.ValidateAuth(authStr); !ok || err != nil {
	//		w.WriteHeader(401)
	//		return
	//	}
	//
	//	handler.ServeHTTP(w, r)
	//})
	return handler
}