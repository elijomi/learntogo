package person

import "strconv"

func NewPerson(name string, age int) *Person {
	if age < 0 { 
		age = 0
	}
	return &Person{
		Name: name, 
		Age: age, 
		Friends: []*Person{},
	}
}

func (p Person) Greet() string {
	return "Hello my name is " + p.Name + " and I am " + strconv.Itoa(p.Age) + " years old."
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// 7: Friends (pointers + slices)
func (p *Person) AddFriend(friend *Person) bool {
	if friend == nil || p == friend {		// req dont allow nil, dont allow self as friend
		return false
	}

	for _, f := range p.Friends {
		if f == friend {	
			return false	// already a friend
		}
	}

	p.Friends = append(p.Friends, friend)
	return true	
}

func (p Person) FriendNames() []string { 
	names := make([]string, len(p.Friends))
	for i, f := range p.Friends {
		names[i] = f.Name
	}

	return names
}
