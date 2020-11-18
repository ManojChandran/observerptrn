package observerptrn

import (
	"fmt"
	"testing"
)

type TestObserver struct {
	ID      int
	Message interface{}
}

func (p *TestObserver) Notify(m interface{}) {
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

	t.Run("Unsubscribe", func(t *testing.T) {
		publisher.Unsubscribe(testObserver2)

		if len(publisher.SubscribersList) != 2 {
			t.Errorf("The size of the observer list is not the "+
				"expected. 3 != %d\n", len(publisher.SubscribersList))
		}

		for _, observer := range publisher.SubscribersList {
			testObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
			}

			if testObserver.ID == 2 {
				t.Fail()
			}
		}
	})

	t.Run("Notify", func(t *testing.T) {
		for _, observer := range publisher.SubscribersList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			if printObserver.Message != "" {
				t.Errorf("The observer's Message field weren't "+"  empty: %s\n", printObserver.Message)
			}
		}
	})
}
