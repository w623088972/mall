package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"myself/mall/model"
	"net/http"
	"strconv"
)

func PrepareContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		//注册来源
		registerSource := c.Request.Header.Get("register_source")
		c.Set("register_source", registerSource)

		//推广渠道
		channel := c.Request.Header.Get("channel")
		c.Set("channel", channel)

		//用户id
		userId, _ := strconv.Atoi(c.Request.Header.Get("x-gateway-user_id"))
		c.Set("user_id", userId)

		//用户角色id
		roleId, _ := strconv.Atoi(c.Request.Header.Get("x-gateway-role_id"))
		c.Set("role_id", roleId)

		projectId, _ := strconv.Atoi(c.Request.Header.Get("x-gateway-project_id"))
		if projectId == -1 {
			pid, _ := strconv.Atoi(c.Request.Header.Get("project_id"))
			if pid > 0 {
				db, err := model.GetDb(pid, 0)
				if err != nil {
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}
				c.Set("project_id", pid)
				c.Set("db", db.DB)
			}
			return
		}
		c.Set("project_id", projectId)

		db, err := model.GetDb(projectId, userId)
		if err != nil {
			log.Printf("failed to get db for projectId:%d userId:%d %s\n", projectId, userId, err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Set("db", db.DB)

	}
}
