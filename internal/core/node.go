package core

import "context"

type Node struct {
	ID int64 `db:"id" json:"id"`
}

type NodeStore interface {
	Find(context.Context, int64) (*Node, error)
	Create(context.Context, *Node) error
	Update(context.Context, *Node) error
	Delete(context.Context, *Node) error
}
