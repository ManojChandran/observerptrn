package observerptrn

import (
	"fmt"
	"testing"
)

type TestObserver struct {
	ID      int
	Message string
}

func (p *TestObserver) Notify(m string) {
	fmt.Printf("Observer %d: message '%s' received \n", p.ID, m)
	p.Message = m
}
func TestObserverptrn(t *testing.T) {
	// Init

	// Execution

	// Validation

	testObserver1 := &TestObserver{1, ""}
	testObserver2 := &TestObserver{2, ""}
	testObserver3 := &TestObserver{3, ""}
	publisher := Publisher{}

	t.Run("Subscribe", func(t *testing.T) {
		publisher.Subscribe(testObserver1)
		publisher.Subscribe(testObserver2)
		publisher.Subscribe(testObserver3)

		if len(publisher.SubscribersList) != 3 {
			t.Fail()
		}
	})
}
