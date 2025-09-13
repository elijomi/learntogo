package person

// 6. Define a struct Person with fields Name (string) and Age (int).
type Person struct {
	Name    string
	Age     int
	Friends []*Person
}
