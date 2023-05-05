package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, -126.3,13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	min := temperatures[0]
	max := temperatures[0]
	for i :=1; i < len(temperatures); i++ {
		if temperatures[i] < min {
			min = temperatures[i]
			continue
		}
		if temperatures[i] > max {
			max = temperatures[i]
			continue
		}
	}

	var minKey int
	// округляем температуру до ближайшего числа, кратного 10 (отрицательные вверх, положительные вниз)
	if dividedMin := min/10.0; dividedMin < 0{
		minKey = int(math.Ceil(dividedMin)) * 10
	} else {
		minKey = int(math.Floor(dividedMin)) * 10
	}
	var maxKey int
	if dividedMax:= max/10.0; dividedMax < 0{
		maxKey = int(math.Ceil(dividedMax)) * 10
	} else {
		maxKey = int(math.Floor(dividedMax)) * 10
	}

	// создаем пустой словарь
	sets := make(map[int][]float64)

	//заполняем словарь пустыми значениями, чтобы сохранить шаг 10 в ключах
	for i := minKey; i <= maxKey; i+=10{
		sets[i] = []float64{}
	}

	// проходим по всей последовательности температур
	for _, temp := range temperatures {
		// округляем температуру до ближайшего числа, кратного 10 (отрицательные вверх, положительные вниз)
		var rounded int
		if dividedTemp:= temp/10.0; dividedTemp < 0{
			rounded = int(math.Ceil(dividedTemp)) * 10
		} else {
			rounded = int(math.Floor(dividedTemp)) * 10
		}
		// добавляем температуру в список в словаре
		sets[rounded] = append(sets[rounded], temp)
	}

	// получаем список ключей словаря
	keys := make([]int, 0, len(sets))
	for k := range sets {
		keys = append(keys, k)
	}

	// сортируем ключи
	sort.Ints(keys)

	// выводим результат по упорядоченным ключам
	for _, key := range keys {
		fmt.Printf("%d: %f\n", key, sets[key])
	}
}
