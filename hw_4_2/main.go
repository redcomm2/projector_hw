package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

type Entity struct {
	Id int
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func main() {

	var entitySlice []Entity
	for i := 0; i < 10; i++ {
		entitySlice = append(entitySlice, Entity{Id: randRange(1, 11)})
	}

	uniqueSlice := uniqueEntities(entitySlice)

	sort.Slice(uniqueSlice, func(i, j int) bool {
		return uniqueSlice[i].Id < uniqueSlice[j].Id
	})

	fmt.Print(uniqueSlice)
}

func uniqueEntities(entities []Entity) []Entity {
	var unique []Entity
	for _, entity := range entities {
		if !contains(unique, entity) {
			unique = append(unique, entity)
		}
	}
	return unique
}

func contains(slice []Entity, entity Entity) bool {
	for _, elem := range slice {
		if elem.Id == entity.Id {
			return true
		}
	}
	return false
}
