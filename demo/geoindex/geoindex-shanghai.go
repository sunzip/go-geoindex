package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	index "github.com/hailocab/go-geoindex"
)

// var geoindexShanghai *index.ClusteringIndex
var geoindexShanghaiMap = make(map[int]*index.ClusteringIndex)

func knearest4GinShanghai(c *gin.Context) {
	p := Point{}
	c.ShouldBind(&p)

	nearest := getIndexShanghai(0).KNearest(index.NewGeoPoint("query", p.Lat, p.Lon), int(p.K), index.Km(5), func(_ index.Point) bool { return true })
	// data, _ := json.Marshal(nearest)
	c.JSON(http.StatusOK, nearest)
}

func points4GinShanghai(c *gin.Context) {
	rec := Rectangle{}
	c.BindQuery(&rec)

	visiblePoints := getIndexShanghai(rec.PointsCount).Range(index.NewGeoPoint("topLeft", rec.TopLeftLat, rec.TopLeftLon),
		index.NewGeoPoint("bottomRight", rec.BottomRightLat, rec.BottomRightLon))

	if false { //test 显示一个点，看看聚合的样式
		visiblePoints = []index.Point{}
		geoPoint := index.GeoPoint{Pid: "id", Plon: 121.2229, Plat: 31.100366}
		onePoint := index.CountPoint{GeoPoint: &geoPoint, Count: 3}
		visiblePoints = append(visiblePoints, onePoint)

		/* geoPoint = index.GeoPoint{Pid: "id2", Plon: 121.2229, Plat: 31.100366}
		twoPoint := index.CountPoint{GeoPoint: &geoPoint, Count: 2}
		visiblePoints = append(visiblePoints, twoPoint) */
	}

	c.JSON(http.StatusOK, visiblePoints)
}

func getIndexShanghai(pointsCount int) *index.ClusteringIndex {
	if v, ok := geoindexShanghaiMap[pointsCount]; !ok {
		geoindexShanghai := index.NewClusteringIndex()
		geoindexShanghaiMap[pointsCount] = geoindexShanghai

		capitals := shanghaiPointsRandom(pointsCount)
		id := 1

		for _, capital := range capitals {
			for i := 0; i < 10; i++ {
				id++

				geoindexShanghai.Add(index.NewGeoPoint(
					fmt.Sprintf("%d", id),
					capital.Lat()+rand.Float64()/30.0*sign(),
					capital.Lon()+rand.Float64()/30.0*sign(),
				))
			}
		}
		return geoindexShanghai
	} else {
		return v
	}
}
func shanghaiPointsRandom(pointsCount int) []index.Point {
	centerLon := 121.2229
	centerLat := 31.100366
	capitals := make([]index.Point, 0)

	var count = 30
	if pointsCount != 0 {
		count = pointsCount
	}
	// ran := rand.Rand
	for i := 0; i < count; i++ {
		id := fmt.Sprint(i)
		rInt := rand.Intn(1000)
		lat := centerLat + (float64(rInt) / 1000 * 0.400 * sign())
		rInt = rand.Intn(1000)
		lon := centerLon + (float64(rInt) / 1000 * 0.400 * sign())

		count := rand.Intn(10)
		capital := index.CountPoint{index.NewGeoPoint(id, lat, lon), count * 10}
		// capital := index.NewGeoPoint(id, lat, lon)
		capitals = append(capitals, capital)
	}

	return capitals
}
