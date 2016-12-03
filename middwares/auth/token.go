package auth

type Token string

func (t Token) isValid() bool{
	return true
}

func (t Token) Encode() *Token{
	return &t
}

func (t Token) Decode() (*Token, error){
	return &t, nil
}
