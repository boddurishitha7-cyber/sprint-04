package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	connString := "host=127.0.0.1 port=5432 user=medapati123DB password=medapati_cloud dbname=medapati_cloud sslmode=disable"

	fmt.Println("Connecting with:", connString)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Println("Connection error:", err)
		return nil, err
	}

	fmt.Println("Database connected successfully!")
	return conn, nil
}
