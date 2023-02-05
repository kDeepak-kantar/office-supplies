package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Input struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
	Env      string
}

func getGormLogger(env string) logger.Interface {
	if env == "prod" {
		return logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,         // Disable color
			},
		)
	} else {
		return logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Warn, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			},
		)
	}
}

func Init(input Input) *gorm.DB {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := retryConnection(6, 10*time.Second, func() (*gorm.DB, error) {
		return establishConnection(ctx, input)
	})
	if err != nil {
		panic(err)
	}

	return conn
}

func establishConnection(ctx context.Context, input Input) (*gorm.DB, error) {

	var dbHost string
	if input.User != "" && input.Password != "" && input.Host != "" {
		dbHost = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", input.User, input.Password, input.Host, input.Port, input.Database)
	}

	dsn := fmt.Sprintf("%s?%s", dbHost, "parseTime=true")

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: getGormLogger(input.Env),
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("err1: ", err)
		return nil, err
	}

	if err = sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}
	return gormDB, nil
}

func retryConnection(attempts int, sleep time.Duration, f func() (*gorm.DB, error)) (*gorm.DB, error) {
	var merr error
	for i := 0; i < attempts; i++ {
		db, err := f()
		if err != nil || db == nil {
			merr = err
			time.Sleep(sleep)
			continue
		}
		if db != nil {
			return db, nil
		}

		log.Println("retrying after error:", err)
	}

	return nil, fmt.Errorf("timea out after %d attempts. finished with error: %v", attempts, merr)
}
