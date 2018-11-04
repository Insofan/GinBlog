/**
 * Created with Goland.
 * Description: 
 * User: Insomnia
 * Date: 2018-11-04
 * Time: 下午10:50
 */

package jwt

import (
	"GinBlog/Service/pkg/e"
	"GinBlog/Service/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
*@Description: jwt token 中间件
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-4
*/
func JWT() gin.HandlerFunc  {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS

		token := c.Query("token")

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
