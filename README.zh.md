# Fxtrategy

[Eng](README.md) | 中文

## 📚 Table of Contents

- [简介](#简介)
- [特性](#特性)
- [GetStart](#get-start)
- [示例](#示例)

---

## 简介

`fxtrategy` 是一个基于 [Uber FX](https://github.com/uber-go/fx) 的策略注入工具，使用泛型与命名策略模式，为 Go 应用程序提供灵活、模块化的策略注册与动态获取能力。

---

## 特性

- ✅ 支持通过接口定义策略行为
- ✅ 基于泛型实现强类型上下文 Context
- ✅ 可按名称注册与获取策略实例
- ✅ 支持多个上下文类型共存（如：`Fool` 和 `Genius`）
- ✅ 与 Uber FX 完美集成

---

## Get Start
```shell
go get github.com/real-uangi/fxtrategy@latest
```

---

## 示例

### 用法
```go
	fx.New(
		fx.Options(fxtrategy.ProvideItem[Fool](newA, "a", "test")),
		fx.Options(fxtrategy.ProvideItem[Fool](newB, "b", "test")),
		fx.Options(fxtrategy.ProvideContext[Fool]("test")),
		fx.Invoke(func(ctx *fxtrategy.Context[Fool]) {
			a, _ := ctx.Get("a")
			a.FoolSpeak()
			b, _ := ctx.Get("b")
			b.FoolSpeak()
		}),
	)
```

### 测试定义
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
