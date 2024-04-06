package v1

import "github.com/gin-gonic/gin"

func AddMember(c *gin.Context) {
	// 获取GET参数
	name := c.Query("name")
	price := c.DefaultQuery("price", "100") //如果请求中不存在这个查询参数，那么它将使用默认值 "100"。
	c.JSON(200, gin.H{
		"v1":    "AddMember",
		"name":  name,
		"price": price,
	})
}
