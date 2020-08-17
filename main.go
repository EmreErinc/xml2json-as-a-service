package main

import (
	_ "context"
	"encoding/json"
	xj "github.com/basgys/goxml2json"
	guuid "github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"runtime"
	"strings"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Routes
	e.POST("/xml", rssHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func rssHandler(c echo.Context) (err error) {
	req := new(RssRequest)
	if err = c.Bind(req); err != nil {
		panic(err)
	}

	res := fetchAndDecodeXML(req)
	return c.JSONPretty(http.StatusOK, res, "	")
}

func fetchAndDecodeXML(src *RssRequest) (raw map[string]interface{}) {
	resp, err := http.Get(src.Link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		xml := strings.NewReader(bodyString)

		// xml to json
		result, err := xj.Convert(xml)
		if err != nil {
			panic("An error occurred while xml parsing")
		}

		if err := json.Unmarshal(result.Bytes(), &raw); err != nil {
			panic(err)
		}

		// if this is true, write response to file
		if src.Write {
			path := RootDirectory() + "/" + guuid.New().String() + ".txt"
			err = ioutil.WriteFile(path, result.Bytes(), 0644)
			if err != nil {
				panic(err)
			}
		}
	}
	return raw
}

type RssRequest struct {
	Link  string `json:"link" form:"link" query:"link"`
	Write bool   `json:"write" form:"write" query:"write"`
}

func RootDirectory() string {
	_, file, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(file))
}
