package factory

import "gopatterns/pkg/static"

// Factory Method Pattern
type CourseFactory interface {
	NewCourse() Course
}

// ChineseFactory
type ChineseFactory struct {
}

func (m *ChineseFactory) NewCourse() Course {
	// 此处封装复杂的对象创建过程
	return newChinese()
}

// EnglishFactory
type EnglishFactory struct {
}

func (m *EnglishFactory) NewCourse() Course {
	// 此处封装复杂的对象创建过程
	return newEnglish()
}

// 为工厂类再创建一个简单工厂，也就是工厂的工厂，用来创建工厂类对象。
func NewFactory(ID int) CourseFactory {
	if ID == static.Chinese {
		return &ChineseFactory{}
	}
	if ID == static.English {
		return &EnglishFactory{}
	}
	return nil
}
