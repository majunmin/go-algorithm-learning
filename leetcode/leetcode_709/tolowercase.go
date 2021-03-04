/**
  @Author: majm@ushareit.com
  @date: 2021/2/18
  @note:
**/
package leetcode_709

//
func toLowerCase(str string) string {
	tempArr := []byte(str)
	for i := range tempArr {
		tempArr[i] |= 32
	}
	return string(tempArr)
}
