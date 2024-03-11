package main

import (
	"fmt"
)

type Zookeeper struct {
	Name string
}

type Cage struct {
	Animal *Animal
}

type Animal struct {
	Name   string
	InCage bool
}

func (keeper Zookeeper) addAnimal(cage *Cage, animal *Animal) {
	if cage.Animal == nil {
		cage.Animal = animal
		animal.InCage = true
		fmt.Printf("Animal %s has been placed to the cage by zookeeper %s\n", animal.Name, keeper.Name)
	}
}

func (keeper Zookeeper) getCage() *Cage {
	return &Cage{}
}

func main() {
	keeper := Zookeeper{
		Name: "Alex",
	}

	dog := Animal{
		Name: "Sharik",
	}

	cat := Animal{
		Name: "Murzik",
	}

	snake := Animal{
		Name: "Sam",
	}

	rat := Animal{
		Name: "Phil",
	}

	cow := Animal{
		Name: "Zorka",
	}

	cage1 := keeper.getCage()
	cage2 := keeper.getCage()
	cage3 := keeper.getCage()
	cage4 := keeper.getCage()
	cage5 := keeper.getCage()

	keeper.addAnimal(cage1, &dog)
	keeper.addAnimal(cage2, &cat)
	keeper.addAnimal(cage3, &snake)
	keeper.addAnimal(cage4, &rat)
	keeper.addAnimal(cage5, &cow)

	fmt.Printf("Is animal %s in cage: %t\n", cat.Name, cat.InCage)
	fmt.Printf("Is animal %s in cage: %t\n", cow.Name, cow.InCage)
}
