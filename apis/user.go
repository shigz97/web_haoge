package apis

import (
	"fmt"
	"net/http"
	"strconv"
	model "web_haoge/models"

	"github.com/gin-gonic/gin"
)

//列数据
func Users(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")

	user.Password = c.Request.FormValue("password")

	res, err := user.Users()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "注册失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": res,
	})
}

//添加数据
func Store(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")

	r, err := user.Finduser()
	if err == nil && r.ID != 0 {
		fmt.Println("exist")
		c.JSON(http.StatusOK, gin.H{
			"code":    -2,
			"message": "user already exits",
		})
		return
	}

	user.Password = c.Request.FormValue("password")

	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    id,
	})
}

//修改数据
func Update(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改成功",
	})
}

//删除数据
func Destroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除成功",
	})
}

//按照用户名查询
func FindUser(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("id")
	res, err := user.Finduser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": res.Password,
	})
}
