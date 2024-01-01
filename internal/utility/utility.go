package utility

import (
	"log"
	"os"
)

func ConnectionStringAndDriver() (string, string) {
	connStr, status := os.LookupEnv("DB_CONNECTION_STRING")
	if !status {
		log.Fatalln("Missing environment variable DB_CONNECTION_STRING")
	}
	DBDriver, statusDriver := os.LookupEnv("DB_DRIVER")
	if !statusDriver {
		log.Fatalln("Missing environment variable DB_CONNECTION_STRING")
	}
	return connStr, DBDriver
}
