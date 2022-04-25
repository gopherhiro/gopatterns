package factory

import (
	"gopatterns/pkg/static"
)

// Chinese
type Chinese struct {
}

func (m *Chinese) GetID() int64 {
	return static.Chinese
}

func (m *Chinese) GetName() string {
	return "Chinese"
}

// English
type English struct {
}

func (m *English) GetID() int64 {
	return static.English
}

func (m *English) GetName() string {
	return "English"
}

// 简单工厂（Simple Factory）- 函数式使用方式
func New(ID int64) Course {
	if ID == static.Chinese {
		return &Chinese{}
	}

	if ID == static.English {
		return &English{}
	}

	return nil
}

// 简单工厂（Simple Factory）- 缓存式使用方式
var CourseMap = map[int64]Course{
	static.Chinese: &Chinese{},
	static.English: &English{},
}

func NewFromCache(ID int64) Course {
	// ok-idiom 模式
	if instance, ok := CourseMap[ID]; ok {
		return instance
	}
	return nil
}
