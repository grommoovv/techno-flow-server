package handler

import (
	"github.com/gin-gonic/gin"
	"server-techno-flow/internal/service"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *gin.Engine {
	handler := gin.Default()

	handler.Use(corsMiddleware)

	api := handler.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-in", h.signIn)
			auth.POST("/sign-out", h.signOut)
			auth.GET("/refresh", h.refresh)
		}

		authenticated := api.Group("")
		{
			users := authenticated.Group("/users")
			{
				users.POST("", h.CreateUser)
				users.GET("", h.GetAllUsers)
				users.GET("/:id", h.GetUserById)
				users.DELETE("/:id", h.DeleteUser)
				users.PATCH("/:id", h.UpdateUser)
			}

			equipment := authenticated.Group("/equipment")
			{
				equipment.POST("", h.CreateEquipment)
				equipment.GET("", h.GetAllEquipment)
				equipment.GET("/:id", h.GetEquipmentById)
				equipment.GET("/usage/:id", h.GetEquipmentReservationDatesById)
				equipment.DELETE("/:id", h.DeleteEquipment)
				equipment.PATCH("/:id", h.UpdateEquipment)
			}

			events := authenticated.Group("/events")
			{
				events.POST("", h.CreateEvent)
				events.GET("", h.GetAllEvents)
				events.GET("/:id", h.GetEventById)
				events.DELETE("/:id", h.DeleteEvent)
				events.PUT("/:id", h.UpdateEvent)
			}

			reports := authenticated.Group("/reports")
			{
				reports.POST("", h.CreateReport)
				reports.GET("", h.GetAllReports)
				reports.GET("/:id", h.GetReportById)
				reports.DELETE("/:id", h.DeleteReport)
				reports.PUT("/:id", h.UpdateReport)
			}

			maintenance := authenticated.Group("/maintenance")
			{
				maintenance.POST("", h.CreateMaintenance)
				maintenance.GET("", h.GetAllMaintenance)
				maintenance.GET("/:id", h.GetMaintenance)
				maintenance.DELETE("/:id", h.DeleteMaintenance)
				maintenance.PUT("/:id", h.UpdateMaintenance)
			}
		}
	}

	return handler
}
