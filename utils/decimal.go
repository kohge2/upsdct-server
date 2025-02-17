package utils

import (
	"math"
	"strconv"
	"strings"
)

// 小数以下を切り捨てで整数*少数の計算 (小数計算の問題の対策で小数を整数に戻して計算)
func MultiplyIntByDecimal(i int, d float64) int {
	s := strconv.FormatFloat(d, 'f', -1, 64)
	if !strings.Contains(s, ".") {
		return 0
	}
	parts := strings.Split(s, ".")
	shift := math.Pow10(len(parts[1]))

	shiftedD := int(d * shift)
	shiftedResult := i * shiftedD

	return int(float64(shiftedResult) / shift)
}
