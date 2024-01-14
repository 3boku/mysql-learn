package database

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
	Id   string
	Pw   string
}

func UserCheck(db *sql.DB, engine *gin.Engine) {
	results, err := db.Query("SELECT * FROM Users")

	if err != nil {
		panic(err)
	}

	engine.GET("/user", func(c *gin.Context) {
		for results.Next() {
			var user User

			err = results.Scan(&user.Name, &user.Id, &user.Pw)
			if err != nil {
				panic(err)
			}

			c.JSON(200, gin.H{
				"name": user.Name,
				"id":   user.Id,
				"pw":   user.Pw, // "message" 필드가 "pong"인 JSON 응답을 반환한다.
			})
		}
	})
}

func UserInsert(db *sql.DB, engine *gin.Engine) {
	engine.POST("/insertData", func(c *gin.Context) {
		var user User

		// JSON 형식의 POST 요청 파싱
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "잘못된 요청입니다.",
			})
			return
		}

		// 데이터 인서트
		_, err := db.Exec("INSERT INTO Users VALUES (?, ?, ?)", user.Id, user.Name, user.Pw)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "데이터베이스 인서트 실패.",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "데이터베이스에 인서트를 완료했습니다.",
		})
	})
}
