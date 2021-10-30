package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	index "github.com/hailocab/go-geoindex"
)

var (
	dir       = "./static"
	dirPrefix = "."
)

func init() {
	// http.HandleFunc("/points", points)
	// http.HandleFunc("/knearest", knearest)
	if exist, _ := PathExists(dir); !exist {
		dirPrefix = ".."
	}
}

// http://127.0.0.1:8090/static/
// http://127.0.0.1:8090/index
func main() {
	if false { //test
		var now time.Time
		fmt.Printf("%+v\n", now)
		if now.IsZero() {
			fmt.Println("now is zero")
		}
		return
		nearest := getIndex().KNearest(index.NewGeoPoint("query", 18.230292850934863, -63.16093794856357), int(1000), index.Km(5), func(_ index.Point) bool { return true })
		fmt.Println(nearest)
		return
	}
	e := gin.Default()
	// e.Static("/index.html", dirPrefix+"/index.html") //不行，browser一直跳到 ip:port//
	e.StaticFile("index", dirPrefix+"/index.html")
	e.Static("/static", dirPrefix+"/static/")
	e.Static("/static-gaode", dirPrefix+"/static-gaode/")
	groupName := ""
	group := e.Group(groupName)
	{
		group.GET("/knearest", knearest4Gin)
		group.GET("/points", points4Gin)

		group.GET("/knearestShanghai", knearest4GinShanghai)
		group.GET("/pointsShanghai", points4GinShanghai)
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
	file, err := os.OpenFile(dirPrefix+"/static/capitals.csv", os.O_RDONLY, 0)

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
	PointsCount    int     `form:"pointsCount"`
	Index          string  `form:"index"`
}

func points4Gin(c *gin.Context) {
	rec := Rectangle{}
	c.BindQuery(&rec)

	visiblePoints := getIndex().Range(index.NewGeoPoint("topLeft", rec.TopLeftLat, rec.TopLeftLon),
		index.NewGeoPoint("bottomRight", rec.BottomRightLat, rec.BottomRightLon))

	c.JSON(http.StatusOK, visiblePoints)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
