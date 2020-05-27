package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rubencougil/geekshubs/elastic/app/user"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"go.elastic.co/apm/module/apmgin"
)

func main() {

	logger := Logger()
	db := Database(logger)
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(ginlogrus.Logger(logger), gin.Recovery())
	r.Use(apmgin.Middleware(r))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"TAG": os.Getenv("TAG"), "ELASTIC_APM_SERVER_URL": os.Getenv("ELASTIC_APM_SERVER_URL"), "ELASTIC_APM_SECRET_TOKEN": os.Getenv("ELASTIC_APM_SECRET_TOKEN")})
	})
	r.POST("/create", user.CreateUserHandler(logger, user.NewUserStore(db, logger)))

	_ = r.Run(":80")
}

func Database(logger *logrus.Logger) *sqlx.DB {
	connectionString := os.Getenv("DATABASE_CONNECTION_STRING")
	db, err := sqlx.Connect("mysql", connectionString) // "user:password@(db:3306)/db"
	if err != nil {
		logger.Fatalf("Cannot connect to the database: %v, %v", err, connectionString)
	}
	return db
}

func Logger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)
	log.SetOutput(logger.Writer())
	logger.SetOutput(io.MultiWriter(os.Stdout))
	return logger
}
