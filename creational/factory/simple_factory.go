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

func newChinese() *Chinese {
	return &Chinese{}
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

func newEnglish() *English {
	return &English{}
}

// 简单工厂（Simple Factory）- 函数式使用方式
func NewCourse(ID int64) Course {
	if ID == static.Chinese {
		return newChinese()
	}

	if ID == static.English {
		return newEnglish()
	}

	return nil
}