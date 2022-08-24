package iterator

type UserCollection struct {
	users []*User
}

func (c *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: c.users,
	}
}
