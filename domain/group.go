package domain

type Groups []Group
type Group struct {
	Name  string
	Items Items
}

type Items []Item
type Item struct {
	Name     Name
	Contents string
}
