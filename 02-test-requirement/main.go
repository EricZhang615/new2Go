package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"test-sample/controller"
	. "test-sample/repository"
)

func main() {
	if err := Init("./"); err != nil {
		os.Exit(-1)
	}
	//fmt.Println(NewTopicDaoInstance().QueryTopicById(1))
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	err := r.Run("172.30.192.1:8888")
	if err != nil {
		return
	}
}
