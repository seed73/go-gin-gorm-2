package main

import (
	"k8sAPI/models"
	"k8sAPI/routes"
)

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @title K8S Swagger Example API
// @version 1.0
// @description GO Gin Gorm
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	// //DB연결
	db := models.SetupDB()

	//Router설정 parm db
	r := routes.SetupRoutes(db)
	//DB자동생성
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.K8s_Pod_List{})
	db.AutoMigrate(&models.K8s_Pod_Detail{})
	db.AutoMigrate(&models.K8s_Pod_Option_List{})
	db.AutoMigrate(&models.K8s_Pod_Option_Detail{})
	db.AutoMigrate(&models.K8s_Service_List{})
	db.AutoMigrate(&models.K8s_Service_Detail{})
	db.AutoMigrate(&models.K8s_Image_List{})

	r.Run()
}
