package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
	index "github.com/hailocab/go-geoindex"
)

func init() {
	// http.HandleFunc("/points", points)
	// http.HandleFunc("/knearest", knearest)
}
func main() {
	if false { //test
		nearest := getIndex().KNearest(index.NewGeoPoint("query", 18.230292850934863, -63.16093794856357), int(1000), index.Km(5), func(_ index.Point) bool { return true })
		fmt.Println(nearest)
		return
	}
	e := gin.Default()
	e.Static("/static", "../static/")
	groupName := ""
	group := e.Group(groupName)
	{
		group.GET("/knearest", knearest4Gin)
		group.GET("/points", points4Gin)
		e.Run(":8090")
	}
}

var geoindex *index.ClusteringIndex

// 获取随机的正负
func sign() float64 {
	if rand.Float64() > 0.5 {
		return 1
	}
	return -1
}

func getIndex() *index.ClusteringIndex {
	if geoindex == nil {
		geoindex = index.NewClusteringIndex()

		capitals := worldCapitals()
		id := 1

		for _, capital := range capitals {
			for i := 0; i < 300; i++ {
				id++

				geoindex.Add(index.NewGeoPoint(
					fmt.Sprintf("%d", id),
					capital.Lat()+rand.Float64()/6.0*sign(),
					capital.Lon()+rand.Float64()/6.0*sign(),
				))
			}
		}
	}

	return geoindex
}

func worldCapitals() []index.Point {
	file, err := os.OpenFile("../static/capitals.csv", os.O_RDONLY, 0)

	if err != nil {
		log.Printf("%v", err)
		return make([]index.Point, 0)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	records, _ := reader.ReadAll()
	capitals := make([]index.Point, 0)

	for _, record := range records {
		id := record[0]
		lat, _ := strconv.ParseFloat(record[3], 64)
		lon, _ := strconv.ParseFloat(record[4], 64)

		capital := index.NewGeoPoint(id, lat, lon)
		capitals = append(capitals, capital)
	}

	return capitals
}

type Point struct {
	Lat float64 `form:"lat"`
	Lon float64 `form:"lon"`
	K   int64   `form:"k"`
}

func knearest4Gin(c *gin.Context) {
	p := Point{}
	c.ShouldBind(&p)

	nearest := getIndex().KNearest(index.NewGeoPoint("query", p.Lat, p.Lon), int(p.K), index.Km(5), func(_ index.Point) bool { return true })
	// data, _ := json.Marshal(nearest)
	c.JSON(http.StatusOK, nearest)
}

type Rectangle struct {
	TopLeftLat     float64 `form:"topLeftLat"`
	TopLeftLon     float64 `form:"topLeftLon"`
	BottomRightLat float64 `form:"bottomRightLat"`
	BottomRightLon float64 `form:"bottomRightLon"`
}

func points4Gin(c *gin.Context) {
	rec := Rectangle{}
	c.BindQuery(&rec)

	visiblePoints := getIndex().Range(index.NewGeoPoint("topLeft", rec.TopLeftLat, rec.TopLeftLon),
		index.NewGeoPoint("bottomRight", rec.BottomRightLat, rec.BottomRightLon))

	c.JSON(http.StatusOK, visiblePoints)
}

func points(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	topLeftLat, _ := strconv.ParseFloat(r.Form["topLeftLat"][0], 64)
	topLeftLon, _ := strconv.ParseFloat(r.Form["topLeftLon"][0], 64)
	bottomRightLat, _ := strconv.ParseFloat(r.Form["bottomRightLat"][0], 64)
	bottomRightLon, _ := strconv.ParseFloat(r.Form["bottomRightLon"][0], 64)

	visiblePoints := getIndex().Range(index.NewGeoPoint("topLeft", topLeftLat, topLeftLon), index.NewGeoPoint("bottomRight", bottomRightLat, bottomRightLon))

	data, _ := json.Marshal(visiblePoints)
	fmt.Fprintln(w, string(data))
}

func knearest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lat, _ := strconv.ParseFloat(r.Form["lat"][0], 64)
	lon, _ := strconv.ParseFloat(r.Form["lon"][0], 64)
	k, _ := strconv.ParseInt(r.Form["k"][0], 10, 32)

	nearest := getIndex().KNearest(index.NewGeoPoint("query", lat, lon), int(k), index.Km(5), func(_ index.Point) bool { return true })
	data, _ := json.Marshal(nearest)
	fmt.Fprintln(w, string(data))
}
