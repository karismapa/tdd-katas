package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

// Sum return total value of given integer string
func Sum(nums string) int {
	if nums == "" {
		return 0
	}

	var arrNums []string
	if len(nums) > 5 && nums[:2] == "//" && nums[3:4] == "\n" {
		arrNums = strings.Split(nums[4:], string(nums[2]))
	} else {
		arrNums = strings.FieldsFunc(nums, split)
	}

	var total int
	var negatives []int
	for _, num := range arrNums {
		val, _ := strconv.Atoi(num)
		if val < 0 {
			negatives = append(negatives, val)
		} else {
			total += val
		}
	}

	if len(negatives) > 0 {
		fmt.Println("negatives not allowed", negatives)
	}

	return total
}

func split(r rune) bool {
	return r == ',' || r == '\n'
}
