package middlewares

import (
    "net/http"
    "github.com/getfider/fider/app/pkg/web"
)

// CORS adds Cross-Origin Resource Sharing response headers
func CORS() web.MiddlewareFunc {
    return func(next web.HandlerFunc) web.HandlerFunc {
        return func(c *web.Context) error {
            c.Response.Header().Set("Access-Control-Allow-Origin", "*")  // Consider specifying domains in production
            c.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            c.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
            
            // Handle preflight requests
            if c.Request.Method == "OPTIONS" {
                c.Response.WriteHeader(http.StatusOK)
                return nil
            }
            
            return next(c)
        }
    }
}

