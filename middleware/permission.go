package middleware

/*
func CheckRolePermission(c *gin.Context) {
	method := c.Request.Method
	path := c.Request.URL.Path
	if method == http.MethodOptions {
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	roleId := c.MustGet("role_id").(int)
	log.Printf("checkRolePermission method:%s path:%s roleId:%d", method, path, roleId)

	permissions, err := systemModel.ListRolePermissions(db, roleId)
	if err != nil && err != gorm.ErrRecordNotFound {
		handler.SendResponse(c, errno.ErrDatabase, "CheckRolePermission", nil)
		c.Abort()
		return
	}
	if err == gorm.ErrRecordNotFound {
		handler.SendResponse(c, errno.ErrPermissionDenied, "CheckRolePermission no permissions found.", nil)
		c.Abort()
		return
	}
	route, err := systemModel.DetailRoute(db, c.Request.URL.Path, method)
	if err != nil {
		handler.SendResponse(c, errno.ErrPermissionDenied, "CheckRolePermission no route."+err.Error(), nil)
		c.Abort()
		return
	}

	permission := systemModel.RolePermissionRecord{}
	for _, p := range permissions {
		if p.ModuleID == route.ModuleId {
			permission = p
			break
		}
	}
	if permission.Id == 0 {
		handler.SendResponse(c, errno.ErrPermissionDenied, "CheckRolePermission permission id.", nil)
		c.Abort()
		return
	}
	modulePermission := uint32(permission.Permission)
	switch method {
	case "GET":
		if !systemModel.IsCapableRead(modulePermission) {
			goto out
		}
	case "PUT":
		if !systemModel.IsCapableUpdate(modulePermission) {
			goto out
		}
	case "POST":
		if !systemModel.IsCapableAdd(modulePermission) {
			goto out
		}
	case "DELETE":
		if !systemModel.IsCapableAdd(modulePermission) {
			goto out
		}
	}
	log.Printf("CheckRolePermission method:%s path:%s roleId:%d",
		method, path, roleId, modulePermission)
	return
out:
	msg := fmt.Sprintf("method: %s path:%s\n", method, c.Request.URL.Path)
	handler.SendResponse(c, errno.ErrPermissionDenied, "CheckRolePermission no permission."+msg, nil)
	c.Abort()
}
*/
