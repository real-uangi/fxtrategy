package fxtrategy_test

import (
	"fmt"
	"github.com/real-uangi/fxtrategy"
	"go.uber.org/fx"
	"testing"
)

type Fool interface {
	FoolSpeak()
	fxtrategy.Nameable
}

func newA() *FoolA {
	return &FoolA{}
}

func newB() *FoolB {
	return &FoolB{}
}

type FoolA struct {
}

func (a *FoolA) FoolSpeak() {
	fmt.Println("i'm a")
}

func (a *FoolA) ItemName() string {
	return "a"
}

type FoolB struct {
}

func (b *FoolB) FoolSpeak() {
	fmt.Println("i'm b")
}

func (b *FoolB) ItemName() string {
	return "b"
}

func TestInjection(t *testing.T) {
	fx.New(
		fx.Options(fxtrategy.ProvideItem[Fool](newA, "test")),
		fx.Options(fxtrategy.ProvideItem[Fool](newB, "test")),
		fx.Options(fxtrategy.ProvideContext[Fool]("test")),
		fx.Invoke(func(ctx *fxtrategy.Context[Fool]) {
			a, _ := ctx.Get("a")
			a.FoolSpeak()
			b, _ := ctx.Get("b")
			b.FoolSpeak()
		}),
	)
}

func TestMultiple(t *testing.T) {
	fx.New(
		fx.Options(fxtrategy.ProvideItem[Fool](newA, "test")),
		fx.Options(fxtrategy.ProvideItem[Fool](newB, "test")),
		fx.Options(fxtrategy.ProvideItem[Genius](newC, "test")),
		fx.Options(fxtrategy.ProvideItem[Genius](newD, "test")),
		fx.Options(fxtrategy.ProvideContext[Fool]("test")),
		fx.Options(fxtrategy.ProvideContext[Genius]("test")),
		fx.Invoke(func(ctx *fxtrategy.Context[Fool]) {
			a, _ := ctx.Get("a")
			a.FoolSpeak()
			b, _ := ctx.Get("b")
			b.FoolSpeak()
			_, ok := ctx.Get("c")
			if ok {
				t.Error("c shouldn't be ok")
			}
		}),
		fx.Invoke(func(ctx *fxtrategy.Context[Genius]) {
			c, _ := ctx.Get("c")
			c.GeniusSpeak()
			d, _ := ctx.Get("d")
			d.GeniusSpeak()
		}),
	)
}

type Genius interface {
	GeniusSpeak()
	fxtrategy.Nameable
}

func newC() *GeniusC {
	return &GeniusC{}
}

func newD() *GeniusD {
	return &GeniusD{}
}

type GeniusC struct {
}

func (c *GeniusC) GeniusSpeak() {
	fmt.Println("i'm c")
}

func (c *GeniusC) ItemName() string {
	return "c"
}

type GeniusD struct {
}

func (d *GeniusD) GeniusSpeak() {
	fmt.Println("i'm d")
}

func (d *GeniusD) ItemName() string {
	return "d"
}
