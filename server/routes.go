package server

import controller "v1/controllers"

func StartRoutes() {
	controller := controller.NewController()
	e.POST("/NewUser", controller.SetUser)
}
