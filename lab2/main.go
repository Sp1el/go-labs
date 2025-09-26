package main

import "fmt"

func main() {

	fmt.Println("Лабораторная работа 2")
	fmt.Println("=====================================")

	for {
		interactiveMenu()

		// Спросить, хочет ли пользователь продолжить
		var continueChoice string
		fmt.Print("\nХотите выбрать другое задание? (y/n): ")
		fmt.Scanln(&continueChoice)
		if continueChoice != "y" && continueChoice != "Y" {
			fmt.Println("Выход из программы. До свидания!")
			break
		}
	}

}

func interactiveMenu() {
	fmt.Println("\n Интерактивное меню: ")
	fmt.Println("1. - Четное или нечетное число")
	fmt.Println("2. - Определитель знака числа")
	fmt.Println("3. - Генератор последовательсти")
	fmt.Println("4. - Длина строки")
	fmt.Println("5. - Площадь прямоугольника ")
	fmt.Println("6. - Среднее значение двух целых чисел ")
	fmt.Println("0. - Выход ")
	fmt.Print("Выберите задание: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		Task1()

	case 2:
		Task2()

	case 3:
		Task3()

	case 4:
		Task4()

	case 5:
		Task5()

	case 6:
		Task6()

	default:
		fmt.Println("Неверный выбор!")

	}

}
