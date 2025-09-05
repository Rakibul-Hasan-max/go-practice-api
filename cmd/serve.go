package cmd

import (
	"go-practice-api/config"
	"go-practice-api/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
