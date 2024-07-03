package db

import (
	"fmt"
	"os"
)

const (
	hostName string = "HOST_NAME"
	dbName   string = "DATABASE_NAME"
	password string = "PASSWORD"
	user     string = "USER"
	port     string = "PORT"
	sslMode  string = "SSL_MODE"
	timeZone string = "TIMEZONE"
)

type ConnectionConfig struct {
	Host         string
	DatabaseName string
	Password     string
	User         string
	Port         string
	SSLMode      string
	TimeZone     string
}

func (c *ConnectionConfig) FromEnv() {
	c.Host = os.Getenv(hostName)
	c.DatabaseName = os.Getenv(dbName)
	c.Password = os.Getenv(password)
	c.User = os.Getenv(user)
	c.Port = os.Getenv(port)
	c.SSLMode = os.Getenv(sslMode)
	c.TimeZone = os.Getenv(timeZone)
}

func (c *ConnectionConfig) getDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.Host, c.User, c.Password, c.DatabaseName, c.Port, c.SSLMode, c.TimeZone,
	)
}
