package routes

import (
	"fmt"
	"net/http"

	"github.com/Deepak/pkg/http/rest"
	"github.com/Deepak/pkg/http/rest/auth"
	"github.com/Deepak/pkg/http/rest/ulists"
	"github.com/Deepak/pkg/http/web"
	"github.com/Deepak/pkg/storage/db/user"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Configure()
}

// Adapter is a structure that contains the necessary dependencies for performing
// the routing.
type repository struct {
	Input
}

type Input struct {
	API      rest.Repository
	Web      web.Repository
	Auth     auth.Repository
	User     user.Repository
	Userlist ulists.Repository
}

// New creates a routing adapter given the necessary repositories.
func Init(input Input) Repository {
	return &repository{
		input,
	}
}

func (r *repository) serverStaticFiles(engine *gin.Engine) {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	engine.Static("/css", "public/css/")
	engine.Static("/js", "public/js/")

	engine.HTMLRender = r.Web.NewMultiTemplate(map[string]map[string][]string{
		"public/views": r.Web.GetViews(),
	})
}

func (r *repository) setupRedirects(engine *gin.Engine) {
	// redirects
	engine.NoRoute(func(c *gin.Context) {
		fmt.Printf("No route found for: %v with method %v \n", c.Request.URL, c.Request.Method)
		c.Redirect(http.StatusFound, "/login")
	})

	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/login")
	})

}

// Configure setups the actual routing for the HTTP API.
func (r *repository) Configure() {
	engine := r.API.Engine()
	engine.Use(cors.New(CORSConfig()))
	r.serverStaticFiles(engine)
	r.setupRedirects(engine)

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// metrics

	// web
	engine.GET("/login", r.Web.GetLogin)
	engine.GET("/overview", r.Web.GetOverview)

	// api
	api := engine.Group("/api")

	// login
	api.POST("/login", r.Auth.Login)
	api.GET("/:id/getallusers", r.Auth.GetAllUsers)
	api.POST("/:id/admin", r.Auth.Admin)
	api.POST("/:id/removeuser", r.Auth.RemoveUser)
	// User LIst
	api.POST("/createlist", r.Userlist.CreateUserList)
	api.GET("/:id/getalluserlist", r.Userlist.GetAllUserLists)
	api.POST("/updateuser", r.Userlist.UpdateUserList)
	api.GET("/:id/getallapproveduser", r.Userlist.GetAllApprovedUserLists)
	api.GET("/:id/getallnotapproved", r.Userlist.GetAllNotApprovedUserLists)
	api.GET("/sendremainder", r.Userlist.SendRemainderrest)

}
func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}
