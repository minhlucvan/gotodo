package models

type Model interface{
	Find(string) (*Model, error)
	FindAll(string) ([]*Model, error)
}