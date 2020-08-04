package server

type Dictionary map[string]*Object

type Database struct {
	Dict   Dictionary
	Expire Dictionary
}
