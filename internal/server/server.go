// Package server assembles HTTP routes and handlers.
package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ravenmk2/baodou/web"
)

// New assembles routes and returns the root HTTP handler.
func New(db *sql.DB) http.Handler {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		if err := db.PingContext(c.Request.Context()); err != nil {
			c.String(http.StatusServiceUnavailable, "db unavailable")
			return
		}
		c.String(http.StatusOK, "ok")
	})

	// SPA fallback: unmatched routes fall back to index.html for the frontend router
	r.NoRoute(gin.WrapH(SPAHandler(web.DistFS())))
	return r
}
