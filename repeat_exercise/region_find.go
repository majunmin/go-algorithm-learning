/**
* Date: 2021/5/25 12:26 上午
* Author: majunmin
**/
package repeat_exercise

import (
	"math"
	"sort"
)

type Range struct {
	Begin uint64
	End   uint64
}

type Points []*Point

type Point struct {
	point uint64
	value []int
}

// 给定区间 对应的 值id, 给一个 区间中的值， 给出对对应的 id
// example:
// [2,7] => 0  [3,6] => 7  [5,9] => 3
func RangeFind(rangeMap map[*Range]int) Points {
	var points []*Point
	existMap := make(map[uint64]int)

	points = append(points, &Point{point: 0})
	points = append(points, &Point{point: math.MaxUint64})
	for r := range rangeMap {
		if _, exist := existMap[r.Begin]; !exist {
			existMap[r.Begin] = 0
			points = append(points, &Point{point: r.Begin})
		}
		if _, exist := existMap[r.End]; !exist {
			existMap[r.End] = 0
			points = append(points, &Point{point: r.End})
		}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].point < points[j].point
	})

	for i, point := range points {
		existMap[point.point] = i
	}

	for r, val := range rangeMap {
		begin, end := r.Begin, r.End
		for i := begin; i < end; i++ {
			if idx, exist := existMap[i]; exist {
				points[idx].value = append(points[idx].value, val)
			}
		}
	}
	for _, point := range points {
		sort.Ints(point.value)
	}
	return points
}

func SearchPoint(search uint64, points Points) []int {
	if len(points) == 0 {
		return []int{}
	}

	left, right := 0, len(points)
	for left < right {
		mid := left + (right-left+1)>>1
		if search < points[mid].point {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return points[left].value
}
