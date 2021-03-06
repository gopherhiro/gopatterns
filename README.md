# 设计模式学习之旅

## go design patterns

1.[理论学习](https://refactoring.guru/design-patterns)

2.[模型举例](https://golangbyexample.com/all-design-patterns-golang/)

3.[运用实战](https://time.geekbang.org/column/intro/100039001)

## 设计模式分类

经典的设计模式分为三类：创建型、结构型、行为型。

- 创建型设计模式：主要解决「对象的创建」问题。
- 结构型设计模式：主要解决「对象的组合或组装」问题。
- 行为型设计模式：主要解决「对象之间的交互」问题。

**设计模式干的事情就是解耦：**

- 创建型模式是将创建和使用代码解耦。
- 结构型模式是将不同功能代码解耦。
- 行为型模式是将不同的行为代码解耦。

借助设计模式，我们能更好的组织代码结构，让其满足开闭原则、高内聚松耦合等特性，以此来应对和控制代码的复杂性，提高代码可扩展性。

## 创建型模式

- 单例模式
- 工厂模式
- 建造者模式

创建型模式主要解决对象的创建问题，封装复杂的创建过程，解耦对象的创建代码和使用代码。

- 单例模式用来创建全局唯一的对象。
- 工厂模式用来创建不同但是相关类型的对象。
- 建造者模式用来创建复杂对象。

## 结构型模式

- 代理模式
- 桥接模式
- 装饰器模式
- 适配器模式
- 门面模式

## 行为型模式

- 观察者模式
- 模板模式
- 策略模式
- 职责链模式
- 状态模式
- 迭代器模式

## 很Nice的想法

- 设计原则和思想其实比设计模式更加普适和重要，掌握了代码的设计原则和思想，我们甚至可以自己创造出来新的设计模式。
- 之所以将某个代码块剥离出来，独立为函数或者类，原因是这个代码块的逻辑过于复杂，剥离之后能让代码更加清晰，更加可读、可维护。
- 业务类最好职责更加单一，只聚焦业务处理。
- 源码之下无秘密。
- 很多设计模式都是试图将庞大的类拆分成更细小的类，然后再通过某种更合理的结构组装在一起。
- 框架的作用：隐藏实现细节，降低开发难度，做到代码复用，解耦业务与非业务代码，聚焦业务开发。
- 在平时的业务开发中，我们要善于抽象非业务的、可复用的功能，并积极地把它们实现成通用的框架。
- 模板模式基于继承来实现，回调基于组合来实现。