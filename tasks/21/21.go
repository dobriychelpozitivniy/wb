package main

// Реализовать паттерн «адаптер» на любом примере.

func main() {
	s := Struct1{}
	a := Adapter{s}
	Check(a)
}

func Check(i Interface) {
	i.CommitChanges()
}

type Interface interface {
	SendRequest() string
	CommitChanges() string
}

type Adapter struct {
	Struct1
}

func (a Adapter) SendRequest() string {
	return a.Send()
}

func (a Adapter) CommitChanges() string {
	return a.Commit()
}

type Struct1 struct {
}

func (s *Struct1) Send() string {
	return "Send"
}

func (s *Struct1) Commit() string {
	return "Commit"
}
