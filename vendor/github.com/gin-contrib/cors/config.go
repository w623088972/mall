package cors

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cors struct {
	allowAllOrigins  bool
	allowCredentials bool
	allowOriginFunc  func(string) bool
	allowOrigins     []string
	exposeHeaders    []string
	normalHeaders    http.Header
	preflightHeaders http.Header
}

var (
	DefaultSchemas = []string{
		"http://",
		"https://",
	}
	ExtensionSchemas = []string{
		"chrome-extension://",
		"safari-extension://",
		"moz-extension://",
		"ms-browser-extension://",
	}
	FileSchemas = []string{
		"file://",
	}
	WebSocketSchemas = []string{
		"ws://",
		"wss://",
	}
)

func newCors(config Config) *Cors {
	if err := config.Validate(); err != nil {
		panic(err.Error())
	}
	return &Cors{
		allowOriginFunc:  config.AllowOriginFunc,
		allowAllOrigins:  config.AllowAllOrigins,
		allowCredentials: config.AllowCredentials,
		allowOrigins:     normalize(config.AllowOrigins),
		normalHeaders:    generateNormalHeaders(config),
		preflightHeaders: generatePreflightHeaders(config),
	}

}

func newCorsWithCors(config Config, cors *Cors) *Cors {

	if err := config.Validate(); err != nil {
		panic(err.Error())
	}

	cors.allowOriginFunc = config.AllowOriginFunc
	cors.allowAllOrigins = config.AllowAllOrigins
	cors.allowCredentials = config.AllowCredentials
	cors.allowOrigins = normalize(config.AllowOrigins)
	cors.normalHeaders = generateNormalHeaders(config)
	cors.preflightHeaders = generatePreflightHeaders(config)

	return cors
}

func (cors *Cors) applyCors(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	if len(origin) == 0 {
		// request is not a CORS request
		return
	}
	host := c.Request.Header.Get("Host")

	if origin == "http://"+host || origin == "https://"+host {
		// request is not a CORS request but have origin header.
		// for example, use fetch api
		return
	}

	if !cors.validateOrigin(origin) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if c.Request.Method == "OPTIONS" {
		cors.handlePreflight(c)
		defer c.AbortWithStatus(http.StatusNoContent) // Using 204 is better than 200 when the request status is OPTIONS
	} else {
		cors.handleNormal(c)
	}

	if !cors.allowAllOrigins {
		c.Header("Access-Control-Allow-Origin", origin)
	}
}

func (cors *Cors) validateOrigin(origin string) bool {
	if cors.allowAllOrigins {
		return true
	}
	for _, value := range cors.allowOrigins {
		if value == origin {
			return true
		}
	}
	if cors.allowOriginFunc != nil {
		return cors.allowOriginFunc(origin)
	}
	return false
}

func (cors *Cors) handlePreflight(c *gin.Context) {
	header := c.Writer.Header()
	for key, value := range cors.preflightHeaders {
		header[key] = value
	}
}

func (cors *Cors) handleNormal(c *gin.Context) {
	header := c.Writer.Header()
	for key, value := range cors.normalHeaders {
		header[key] = value
	}
}

func (cors *Cors) ChangeTestOrigin(index int, oUrl string) {
	//log.Printf("cors.allowOrigins", cors.allowOrigins)
	//log.Printf("oUrl", oUrl)
	cors.allowOrigins[index] = oUrl
	log.Printf("cors.allowOrigins after change:", cors.allowOrigins)
}
