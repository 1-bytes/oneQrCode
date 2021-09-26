package bootstrap

import (
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

var g errgroup.Group

// SetupHttpService 初始化并启动 HTTP 服务.
func SetupHttpService(server *http.Server) {
	// run server
	g.Go(func() error {
		return server.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
