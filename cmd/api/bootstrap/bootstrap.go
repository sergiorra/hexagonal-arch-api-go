package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	mooc "github.com/sergiorra/hexagonal-arch-api-go/internal"
	"github.com/sergiorra/hexagonal-arch-api-go/internal/increasing"
	"time"

	"github.com/sergiorra/hexagonal-arch-api-go/internal/creating"
	"github.com/sergiorra/hexagonal-arch-api-go/internal/platform/bus/inmemory"
	"github.com/sergiorra/hexagonal-arch-api-go/internal/platform/server"
	"github.com/sergiorra/hexagonal-arch-api-go/internal/platform/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser    = "codely"
	dbPass    = "codely"
	dbHost    = "localhost"
	dbPort    = "3306"
	dbName    = "codely"
	dbTimeout = 5 * time.Second
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}