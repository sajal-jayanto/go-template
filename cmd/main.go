package main

import (
	"log"

	"github.com/sajal-jayanto/go-template/cmd/api"
	"github.com/sajal-jayanto/go-template/configs"
	"github.com/sajal-jayanto/go-template/db"
)

func main(){
	DBconfig := db.DBConfig{
		User:      configs.Envs.DBUser,
		Password:  configs.Envs.DBPassword,
		Host:      configs.Envs.DBHost,
		Port:      configs.Envs.DBPort,             
		DBName:    configs.Envs.DBName,
		SSLMode:   configs.Envs.DB_SSLMode,
	}
	
	connStr := DBconfig.ConnStr()

	err := db.NewGormDBConnection(connStr)
	if err != nil {
		log.Fatal(err)
	}

	addr := ":" + configs.Envs.Port
	server := api.NewAPIServer(addr)
	if err := server.Run(); err != nil{
		log.Fatal(err)
	}
}