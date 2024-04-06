package common

import (
	"GinStudy/src/main/config"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// 打印
func print(i interface{}) {
	fmt.Println("---")
	fmt.Println(i)
	fmt.Println("---")
}

// 返回JSON
func RetJson(code, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // 使用gin.H（它是map[string]interface{}的别名）
		"code": code,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// MD5方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil)) //这是 Go 标准库 encoding/hex 中的一个函数。
	// 它接受一个字节切片作为参数，并返回这个字节切片的十六进制字符串表示
}

// 生成签名
func CreateSign(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	// 自定义签名算法
	sign := MD5(MD5(str) + MD5(config.APP_NAME+config.APP_SECRET))
	return sign
}

// 验证签名
func VerifySign(c *gin.Context) {
	var method = c.Request.Method
	var ts int64
	var sn string
	//url.Values 是 map[string][]string 的一个别名，它允许你存储多个值对应一个键。
	//例如，如果你有一个URL像 http://example.com/search?q=apple&q=banana，那么它的查询参数可以表示为一个 url.Values 对象如下：
	//req := url.Values{}
	//req.Set("q", "apple")
	//req.Add("q", "banana")
	var req url.Values
	if method == "GET" {
		//例如，如果请求的 URL 是 http://example.com/search?key1=value1&key2=value2，那么 req 将包含两个键值对：
		//"key1" 对应 "value1"，"key2" 对应 "value2"。你可以通过 req.Get("key1") 来获取 "key1" 的值，即 "value1"。
		req = c.Request.URL.Query()
		sn = c.Query("sn")
		ts, _ = strconv.ParseInt(c.Query("sn"), 10, 64) // strconv就是一个处理数字和字符串类型转换的类
	} else if method == "POST" {
		c.Request.ParseForm() //解析表单数据
		req = c.Request.PostForm
		sn = c.PostForm("sn")
		ts, _ = strconv.ParseInt(c.PostForm("sn"), 10, 64)
	} else {
		RetJson("500", "Illegal requests", "", c)
		return
	}
	exp, _ := strconv.ParseInt(config.API_EXPIRY, 10, 64)

	// 验证过期时间
	if ts > GetTimeUnix() || GetTimeUnix()-ts >= exp {
		RetJson("500", "Ts Error", "", c)
		return
	}

	// 验证签名
	if sn == "" || sn != CreateSign(req) {
		RetJson("500", "Sn Error", "", c)
		return
	}
}
