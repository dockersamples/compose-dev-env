package app

import (
	"aquafarm-management/app/config"
	"aquafarm-management/app/handler"
	"aquafarm-management/app/repository"
	"aquafarm-management/app/usecase"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func New(cfg *config.Config) *App {
	a := &App{
		Router: gin.Default(),
	}

	db := repository.NewDB(cfg)

	itemRepository := repository.NewItemRepository(db)
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	itemHandler := handler.NewItemHandler(itemUsecase)

	v1 := a.Router.Group("/v1")
	{
		pond := v1.Group("/pond")
		{
			pond.GET("/", itemHandler.Fetch)
			pond.GET("/:id", itemHandler.Get)
			pond.POST("/", itemHandler.Store)
			pond.PUT("/:id", itemHandler.Update)
			pond.DELETE("/:id", itemHandler.Delete)
		}

		farm := v1.Group("/farm")
		{
			farm.GET("/", itemHandler.Fetch)
			farm.GET("/:id", itemHandler.Get)
			farm.POST("/", itemHandler.Store)
			farm.PUT("/:id", itemHandler.Update)
			farm.DELETE("/:id", itemHandler.Delete)
		}

		logs := v1.Group("/logs")
		{
			logs.GET("/", itemHandler.Fetch)
			logs.GET("/statistics", itemHandler.Fetch)
		}
	}

	return a
}

func (a *App) Run(cfg *config.Config) error {
	return a.Router.Run(":" + cfg.Server.Port)
}
