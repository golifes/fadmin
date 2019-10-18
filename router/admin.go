package router

func (e *Engine) admin() {
	r := e.Group("/admin")
	{
		//r.POST("/register", e.Register)
		//r.POST("/login", e.Login)
		domain := r.Group("domain")
		{
			domain.POST("/add", e.AddDomain)
			domain.POST("/delete", e.DeleteDomain)
			domain.GET("/find", e.FindDomain)
		}

	}
}
