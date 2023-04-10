package router

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/internal/domain"
	"desafio-go-web/internal/tickets"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	list   []domain.Ticket
}

func NewRouter(engine *gin.Engine, list []domain.Ticket) *Router {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return &Router{
		engine: engine,
		list:   list,
	}
}

func (r *Router) MapRoutes() {
	repository := tickets.NewRepository(r.list)
	service := tickets.NewService(repository)
	productHandler := handler.NewService(service)

	productGroup := r.engine.Group("/ticket")
	{
		productGroup.GET("getByCountry/:dest", productHandler.GetTicketsByCountry())
		productGroup.GET("/getAverage/:dest", productHandler.AverageDestination())
	}
}
