package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type articles struct {
	Url        string `gorm:"type:varchar(100);UNIQUE;PRIMARY_KEY;NOT NULL"`
	Board      string `gorm:"type:varchar(50);INDEX;NOT NULL"`
	Title      string `gorm:"type:varchar(50)"`
	Author     string `gorm:"type:varchar(50);INDEX"`
	Contains   string `gorm:"type:longtext"`
	Reply      string `gorm:"type:longtext"`
	Date       time.Time
	Updated_at time.Time
}

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"result": "success"})
}

func SearchArticles(ctx *gin.Context) {
	articles := []articles{}
	switch {
	case len(ctx.PostForm("board")) > 0:
		if db.Limit(10).Table("articles").Where("board =?", ctx.PostForm("board")).Find(&articles).RowsAffected == 0 {
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"result": "error", "error": "查無查詢條件內文章"})
		} else {
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"result": "success", "articles": &articles})
		}
	case len(ctx.PostForm("author")) > 0:
		if db.Limit(10).Table("articles").Where("author like ?", ctx.PostForm("author")+"%").Find(&articles).RowsAffected == 0 {
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"result": "error", "error": "查無查詢條件內文章"})
		} else {
			ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"result": "success", "articles": &articles})
		}
	default:
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"result": "error", "error": "請輸入查詢條件"})
	}
}
