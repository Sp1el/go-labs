package main

import "fmt"

func main() {

	fmt.Println("\n\n=== Лабораторная работа №4 ===")

	fmt.Print("\nВвести карту вручную (1) или использовать значения по умолчанию (2)? (1/2)")
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		inputCustomMap()
	} else {
		initMapPerson()
	}

	interactiveMenu()
}

func interactiveMenu() {

	for {
		fmt.Println("\nГЛАВНОЕ МЕНЮ")
		fmt.Printf("Текущий размер карты: %d записей\n", len(mapPerson))
		fmt.Println("1. Добавить нового человека")
		fmt.Println("2. Средний возраст из карты")
		fmt.Println("3. Удаление из карты по имени")
		fmt.Println("4. Конвертер в верхний регистр")
		fmt.Println("5. Сумматор чисел")
		fmt.Println("6. Реверс массива чисел")
		fmt.Println("0. Выход")
		fmt.Print("Выберите задание: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			task1()
		case 2:
			task2()
		case 3:
			task3()
		case 4:
			task4()
		case 5:
			task5()
		case 6:
			task6()
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор!")
		}
	}
}
