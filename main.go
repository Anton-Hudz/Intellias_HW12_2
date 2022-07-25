package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	// на основании кода из первого задания:
	var wg sync.WaitGroup
	result := 0
	for i := 0; i < len(n); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result += sum(n[i])
		}(i)
	}
	wg.Wait()
	fmt.Println("result:", result)
}

func sum(n []int) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}
