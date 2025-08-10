package db

import (
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  string
}

func (config DBConfig) ConnStr() string {
	user := url.QueryEscape(config.User)
	pass := url.QueryEscape(config.Password)

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s", 
		user, 
		pass, 
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)
}

var Connection *gorm.DB

func NewGormDBConnection(connStr string) error {
	var err error
	Connection, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
		return err
	}
	log.Println("DB: Successfully connected!")
	return nil
}
