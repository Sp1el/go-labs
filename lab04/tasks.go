package main

import (
	"fmt"
	"strconv"
	"strings"
)

var mapPerson map[string]int

func initMapPerson() {
	mapPerson = map[string]int{
		"Алексей": 25,
		"Мария":   30,
		"Иван":    28,
		"Ольга":   35,
		"Петр":    32,
	}
}

func inputCustomMap() {
	mapPerson = make(map[string]int)

	fmt.Println("\n=== Создание пользовательской карты ===")

	var count int
	fmt.Print("Сколько человек хотите добавить? ")
	fmt.Scanln(&count)

	if count <= 0 {
		fmt.Println("Неверное количество!\nКарта использована по умолчанию.")
		initMapPerson()
		return
	}

	for i := 0; i < count; i++ {
		var name string
		var age int

		fmt.Printf("\nЧеловек %d:\n", i+1)
		fmt.Print(" Имя: ")
		fmt.Scanln(&name)

		fmt.Print(" Возраст: ")
		fmt.Scanln(&age)

		mapPerson[name] = age
		fmt.Printf("Добавлено: %s - %d лет\n", name, age)
	}

	fmt.Printf("\n Готово! Добавлено %d человекa.\n", len(mapPerson))
}

func outputMap() {
	for name, age := range mapPerson {
		fmt.Printf(" %s: %d лет\n", name, age)
	}
}

func averageAge() {
	sum := 0
	for _, age := range mapPerson {
		sum += age
	}

	average := float64(sum) / float64(len(mapPerson))
	fmt.Printf("\nСредний возраст: %2.f лет\n", average)
}

func task1() {
	fmt.Println("=== Задание №1 : Добавление нового человека.")

	fmt.Println("\nТекущая карта: ")
	outputMap()

	fmt.Print("Имя нового человека: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Возраст нового человека: ")
	var age int
	fmt.Scanln(&age)

	mapPerson[name] = age

	fmt.Println("\nПосле добавления: ")
	outputMap()

}

func task2() {

	fmt.Println("\n=== Задание №2 : Средний возраст.")

	if len(mapPerson) == 0 {
		fmt.Println("Карта пустая! Выполните задание 1.")
		return
	}

	fmt.Println("Текущая карта: ")
	outputMap()

	averageAge()

}

func task3() {

	fmt.Println("\n=== Задание 3: Удаление по имени ===")

	if len(mapPerson) == 0 {
		fmt.Println("Карта пуста! Сначала выполните Задание 1.")
		return
	}

	fmt.Println("Текущая карта: ")
	outputMap()

	var name string
	fmt.Printf("\nВведите имя для удаления: ")
	fmt.Scanln(&name)

	if _, exists := mapPerson[name]; exists {
		delete(mapPerson, name)
		fmt.Printf("Удален: %s\n", name)
	} else {
		fmt.Printf(" Имя '%s' не найдено!\n", name)
	}

	fmt.Println("\nОбновленная карта: ")
	outputMap()
}

func task4() {

	fmt.Println("\n=== Задание 4: Строка в верхнем регистре ===")

	fmt.Print("Введите строку: ")
	var input string
	fmt.Scanln(&input)

	fmt.Printf("\nСтрока в верхнем регистре: %s", strings.ToUpper(input))
}

func task5() {

	fmt.Println("\n=== Задание 5: Сумма чисел ===")
	fmt.Println("Введите несколько чисел через пробел: ")

	var input string
	fmt.Scanln(&input)

	parts := strings.Fields(input)
	sum := 0

	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		sum += num
	}

	fmt.Printf("Сумма: %d\n", sum)
}

func printArray(arr []int) {
	fmt.Print("[")
	for i, num := range arr {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(num)
	}
	fmt.Println("]")
}

func reverseArray(arr []int) []int {
	reversed := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		reversed[i] = arr[len(arr)-1-i]
	}
	return reversed
}

func task6() {

	fmt.Println("\n=== Задание 6: Реверс массива чисел ===")

	var size int
	fmt.Printf("Введите размер массива: ")
	fmt.Scanln(&size)

	if size <= 0 {
		fmt.Println("Размер массива должен быть больше 0!")
		return
	}

	arr := make([]int, size)

	fmt.Printf("Введите %d целых чисел: \n", size)
	for i := 0; i < size; i++ {
		fmt.Printf("Элемент [%d]: ", i)
		fmt.Scanln(&arr[i])
	}

	fmt.Println("\nИсходный массив: ")
	printArray(arr)

	reversed := reverseArray(arr)
	printArray(reversed)
}
