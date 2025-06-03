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
