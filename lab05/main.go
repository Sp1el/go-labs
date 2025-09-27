package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n=== Лабораторная работа № 5 ===")

	interactiveMenu()
}

func interactiveMenu() {
	for {
		fmt.Println("\n ГЛАВНОЕ МЕНЮ")
		fmt.Println("1. Задание 1 - Структура Person и методы")
		fmt.Println("2. Задание 2 - Метод birthday()")
		fmt.Println("3. Задание 3 - Структура Circle и площадь")
		fmt.Println("4. Задание 4 - Интерфейс Shape")
		fmt.Println("5. Задание 5 - Срез интерфейсов Shape")
		fmt.Println("6. Задание 6 - Интерфейс Stringer для Book")
		fmt.Println("0. Выход")
		fmt.Print("Ваш выбор: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("❌ Ошибка ввода! Введите число.")
			// Очищаем буфер ввода
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choice {
		case 1:
			clearScreen()
			task1()
		case 2:
			clearScreen()
			task2()
		case 3:
			clearScreen()
			task3()
		case 4:
			clearScreen()
			task4()
		case 5:
			clearScreen()
			task5()
		case 6:
			clearScreen()
			task6()
		case 0:
			fmt.Println("\nДо свидания!")
			return
		default:
			fmt.Println("Неверный выбор! Попробуйте снова.")
		}

		fmt.Print("\nНажмите Enter для продолжения...")
		fmt.Scanln()
		clearScreen()
	}
}

// Функция для очистки экрана (кроссплатформенная)
func clearScreen() {
	// Простая очистка выводом пустых строк
	fmt.Print("\033[2J") // ANSI код очистки экрана
	fmt.Print("\033[H")  // Перемещение курсора в начало
}
