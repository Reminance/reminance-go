package v1

import (
	"github.com/reminance/reminance-go/xinliangnote-Go/01-Gin框架/codes/04-数据绑定和验证/ginDemo/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMember(c *gin.Context) {

	res := entity.Result{}
	mem := entity.Member{}

	if err := c.ShouldBind(&mem); err != nil {
		res.SetCode(entity.CODE_ERROR)
		res.SetMessage(err.Error())
		c.JSON(http.StatusForbidden, res)
		c.Abort()
		return
	}

	// 处理业务(下次再分享)

	data := map[string]interface{}{
		"name" : mem.Name,
		"age"  : mem.Age,
	}
	res.SetCode(entity.CODE_SUCCESS)
	res.SetData(data)
	c.JSON(http.StatusOK, res)
}
