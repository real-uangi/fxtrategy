# Fxtrategy

Eng | [中文](README.zh.md)

## 📚 Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [GetStart](#get-start)
- [Example](#example)

---

## Introduction

`fxtrategy` is a strategy injection utility based on [Uber FX](https://github.com/uber-go/fx), leveraging generics and named strategy patterns to enable flexible, modular registration and dynamic retrieval of strategies in Go applications.

---

## Features

- ✅ Support defining strategy behavior via interfaces  
- ✅ Strongly-typed context via Go generics  
- ✅ Name-based strategy registration and retrieval  
- ✅ Supports multiple contexts (e.g., `Fool` and `Genius`)  
- ✅ Fully integrated with Uber FX  

---

## Get Start
```shell
go get github.com/real-uangi/fxtrategy@latest
```

---

## Example

### Usage
```go
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
```

### TestSamples
```go
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
```
