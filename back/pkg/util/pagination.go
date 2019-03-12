/**
 * Created with Goland.
 * Description: 
 * User: Insomnia
 * Date: 2018-11-01
 * Time: 上午12:20
 */

package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"GinBlog/Service/pkg/setting"
)

func GetPage(c *gin.Context) int  {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
