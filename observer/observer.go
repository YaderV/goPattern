package observer

import "time"

// Topic is a representation of a element that triggers
// an event
type Topic interface {
	register(o Observer)
	broadcast()
}

// Observer is a representation of a elements that is triggered
// when an event happens
type Observer interface {
	trigger()
}

// Item is a element that trigger an event
type Item struct {
	observers []Observer
	Name      string
}

// Register an observer
func (i *Item) Register(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *Item) broadcast() {
	for _, o := range i.observers {
		o.trigger()
	}
}

// Update ...
func (i *Item) Update() {
	i.broadcast()
}

// Email is an example for something that might be triggered
// after an event
type Email struct {
	Name          string
	LastExecution time.Time
}

func (e *Email) trigger() {
	e.LastExecution = time.Now()
}
