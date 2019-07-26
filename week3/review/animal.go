package main

import (
	"fmt"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() string{
	return a.food
}

func (a Animal) Move() string{
	return a.locomotion
}

func (a Animal) Speak() string{
	return a.noise
}

func main() {
	cow := Animal{"grass","walk","moo"}
	bird := Animal{"worms","fly","peep"}
	snake := Animal{"mice","slither","hsss"}
	
	var animal, data string

	for {
		fmt.Println("Please type 1 of these animals: cow, bird or snake and 1 of these datas: eat, move or speak")
		fmt.Scanln(&animal,&data)
		switch animal{
		case "cow":
			switch data{
			case "eat":
				fmt.Println(cow.Eat())
			case "move":
				fmt.Println(cow.Move())
			case "speak":
				fmt.Println(cow.Speak())
			}
		case "bird":
			switch data{
			case "eat":
				fmt.Println(bird.Eat())
			case "move":
				fmt.Println(bird.Move())
			case "speak":
				fmt.Println(bird.Speak())
			}
		case "snake":
			switch data{
			case "eat":
				fmt.Println(snake.Eat())
			case "move":
				fmt.Println(snake.Move())
			case "speak":
				fmt.Println(snake.Speak())
			}
		}
		fmt.Println("")
	}
}
