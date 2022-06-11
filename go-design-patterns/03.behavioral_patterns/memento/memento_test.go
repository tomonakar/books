package memento

import "testing"

func TestCareTaker_Add(t *testing.T) {
	originator := &originator{}
	originator.state = State{Description: "Idle"}

	careTaker := careTaker{}

	mem := originator.NewMemento()
	if mem.state.Description != "Idle" {
		t.Errorf("Expected %s, got %s", "Idle", mem.state.Description)
	}

	currentLen := len(careTaker.mementoList)
	careTaker.Add(mem)

	if len(careTaker.mementoList) != currentLen+1 {
		t.Error("No new elements were added on the lists")
	}
}

func TestCareTaker_Memento(t *testing.T) {
	originator := originator{}
	careTaker := careTaker{}

	originator.state = State{"Idle"}
	careTaker.Add(originator.NewMemento())

	mem, err := careTaker.Memento(0)
	if err != nil {
		t.Fatal(err)
	}

	if mem.state.Description != "Idle" {
		t.Error("Unexpected state")
	}

	mem, err = careTaker.Memento(-1)
	if err == nil {
		t.Fatalf("An error is expected when asking for a negative number no error was found")
	}
}

func TestOriginator_ExtractAndStoreState(t *testing.T) {
	originator := originator{state: State{"Idle"}}
	idleMemento := originator.NewMemento()

	originator.state = State{"Working"}

	originator.ExtractAndStoreState(idleMemento)
	if originator.state.Description != "Idle" {
		t.Error("Unexpected state found")
	}
}
