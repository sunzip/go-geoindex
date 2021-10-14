package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	index "github.com/hailocab/go-geoindex"
)

var geoindexShanghai *index.ClusteringIndex

func knearest4GinShanghai(c *gin.Context) {
	p := Point{}
	c.ShouldBind(&p)

	nearest := getIndexShanghai().KNearest(index.NewGeoPoint("query", p.Lat, p.Lon), int(p.K), index.Km(5), func(_ index.Point) bool { return true })
	// data, _ := json.Marshal(nearest)
	c.JSON(http.StatusOK, nearest)
}

func points4GinShanghai(c *gin.Context) {
	rec := Rectangle{}
	c.BindQuery(&rec)

	visiblePoints := getIndexShanghai().Range(index.NewGeoPoint("topLeft", rec.TopLeftLat, rec.TopLeftLon),
		index.NewGeoPoint("bottomRight", rec.BottomRightLat, rec.BottomRightLon))

	c.JSON(http.StatusOK, visiblePoints)
}

func getIndexShanghai() *index.ClusteringIndex {
	if geoindexShanghai == nil {
		geoindexShanghai = index.NewClusteringIndex()

		capitals := shanghaiPointsRandom()
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
	}

	return geoindexShanghai
}
func shanghaiPointsRandom() []index.Point {
	centerLon := 121.2229
	centerLat := 31.100366
	capitals := make([]index.Point, 0)

	// ran := rand.Rand
	for i := 0; i < 30; i++ {
		id := fmt.Sprint(i)
		rInt := rand.Intn(1000)
		lat := centerLat + (float64(rInt) / 1000 * 0.400 * sign())
		rInt = rand.Intn(1000)
		lon := centerLon + (float64(rInt) / 1000 * 0.400 * sign())

		capital := index.NewGeoPoint(id, lat, lon)
		capitals = append(capitals, capital)
	}

	return capitals
}
