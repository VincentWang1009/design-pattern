package iterator

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {

	user1 := &User{"tom", 15}
	user2 := &User{"jack", 16}
	user3 := &User{"alice", 18}

	userCollec := &UserCollection{
		users: []*User{user1, user2, user3},
	}

	iter := userCollec.createIterator()

	for iter.hasNext() {
		user := iter.getNext()
		fmt.Printf("user is %+v", user)
	}

}
