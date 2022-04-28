package iterator

// Collection
// 集合
type Collection interface {
	newIterator() Iterator
}

// UserCollection
// 具体集合
type user struct {
	name string
	age  int
}

type UserCollection struct {
	users []*user
}

// 一个集合可以对应多个迭代器
// 可以结合工厂模式来返回不同的迭代器
func (u *UserCollection) newIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

// Iterator
// 迭代器
type Iterator interface {
	hasMore() bool
	next() *user
}

// UserIterator
// 具体迭代器
type UserIterator struct {
	index int
	users []*user
}

func (u *UserIterator) hasMore() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) next() *user {
	if u.hasMore() {
		elem := u.users[u.index]
		u.index++
		return elem
	}
	return nil
}
