# 状态模式
状态是一种行为设计模式，当一个对象的内部状态发生变化时，可以改变该对象的行为。

即不同的状态对应不同的行为，状态模式可以让状态和行为保持统一。

State is a behavioral design pattern that allows an object to change the behavior when its internal state changes.

![image](https://user-images.githubusercontent.com/65383410/165517171-0a980878-8ad4-4515-acfb-b014bd839e13.png)

状态模式建议为对象的所有可能状态新建一个类，然后将所有状态的对应行为抽取到这些类中。

状态模式是状态机的一种实现方式。

状态机又叫有限状态机，它有 3 个部分组成：
- 状态
- 事件
- 动作

其中，事件也称为转移条件。

事件触发状态的转移及动作的执行。不过，动作不是必须的，也可能只转移状态，不执行任何动作。

## Real-World Analogy
The buttons and switches in your smartphone behave differently depending on the current state of the device:
- When the phone is unlocked, pressing buttons leads to executing various functions.
- When the phone is locked, pressing any button leads to the unlock screen.
- When the phone’s charge is low, pressing any button shows the charging screen.


## 实现

```go

```

## 用法

```go

```

## 应用场景

## 优点

## 缺点
