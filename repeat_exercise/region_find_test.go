/**
* Date: 2021/5/25 8:12 上午
* Author: majunmin
**/
package repeat_exercise

import "testing"

func Test_RegionFind(t *testing.T) {

	rangeMap := map[*Range]int{
		&Range{Begin: 2, End: 7}: 0,
		&Range{Begin: 3, End: 6}: 7,
		&Range{Begin: 5, End: 9}: 3,
	}

	result := RangeFind(rangeMap)
	t.Log(SearchPoint(0, result))
	t.Log(SearchPoint(1, result))
	t.Log(SearchPoint(2, result))
	t.Log(SearchPoint(3, result))
	t.Log(SearchPoint(4, result))
	t.Log(SearchPoint(5, result))
	t.Log(SearchPoint(6, result))
	t.Log(SearchPoint(7, result))
	t.Log(SearchPoint(8, result))
	t.Log(SearchPoint(9, result))

}
