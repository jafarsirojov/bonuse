package main

import "fmt"

func main() {
	sales := []int{12_000,8_000,15_000,8_000}
	for _, sale := range sales{
		result :=bonuses(sale)
		fmt.Println(result)
	}

}

func bonuses(sale int)  int{
	const border=10_000
	const percent=5
	total :=0
	if sale>border{
		total=(sale-border)*percent/100
	}
	return total

}