package entity

import "fmt"

// Product - Сущность продукты. Здесь хронятся данные по продуктам нашего сервиса
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

	Processe EventStatus = iota
	Deferred
	Processed
)

// ProductEvent - Сущность событий над Product
type ProductEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Product
}

func (p *Product) String() string {
	return fmt.Sprintf("%s - %s", p.Title, p.Description)
}
