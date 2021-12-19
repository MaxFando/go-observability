package main

import "github.com/sirupsen/logrus"

const (
	DatabaseURL = "postgres://usr:pwd@localhost:5432/example?sslmode=disable"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	//logger.SetFormatter(&logrus.JSONFormatter{})

}
