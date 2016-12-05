package auth

import "github.com/kataras/go-errors"

type Token string

func (t Token) isValid() bool{
	return true
}

func (t Token) Encode() *Token{
	return &t
}

func (t Token) Decode() (*Token, error){
	if t.isValid() == false {
		return nil, errors.New("token not valid.")
	}
	return &t, nil
}
