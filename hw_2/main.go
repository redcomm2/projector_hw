package main

import "fmt"

type Zookeeper struct {
	Name string
}

type Cage struct {
	isOccupied bool
	Animal     *Animal
}

type Animal struct {
	Name   string
	InCage bool
}

func (keeper Zookeeper) addAnimal(cage *Cage, animal *Animal) error {
	if cage.Animal == nil {
		cage.Animal = animal
		cage.isOccupied = true
		animal.InCage = true
		fmt.Printf("Animal %s has been placed to the cage by zookeeper %s\n", animal.Name, keeper.Name)
	}
	return nil
}

func (keeper Zookeeper) getCage() *Cage {
	return &Cage{
		isOccupied: false,
	}
}

func main() {
	keeper := Zookeeper{
		Name: "Alex",
	}

	dog := Animal{
		Name:   "Sharik",
		InCage: false,
	}

	cat := Animal{
		Name:   "Murzik",
		InCage: false,
	}

	snake := Animal{
		Name:   "Sam",
		InCage: false,
	}

	rat := Animal{
		Name:   "Phil",
		InCage: false,
	}

	cow := Animal{
		Name:   "Zorka",
		InCage: false,
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
}
