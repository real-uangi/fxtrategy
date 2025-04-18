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

### Constructors
```go
func newA() fxtrategy.Strategy[Fool] {
	return fxtrategy.Strategy[Fool]{
		NS: fxtrategy.NamedStrategy[Fool]{
			Name: "a",
			Item: &FoolA{},
		},
	}
}
```

### Constructors with dependencies
```go
func newA(dependencyA any, dependencyB any) fxtrategy.Strategy[Fool] {
	return fxtrategy.Strategy[Fool]{
		NS: fxtrategy.NamedStrategy[Fool]{
			Name: "a",
			Item: &FoolA{
				depA: dependencyA,
				depB: dependencyB,
            },
		},
	}
}
```

### Usage
```go
fx.New(
    fx.Provide(newA, newB, newC, newD),
    fx.Provide(fxtrategy.NewContext[Fool]),
    fx.Provide(fxtrategy.NewContext[Genius]),

    fx.Invoke(func(ctx *fxtrategy.Context[Fool]) {
        a, _ := ctx.Get("a")
        a.Speak() // Output: i'm a
        b, _ := ctx.Get("b")
        b.Speak() // Output: i'm b
    }),

    fx.Invoke(func(ctx *fxtrategy.Context[Genius]) {
        c, _ := ctx.Get("c")
        c.Speak() // Output: i'm c
        d, _ := ctx.Get("d")
        d.Speak() // Output: i'm d
    }),
)
```
