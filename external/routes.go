package webapp

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/ooaklee/ghatd/external/router"
)

// webAppHandler expected methods for valid web app handler
type webAppHandler interface {
}

// AttachRoutesRequest holds everything needed to attach web app
// routes to router
type AttachRoutesRequest struct {
	// Router main router being served by server
	Router *router.Router

	// Handler valid web app handler
	Handler webAppHandler

	// WebAppFileSystem the file system that holds files utilised
	// by the web app
	WebAppFileSystem fs.FS

	// EmbeddedContentFilePathPrefix the prefix used to access the embedded files
	EmbeddedContentFilePathPrefix string
}

// AttachRoutes attaches webApp handler to corresponding
// routes on router
func AttachRoutes(request *AttachRoutesRequest) {

	// Create filesystem only holding static assets
	staticSubFS, err := fs.Sub(request.WebAppFileSystem, fmt.Sprintf("%sdist", request.EmbeddedContentFilePathPrefix))
	if err != nil {
		log.Default().Panicln("unable-to-create-file-system-for-static-assets", err)
		return
	}

	httpRouter := request.Router.GetRouter()

	// Create path for handling static assets
	httpRouter.PathPrefix("/").Handler(http.FileServer(http.FS(staticSubFS)))
	http.Handle("/", http.FileServer(http.Dir("/")))
}
