package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	index "github.com/hailocab/go-geoindex"
)

// 废弃
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

// 废弃
func knearest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lat, _ := strconv.ParseFloat(r.Form["lat"][0], 64)
	lon, _ := strconv.ParseFloat(r.Form["lon"][0], 64)
	k, _ := strconv.ParseInt(r.Form["k"][0], 10, 32)

	nearest := getIndex().KNearest(index.NewGeoPoint("query", lat, lon), int(k), index.Km(5), func(_ index.Point) bool { return true })
	data, _ := json.Marshal(nearest)
	fmt.Fprintln(w, string(data))
}
