package admin

import "github.com/gin-gonic/gin"

type Blog struct {
}

func NewBlog() Blog {
	return Blog{}
}

func (b Blog) Get(c *gin.Context) {

}

func (b Blog) GetList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "200",
	})

}
func (b Blog) Delete(c *gin.Context) {

}
func (b Blog) Update(c *gin.Context) {

}