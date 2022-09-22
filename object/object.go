package object

import "fmt"

type ObjectType string

const (
	INTEGAR_OBJ = "INTEGAR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGAR_OBJ
}
