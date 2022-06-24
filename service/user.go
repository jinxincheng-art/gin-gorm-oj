package service

import (
	"gin-gorm-oj/models"
	"gin-gorm-oj/utils"
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
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户唯一标识不能为空",
		})
		return
	}

	data := new(models.UserBasic)

	err := models.DB.Omit("password").Where("identity = ?", identity).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户不存在",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User Detail Error:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /login [post]
func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == "" || password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg": "用户名或密码为空",
		})
		return
	}

	password = utils.GetMd5(password)

	data := new(models.UserBasic)

	err := models.DB.Where("name = ? and password = ?",username,password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg": "用户名或密码错误",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg": "Get UserBasic Error:" + err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(data.Identity, data.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg": "Generate Token Error:" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
