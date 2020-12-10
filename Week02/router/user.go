package router

import (
	"Week02/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Detail(ctx *gin.Context) {
	uid := ctx.Param(`id`)

	id, err := strconv.Atoi(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			`code`: 400001,
			`msg`:  `参数错误`,
		})
		return
	}

	user, err := service.GetDetail(uint(id))
	if err != nil {
		log.Printf("ERROR: query user failed, uid=%d, err=%v", id, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			`code`: 500001,
			`msg`:  `系统错误`,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		`code`: 0,
		`msg`:  `成功`,
		`data`: user,
	})
	return
}
