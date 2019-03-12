/**
 * Created with Goland.
 * Description: 
 * User: Insomnia
 * Date: 2018-11-04
 * Time: 下午10:54
 */

package models

type Auth struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

/**
*@Description: db check auth
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-4
*/
func CheckAuth(username, password string) bool  {
	var auth Auth

	//这里的select应用 非常有意思
	db.Select("id").Where(Auth{Username: username, Password:password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}