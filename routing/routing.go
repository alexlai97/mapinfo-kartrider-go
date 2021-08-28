package routing

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"

	"strconv"

	"github.com/alexlai97/mapinfo-kartrider/maps"
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

			c.HTML(http.StatusOK, "singlemap", gin.H{
				"title":   "Single map",
				"details": maps.GetSingleMapFromDB(id),
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
	}
}

// serve the app at ipaddr
func ServeRouter(ipaddr string) {
	router.Run(ipaddr)
}

// show all maps
// reuse code
func routingToAllMaps(c *gin.Context) {
	c.HTML(http.StatusOK, "allmaps", gin.H{
		"title": "Maps",
		"maps":  maps.GetAllMapsFromDB(),
	})
}
