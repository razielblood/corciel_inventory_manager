package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/razielblood/corciel_inventory_manager/types"
)

func (s APIServer) handleGetManufacturers(c *gin.Context) {
	manufacturers, err := s.store.GetManufacturers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, manufacturers)
}

func (s APIServer) handleGetManufacturerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manufacturer, err := s.store.GetManufacturerByID(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, manufacturer)

}

func (s APIServer) handlePostManufacturer(c *gin.Context) {
	newManufacturerReq := new(types.CreateManufacturerRequest)

	if err := c.BindJSON(newManufacturerReq); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newManufacturer := types.CreateManufacturer(newManufacturerReq.Name)

	err := s.store.CreateManufacturer(newManufacturer)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newManufacturer)
}

func (s APIServer) handlePutManufacturer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedManufacturerReq := new(types.Manufacturer)

	if err := c.BindJSON(updatedManufacturerReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedManufacturer := types.CreateManufacturer(updatedManufacturerReq.Name)
	updatedManufacturer.ID = id

	s.store.UpdateManufacturer(updatedManufacturer)

	c.IndentedJSON(http.StatusOK, updatedManufacturer)
}
