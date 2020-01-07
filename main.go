package main

func main() {

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