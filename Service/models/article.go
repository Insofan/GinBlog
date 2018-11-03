/**
 * Created with Goland.
 * Description:
 * User: Insomnia
 * Date: 2018-11-03
 * Time: 下午10:24
 */

package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	/*这里和gorm 的的外键有关系, 约定俗成TagId作为Tag的外键*/
	TagId int `json:"tag_id" gorm:"index""`
	/*这里tag是嵌套, 很有意义*/
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/**
*@Description: check if exist article
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-3
 */
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.Id > 0 {
		return true
	}
	return false
}

/**
*@Description: Get articles total
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-3
 */
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

/**
*@Description: Get all articles
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-3
 */
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	//offset 这里
	/*
		Preload就是一个预加载器，它会执行两条SQL，分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询
	*/
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

/**
*@Description: 这部分演示了gorm, article如何关联到tag
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-3
 */
func GetArticle(id int) (article Article) {
	/*首先是gorm本身做了大量的约定俗成*/
	db.Where("id = ?", id).First(&article)
	/*
		Article有一个结构体成员是TagID，就是外键。gorm会通过类名+ID的方式去找到这两个类之间的关联关系
		Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
	*/
	db.Model(&article).Related(&article.Tag)
	return
}

/**
*@Description: Edit article
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-3
 */
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

/**
*@Description: 添加文章
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-3
 */
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagId:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

/**
*@Description: 删除
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-3
 */
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

/**
*@Description: 实现gorm的回调
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-3
 */
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

/**
*@Description: 实现gorm的回调
*@Param:
*@return:
*@Author: Insomnia
*@date: 18-11-3
 */

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
