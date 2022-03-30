package utils

import (
	"fmt"
)

type ComponentKind string

const (
	ComponentTiDB ComponentKind = "tidb"
	ComponentTiKV ComponentKind = "tikv"
)

type Component struct {
	Kind ComponentKind
	Addr string // host:port
}

func (c Component) String() string {
	return fmt.Sprintf("%s://%s", c.Kind, c.Addr)
}
