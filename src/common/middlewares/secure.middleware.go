package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func SecureMiddleware() gin.HandlerFunc {

	secureMiddleware := secure.New(secure.Options{
		FrameDeny:                 true,
		BrowserXssFilter:          true,
		ContentTypeNosniff:        true,
		ContentSecurityPolicy:     "default-src 'self'",
		ReferrerPolicy:            "same-origin",
		PermissionsPolicy:         "fullscreen=(), geolocation=()",
		CrossOriginOpenerPolicy:   "same-origin",
		CrossOriginEmbedderPolicy: "require-corp",
		CrossOriginResourcePolicy: "same-origin",
		FeaturePolicy:             "vibrate 'none';",
	})

	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatus(400)
			return
		}
		c.Next()
	}
}
