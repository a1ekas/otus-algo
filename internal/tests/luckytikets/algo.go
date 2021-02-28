package luckytickets

import "math"

// luckyTicketsAlgo ...
type luckyTicketsAlgo struct {
	maxSumDigits int
	countNumbers int
	digitsSums   map[int]int64
}

// newLuckyTicketsAlgo констуктор алгоритма подсчета счастливых билетов
func newLuckyTicketsAlgo(N int) *luckyTicketsAlgo {
	a := luckyTicketsAlgo{
		maxSumDigits: 9 * N,
		countNumbers: int(math.Pow(10, float64(N))),
		digitsSums:   map[int]int64{},
	}

	return &a
}

// count ...
func (a *luckyTicketsAlgo) count() int64 {
	a.initSums()
	a.prepareDigitsSumForNumbers()
	return a.getTotalCount()
}

func (a *luckyTicketsAlgo) initSums() {
	for i := 0; i <= a.maxSumDigits; i++ {
		a.digitsSums[i] = 0
	}
}

func (a *luckyTicketsAlgo) prepareDigitsSumForNumbers() {
	for i := 0; i < a.countNumbers; i++ {
		a.countDigitsSumForNumber(i)
	}
}

// digitsSum - вычисляем сумму цифр в числе
func (a *luckyTicketsAlgo) digitsSum(number int) int {
	if number == 0 {
		return number
	}
	return int((number % 10) + a.digitsSum(number/10))
}

// countDigitsSumForNumber - подсчитываем количество одинаковых сум ()
func (a *luckyTicketsAlgo) countDigitsSumForNumber(number int) {
	sum := a.digitsSum(number)
	//Учитываем в массиве сумм
	if sum <= a.maxSumDigits {
		a.digitsSums[sum]++
	}
}

func (a *luckyTicketsAlgo) getTotalCount() int64 {
	var count int64
	count = 0
	for i := 0; i <= a.maxSumDigits; i++ {
		count += a.digitsSums[i] * a.digitsSums[i]
	}
	return count
}
