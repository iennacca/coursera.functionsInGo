package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
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

func parseInput(cmds string) (string, string) {
	cmd := strings.Fields(cmds)
	return strings.ToLower(cmd[0]), strings.Title(strings.ToLower(cmd[1]))
}

func main() {
	m := initData()

	fmt.Println("At the prompt, enter a string consisting of one of the following animals \n    (cow, bird, snake)\n followed by one of these methods \n    (eat, move, speak):")

	for true {
		cmds, _ := getString("> ")
		animalName, methodName := parseInput(cmds)

		a := m[animalName]
		result := reflect.ValueOf(&a).MethodByName(methodName).Call([]reflect.Value{})
		fmt.Printf("%s.%s: %s\n", animalName, methodName, result)
	}
}
