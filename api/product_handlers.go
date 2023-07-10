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
		c.IndentedJSON(http.StatusBadRequest, types.APIError{Message: err.Error()})
		return
	}

	product, err := s.store.GetProductByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, types.APIError{Message: err.Error()})
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
		return
	}

	newProduct := types.CreateProduct(newProductReq.Name, newProductReq.Description, newProductReq.WeightInKG, newProductReq.PiecesPerPackage, newProductReq.Image, newProductReq.Manufacturer, newProductReq.Category)

	if err := s.store.CreateProduct(newProduct); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newProduct)
}
