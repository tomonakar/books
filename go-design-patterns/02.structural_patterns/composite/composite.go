package composition

import "fmt"

type Athlete struct {}

func (a *Athlete) Train() {
	fmt.Println("Athlete trains")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim *func()
}

func Swim() {
	fmt.Println("Swimming")
}


type Animal struct {}

func (r *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImplementor struct {}

func (s *SwimmerImplementor) Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmerB struct {
	Athlete
	Swimmer
}


// ------- BinaryTree -----
type Tree struct {
	LeafValue int
	Right *Tree
	Left *Tree
}


// -------- Composition vs Inheritance -----

type Parent struct {
	SomeField int
}

// embedding
type Son struct {
	Parent
}

func GetParentFieldEmbedded(p *Parent) int {
	return p.SomeField
}

// composition
type Son2 struct {
	P Parent
}

func GetParentFieldComposition(p Parent) int {
	return p.SomeField

}