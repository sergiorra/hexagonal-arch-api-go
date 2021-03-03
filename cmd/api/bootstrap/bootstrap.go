package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/sergiorra/hexagonal-arch-api-go/internal/platform/server"
	"github.com/sergiorra/hexagonal-arch-api-go/internal/platform/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "test"
	dbPass = "test"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "test"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
