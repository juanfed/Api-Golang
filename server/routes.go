package server

import (
	controller "v1/controllers"
	"v1/dal"
)

func StartRoutes() {
	mysql := dal.NewDatabaseSql()
	userController := controller.NewController(mysql)
	e.POST("/user", userController.Set)
	e.DELETE("/deleteUser/", userController.Delete)
	e.PUT("/user", userController.Update)
	e.GET("/user/:id", userController.Get)

	e.GET("/allUser", userController.GetAll)

}
