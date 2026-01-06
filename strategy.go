/*
 * Copyright 2026 Uangi. All rights reserved.
 * @author uangi
 * @date 2026/1/6 11:52
 */

// Package fxtrategy

package fxtrategy

import "go.uber.org/fx"

type Nameable interface {
	ItemName() string
}

func ProvideItem[T Nameable](constructor any, group string) fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				constructor,
				fx.As(new(T)),
				fx.ResultTags(`group:"`+group+`"`),
			),
		),
	)
}

// Context to inject and use Items
type Context[T Nameable] struct {
	mapping map[string]T
}

func ProvideContext[T Nameable](group string) fx.Option {
	return fx.Provide(
		fx.Annotate(
			newContext[T],
			fx.ParamTags(`group:"`+group+`"`),
		),
	)
}

func newContext[T Nameable](items []T) *Context[T] {
	itemMap := make(map[string]T, len(items))
	for _, item := range items {
		itemMap[item.ItemName()] = item
	}
	return &Context[T]{
		mapping: itemMap,
	}
}

func (c *Context[T]) Get(name string) (T, bool) {
	t, ok := c.mapping[name]
	return t, ok
}

type TravelFunc[T any] func(string, T) error

// ForEach execute fn on each item
func (c *Context[T]) ForEach(fn TravelFunc[T]) error {
	for k, v := range c.mapping {
		err := fn(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Names returns all item's name
func (c *Context[T]) Names() []string {
	names := make([]string, 0, len(c.mapping))
	for name := range c.mapping {
		names = append(names, name)
	}
	return names
}
