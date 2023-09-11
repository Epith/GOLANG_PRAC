package main

import "fmt"

func main(){
	//var ages [3]int = [3]int{20, 25, 39}
	var ages = [3]int{20, 25, 39}

	names := [4]string{"james","daniel","jake","vasco"}
	names[1] = "Jay"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slice
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	//slice range
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "zack")

	fmt.Println(rangeOne)
}