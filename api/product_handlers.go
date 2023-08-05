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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func (s APIServer) handleGetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := s.store.GetProductByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct := types.CreateProduct(newProductReq.Name, newProductReq.Description, newProductReq.WeightInKG, newProductReq.PiecesPerPackage, newProductReq.Image, newProductReq.Brand, newProductReq.Category)

	if err := s.store.CreateProduct(newProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var err error
	newProduct.Category, err = s.store.GetCategoryByID(newProduct.Category.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newProduct.Brand, err = s.store.GetBrandByID(newProduct.Brand.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func (s APIServer) handlePutProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProductReq := new(types.UpdateProductRequest)

	if err := c.BindJSON(updateProductReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProduct, err := s.store.GetProductByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := s.store.GetCategoryByID(updateProductReq.Category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBrand, err := s.store.GetBrandByID(updateProductReq.Brand)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProduct.Name = updateProductReq.Name
	updateProduct.Description = updateProductReq.Description
	updateProduct.WeightInKG = updateProductReq.WeightInKG
	updateProduct.PiecesPerPackage = updateProductReq.PiecesPerPackage
	updateProduct.Image = updateProductReq.Image
	updateProduct.Category = updatedCategory
	updateProduct.Brand = updatedBrand

	if err := s.store.UpdateProduct(updateProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updateProduct)
}
