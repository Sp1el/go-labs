package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

func (p Person) PrintInfo() {
	fmt.Printf("Персона: %s, %d лет\n", p.name, p.age)
}

func addPerson(people *[]Person) {
	var person Person

	fmt.Print("Введите имя: ")
	fmt.Scanln(&person.name)

	fmt.Print("Введите возраст: ")
	fmt.Scanln(&person.age)

	*people = append(*people, person)
	fmt.Println("Персона добавлена!")
}

func (p *Person) Birthday() {
	p.age++
	fmt.Printf("У %s день рождение! Теперь ему %d лет\n", p.name, p.age)
}

func task1() {
	fmt.Println("\n=== Задание №1 : Структура Person ===")

	var person Person

	fmt.Print("Введите имя: ")
	fmt.Scanln(&person.name)

	fmt.Print("Введите возраст: ")
	fmt.Scanln(&person.age)

	fmt.Println("\nПерсона создана: ")
	person.PrintInfo()
}

func task2() {

	fmt.Println("\n=== Задание 2: Метод birthday() ===")

	var person Person

	fmt.Print("Введите имя: ")
	fmt.Scanln(&person.name)

	fmt.Print("Введите возраст: ")
	fmt.Scanln(&person.age)

	fmt.Println("\nПерсона создана: ")
	person.PrintInfo()

	fmt.Print("\nДо дня рождения: ")
	person.PrintInfo()

	person.Birthday()

	fmt.Print("После дня рождения: ")
	person.PrintInfo()
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func task3() {

	fmt.Println("\n=== Задание №3: Структура Circle ===")

	var radius float64
	fmt.Print("\nВведите радиус: ")
	fmt.Scanln(&radius)

	circle := Circle{radius}
	area := circle.Area()

	fmt.Printf("Круг с радиусом %.2f\n", circle.radius)
	fmt.Printf("Площадь: %.2f\n", area)

}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func task4() {
	fmt.Println("\n=== Задание №4 : Интерфейс Shape ===")

	var width int
	var height int
	var radius int

	fmt.Print("\nВведите ширину: ")
	fmt.Scanln(&width)

	fmt.Print("Введите высоту: ")
	fmt.Scanln(&height)

	fmt.Print("Введите радиус: ")
	fmt.Scanln(&radius)

	var shape1 Shape = Rectangle{float64(width), float64(height)}
	var shape2 Shape = Circle{float64(radius)}

	fmt.Printf("\nПрямоугольник %dx%d: площадь = %.2f\n", width, height, shape1.Area())
	fmt.Printf("Круг радиусом %d: площадь = %.2f\n", radius, shape2.Area())
}

func printAreas(shapes []Shape) {

	fmt.Println("Площади фигур: ")
	for i, shape := range shapes {
		fmt.Printf("%d. Площадь = %.2f\n", i+1, shape.Area())
	}
}

func task5() {

	fmt.Println("\n=== Задание №5 : Срез интерфейсов Shape ===")

	shapes := []Shape{
		Rectangle{width: 2, height: 3},
		Circle{radius: 2.5},
		Rectangle{width: 5, height: 5},
		Circle{radius: 1},
	}

	printAreas(shapes)
}

type Stringer interface {
	String() string
}

type Book struct {
	title  string
	author string
	year   int
}

func (b Book) String() string {
	return fmt.Sprintf(" \"%s\" - %s (%d год)", b.title, b.author, b.year)
}

func task6() {

	fmt.Println("\n=== Задание №6 : Интерфейс Stringer ===")

	var titleBook string
	var authorBook string
	var yearBook int

	fmt.Print("\nВведите название книги: ")
	fmt.Scanln(&titleBook)

	fmt.Print("Введите автора книги: ")
	fmt.Scanln(&authorBook)

	fmt.Print("Введите год публикации: ")
	fmt.Scanln(&yearBook)

	book := Book{
		title:  titleBook,
		author: authorBook,
		year:   yearBook,
	}

	fmt.Println("\nИнформация о книге: ")
	fmt.Println(book)

	var stringer Stringer = book
	fmt.Println("\nЧерез интерфейс: ")
	fmt.Println(stringer)
}
