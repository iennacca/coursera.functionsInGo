package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Animal ------------------------------

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func (a *Animal) Initialize(food, locomotion, noise string) Animal {
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
	return *a
}

// AnimalFactory ------------------------------

type AnimalFactory struct {
	types map[string]Animal
	list  map[string]Animal
}

func (af *AnimalFactory) InitTemplate() {
	m := make(map[string]Animal)

	var a Animal
	m["cow"] = a.Initialize("grass", "walk", "moo")
	m["bird"] = a.Initialize("worms", "fly", "peep")
	m["snake"] = a.Initialize("mice", "slither", "hiss")
	af.types = m
	af.list = make(map[string]Animal)
}

func (af *AnimalFactory) CreateAnimal(args []string) (Animal, error) {
	var a Animal

	if len(args) < 2 {
		return a, errors.New("Error: Invalid number of arguments for newanimal")
	}
	animalName := args[0]
	animalType := strings.ToLower(args[1])

	t := af.types
	a, ok := t[animalType]
	if !ok {
		return a, fmt.Errorf("Error: unknown animal type %s", animalType)
	}

	// af.list[animalName] = Animal{a.food, a.locomotion, a.noise}
	af.list[animalName] = a
	return a, nil
}

func (af *AnimalFactory) Query(args []string) error {
	if len(args) < 2 {
		return errors.New("Error: Invalid number of arguments for query")
	}
	animalName := args[0]
	command := strings.Title(strings.ToLower(args[1]))

	a, ok := af.list[animalName]
	if !ok {
		return fmt.Errorf("Error: unknown animal name %s", animalName)
	}

	v := reflect.ValueOf(&a).MethodByName(command)
	if !v.IsValid() {
		return fmt.Errorf("Error: unknown function name %s", command)
	}

	v.Call([]reflect.Value{})
	return nil
}

func (af *AnimalFactory) List() map[string]Animal {
	return af.list
}

// Input ------------------------------

func getString(prompt string) (string, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", errors.New("Buffer empty")
}

func parseInput(cmds string) (string, []string, error) {
	cmd := strings.Fields(cmds)
	if len(cmd) == 0 {
		return "", nil, errors.New("Error: Cannot parse input string")
	}
	return strings.ToLower(cmd[0]), cmd[1:], nil
}

func showPrompt() {
	fmt.Println(
		"At the prompt, enter a string consisting of " +
			"one of the following commands:\n" +
			"    newanimal <name> <animaltype>\n" +
			"    query <name> <command>\n" +
			"where\n" +
			"    <name> is any string\n" +
			"    <animaltype> is either cow, bird, or snake\n" +
			"    <command> is either ear, move or speak\n" +
			"Examples: \n" +
			"    newanimal Bessie cow\n" +
			"    query Bessie eat\n")
}

// Main ------------------------------

func main() {
	var af AnimalFactory
	showPrompt()

	af.InitTemplate()

	for true {
		cmds, _ := getString("> ")
		commandName, args, err := parseInput(cmds)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch commandName {
		case "newanimal":
			a, err := af.CreateAnimal(args)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("[%s (%s)] created it!\n", args[0], a)
		case "query":
			err := af.Query(args)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "list":
			fmt.Println(af.List())
		default:
			fmt.Println("Error: Unrecognized command")
		}
	}
}
