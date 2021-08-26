package main

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"

	"strconv"
)

// pointer to gin.Engine
var router *gin.Engine

// custom load templates function
func loadTemplates(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("allmaps", templateDir+"/layouts/base.html", templateDir+"/includes/maps.html")
	r.AddFromFiles("singlemap", templateDir+"/layouts/base.html", templateDir+"/includes/singlemap.html")

	return r
}

// setup routing scheme
func initRoutingScheme() {
	r := gin.Default()
	router = r // set global variable

	// load templates
	r.HTMLRender = loadTemplates("./templates")

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

			c.HTML(http.StatusOK, "singlemap", gin.H{
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
	c.HTML(http.StatusOK, "allmaps", gin.H{
		"title": "Maps",
		"maps":  getAllMapsFromDB(),
	})
}
