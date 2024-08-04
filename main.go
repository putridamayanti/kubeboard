package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"kubeboard/controllers"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	controllers.ClusterConnection()

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())

	api := router.Group("/api")
	{
		api.GET("/health", controllers.CheckHealth)

		api.GET("/switch-cluster/:filename", controllers.SwitchConnection)

		api.GET("/namespace", controllers.GetNamespaces)
		api.GET("/namespace/:name", controllers.GetNamespaceByName)
		api.POST("/namespace", controllers.CreateNamespace)
		api.PATCH("/namespace", controllers.UpdateNamespace)
		api.DELETE("/namespace/:name", controllers.DeleteNamespace)

		api.GET("/node", controllers.GetNodes)
		api.GET("/node/:name", controllers.GetNodeByName)

		api.GET("/crd", controllers.GetCustomResourceDefinitions)
		api.GET("/crd/:name", controllers.GetCustomResourceDefinitionByName)
		api.POST("/crd", controllers.CreateCustomResourceDefinition)
		api.PATCH("/crd", controllers.UpdateCustomResourceDefinition)
		api.DELETE("/crd/:name", controllers.DeleteCustomResourceDefinition)

		app := api.Group("/:namespace")
		{
			app.GET("/config-map", controllers.GetConfigMaps)
			app.GET("/config-map/:name", controllers.GetConfigMapByName)
			app.POST("/config-map", controllers.CreateConfigMap)
			app.PATCH("/config-map", controllers.UpdateConfigMap)
			app.DELETE("/config-map/:name", controllers.DeleteConfigMap)

			app.GET("/deployment", controllers.GetDeployments)
			app.GET("/deployment/:name", controllers.GetDeploymentByName)
			app.POST("/deployment", controllers.CreateDeployment)
			app.PATCH("/deployment", controllers.UpdateDeployment)
			app.DELETE("/deployment/:name", controllers.DeleteDeployment)

			app.GET("/pod", controllers.GetPods)
			app.GET("/pod/:name", controllers.GetPodByName)
			app.POST("/pod", controllers.CreatePod)
			app.PATCH("/pod", controllers.UpdatePod)
			app.DELETE("/pod/:name", controllers.DeletePod)

			app.GET("/secret", controllers.GetSecrets)
			app.GET("/secret/:name", controllers.GetSecretByName)
			app.POST("/secret", controllers.CreateSecret)
			app.PATCH("/secret", controllers.UpdateSecret)
			app.DELETE("/secret/:name", controllers.DeleteSecret)

			app.GET("/service", controllers.GetServices)
			app.GET("/service/:name", controllers.GetServiceByName)
			app.POST("/service", controllers.CreateService)
			app.PATCH("/service", controllers.UpdateService)
			app.DELETE("/service/:name", controllers.DeleteService)
		}
	}

	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	err = router.Run(":" + port)
	if err != nil {
		return
	}
}
