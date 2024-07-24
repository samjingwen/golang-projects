package qns3_6

import (
	"testing"
)

func TestAnimalShelter(t *testing.T) {
	queue := AnimalShelter{}
	queue.enqueue(&Cat{})
	queue.enqueue(&Cat{})
	queue.enqueue(&Dog{})
	queue.enqueue(&Dog{})

	actual, _ := queue.dequeAny()
	if _, ok := actual.(*Cat); !ok {
		t.Errorf("actual = %T, expected = %T", actual, &Cat{})
	}

	actual, _ = queue.dequeDog()
	if _, ok := actual.(*Dog); !ok {
		t.Errorf("actual = %T, expected = %T", actual, &Dog{})
	}

	actual, _ = queue.dequeAny()
	if _, ok := actual.(*Cat); !ok {
		t.Errorf("actual = %T, expected = %T", actual, &Cat{})
	}

	actual, _ = queue.dequeAny()
	if _, ok := actual.(*Dog); !ok {
		t.Errorf("actual = %T, expected = %T", actual, &Dog{})
	}
}
