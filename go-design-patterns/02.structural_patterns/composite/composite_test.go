package composition

import "testing"

func TestAthlete_Train(t *testing.T) {
	athlete := Athlete{}
	athlete.Train()
}

func TestCompositeSwimmerA_Train(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: &localSwim,
	}

	swimmer.MyAthlete.Train()
	(*swimmer.MySwim)()
}


func TestShark_Eat(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}

	// fish.Animal.Eat() としなくても、embedしてるので直接呼び出せる
	fish.Eat()
	fish.Swim()
}

func TestTree(t *testing.T) {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
		},
		Left: &Tree{4, nil, nil},
	}

	println(root.Right.Right.LeafValue)
}

func TestSon_GetParentField_Composition(t *testing.T) {
	son2 := Son2{}
	GetParentFieldComposition(son2.P)
}

func TestSon_GetParentField_embedded(t *testing.T) {
	son := Son{}
	GetParentFieldEmbedded(&son.Parent)
}

