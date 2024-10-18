package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// fmt.Println("test")
	// fmt.Println(time.Now())

	//　変数定義
	// var name string = "Tom"
	// fmt.Println(name)
	// var num int = 1
	// fmt.Println(num + 1)

	// 別関数を呼ぶ
	getJson()


}

func getJson () {
		r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
    
    r.Run()
}