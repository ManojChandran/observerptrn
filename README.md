# Observer Design Pattern
This is simple GO package to implement Observer pattern. Observer pattern, also known as publish/subscriber or publish/listener. It is useful when we need to trigger many actions to be performed after an event.

Only interface we need to use the package is:
```go
type Observer interface { 
  Notify(string) 
} 
```

We just need a structure called Publisher with three methods:
```go
type Publisher struct { 
  ObserversList []Observer 
} 
 
func (s *Publisher) Subscribe(o Observer) {} 
 
func (s *Publisher) Unsubscribe(o Observer) {} 
 
func (s *Publisher) Publish(m string) {} 
```

# Get Package
```
$ go get github.com/ManojChandran/observerptrn
```

# Sample code to implement 

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

//
type FireService struct{}

//
func (d *FireService) Notify(data string) {
	fmt.Printf("\n A Fire service has been called for %s",
		data)
}

//
type AmbulanceService struct{}

//
func (d *AmbulanceService) Notify(data string) {
	fmt.Printf("\n A Ambulance service has been called for %s",
		data)
}

func main() {
	myHouse := NewHouse("Meparambu")
	fire := &FireService{}
	ambulance := &AmbulanceService{}

	myHouse.Observer.Subscribe(fire)
	myHouse.Observer.Subscribe(ambulance)
	fmt.Println("hello")

	myHouse.catchFire()
	myHouse.Observer.Unsubscribe(ambulance)
	myHouse.catchFire()
}
```