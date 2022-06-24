package service

import (
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)


// GetUserDetail
// @Tags 公共方法
// @Summary 用户详情
// @Param user_identity query string false "user_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user-detail [get]
func GetUserDetail(ctx *gin.Context) {
	identity := ctx.Query("user_identity")

	if identity == "" {
		ctx.JSON(http.StatusOK,gin.H{
			"code": -1,
			"msg": "用户唯一标识不能为空",
		})
		return
	}

	data := new(models.UserBasic)

	err := models.DB.Where("identity = ?",identity).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			ctx.JSON(http.StatusOK,gin.H{
				"code": -1,
				"msg": "用户不存在",
			})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{
			"code": -1,
			"msg": "Get User Detail Error:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"data": data,
	})
}
