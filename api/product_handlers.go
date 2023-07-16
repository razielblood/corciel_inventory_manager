package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/razielblood/corciel_inventory_manager/types"
)

func (s APIServer) handleGetProducts(c *gin.Context) {
	products, err := s.store.GetProducts()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func (s APIServer) handleGetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product, err := s.store.GetProductByID(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

// postProduct adds a product from JSON received in the request body.
func (s APIServer) handlePostProduct(c *gin.Context) {
	newProductReq := new(types.CreateProductRequest)

	// Call BindJSON to bind the received JSON to
	// newProduct.
	if err := c.BindJSON(newProductReq); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newProduct := types.CreateProduct(newProductReq.Name, newProductReq.Description, newProductReq.WeightInKG, newProductReq.PiecesPerPackage, newProductReq.Image, newProductReq.Manufacturer, newProductReq.Category)

	if err := s.store.CreateProduct(newProduct); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newProduct)
}

func (s APIServer) handlePutProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updateProductReq := new(types.UpdateProductRequest)

	if err := c.BindJSON(updateProductReq); err != nil {
		return
	}

	updateProduct, err := s.store.GetProductByID(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedCategory, err := s.store.GetCategoryByID(updateProductReq.Category)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedManufacturer, err := s.store.GetManufacturerByID(updateProductReq.Manufacturer)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updateProduct.Name = updateProductReq.Name
	updateProduct.Description = updateProductReq.Description
	updateProduct.WeightInKG = updateProductReq.WeightInKG
	updateProduct.PiecesPerPackage = updateProductReq.PiecesPerPackage
	updateProduct.Image = updateProductReq.Image
	updateProduct.Category = updatedCategory
	updateProduct.Manufacturer = updatedManufacturer

	if err := s.store.UpdateProduct(updateProduct); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, updateProduct)
}
