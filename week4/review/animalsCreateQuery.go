package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type  Cow struct {
	food, move,speak string
}

type  Bird struct {
	food, move,speak string
}

type  Snake struct {
	food, move, speak string
}

func (c *Cow) Eat(){
	fmt.Println(c.food)
}
func (c *Bird) Eat(){
	fmt.Println(c.food)
}
func (c *Snake) Eat(){
	fmt.Println(c.food)
}


func (c *Cow) Move(){
	fmt.Println(c.move)
}
func (c *Bird) Move(){
	fmt.Println(c.move)
}
func (c *Snake) Move(){
	fmt.Println(c.move)
}

func (c *Cow) Speak(){
	fmt.Println(c.speak)
}
func (c *Bird) Speak(){
	fmt.Println(c.speak)
}
func (c *Snake) Speak(){
	fmt.Println(c.speak)
}


func main() {
	cow := Cow {"grass","walk","moo"	}
	bird:= Bird {	"worms",	"fly",	"peep"}
	snake:= Snake {	"mice",	"slither",	"hsss"}


	animals :=make([]map[string]interface{},0)
	for {
		//Get Datas query or create , name of animal , and its act
		query_create , animals_name ,a_name := getInput()
		nameIsCorrect,animalT := check_animal(a_name)
		var tmpAnimal interface{}
		if query_create == "query" {
			nameIsvalid := false
			for _, Value := range animals {
				if v, found := Value[animals_name]; found {
					tmpAnimal = v
					check_animal_behaviour(a_name ,tmpAnimal)
					nameIsvalid = true
					break
				}
			}
			if !nameIsvalid {
				fmt.Println(animals_name,"is invalid animal in list")
			}
		} else if query_create == "newanimal" {
			if nameIsCorrect {
				nameIsDuplicated := true
				for _, Value := range animals {
					if _, found := Value[animals_name]; found {
						nameIsDuplicated = false
						fmt.Println(animals_name ," is exsit in list")
						break
					}
				}
				if nameIsDuplicated {
					switch animalT.(type) {
						case Cow:
							tmpAnimal = cow
						case Bird:
							tmpAnimal = bird
						case Snake:
							tmpAnimal = snake
					}
					animal := make(map[string]interface{})
					animal[animals_name] = tmpAnimal
					animals = append(animals, animal)
					fmt.Println("Created.")
				}
			}else{
				fmt.Println("this type of animal is unkhown!!!")
			}
		} else {
			fmt.Println("command should start with query or newanimal!!!")
		}

	}

}


func getInput() (query_create string,name string,behavior string) {
	reader := bufio.NewReader(os.Stdin)
	isOk:= true

	for isOk {
		fmt.Print(">")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err!",err)
		}else{
			s = strings.TrimSuffix(s, "\n")
			words := strings.Fields(s)
			if len(words) != 3 {
				fmt.Println("error 3 words should entered!!")
			} else {
				isOk = false
				query_create = words[0]
				name = words[1]
				behavior = words[2]
			}
		}
	}

	return query_create , name , behavior
}

func check_animal(a_name string) (nameIsCorrect bool,a_type interface{}){
	nameIsCorrect = true
	if a_name == "cow" 	{
		a_type =Cow{}
	}else if a_name == "bird"{
		a_type =Bird{}
	}else if a_name == "snake" {
		a_type =Snake{}
	} else {
		nameIsCorrect = false
	}
	return nameIsCorrect , a_type
}

func check_animal_behaviour (a_behaviour string,data interface{}){

	f, ok := data.(Cow)
	f2, ok2 := data.(Bird)
	f3, ok3 := data.(Snake)

	if a_behaviour == "eat" {
		if ok {
			f.Eat()
			return
		} else if ok2 {
			f2.Eat()
			return
		} else if ok3 {
			f3.Eat()
			return
		} else {
			fmt.Println("this animal is unkhown!!!")
		}
	} else if a_behaviour == "move" {
		if ok {
			f.Move()
			return
		} else if ok2 {
			f2.Move()
			return
		} else if ok3 {
			f3.Move()
			return
		} else {
			fmt.Println("this animal is unkhown!!!")
		}
	} else if a_behaviour == "speak" {
		if ok {
			f.Speak()
			return
		} else if ok2 {
			f2.Speak()
			return
		} else if ok3 {
			f3.Speak()
			return
		} else {
			fmt.Println("this animal is unkhown!!!")
		}
	} else {
		fmt.Println("this act of animal is unkhown!!!")
	}
}

