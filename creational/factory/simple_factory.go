package factory

import (
	"gopatterns/pkg/static"
)

func newChinese() *Chinese {
	return &Chinese{}
}

func newEnglish() *English {
	return &English{}
}

// NewCourse 简单工厂（Simple Factory）- 函数式使用方式
func NewCourse(ID int64) Course {
	if ID == static.Chinese {
		return newChinese()
	}

	if ID == static.English {
		return newEnglish()
	}

	return nil
}