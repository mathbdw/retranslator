package sender

import (
	"github.com/mathbdw/retranslator/internal/entity"
)

type EventSender interface {
	Send(product *entity.ProductEvent) error
}
