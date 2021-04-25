package main

import (
	imageCodeContext "context"
	proto "gitee.com/hanwei66/GetImageCode/proto"
	"github.com/gin-gonic/gin"
	"github.com/micro/micro/v3/service"
	"github.com/prometheus/common/log"
	"net/http"
)

func GetImageCode(context *gin.Context) {
	srv := service.New()
	client := proto.NewGetImageCodeService("", srv.Client())
	rsp, err := client.GetImageCode(imageCodeContext.TODO(), &proto.Request{Uuid: "xxxxxx"})
	if err != nil {
		log.Info("client.GetImageCode err:", err)
	}

	log.Info(rsp.String())

	posts := rsp.String()
	/*	for _, post := range posts {
		log.Println("rsp:",post )
	}*/
	context.JSON(http.StatusOK, posts)

}

func main() {
	router := gin.Default()
	// router.GET("/getAll", CallRemote)
	router.GET("/getImageCode", GetImageCode)

	router.Run(":9000")
}
