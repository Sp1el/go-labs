package main

import (
	"fmt"
	"time"
)

func Task1() {

	fmt.Println("=== Задание 1: Текущее время и дата ===")

	currentTime := time.Now()
	fmt.Println("\n\tТекущее время и дата: ", currentTime)
	fmt.Println("\tотформатированная дата: ", currentTime.Format("02.01.2006 15:04:05"))

}

func Task2() {

	fmt.Println("\n=== Задание 2: Переменные различных типов ===")

	var numberInt int = 10
	var numberFloat float64 = 15.4
	var message string = "Hello"
	var isActive bool = true

	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Println("\n\t", numberInt, numberFloat, message, isActive)
	fmt.Println("\t", defaultInt, defaultFloat, defaultString, defaultBool)

}

func Task3() {

	fmt.Println("\n=== Задание 3: Краткая форма объявления ===")

	name := "student"
	age := 20
	salary := 100.00
	isProgram := true

	fmt.Println("\n\t", name, age, salary, isProgram)

}

func Task4() {

	fmt.Println("\n=== Задание 4: Арифметические операции ===")

	num1 := 10
	num2 := 2

	fmt.Printf("\n\t%d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("\t%d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("\t%d * %d = %d\n", num1, num2, num1*num2)
	fmt.Printf("\t%d / %d = %d\n", num1, num2, num1/num2)
	fmt.Printf("\t%d %% %d = %d\n", num1, num2, num1%num2)
	// %d - целое число, %s - строка, %f - float, %t - bool, %v - универсальный спецификатор.
	// Чтобы вывести знак %, нужно удвоить: %%

}

func Task5() {

	fmt.Println("\n=== Задание 5: Функции с числами с плавающей точкой ===")

	x := 15.3
	y := 4.2

	resultSum, resultDif := calculateFloat(x, y)

	fmt.Printf("\n\tСумма %.2f\n", resultSum)
	fmt.Printf("\tРазность %.2f\n", resultDif)

}

func calculateFloat(a, b float64) (float64, float64) {
	return a + b, a - b
}

func Task6() {

	fmt.Println("\n=== Задание 6: Среднее значение трех чисел ===")
	a := 5
	b := 10
	c := 15
	fmt.Printf("\n\tЧисла %d, %d, %d\n", a, b, c)
	fmt.Printf("\tСреднее значение трех чисел: %f\n", float64((a+b+c)/3))

}
