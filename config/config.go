package config

import (
	"fmt"
	"os"
)

func GetDSN() string {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")
	sslCert := os.Getenv("DB_SSL_CERT")

	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s sslrootcert=%s",
		user,
		pass,
		host,
		port,
		dbName,
		sslMode,
		sslCert,
	)
}
