package entity

import "fmt"

// Product - entity
type Product struct {
	ID          uint64
	Title       string `json:"title"`
	Description string `json:"description"`
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type ProductEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Product
}

func (p *Product) String() string {
	return fmt.Sprintf("%s - %s", p.Title, p.Description)
}
