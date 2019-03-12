/**
 * Created with Goland.
 * Description: 
 * User: Insomnia
 * Date: 2018-11-04
 * Time: 下午10:13
 */

package util

import (
	"GinBlog/Service/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

/**
*@Description: 生成token
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-4
*/
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	//这个有三种加密方案 hs256, hs384, hs512
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//方法内部生成签名字符串, 再用于获取完整, 已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

/**
*@Description: 解析token
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-4
*/
func ParseToken(token string) (*Claims, error)  {
	//用来解析声明, 主要负责具体的解码校验过程, 最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
