package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/razielblood/corciel_inventory_manager/types"
)

func (s APIServer) handleGetCategories(c *gin.Context) {
	categories, err := s.store.GetCategories()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, categories)
}

func (s APIServer) handleGetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid id '%v'", c.Param("id"))})
		return
	}

	category, err := s.store.GetCategoryByID(categoryID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, category)

}

func (s APIServer) handlePostCategory(c *gin.Context) {
	createCategoryReq := new(types.CreateCategoryRequest)

	if err := c.BindJSON(createCategoryReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCategory := types.CreateCategory(createCategoryReq.Name, createCategoryReq.Description)

	if err := s.store.CreateCategory(newCategory); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, newCategory)
}

func (s APIServer) handlePutCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid id '%v'", c.Param("id"))})
		return
	}
	updateCategoryReq := new(types.CreateCategoryRequest)

	if err := c.BindJSON(updateCategoryReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory := types.CreateCategory(updateCategoryReq.Name, updateCategoryReq.Description)
	updatedCategory.ID = categoryID

	err = s.store.UpdateCategory(updatedCategory)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedCategory)
}
