package cmd

import (
	"github.com/3boku/mysql-learn/MySQL-with-go/database"
	"github.com/3boku/mysql-learn/MySQL-with-go/network"
)

type Cmd struct {
	network  *network.Network
	database *database.Database
}

var databaSource = "root:1234@tcp(127.0.0.1:3306)/mydatabase"

var port = ":8080"

func NewCmd() *Cmd {
	c := new(Cmd)

	network.NewNetwork(databaSource)
	return c
}
