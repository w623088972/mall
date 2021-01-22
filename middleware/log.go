package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"myself/mall/conf"
)

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		/* 				reg := regexp.MustCompile("(/v1/user|/login)")
		   				if !reg.MatchString(path) {
		   					return
		   				} */

		// Skip for the health check requests.
		//if path == "/sd/health" || path == "/s/ram" || path == "/sd/cpu" || path == "/sd/disk" {

		if path == "/" || path == "/health" {
			return
		}

		// The basic informations.
		method := c.Request.Method
		ip := c.ClientIP()

		//渠道 平台 设备码
		/* 	channel := c.Request.Header.Get("channel")
		log.Println("request register_source", source)
		c.Set("channel", channel) */

		//uuid
		requestId := uuid.NewV4()
		c.Set("requestId", requestId.String())

		/* 		conf.LOG.Self.WithFields(logrus.Fields{
			"ip":        ip,
			"method":    method,
			"path":      path,
			"requestID": requestID,
		}).Info("GIN route start") */

		//语言
		language := c.Request.Header.Get("language")
		if language == "" {
			language = "en" //en英语，chs简中，cht繁中，默认英语
		}
		c.Set("language", language)

		//平台
		platform := c.Request.Header.Get("platform")
		if platform == "" {
			platform = "unknown" //ios android
		}
		c.Set("platform", platform)

		//版本
		version := c.Request.Header.Get("version")
		c.Set("version", version)

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		conf.LOG.Self.WithFields(logrus.Fields{
			"latency":   latency,
			"ip":        ip,
			"method":    method,
			"path":      path,
			"language":  language,
			"version":   version,
			"requestId": requestId,
			"platform":  platform,
		}).Info("GIN route end")

		/*
			x_forwarded_for := c.Request.Header.Get("x-forwarded-for")
			Proxy_Client_IP := c.Request.Header.Get("Proxy-Client-IP")
			WL_Proxy_Client_IP := c.Request.Header.Get("WL-Proxy-Client-IP")
			HTTP_CLIENT_IP := c.Request.Header.Get("HTTP_CLIENT_IP")
			HTTP_X_FORWARDED_FOR := c.Request.Header.Get("HTTP_X_FORWARDED_FOR")

			X_Forwarded_For := c.Request.Header.Get("X-Forwarded-For")
			X_Real_Ip := c.Request.Header.Get("X-Real-Ip")
			X_Appengine_Remote_Addr := c.Request.Header.Get("X-Appengine-Remote-Addr")
			RemoteAddr := c.Request.RemoteAddr
			x_Original_Forwarde_For := c.Request.Header.Get("x-Original-Forwarded-For")

			conf.LOG.Self.WithFields(logrus.Fields{
				"x_forwarded_for":         x_forwarded_for,
				"Proxy_Client_IP":         Proxy_Client_IP,
				"WL_Proxy_Client_IP":      WL_Proxy_Client_IP,
				"HTTP_CLIENT_IP":          HTTP_CLIENT_IP,
				"HTTP_X_FORWARDED_FOR":    HTTP_X_FORWARDED_FOR,
				"X_Forwarded_For":         X_Forwarded_For,
				"X_Real_Ip":               X_Real_Ip,
				"X_Appengine_Remote_Addr": X_Appengine_Remote_Addr,
				"X_RRemoteAddral_Ip":      RemoteAddr,
				"x_Original_Forwarde_For": x_Original_Forwarde_For,
			}).Info("GIN route end")
		*/
	}
}
