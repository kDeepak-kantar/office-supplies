package main

import (
	"fmt"
	"os"

	authDom "github.com/Deepak/pkg/domain/auth"
	ulistDom "github.com/Deepak/pkg/domain/ulist"
	"github.com/Deepak/pkg/http/rest"
	"github.com/Deepak/pkg/http/rest/auth"
	"github.com/Deepak/pkg/http/rest/health"
	"github.com/Deepak/pkg/http/rest/ulists"
	"github.com/Deepak/pkg/http/routes"
	"github.com/Deepak/pkg/logger"
	db "github.com/Deepak/pkg/storage"
	userDb "github.com/Deepak/pkg/storage/db/user"

	userlistDb "github.com/Deepak/pkg/storage/db/userlist"
	"gorm.io/gorm"
)

type Storage struct {
	User     userDb.Repository
	UserList userlistDb.Repository
}

func initStorage(conn *gorm.DB) Storage {
	logGroup := "Init Storage"

	userlisDb, err := userlistDb.Init(conn)
	if err != nil {
		logger.LogCriticalError(logGroup, fmt.Errorf("error: failed to userlist DB storage"))
		panic(err)
	}

	userDb, err := userDb.Init(conn)
	if err != nil {
		logger.LogCriticalError(logGroup, fmt.Errorf("error: failed to init user storage"))
		panic(err)
	}

	return Storage{
		User:     userDb,
		UserList: userlisDb,
	}
}

type Domains struct {
	UList ulistDom.Domain
	Auth  authDom.Domain
}

func initDomains(s Storage) Domains {
	logGroup := "Init Domain"

	auth, err := authDom.Init(authDom.Input{
		User: s.User,
	})
	if err != nil {
		logger.LogCriticalError(logGroup, fmt.Errorf("error: failed to init auth domain"))
		panic(err)
	}
	ulis, err := ulistDom.Init(ulistDom.Input{
		User:     s.User,
		UserList: s.UserList,
	})

	if err != nil {
		logger.LogCriticalError(logGroup, fmt.Errorf("error: failed to init userlist domain"))
		panic(err)
	}

	return Domains{
		Auth:  auth,
		UList: ulis,
	}
}

type API struct {
	Auth     auth.Repository
	Rest     rest.Repository
	UserList ulists.Repository
	Health   health.Repository
}

func initAPIServices(d Domains) API {
	logGroup := "rest"
	// init rest service
	config := rest.Init(&rest.Configuration{
		Env:  getenv("ENVIRONMENT", "dev"),
		Host: getenv("HTTP_HOST", "localhost"),
		Port: 7068,
	})

	auth := auth.Init(auth.Input{
		Auth: d.Auth,
	})

	userlist, err := ulists.Init(ulists.Input{
		Ulist: d.UList,
	})
	if err != nil {
		logger.LogCriticalError(logGroup, fmt.Errorf("error: failed to init coffee date rest"))
		panic(err)
	}
	return API{
		Auth:     auth,
		UserList: userlist,
		Rest:     config,
		Health:   health.Init(),
	}
}

func initWebServices(d Domains) {

}

func main() {
	logGroup := "main"
	logger.LogInfo(logGroup, "Starting up")

	dbConnection := db.Init(db.Input{
		Host:     getenv("DB_HOST", "0.0.0.0"),
		Port:     getenv("DB_PORT", "3312"),
		User:     getenv("DB_USER", "coffeetalk"),
		Password: getenv("DB_PASSWORD", "coffeetalk"),
		Database: getenv("DB_NAME", "coffeetalk"),
		Env:      getenv("ENVIRONMENT", "dev"),
	})

	storage := initStorage(dbConnection)
	doms := initDomains(storage)
	go doms.Auth.Scheduler()
	api := initAPIServices(doms)

	// init api routes
	routes := routes.Init(routes.Input{
		API:      api.Rest,
		Auth:     api.Auth,
		Userlist: api.UserList,
		Health:   api.Health,
	})
	routes.Configure()

	err := api.Rest.Run()
	if err != nil {
		logger.Log("main", logger.SeverityError, "server terminated with error", err, nil)
	}
}

func getenv(key, fb string) string {
	v := os.Getenv(key)

	if v == "" {
		return fb
	}

	return v
}
