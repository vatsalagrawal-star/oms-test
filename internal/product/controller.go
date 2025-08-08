package product

import (
	"encoding/json"
	"io"
	"net/http"

	"oms-test/models"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ser ProductService
}

func (ctrl *ProductController) ProductFromId(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := ctrl.ser.getProduct(id)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (ctrl *ProductController) SearchProduct(ctx *gin.Context) {
	nameOrCategory := ctx.Query("nameOrCategory")
	products, err := ctrl.ser.searchProducts(nameOrCategory)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while searching products"})
		return
	}
	if len(*products) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (ctrl *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := ctrl.ser.createProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdProduct)
}

func (ctrl *ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingProduct, err := ctrl.ser.getProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	product.ID = existingProduct.ID // Ensure the ID is set for the update
	err = ctrl.ser.updateProduct(existingProduct, &product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (ctrl *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	existingProduct, err := ctrl.ser.getProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctrl.ser.deleteProduct(existingProduct.ID)
	ctx.Status(http.StatusNoContent)
}

func (ctrl *ProductController) InflowProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	body, bodyErr := io.ReadAll(ctx.Request.Body)
	if bodyErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var bodyMap map[string]any
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	quantity, ok := bodyMap["quantity"].(uint)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing quantity"})
		return
	}

	product, prodErr := ctrl.ser.getProduct(id)
	if prodErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	inflowErr := ctrl.ser.inflowProduct(product, uint(quantity))
	if inflowErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to inflow product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product inflow successful"})
}

func (ctrl *ProductController) OutflowProduct(c *gin.Context) {
	id := c.Param("id")
	body, bodyErr := io.ReadAll(c.Request.Body)
	if bodyErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var bodyMap map[string]any
	if err := json.Unmarshal(body, &bodyMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	quantity, ok := bodyMap["quantity"].(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing quantity"})
		return
	}

	product, prodErr := ctrl.ser.getProduct(id)
	if prodErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	outflowErr := ctrl.ser.outflowProduct(product, uint(quantity))
	if outflowErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to outflow product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product outflow successful"})
}

func (ctrl *ProductController) GetAllProducts(c *gin.Context) {
	products, err := ctrl.ser.getAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	if len(*products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
		return
	}

	c.JSON(http.StatusOK, products)
}
