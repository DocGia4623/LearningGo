package controller

import (
	"net/http"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/models"
	"vietanh/gin-gorm-rest/utils"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

// GetDevices godoc
// @Summary Get all devices
// @Description Get all devices
// @Tags devices
// @Accept  json
// @Produce  json
// @Success 200 {array} utils.DeviceResponse
// @Router /device/ [get]
func GetDevices(c *gin.Context) {
	devices := []models.Device{}
	if err := config.DB.Preload("User").Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve devices"})
		return
	}

	var deviceResponses []utils.DeviceResponse

	for _, device := range devices {
		var response utils.DeviceResponse
		// Map the device model to the DeviceResponse struct
		err := mapstructure.Decode(device, &response)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to map device to response"})
			return
		}
		deviceResponses = append(deviceResponses, response)
	}

	c.JSON(http.StatusOK, deviceResponses)
}

// CreateDevice godoc
// @Summary Create a device
// @Description Create a new device
// @Tags devices
// @Accept  json
// @Produce  json
// @Param device body models.Device true "Device"
// @Success 200 {object} models.Device
// @Router /device/ [post]
func CreateDevice(c *gin.Context) {
	var device models.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&device)
	config.DB.Preload("User").First(&device, device.ID)
	c.JSON(200, &device)
}

// DeleteDevice godoc
// @Summary Delete a device
// @Description Delete a device by ID
// @Tags devices
// @Accept  json
// @Produce  json
// @Param id path int true "Device ID"
// @Success 200 {object} models.Device
// @Router /device/{id} [delete]
func DeleteDevice(c *gin.Context) {
	var device models.Device
	config.DB.Where("id = ?", c.Param("id")).Delete(&device)
	c.JSON(200, &device)
}

// UpdateDevice godoc
// @Summary Update a device
// @Description Update a device by ID
// @Tags devices
// @Accept  json
// @Produce  json
// @Param id path int true "Device ID"
// @Param device body models.Device true "Device"
// @Success 200 {object} models.Device
// @Router /device/{id} [put]
func UpdateDevice(c *gin.Context) {
	var device models.Device
	if err := c.BindJSON(&device); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&device)
	config.DB.Preload("User").First(&device, device.ID)
	c.JSON(200, &device)
}

// GetDevice godoc
// @Summary Get a device
// @Description Get a device by ID
// @Tags devices
// @Accept  json
// @Produce  json
// @Param id path int true "Device ID"
// @Success 200 {object} utils.DeviceResponse
// @Router /device/{id} [get]
func GetDevice(c *gin.Context) {
	var device models.Device

	if err := config.DB.Preload("User").Where("id = ?", c.Param("id")).First(&device).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseDevice utils.DeviceResponse
	mapstructure.Decode(device, &responseDevice)

	c.JSON(200, &responseDevice)
}
