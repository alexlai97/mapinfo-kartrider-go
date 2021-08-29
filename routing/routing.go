package routing

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"

	"strconv"

	"github.com/alexlai97/mapinfo-kartrider/maps"
	"github.com/alexlai97/mapinfo-kartrider/model"
	"github.com/alexlai97/mapinfo-kartrider/users"
)

// pointer to gin.Engine
var router *gin.Engine

// custom load templates function
func loadTemplates(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("allmaps", templateDir+"/layouts/base.html", templateDir+"/includes/maps/maps.html")
	r.AddFromFiles("singlemap", templateDir+"/layouts/base.html", templateDir+"/includes/maps/singlemap.html")
	r.AddFromFiles("login", templateDir+"/layouts/base.html", templateDir+"/includes/auth/login.html")
	r.AddFromFiles("register", templateDir+"/layouts/base.html", templateDir+"/includes/auth/register.html")

	return r
}

// setup routing scheme
func InitRoutingScheme() {
	r := gin.Default()
	router = r // set global variable

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// load templates
	r.HTMLRender = loadTemplates("./templates")

	r.Static("/assets", "./assets")

	// "api/v1/maps" json
	group_v1 := r.Group("api/v1/maps")
	{
		group_v1.GET("/", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, maps.GetAllMapsFromDB())
		})
		group_v1.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.String(http.StatusNotAcceptable, err.Error())
			}
			c.IndentedJSON(http.StatusOK, maps.GetSingleMapFromDB(id))
		})
	}

	// "/" and "/maps" html
	r.GET("/", routingToAllMaps)
	group_map := r.Group("maps")
	{
		group_map.GET("/", routingToAllMaps)
		group_map.GET("/:id", func(c *gin.Context) {
			id_str := c.Param("id")
			id, err := strconv.Atoi(id_str)
			if err != nil {
				c.String(http.StatusNotAcceptable, err.Error())
			}

			user, loggedin := users.GetLoggedInUser(c)
			log.Println("username: ", user.Username, "loggedin?: ", loggedin)
			c.HTML(http.StatusOK, "singlemap", gin.H{
				"loggedin": loggedin,
				"username": user.Username,
				"details":  maps.GetSingleMapFromDB(id),
			})
		})
	}

	group_auth := r.Group("auth")
	{
		group_auth.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register", gin.H{})
		})
		group_auth.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login", gin.H{})
		})
		group_auth.POST("/register", func(c *gin.Context) {
			var err error
			request := model.User{}
			err = c.ShouldBind(&request)
			if err != nil {
				// TODO: better logging
				log.Println("register c.ShouldBind failed", err.Error())
			}
			// TODO: store hash of password

			// TODO:
			// if user exists, "username exists"
			// if empty form return "username/password needed"

			err = users.RegisterUser(request)
			if err != nil {
				log.Println("register User failed", err.Error())
			}

			c.Redirect(http.StatusFound, "/auth/login")
		})
		group_auth.POST("/login", func(c *gin.Context) {
			var err error
			request := model.User{}
			err = c.ShouldBind(&request)
			if err != nil {
				log.Println("login c.ShouldBind failed", err.Error())
			}

			user := users.GetUserByUsername(request.Username)
			if request.Password != user.Password {
				err = errors.New("incorrect password")
			}

			if err != nil {
				log.Println(err.Error())
				// TODO: flash error
			} else {
				session := sessions.Default(c)
				session.Clear()
				session.Set("account_id", user.ID)
				session.Save()
				log.Println("session saved accound_id: ", user.ID)
				user, loggedin := users.GetLoggedInUser(c)
				log.Println("username: ", user.Username, "loggedin?: ", loggedin)
				c.Redirect(http.StatusFound, "/maps")
			}
		})
		group_auth.Any("/logout", func(c *gin.Context) {
			session := sessions.Default(c)
			session.Delete("account_id")
			session.Save()
			c.Redirect(http.StatusFound, "/maps")
		})
	}
}

// serve the app at ipaddr
func ServeRouter(ipaddr string) {
	router.Run(ipaddr)
}

// show all maps
// reuse code
func routingToAllMaps(c *gin.Context) {
	user, loggedin := users.GetLoggedInUser(c)
	log.Println("username: ", user.Username, "loggedin?: ", loggedin)
	c.HTML(http.StatusOK, "allmaps", gin.H{
		"loggedin": loggedin,
		"username": user.Username,
		"maps":     maps.GetAllMapsFromDB(),
	})
}
