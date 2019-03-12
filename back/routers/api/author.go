/**包
 * Created with Goland.
 * Description: 
 * User: Insomnia
 * Date: 2018-11-04
 * Time: 下午10:58
 */

package api

import (
	"GinBlog/Service/models"
	"GinBlog/Service/pkg/e"
	"GinBlog/Service/pkg/util"
	"GinBlog/Service/pkg/logging"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}
// @Summary 获取Token
// @Produce  json
// @Param username query string true "UserName"
// @Param password query string true "Password"
// @Success 200 {string} json "{"code":200,"data":{"token": },"msg":"ok"}"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok , _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}
