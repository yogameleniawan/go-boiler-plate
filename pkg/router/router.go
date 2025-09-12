package router

import (
	"compress/zlib"
	"net/http"

	"github.com/absendulu-project/backend/internal/health"
	"github.com/absendulu-project/backend/pkg/middleware"
	"github.com/absendulu-project/backend/pkg/response"
	"github.com/go-chi/chi/v5"
	cmiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/unrolled/secure"
)

// SetupRoutes this function for centralize setup all route in this app.
// why wee need to centralize?, it's for easies debugging if any issue
//
// parameters: all interface handlers we need to expose with rest
func SetupRoutes(
	health health.Handler,
) *chi.Mux {
	mux := chi.NewRouter()

	// chi middleware
	mux.Use(cmiddleware.Logger)
	mux.Use(cmiddleware.Recoverer)
	mux.Use(cmiddleware.RealIP)
	mux.Use(cmiddleware.NoCache)
	mux.Use(cmiddleware.GetHead)
	mux.Use(cmiddleware.Compress(zlib.BestCompression))
	mux.Use(cmiddleware.AllowContentType("application/json"))
	mux.Use(secure.New(secure.Options{
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		STSIncludeSubdomains: true,
		STSPreload:           true,
		STSSeconds:           900,
	}).Handler)

	mux.MethodNotAllowed(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := response.JSON{Code: http.StatusMethodNotAllowed, Message: "Route method not allowed"}
		response.ResponseJSON(w, res.Code, res)
	}))

	mux.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := response.JSON{Code: http.StatusNotFound, Message: "Route not found"}
		response.ResponseJSON(w, res.Code, res)
	}))

	// set cors middleware
	mux.Use(middleware.Cors())
	// set middleware rate limiter
	mux.Use(middleware.RateLimit(1000, 10))

	// set prefix v1
	mux.Route("/", func(r chi.Router) {
		// ROUTE TERITORY
		r.Get("/", health.Health)
	})

	return mux
}
