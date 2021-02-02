package lkytkts

import "math"

// LuckyTicketsAlgo ...
type LuckyTicketsAlgo struct {
	maxSumDigits int
	countNumbers int
	digitsSums   map[int]int64
}

// NewLuckyTicketsAlgo констуктор алгоритма подсчета счастливых билетов
func NewLuckyTicketsAlgo(N int) *LuckyTicketsAlgo {
	a := LuckyTicketsAlgo{
		maxSumDigits: 9 * N,
		countNumbers: int(math.Pow(10, float64(N))),
		digitsSums:   map[int]int64{},
	}

	return &a
}

// Count ...
func (a *LuckyTicketsAlgo) Count() int64 {
	a.initSums()
	a.prepareDigitsSumForNumbers()
	return a.getTotalCount()
}

func (a *LuckyTicketsAlgo) initSums() {
	for i := 0; i <= a.maxSumDigits; i++ {
		a.digitsSums[i] = 0
	}
}

func (a *LuckyTicketsAlgo) prepareDigitsSumForNumbers() {
	for i := 0; i < a.countNumbers; i++ {
		a.countDigitsSumForNumber(i)
	}
}

// digitsSum - вычисляем сумму цифр в числе
func (a *LuckyTicketsAlgo) digitsSum(number int) int {
	if number == 0 {
		return number
	}
	return int((number % 10) + a.digitsSum(number/10))
}

// countDigitsSumForNumber - подсчитываем количество одинаковых сум ()
func (a *LuckyTicketsAlgo) countDigitsSumForNumber(number int) {
	sum := a.digitsSum(number)
	//Учитываем в массиве сумм
	if sum <= a.maxSumDigits {
		a.digitsSums[sum]++
	}
}

func (a *LuckyTicketsAlgo) getTotalCount() int64 {
	var count int64
	count = 0
	for i := 0; i <= a.maxSumDigits; i++ {
		count += a.digitsSums[i] * a.digitsSums[i]
	}
	return count
}
