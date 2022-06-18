package service

import (
	"gin-gorm-oj/define"
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func GetProblemList(ctx *gin.Context) {
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
	keyword := ctx.Query("keyword")

	categoryIdentity := ctx.Query("category_identity")

	var totalCount int64

	list := make([]*models.ProblemBasic, 0)

	tx := models.GetProblemList(keyword, categoryIdentity)

	err = tx.Debug().Count(&totalCount).Omit("content").Offset((page - 1) * size).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem Error:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"data":       list,
			"totalCount": totalCount,
		},
	})
}
