package routes

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/juanmalv/go-crud-pod/cmd/webservice/routes/handlers"
)

type Router interface {
	Configure()
	Register()
	Run() error
}

type GinRouter struct {
	engine *gin.Engine
}

//todo: Check if is a write method. If not, change type to non-pointer.
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
	// r.engine.GET("/income", handlers.GetIncome)
	// r.engine.POST("/income", handlers.NewIncome)
	// r.engine.GET("/expense", handlers.GetExpenses)
	// r.engine.POST("/expense", handlers.NewExpense)
}

func (r *GinRouter) Run() error {
	r.Configure()
	r.Register()

	err := r.engine.Run(":8080")

	if err != nil {
		log.Fatalf("error running server", err)
	}
	return err
}
