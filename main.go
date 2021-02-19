package main

import (
	_ "github.com/MiracleWong/go-blog-service/docs"
	"github.com/MiracleWong/go-blog-service/internal/routers"
	"net/http"
	"time"
)



// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout:  10* time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//r := gin.Default()
	//
	//c := controller.NewController()
	//
	//v1 := r.Group("/api/v1")
	//{
	//	accounts := v1.Group("/accounts")
	//	{
	//		accounts.GET(":id", c.ShowAccount)
	//		accounts.GET("", c.ListAccounts)
	//		accounts.POST("", c.AddAccount)
	//		accounts.DELETE(":id", c.DeleteAccount)
	//		accounts.PATCH(":id", c.UpdateAccount)
	//		accounts.POST(":id/images", c.UploadAccountImage)
	//	}
	//	//...
	//}
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.ListenAndServe()
}