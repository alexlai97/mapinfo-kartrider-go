package main

import (
	"net/http"

	// "github.com/stnc/pongo4gin"
	// "github.com/flosch/pongo2"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"

	"strconv"
)

// pointer to gin.Engine
var router *gin.Engine

// FIXME: doesn't work
func loadTemplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("allmaps", "templates/include/layout.html", "templates/maps/maps.html")
	r.AddFromFiles("singlemap", "templates/include/layout.html", "templates/maps/singlemap.html")
	return r
}

// setup routing scheme
func initRoutingScheme() {
	r := gin.Default()
	router = r // set global variable

	// load templates
	r.LoadHTMLGlob("templates/**/*.html")
	// r.HTMLRender = loadTemplates() // FIXME: doesn't work

	r.Static("/assets", "./assets")

	// json
	group_v1 := r.Group("api/v1/maps")
	{
		group_v1.GET("/", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, getAllMapsFromDB())
		})
		group_v1.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.String(http.StatusNotAcceptable, err.Error())
			}
			c.IndentedJSON(http.StatusOK, getSingleMapFromDB(id))
		})
	}

	// html
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

			c.HTML(http.StatusOK, "singlemap.html", gin.H{
				"title":   "Single map",
				"details": getSingleMapFromDB(id),
			})
		})
	}

}

// serve the app at ipaddr
func serveRouter(ipaddr string) {
	router.Run(ipaddr)
}

// show all maps
func routingToAllMaps(c *gin.Context) {
	c.HTML(http.StatusOK, "maps.html", gin.H{
		"title": "Maps",
		"maps":  getAllMapsFromDB(),
	})
}
