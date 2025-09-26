package main

import "fmt"

func Task1() {
	fmt.Println("\n === Задание 1: Четное или нечетное число ===")
	var num int

	fmt.Print("Введите число: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("Число %d - четное\n", num)
	} else {
		fmt.Printf("Число %d - нечетное\n", num)
	}
}

func Task2() {
	fmt.Println("\n === Задание 2: Определитель знака числа ===")
	var num int

	fmt.Print("Введите число: ")
	fmt.Scanln(&num)

	switch {
	case num > 0:
		fmt.Println("Число - Positive")
	case num < 0:
		fmt.Println("Число - Negative")
	default:
		fmt.Println("Число - Zero")
	}
}

func Task3() {
	fmt.Println("\n === Задание 3: Генератор последовательности ===")
	fmt.Println("Используя конструкцию for")
	fmt.Println("Числа от 1 до 10: ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func Task4() {
	fmt.Println("\n === Задание 4: Длина строки ===")
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	lenght := len(input)
	fmt.Printf("Длина строки '%s' = %d символов\n", input, lenght)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height

}
func Task5() {
	fmt.Println("\n=== Задание 5: Площадь прямоугольника ===")
	var w, h float64
	fmt.Print("Введите ширину и высоту прямоугольника: ")
	fmt.Scanln(&w, &h)

	rect := Rectangle{Width: w, Height: h}
	area := rect.Area()

	fmt.Printf("Площадь прямоугольника %.2f x %.2f = %.2f\n", w, h, area)
}

func Task6() {
	fmt.Println("\n=== Задание 6: Среднее значение ===")
	var num1, num2 int
	fmt.Print("Введите два целых числа: ")
	fmt.Scanln(&num1, &num2)

	average := (float64(num1) + float64(num2)) / 2.0
	fmt.Printf("Среднее значение чисел %d и %d = %.2f\n", num1, num2, average)
}
