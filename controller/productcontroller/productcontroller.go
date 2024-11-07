package productcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/riz-it/golang-gin/model"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []model.Product
	model.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	model.DB.Create(&product)
	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func Show(c *gin.Context) {
	var product model.Product
	id := c.Param("id")

	if err := model.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func Update(c *gin.Context) {
	var product model.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})

}

func Delete(c *gin.Context) {
	var product model.Product
	id := c.Param("id")

	param, _ := strconv.ParseInt(id, 10, 64)
	if model.DB.Delete(&product, param).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})

}
