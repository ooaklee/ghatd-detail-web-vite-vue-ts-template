package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ooaklee/ghatd/external/logger"
	loggerMiddleware "github.com/ooaklee/ghatd/external/logger/middleware"
	"github.com/ooaklee/ghatd/external/middleware/contenttype"
	"github.com/ooaklee/ghatd/external/response"
	"github.com/ooaklee/ghatd/external/router"

	//>ghatd {{ define "WebDetailImports" }}
	webapp "github.com/ooaklee/ghatd-detail-web-vite-vue-ts-template/external"
	//>ghatd {{ end }}
)

// content holds our static web server content.
//
//go:embed web/*
var content embed.FS

const serverPort = ":4044"

func main() {

	// Initialise detail logger
	appLogger, err := logger.NewLogger(
		"info",
		"local",
		"ghatd-detail-api",
	)
	if err != nil {
		log.Default().Panicf("server/unable-to-initiate-logger - %v", err)
	}

	tempRouterMiddlewares := []mux.MiddlewareFunc{loggerMiddleware.NewLogger(appLogger).HTTPLogger, contenttype.NewContentType}

	// Initialise router
	httpRouter := router.NewRouter(response.GetResourceNotFoundError, response.GetDefault200Response, tempRouterMiddlewares...)

	// Prep web detail
	NewWebDetail(httpRouter, content, "web/")

	// Define server
	srv := &http.Server{
		Addr:    serverPort,
		Handler: httpRouter.GetRouter(),
	}

	// Start listening
	fmt.Printf("\nServer is listening on port - %s\n", serverPort)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Default().Panicf("server/unable-to-start-server - %v", err)
	}
}

func NewWebDetail(httpRouter *router.Router, embeddedContent fs.FS, embeddedContentFilePathPrefix string) {

	//>ghatd {{ define "WebDetailInit" }}

	// Attach routes
	webapp.AttachRoutes(&webapp.AttachRoutesRequest{
		Router:                        httpRouter,
		WebAppFileSystem:              embeddedContent,
		EmbeddedContentFilePathPrefix: embeddedContentFilePathPrefix,
	})

	//>ghatd {{ end }}
}
