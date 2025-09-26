package main

import (
	"fmt"
	"lab3/mathutils"
	"lab3/stringutils"
)

func tasks1_2() {
	fmt.Println("\n=== Задание 1 - 2 ===")
	fmt.Println("Был создан пакет (папка) mathutils в каталоге lab3")
	fmt.Println("Прописан алгоритм для нахождение факториала для целого числа")

	var num int

	fmt.Print("\nВведите число: ")
	fmt.Scanln(&num)

	fmt.Printf("Факториал числа %d = %d\n", num, mathutils.Factorial(num))
}

func tasks3() {
	fmt.Println("\n=== Задание 3 ===")
	fmt.Println("Был создан пакет (папка) stringutils в каталоге lab3")
	fmt.Println("Прописан алгоритм для переворота строки")

	var input string

	fmt.Print("\nВведите строку: ")
	fmt.Scanln(&input)

	fmt.Printf("Строка до переворота %s и после %s\n", input, stringutils.Reverse(input))

}

func tasks4() {
	fmt.Println("\n=== Задание 4 ===")
	fmt.Println("Создать массив из 5 целых чисел и вывести на экран")

	var numbers [5]int
	var choice int

	fmt.Println("Выберите способ заполнения массива:")
	fmt.Println("1. Заполнить значениями по умолчанию (1, 2, 3, 4, 5)")
	fmt.Println("2. Ввести значения вручную")
	fmt.Print("Ваш выбор: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		numbers = [5]int{1, 2, 3, 4, 5}

	case 2:
		fmt.Println("Введите 5 целых чисел: ")
		for i := 0; i < 5; i++ {
			fmt.Printf("nubmers[%d] = ", i)
			fmt.Scanln(&numbers[i])
		}
		fmt.Println("Массив заполнен вручную")

	default:
		fmt.Println("Неверный выбор, используется вариант по умолчанию")
		numbers = [5]int{1, 2, 3, 4, 5}
	}

	fmt.Println("\nПолученный массив: ")
	for i, num := range numbers {
		fmt.Printf("numbers[%d] = %d\n", i, num)
	}
}

func tasks5() {

	fmt.Println("\n=== Задание 5 ===")
	fmt.Println("Выполнить операции со срезами, добавлением и удалением")

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Исходный массив: %v\n", arr)

	slice := arr[1:4]
	fmt.Printf("Срез [1:4]: %v\n", slice)

	slice = append(slice, 6)
	fmt.Printf("После добавления 6: %v\n", slice)

	slice = append(slice[:1], slice[2:]...)
	fmt.Printf("После удаления элемента с индексом 1: %v\n", slice)

}

func findLongestString(strings []string) (string, int) {
	if len(strings) == 0 {
		return "", 0
	}

	longest := strings[0]
	maxLenght := len(strings[0])

	for i := 1; i < len(strings); i++ {
		currecntLenght := len(strings[i])
		if currecntLenght > maxLenght {
			longest = strings[i]
			maxLenght = currecntLenght
		}
	}

	return longest, maxLenght
}

func tasks6() {

	fmt.Println("\n=== Задание 6 ===")
	fmt.Println("Самая длинная строка")

	var count int

	fmt.Print("Сколько строк вы хотите ввести? ")
	fmt.Scanln(&count)

	if count <= 0 {
		fmt.Println("Неверное количество строк!")
		return
	}

	strings := make([]string, count)

	fmt.Println("Введите строки: ")
	for i := 0; i < count; i++ {
		fmt.Printf("Строка %d: ", i+1)
		fmt.Scanln(&strings[i])
	}

	fmt.Println("\nВведеные строки: ", strings)

	longest, length := findLongestString(strings)

	fmt.Printf("Самая длинная строка: '%s' (Длина: %d символов)", longest, length)

}
