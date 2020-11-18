# Observer Design Pattern
This is simple GO package to implement Observer pattern. Observer pattern, also known as publish/subscriber or publish/listener. It is useful when we need to trigger many actions to be performed after an event.

# Implementation details
Main parts of the Observer pattern implementations are as mentioned below.

#### A Observer interface, for the subscribers to implement.
```go
type Observer interface { 
  Notify(string) 
} 
```
#### Struct to hold the Observer list:
```go
type Publisher struct { 
  ObserversList []Observer 
} 
```
#### Publisher has three main methods:
1) A Method to add new subscribers to the publisher
2) A Method to remove subscribers from the publisher
3) A Method to accept message and publish 
```go 
func (s *Publisher) Subscribe(o Observer) {} 
 
func (s *Publisher) Unsubscribe(o Observer) {} 
 
func (s *Publisher) Publish(m string) {} 
```

# Get Package
```
$ go get github.com/ManojChandran/observerptrn
```

# Example code to implement 

Below code provide sample implementation of observer pattern

```go
package main

import (
	"fmt"

	"github.com/ManojChandran/observerptrn"
)

//
type House struct {
	Observer observerptrn.Publisher
	Name     string
}

//
func NewHouse(name string) *House {
	return &House{
		Observer: observerptrn.Publisher{},
		Name:     name,
	}
}

func (h *House) catchFire() {
	h.Observer.Publish(h.Name)
}

// Dummy fire service 
type FireService struct{}

// Notify interface implementation
func (d *FireService) Notify(data string) {
	fmt.Printf("\n A Fire service has been called for %s",
		data)
}

// Dummy ambulance service
type AmbulanceService struct{}

// Notify interface implementation
func (d *AmbulanceService) Notify(data string) {
	fmt.Printf("\n A Ambulance service has been called for %s",
		data)
}

func main() {
	myHouse := NewHouse("Meparambu")
	fire := &FireService{}
	ambulance := &AmbulanceService{}

	// Subscribing the service
	myHouse.Observer.Subscribe(fire)
	myHouse.Observer.Subscribe(ambulance)

	// Event
	myHouse.catchFire()

	// Un subscribing the service
	myHouse.Observer.Unsubscribe(ambulance)
	
	// Event
	myHouse.catchFire()
}
```