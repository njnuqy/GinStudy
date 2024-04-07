package sign

import (
	"GinStudy/src/main/common"
	"GinStudy/src/main/config"
	"GinStudy/src/main/entity"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := entity.Result{}
		sign, err := verifySign(c)
		if sign != nil {
			res.SetCode(entity.CODE_ERROR)
			res.SetMessage("Debug sign")
			res.SetData(sign)
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		if err != nil {
			res.SetCode(entity.CODE_ERROR)
			res.SetMessage(err.Error())
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		c.Next()
	}
}

// 验证签名
func verifySign(c *gin.Context) (res map[string]string, err error) {
	var method = c.Request.Method
	var ts int64
	var sn string
	var req url.Values
	if method == "GET" {
		req = c.Request.URL.Query()
		sn = c.Query("sn")
		ts, _ = strconv.ParseInt(c.Query("ts"), 10, 64)
	} else if method == "POST" {
		c.Request.ParseForm()
		req = c.Request.PostForm
		sn = c.Query("sn")
		ts, _ = strconv.ParseInt(c.Query("ts"), 10, 64)
	} else {
		err = errors.New("非法请求")
		return nil, err
	}
	exp, _ := strconv.ParseInt(config.API_EXPIRY, 10, 64)
	// 验证过期时间
	timestamp := time.Now().Unix()
	if ts > timestamp || timestamp-ts > exp {
		err = errors.New("Ts Error")
		return nil, err
	}
	// 验证签名
	if sn == "" || sn != common.CreateSign(req) {
		err = errors.New("sn Error")
		return nil, err
	}
	return
}
