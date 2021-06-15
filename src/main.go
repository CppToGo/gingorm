package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Result struct {
	Dataid uint32
	Index1 uint64
	Index2 uint32
	Type   uint16
	Version uint32
	Buffer  []byte
	Createtime uint32
	Updatetime uint32
}


func main() {
	// 获取gin路由
	// 默认路由已经加载 logger, recovery 中间件
	router := gin.Default()
	// 加载模板文件
	router.LoadHTMLGlob("template/*")

	db := router.Group("/db")
	{
		db.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "db_index.html", gin.H{})
		})

		db.POST("/datainfo", func(c *gin.Context) {
			UID := c.DefaultPostForm("UID", "*")
			DataID := c.DefaultPostForm("DataID", "*")
			Index2 := c.DefaultPostForm("Index2", "*")
			uid, _ := strconv.Atoi(UID)
			tableName := fmt.Sprintf("tbl_num_index_%03d",uid%512 + 1)
			result := make([]Result, 10)
			dbrl := DataBase.Table(tableName).Where("dataid = ? and index1= ? and index2= ?", DataID, UID, Index2).Find(&result)
			c.HTML(http.StatusOK, "db_page.html", gin.H{"title": uid,"p1":result, "tableName":tableName, "Error":dbrl.Error})
		})

		db.GET("/datainfo/:uid", func(c *gin.Context) {
			UID := c.Param("uid")
			uid, _ := strconv.Atoi(UID)
			tableName := fmt.Sprintf("tbl_num_index_%03d",uid%512 + 1)
			result := make([]Result, 10)
			dbrl := DataBase.Table(tableName).Where("index1= ?", UID).Find(&result)
			c.HTML(http.StatusOK, "db_page.html", gin.H{"title": uid,"p1":result, "tableName":tableName, "Error":dbrl.Error})
		})

		db.GET("/datainfo/:uid/:dataid", func(c *gin.Context) {
			UID := c.Param("uid")
			DataID := c.Param("dataid")
			uid, _ := strconv.Atoi(UID)
			tableName := fmt.Sprintf("tbl_num_index_%03d",uid%512 + 1)
			result := make([]Result, 10)
			dbrl := DataBase.Table(tableName).Where("dataid = ? and index1= ?", DataID, UID).Find(&result)
			c.HTML(http.StatusOK, "db_page.html", gin.H{"title": uid,"p1":result, "tableName":tableName, "Error":dbrl.Error})
		})

		db.GET("/datainfo/:uid/:dataid/:index2", func(c *gin.Context) {
			UID := c.Param("uid")
			DataID := c.Param("dataid")
			Index2 := c.Param("index2")
			uid, _ := strconv.Atoi(UID)
			tableName := fmt.Sprintf("tbl_num_index_%03d",uid%512 + 1)
			result := make([]Result, 10)
			dbrl := DataBase.Table(tableName).Where("dataid = ? and index1= ? and index2= ?", DataID, UID, Index2).Find(&result)
			c.HTML(http.StatusOK, "db_page.html", gin.H{"title": uid,"p1":result, "tableName":tableName, "Error":dbrl.Error})
		})
	}

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
}
