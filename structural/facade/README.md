# 门面模式
门面模式，也叫外观模式，英文全称是 Facade Design Pattern。

Provide a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.

门面模式为子系统提供一组统一的接口，定义一组高层接口「让子系统更易用」。

## 举例理解
假设有一个系统 A，提供了 a、b、c、d 四个接口。系统 B 完成某个业务功能，需要调用 A 系统的 a、b、d 接口。利用门面模式，我们提供一个包裹 a、b、d 接口调用的门面接口 x，给系统 B 直接使用。

## 实现

```go

```

## 用法

```go

```

## 应用场景
门面模式应用场景比较明确，主要在接口（粒度）设计方面使用。
- 解决易用性问题
- 解决性能问题
- 解决分布式事务问题



## 优点

## 缺点
