package v2

import "github.com/gin-gonic/gin"

func AddProduct(c *gin.Context) {
	// 获取GET参数
	name := c.Query("name")
	price := c.DefaultQuery("price", "100")
	c.JSON(200, gin.H{
		"v2":    "AddProduct",
		"name":  name,
		"price": price,
	})
}
