package material

import (
	"time"
)

const (
	TypeRaw          = "raw"
	TypeSemiFinished = "semi-finished"
	TypeFinished     = "finished"
)

type Material struct {
	ID        int
	Name      string
	Type      string
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
