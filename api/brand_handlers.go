package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/razielblood/corciel_inventory_manager/types"
)

func (s APIServer) handleGetBrands(c *gin.Context) {
	brands, err := s.store.GetBrands()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, brands)
}

func (s APIServer) handleGetBrandByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand, err := s.store.GetBrandByID(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, brand)

}

func (s APIServer) handlePostBrand(c *gin.Context) {
	newBrandReq := new(types.CreateBrandRequest)

	if err := c.BindJSON(newBrandReq); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBrand := types.CreateBrand(newBrandReq.Name, newBrandReq.Manufacturer)

	err := s.store.CreateBrand(newBrand)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newBrand)
}

func (s APIServer) handlePutBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBrandReq := new(types.Brand)

	if err := c.BindJSON(updatedBrandReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedBrand := types.CreateBrand(updatedBrandReq.Name, updatedBrandReq.Manufacturer.ID)
	updatedBrand.ID = id

	s.store.UpdateBrand(updatedBrand)

	c.IndentedJSON(http.StatusOK, updatedBrand)
}
