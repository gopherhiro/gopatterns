package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	user1 := &user{
		name: "a",
		age:  1,
	}
	user2 := &user{
		name: "b",
		age:  2,
	}
	uc := &UserCollection{
		users: []*user{user1, user2},
	}

	iterator := uc.newIterator()
	for iterator.hasMore() {
		u := iterator.current()
		fmt.Printf("user:%+v\n", u)
		iterator.next()
	}
}
