package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func (a *Animal) Initialize(food, locomotion, noise string) Animal {
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
	return *a
}

func initData() map[string]Animal {
	var a Animal
	m := make(map[string]Animal)
	m["cow"] = a.Initialize("grass", "walk", "moo")
	m["bird"] = a.Initialize("worms", "fly", "peep")
	m["snake"] = a.Initialize("mice", "slither", "hiss")
	return m
}

func getString(prompt string) (string, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", errors.New("Buffer empty")
}

func main() {
	m := initData()

	name, _ := getString("Enter an animal: ")
	a := m[name]
	println(a.Eat())
}
