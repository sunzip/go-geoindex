package geoindex

var (
	minLon          = -180.0
	minLat          = -90.0
	latDegreeLength = Km(111.0)
	lonDegreeLength = Km(85.0)
)

type Meters float64

func Km(km float64) Meters {
	return Meters(km * 1000)
}

func Meter(meters float64) Meters {
	return Meters(meters)
}

type cell struct {
	x int
	y int
}

// 分辨率似乎是地面长度
//  x,y 组成的 cell 只是抽象的不同分辨率层级上的离散点。不是坐标
//  这个抽象的离散点，就起到聚合作用。 多个坐标 point 可以抽象到一个离散点
//  x 是维度，所以，右边数值小
//  y 是经度，所以，右边数值小
func cellOf(point Point, resolution Meters) cell {
	x := int((-minLat + point.Lat()) * float64(latDegreeLength) / float64(resolution))
	y := int((-minLon + point.Lon()) * float64(lonDegreeLength) / float64(resolution))

	return cell{x, y}
}

type geoIndex struct {
	resolution Meters
	// 一个 cell 一个点。起不到聚合的作用
	index    map[cell]interface{}
	newEntry func() interface{}
}

// Creates new geo index with resolution a function that returns a new entry that is stored in each cell.
//  分辨率似乎是地面长度 （ 1 pixel 像素map到地面的长度？ ）
func newGeoIndex(resolution Meters, newEntry func() interface{}) *geoIndex {
	return &geoIndex{resolution, make(map[cell]interface{}), newEntry}
}

func (i *geoIndex) Clone() *geoIndex {
	clone := &geoIndex{
		resolution: i.resolution,
		index:      make(map[cell]interface{}, len(i.index)),
		newEntry:   i.newEntry,
	}
	for k, v := range i.index {
		set, ok := v.(set)
		if !ok {
			panic("Cannot cast value to set")
		}
		clone.index[k] = set.Clone()
	}

	return clone
}

// AddEntryAt adds an entry if missing, returns the entry at specific position.
func (geoIndex *geoIndex) AddEntryAt(point Point) interface{} {
	square := cellOf(point, geoIndex.resolution)

	if _, ok := geoIndex.index[square]; !ok {
		geoIndex.index[square] = geoIndex.newEntry()
	}

	return geoIndex.index[square]
}

// GetEntryAt gets an entry from the geoindex, if missing returns an empty entry without changing the index.
func (geoIndex *geoIndex) GetEntryAt(point Point) interface{} {
	square := cellOf(point, geoIndex.resolution)

	entries, ok := geoIndex.index[square]
	if !ok {
		return geoIndex.newEntry()
	}

	return entries
}

// Range returns the index entries within lat, lng range.
func (geoIndex *geoIndex) Range(topLeft Point, bottomRight Point) []interface{} {
	topLeftIndex := cellOf(topLeft, geoIndex.resolution)
	bottomRightIndex := cellOf(bottomRight, geoIndex.resolution)

	//右下的x怎么是最小的呢？ 因为x 是维度，右边数值小
	return geoIndex.get(bottomRightIndex.x, topLeftIndex.x, topLeftIndex.y, bottomRightIndex.y)
}

// 获取entry
func (geoIndex *geoIndex) get(minx int, maxx int, miny int, maxy int) []interface{} {
	entries := make([]interface{}, 0, 0)

	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			if indexEntry, ok := geoIndex.index[cell{x, y}]; ok {
				entries = append(entries, indexEntry)
			}
		}
	}

	return entries
}

func (g *geoIndex) getCells(minx int, maxx int, miny int, maxy int) []cell {
	indices := make([]cell, 0)

	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			indices = append(indices, cell{x, y})
		}
	}

	return indices
}
