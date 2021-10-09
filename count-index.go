package geoindex

import (
	"fmt"
)

type CountIndex struct {
	index           *geoIndex
	currentPosition map[string]Point
}

type CountPoint struct {
	*GeoPoint
	Count interface{}
}

func (p *CountPoint) String() string {
	return fmt.Sprintf("%f %f %d", p.Lat(), p.Lon(), p.Count)
}

// NewCountIndex creates an index which counts the points in each cell.
func NewCountIndex(resolution Meters) *CountIndex {
	newCounter := func() interface{} {
		return &singleValueAccumulatingCounter{}
	}

	return &CountIndex{newGeoIndex(resolution, newCounter), make(map[string]Point)}
}

// NewExpiringCountIndex creates an index, which maintains an expiring counter for each cell.
//  expiringCounter.Remove 并没有实现。所以 ， 该构造方法不能用？
func NewExpiringCountIndex(resolution Meters, expiration Minutes) *CountIndex {
	newExpiringCounter := func() interface{} {
		return newExpiringCounter(expiration)
	}

	return &CountIndex{newGeoIndex(resolution, newExpiringCounter), make(map[string]Point)}
}

func (index *CountIndex) Clone() *CountIndex {
	clone := &CountIndex{}

	// Copy all entries from current positions
	clone.currentPosition = make(map[string]Point, len(index.currentPosition))
	// 取v的值没问题，&v是不变的
	for k, v := range index.currentPosition {
		clone.currentPosition[k] = v
	}

	// Copying underlying geoindex data
	clone.index = index.index.Clone()

	return clone
}

// Add adds a point.
func (countIndex *CountIndex) Add(point Point) {
	countIndex.Remove(point.Id())
	countIndex.currentPosition[point.Id()] = point
	// GEO 的 index 添加 EntryAt （ 就是根据分辨率抽象的离散点，关联到该 point 上 ）
	/* 多个点对应一个cell，因此，需要先清理掉之前的关联关系 */
	countIndex.index.AddEntryAt(point).(counter).Add(point)
}

// Remove removes a point.
func (countIndex *CountIndex) Remove(id string) {
	if prev, ok := countIndex.currentPosition[id]; ok {
		countIndex.index.GetEntryAt(prev).(counter).Remove(prev)
		// expiringCounter.Remove 并没有实现
		// singleValueAccumulatingCounter.Remove 实现了
		delete(countIndex.currentPosition, id)
	}
}

// Range returns the counters within some lat, lng range.
func (countIndex *CountIndex) Range(topLeft Point, bottomRight Point) []Point {
	// 获取的是cell的集合。
	/*
		为什么不使用 countIndex.CurrentPoints ，然后进行判断
		感觉，类似，之前转成离散的int点，然后判断
	*/
	counters := countIndex.index.Range(topLeft, bottomRight)

	points := make([]Point, 0)

	for _, c := range counters {
		if c.(counter).Point() != nil {
			points = append(points, c.(counter).Point())
		}
	}

	return points
}

// KNearest just to satisfy an interface. Doesn't make much sense for count index.
//  对计数索引没有多少意义
func (index *CountIndex) KNearest(point Point, k int, maxDistance Meters, accept func(p Point) bool) []Point {
	panic("Unsupported operation")
}
