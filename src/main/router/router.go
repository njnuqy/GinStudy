package router

import (
	"GinStudy/src/main/common"
	v1 "GinStudy/src/main/controller/v1"
	v2 "GinStudy/src/main/controller/v2"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

func InitRouter(r *gin.Engine) {
	r.GET("/sn", signDemo)

	//v1版本
	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/product/add", v1.AddProduct)
		GroupV1.Any("/member/add", v1.AddMember)
	}

	// v2版本
	GroupV2 := r.Group("/v2")
	{
		GroupV2.Any("/product/add", v2.AddProduct)
		GroupV2.Any("/member/add", v2.AddMember)
	}
}

// 在 Gin 的中间件链中，每个中间件都会接收一个 gin.Context 对象，并可能修改它（比如设置新的响应头或响应体），
// 然后将它传递给下一个中间件。最终，当中间件链处理完毕，Gin 会发送一个 HTTP 响应给客户端。
// 因此，当你看到函数签名中有 c *gin.Context 这样的参数时，它表示该函数接收一个指向 gin.Context 结构体的指针作为参数，
// 通过这个指针，函数可以访问和修改请求的上下文信息。
func signDemo(c *gin.Context) {
	ts := strconv.FormatInt(common.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name":  []string{"a"},
		"price": []string{"10"},
		"ts":    []string{ts},
	}
	res["sn"] = common.CreateSign(params)
	res["ts"] = ts
	common.RetJson("200", "", res, c)
}
