package network

import (
	"database/sql"
	"github.com/3boku/mysql-learn/MySQL-with-go/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var sectionIcons = []string{
	"🍚", "🍿", "🍜", "🍣", "🥩", "☕", "🍰",
}

type Network struct {
	engine *gin.Engine
}

func NewNetwork(dataSourceName string) {
	r := gin.New()

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	database.UserCheck(db, r)
	database.UserInsert(db, r)

	r.Run()
}
