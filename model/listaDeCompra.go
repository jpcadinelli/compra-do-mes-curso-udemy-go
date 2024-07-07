package model

import (
	"time"
)

type ListaDeCompra struct {
	Data    time.Time
	Lista   []Item
	Mercado Mercado
}
