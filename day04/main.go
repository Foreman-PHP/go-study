package main

import "fmt"

var (
	coins        = 5000
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func main() {
	dispatchCoin()
	sum := sum(distribution)
	fmt.Println(coins - sum)
	fmt.Println(distribution)
}

func dispatchCoin() {
	for _, name := range users {
		for _, v := range name {
			switch v {
			case 'e', 'E':
				distribution[name]++
			case 'i', 'I':
				distribution[name] += 2
			case 'o', 'O':
				distribution[name] += 3
			case 'u', 'U':
				distribution[name] += 4
			}
		}
	}
}

func sum(dis map[string]int) (coin int) {
	coin = 0
	for _, v := range dis {
		coin += v
	}
	return
}
