package observerptrn

// Subscriber needs to implement this function
type Subscriber interface {
	Notify(string)
}

// Publisher hold the list of subscriber
type Publisher struct {
	SubscribersList []Subscriber
}

// Subscribe method to subscribe for an event
func (s *Publisher) Subscribe(o Subscriber) {
	s.SubscribersList = append(s.SubscribersList, o)
}

// Unsubscribe method to to un subscribe from an event
func (s *Publisher) Unsubscribe(o Subscriber) {
	var indexToRemove int
	for i, Subscriber := range s.SubscribersList {
		if Subscriber == o {
			indexToRemove = i
			break
		}
	}
	s.SubscribersList = append(s.SubscribersList[:indexToRemove], s.SubscribersList[indexToRemove+1:]...)
}

// Publish method to publish an event or message
func (s *Publisher) Publish(m string) {
	for _, Subscriber := range s.SubscribersList {
		Subscriber.Notify(m)
	}
}
