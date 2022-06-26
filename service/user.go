package service

import (
	"gin-gorm-oj/models"
	"gin-gorm-oj/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
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
			"msg":  "用户名或密码为空",
		})
		return
	}

	password = utils.GetMd5(password)

	data := new(models.UserBasic)

	err := models.DB.Where("name = ? and password = ?", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserBasic Error:" + err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(data.Identity, data.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Generate Token Error:" + err.Error(),
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

// SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @Param email formData string false "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send-code [post]
func SendCode(ctx *gin.Context) {
	email := ctx.PostForm("email")

	if email == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不能为空",
		})
		return
	}
	emails := make([]string, 1)
	emails[0] = email

	code, err := utils.SendEmailValidate(emails)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Send Email Validate Error:" + err.Error(),
		})
		return
	}
	models.RDB.Set(email, code, time.Second*300)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

// Register
// @Tags 公共方法
// @Summary 用户注册
// @Param mail formData string true "mail"
// @Param code formData string true "code"
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /register [post]
func Register(ctx *gin.Context) {
	mail := ctx.PostForm("mail")
	code := ctx.PostForm("code")
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")

	//参数校验
	if mail == "" || code == "" || name == "" || password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填信息不能为空",
		})
		return
	}

	// 校验用户是否已经注册
	var count int64
	models.DB.Model(new(models.UserBasic)).Where("mail = ?",mail).Count(&count)
	if count != 0 {
		ctx.JSON(http.StatusOK,gin.H{
			"code": -1,
			"msg": "该邮箱已注册",
		})
		return
	}

	// 校验验证码
	sysCode := models.RDB.Get(mail).Val()
	if code != sysCode {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码错误",
		})
		return
	}

	// 插入数据
	identity := utils.GenerateUUID()
	data := models.UserBasic{
		Identity:identity,
		Name:     name,
		Password: password,
		Mail:     mail,
		Phone:    phone,
	}

	err := models.DB.Create(&data).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create User Error" + err.Error(),
		})
		return
	}

	// 生成token
	token, err := utils.GenerateToken(identity, name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Generate Token Error" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
