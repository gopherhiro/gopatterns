package factory

import "gopatterns/pkg/static"

func newChineseFactory() CourseFactory {
	return &ChineseFactory{}
}

func newEnglishFactory() CourseFactory {
	return &EnglishFactory{}
}

// 为工厂类再创建一个简单工厂，也就是工厂的工厂，用来创建工厂类对象。
func NewFactory(ID int) CourseFactory {
	if ID == static.Chinese {
		return newChineseFactory()
	}
	if ID == static.English {
		return newEnglishFactory()
	}
	return nil
}
