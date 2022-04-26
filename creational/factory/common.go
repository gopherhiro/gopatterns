package factory

import "gopatterns/pkg/static"

type Course interface {
	GetID() int64
	GetName() string
}

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
