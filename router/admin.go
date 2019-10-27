package router

func (e *Engine) admin() {
	r := e.Group("/admin")
	{
		r.POST("/register", e.Register)
		r.POST("/login", e.Login)
		domain := r.Group("domain")
		{
			domain.POST("/add", e.AddDomain)
			domain.POST("/del", e.DeleteDomain)
			domain.GET("/find", e.FindDomain)
			domain.POST("/update", e.UpdateDomain)
		}
		app := r.Group("app")
		{
			app.POST("/add", e.AddApp)
			app.POST("/del", e.DeleteApp)
			app.GET("/find", e.FindApp)
			app.POST("/update", e.UpdateApp)
		}
		role := r.Group("role")
		{
			role.POST("/add", e.AddRole)
			role.POST("/del", e.DeleteRole)
			role.GET("/find", e.FindRole)
			role.POST("/app", e.UpdateRole) //给app重新调整角色
			role.POST("/name", e.UpdateRoleName)
		}

		user := r.Group("/user")
		{
			user.POST("/forbid", e.ForbidUser)
			user.POST("/pwd", e.UpdatePwd)
			user.POST("/phone", e.UpdatePhone)
			user.POST("/loginPhone", e.LoginPhone) //手机号码登录
			user.POST("/gid", e.UserGroup)         //给用户分配组
		}
		group := r.Group("/group")
		{
			group.POST("/add", e.AddGroup)
			group.POST("/del", e.DeleteGroup)
			group.POST("/update", e.UpdateGroup)
			group.GET("/find", e.FindGroup) //查询组
		}
	}
}
