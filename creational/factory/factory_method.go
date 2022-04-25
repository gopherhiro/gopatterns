package factory

import "gopatterns/pkg/static"

// Factory Method Pattern
type CourseFactory interface {
	New() Course
}

// ChineseFactory
type ChineseFactory struct {
}

func (m *ChineseFactory) New() Course {
	return &Chinese{}
}

// EnglishFactory
type EnglishFactory struct {
}

func (m *EnglishFactory) New() Course {
	return &English{}
}

// 为工厂类再创建一个简单工厂，也就是工厂的工厂，用来创建工厂类对象。
func NewInstance(ID int) CourseFactory {
	if ID == static.Chinese {
		return &ChineseFactory{}
	}
	if ID == static.English {
		return &EnglishFactory{}
	}
	return nil
}
