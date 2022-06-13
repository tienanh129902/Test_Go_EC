package main

// import "fmt"

type Reptile interface {
	Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
	CreateReptile ReptileCreator
	hatchCount    int
}

func (egg *ReptileEgg) Hatch() Reptile {
	if egg.hatchCount > 1 {
		return nil
	}
	egg.hatchCount++
	return egg.CreateReptile()
}

type FireDragon struct {
	Reptile
}

func (fireDragon *FireDragon) Lay() {

}
