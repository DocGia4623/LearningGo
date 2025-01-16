package routes

import (
	"vietanh/gin-gorm-rest/controller"

	"github.com/gin-gonic/gin"
)

// DeviceRoute defines the device routes
func DeviceRoute(router *gin.Engine) {
	deviceRoutes := router.Group("/device")
	{
		// @Summary Get all devices
		// @Description Get all devices
		// @Tags devices
		// @Accept  json
		// @Produce  json
		// @Success 200 {array} models.Device
		// @Router /device/ [get]
		deviceRoutes.GET("/", controller.GetDevices)

		// @Summary Create a device
		// @Description Create a new device
		// @Tags devices
		// @Accept  json
		// @Produce  json
		// @Param device body models.Device true "Device"
		// @Success 200 {object} models.Device
		// @Router /device/ [post]
		deviceRoutes.POST("/", controller.CreateDevice)

		// @Summary Delete a device
		// @Description Delete a device by ID
		// @Tags devices
		// @Accept  json
		// @Produce  json
		// @Param id path int true "Device ID"
		// @Success 200 {object} models.Device
		// @Router /device/{id} [delete]
		deviceRoutes.DELETE("/:id", controller.DeleteDevice)

		// @Summary Update a device
		// @Description Update a device by ID
		// @Tags devices
		// @Accept  json
		// @Produce  json
		// @Param id path int true "Device ID"
		// @Param device body models.Device true "Device"
		// @Success 200 {object} models.Device
		// @Router /device/{id} [put]
		deviceRoutes.PUT("/:id", controller.UpdateDevice)

		// @Summary Get a device
		// @Description Get a device by ID
		// @Tags devices
		// @Accept  json
		// @Produce  json
		// @Param id path int true "Device ID"
		// @Success 200 {object} models.Device
		// @Router /device/{id} [get]
		deviceRoutes.GET("/:id", controller.GetDevice)
	}
}
