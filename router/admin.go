package router

func admin(e *Engine) {
	r := e.Group("/admin")
	{
		//r.POST("/register", e.Register)
		//r.POST("/login", e.Login)
		domain := r.Group("domain")
		{
			domain.POST("/domain", e.AddDomain)
			domain.POST("/deleteDomain", e.DeleteDomain)
			domain.GET("/domain", e.FindDomain)
		}

	}
}
