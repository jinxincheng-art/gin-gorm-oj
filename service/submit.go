package service

import (
	"gin-gorm-oj/define"
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetSubmitList
// @Tags 公共方法
// @Summary 提交列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param problem_identity query string false "problem_identity"
// @Param user_identity query string false "user_identity"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /submit-list [get]
func GetSubmitList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv Error:", err)
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("GetProblemList Size strconv Error:", err)
		return
	}
	var submitCount int64

	problemIdentity := ctx.Query("problem_identity")

	userIdentity := ctx.Query("user_identity")

	status, err := strconv.Atoi(ctx.Query("status"))
	if err != nil {
		log.Println("GetProblemList Status strconv Error:", err)
	}

	list := make([]*models.SubmitBasic, 0)

	tx := models.GetSubmitList(problemIdentity, userIdentity, status)
	err = tx.Count(&submitCount).Offset((page - 1) * size).Limit(size).Find(&list).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Submit List Error:" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": submitCount,
		},
	})
}
