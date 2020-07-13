package routes

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/juanmalv/go-crud-pod/cmd/webservice/routes/handlers"
	"log"
	"time"
)

type GinRouter struct {
	engine *gin.Engine
}

func (r *GinRouter) Configure() {
	r.engine = gin.New()
	r.engine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
}

func (r *GinRouter) Register() {
	r.engine.GET("/ping", handlers.Ping)
	r.engine.GET("/income", handlers.GetIncome)
	r.engine.POST("/income", handlers.NewIncome)
	r.engine.PUT("income", handlers.UpdateIncome)
	r.engine.DELETE("income", handlers.DeleteIncome)
}

func (r *GinRouter) Run() error {
	r.Configure()
	r.Register()

	err := r.engine.Run(":8080")

	if err != nil {
		log.Fatalf("%v",err)
	}

	return err
}
