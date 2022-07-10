package routes

import (
	"fmt"
	"k8sAPI/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// parm db *gorm.DB
func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	fmt.Println("실행 routes")

	r.POST("/k8s/create-pod-service", controllers.CreateKubeService)
	r.POST("/k8s/start-service", controllers.StartService)
	r.POST("/k8s/test-restart", controllers.ReStart)
	r.POST("/k8s/test-new-namespace", controllers.NewNamespace)
	r.DELETE("/k8s/delete-pod-service", controllers.DeleteKubeService)
	r.DELETE("/k8s/stop-service", controllers.StopService)
	r.GET("/k8s/list-all", controllers.ListAll)
	r.GET("/k8s/list-pod", controllers.ListPod)
	r.GET("/k8s/list-service", controllers.ListService)
	r.GET("/k8s/list-deployment", controllers.ListDeplyoment)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
