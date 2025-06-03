package repo

import (
	"github.com/mathbdw/retranslator/internal/entity"
)

type EventRepo interface {
	Lock(n uint64) ([]entity.ProductEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []entity.ProductEvent) error
	Remove(eventIDs []uint64) error
}

type Cache struct {
	prodsEvent *[]entity.ProductEvent
}

func (c *Cache) Lock(n uint64) ([]entity.ProductEvent, error) {
	var max, i uint64
	if uint64(len(*c.prodsEvent)) > n {
		max = n
	} else {
		max = uint64(len(*c.prodsEvent))
	}
	outEvent := make([]entity.ProductEvent, max)
	for i = 0; i < max; i++ {
		if (*c.prodsEvent)[i].Status != entity.Processe && (*c.prodsEvent)[i].Status != entity.Deferred && (*c.prodsEvent)[i].Status != entity.Processed {
			switch (*c.prodsEvent)[i].Type {
			case entity.Created:
				(*c.prodsEvent)[i].Status = entity.Processe
				outEvent = append(outEvent, (*c.prodsEvent)[i])
			}
		}
	}

	return outEvent, nil
}

func (c *Cache) Unlock(eventIDs []uint64) error {
	for i, _ := range eventIDs {
		switch (*c.prodsEvent)[i].Type {
		case entity.Created:
			if (*c.prodsEvent)[i].Status != entity.Processe {
				(*c.prodsEvent)[i].Status = entity.Processed
			} else {
				(*c.prodsEvent)[i].Status = entity.Deferred
			}
		}
	}

	return nil
}

func (c *Cache) Add(event []entity.ProductEvent) error {
	for i, _ := range event {
		event[i].Status = entity.Processe
		*c.prodsEvent = append(*c.prodsEvent, event[i])
	}

	return nil
}

func (c *Cache) Remove(eventIDs []uint64) error {
	var id uint64
	for _, id = range eventIDs {
		switch {
		case id == (*c.prodsEvent)[0].ID:
			*c.prodsEvent = (*c.prodsEvent)[1:]
		case id == (*c.prodsEvent)[len(*c.prodsEvent)].ID:
			*c.prodsEvent = (*c.prodsEvent)[:len(*c.prodsEvent)]
		default:
			for i := 0; i < len(*c.prodsEvent); i++ {
				if (*c.prodsEvent)[i].ID == id {
					*c.prodsEvent = append((*c.prodsEvent)[:i], (*c.prodsEvent)[i+1:]...)
					break
				}
			}
		}
	}
	return nil
}
