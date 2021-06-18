package main

import (
	"fmt"
	"gingorm/src/Framwork"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


type NumResult struct {
	Dataid uint32
	Index1 uint64
	Index2 uint32
	Type   uint16
	Version uint32
	Buffer  []byte
	Createtime uint32
	Updatetime uint32
}

type StrResult struct {
	Dataid uint32
	Index1 string
	Type   uint16
	Version uint32
	Buffer  []byte
	Createtime uint32
	Updatetime uint32
}

func OneStringIndexHashFunc(num int, input interface{}) int {
	if num <= 0 {
		return 0
	}

	var result uint64
	if str , ok := input.([]byte); ok {
		for _, v := range str {
			result += uint64(v)
		}
	}

	return int(result % uint64(num)) +1
}

type CoredataQueryHandle struct {
	Framwork.Handle
}

func (h *CoredataQueryHandle) SetRootRouter(){
	h.RootRouter = "/db"
}

func (h *CoredataQueryHandle) Process(group *gin.RouterGroup) error {
	statuCode := http.StatusOK

	group.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "db_index.html", gin.H{})
	})

	group.POST("/num", func(c *gin.Context) {
		UID := c.DefaultPostForm("UID", "*")
		DataID := c.DefaultPostForm("DataID", "*")
		Index2 := c.DefaultPostForm("Index2", "*")
		uid, _ := strconv.Atoi(UID)
		tableName := fmt.Sprintf("tbl_num_index_%03d",uid%512 + 1)
		result := make([]NumResult, 10)
		dbrl := DataBase.Table(tableName).Where("dataid = ? and index1= ? and index2= ?", DataID, UID, Index2).Find(&result)
		if dbrl.Error != nil {
			statuCode = http.StatusRequestTimeout
		}else{
			statuCode = http.StatusOK
		}
		c.HTML(statuCode, "db_page.html", gin.H{"title": uid,"p1":result, "tableName":tableName, "Error":dbrl.Error})
	})

	group.GET("/num/:uid", func(c *gin.Context) {
		UID := c.Param("uid")
		uid, _ := strconv.Atoi(UID)
		tableName := fmt.Sprintf("tbl_num_index_%03d",uid%512 + 1)
		result := make([]NumResult, 10)
		dbrl := DataBase.Table(tableName).Where("index1= ?", UID).Find(&result)
		if dbrl.Error != nil {
			statuCode = http.StatusRequestTimeout
		}else{
			statuCode = http.StatusOK
		}
		c.HTML(statuCode, "db_page.html", gin.H{"title": uid,"p1":result, "tableName":tableName, "Error":dbrl.Error})
	})

	group.GET("/num/:uid/:dataid", func(c *gin.Context) {
		UID := c.Param("uid")
		DataID := c.Param("dataid")
		uid, _ := strconv.Atoi(UID)
		tableName := fmt.Sprintf("tbl_num_index_%03d",uid%512 + 1)
		result := make([]NumResult, 10)
		dbrl := DataBase.Table(tableName).Where("dataid = ? and index1= ?", DataID, UID).Find(&result)
		if dbrl.Error != nil {
			statuCode = http.StatusRequestTimeout
		}else{
			statuCode = http.StatusOK
		}
		c.HTML(statuCode, "db_page.html", gin.H{"title": uid,"p1":result, "tableName":tableName, "Error":dbrl.Error})
	})

	group.GET("/num/:uid/:dataid/:index2", func(c *gin.Context) {
		UID := c.Param("uid")
		DataID := c.Param("dataid")
		Index2 := c.Param("index2")
		uid, _ := strconv.Atoi(UID)
		tableName := fmt.Sprintf("tbl_num_index_%03d",uid%512 + 1)
		result := make([]NumResult, 10)
		dbrl := DataBase.Table(tableName).Where("dataid = ? and index1= ? and index2= ?", DataID, UID, Index2).Find(&result)
		if dbrl.Error != nil {
			statuCode = http.StatusRequestTimeout
		}else{
			statuCode = http.StatusOK
		}
		c.HTML(statuCode, "db_page.html", gin.H{"title": uid,"p1":result, "tableName":tableName, "Error":dbrl.Error})
	})


	group.GET("/str/:uid", func(c *gin.Context) {
		UID := c.Param("uid")
		Index1 := OneStringIndexHashFunc(512, []byte(UID))
		tableName := fmt.Sprintf("tbl_str_index_%03d",Index1)
		result := make([]StrResult, 10)
		dbrl := DataBase.Table(tableName).Where("index1 = ?", UID).Find(&result)
		if dbrl.Error != nil {
			statuCode = http.StatusRequestTimeout
		}else{
			statuCode = http.StatusOK
		}
		c.HTML(statuCode, "db_page.html", gin.H{"title": UID,"p1":result, "tableName":tableName, "Error":dbrl.Error})
	})

	group.GET("/str/:uid/:dataid", func(c *gin.Context) {
		UID := c.Param("uid")
		DataID := c.Param("dataid")
		Index1 := OneStringIndexHashFunc(512, []byte(UID))
		tableName := fmt.Sprintf("tbl_str_index_%03d",Index1)
		result := make([]StrResult, 10)
		dbrl := DataBase.Table(tableName).Where("dataid = ? and index1 = ?", DataID, UID).Find(&result)
		if dbrl.Error != nil {
			statuCode = http.StatusRequestTimeout
		}else{
			statuCode = http.StatusOK
		}
		c.HTML(statuCode, "db_page.html", gin.H{"title": UID,"p1":result, "tableName":tableName, "Error":dbrl.Error})
	})

	return nil
}



