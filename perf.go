package main

import (
	"github.com/gin-gonic/gin"
	"math"
)

func main() {
	router := gin.Default()
	router.POST("/print/subsets", func(context *gin.Context) {
		var numbers []int
		if context.ShouldBind(&numbers) == nil {
			pow := int(math.Pow(2, float64(len(numbers))))
			var result = make([][]int, pow)
			PrintAllSubsets(numbers, result, pow)
			context.JSON(200, result)
		}
	})
	_ = router.Run()
}

func PrintAllSubsets(nums []int, result [][]int, pow int) {
	len := len(nums)
	for i := 0; i < pow; i++ {
		temp := i
		count := 0
		result[i] = make([]int, len)
		for temp > 0 {
			bit := temp & 1
			if bit == 1 {
				result[i][count] = nums[count]
			}
			count++
			temp >>= 1
		}
	}
}
