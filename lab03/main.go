package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nЛабораторная работа 3")
	fmt.Println("=========================")

	for {
		interactiveMenu()

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
	fmt.Println("1 - 2. - Создание пакета mathutils / Нахождение факториала")
	fmt.Println("3. - Создание пакета stringutils / Переворот строки")
	fmt.Println("4. - Статистический массив")
	fmt.Println("5. - Динамический срез ")
	fmt.Println("6. - Самая длинная строка ")
	fmt.Println("0. - Выход ")
	fmt.Print("Выберите задание: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1, 2:
		tasks1_2()
	case 3:
		tasks3()
	case 4:
		tasks4()
	case 5:
		tasks5()
	case 6:
		tasks6()
	default:
		fmt.Println("Неверный выбор")

	}

}
