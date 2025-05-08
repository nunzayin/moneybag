package main

type Bags map[string]uint
type Tags map[string][]string
type State struct {
	bags Bags
	tags Tags
}
type Transfer struct {
	source string
	destination string
	amount int
	comment string
}
type Tagging struct {
	tag string
	names []string
}
type Modifier interface {
	eval(state State)
}
type Book []Modifier

func main() {

}
