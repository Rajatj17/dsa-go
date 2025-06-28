package stockspan

import basic "DSA/Stack/a_basic"

type StockItem struct {
	index int
	value int
}

func StockSpan(prices []int) []int {
	stack := basic.NewStack[StockItem]()
	n := len(prices)

	result := []int{1}

	for i := 1; i < n; i++ {
		top, exists := stack.Top()
		for exists && top.value <= prices[i] {
			stack.Pop()
			top, exists = stack.Top()
		}

		span := i - top.index
		result = append(result, span)

		stack.Push(StockItem{
			index: i,
			value: top.value,
		})
	}

	return result
}
