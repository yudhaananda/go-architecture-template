package formatter

type Auth[T comparable] struct {
	Data  T      `json:"profile"`
	Token string `json:"token"`
}

func (f *Auth[T]) Format(data T, token string) {
	f.Data = data
	f.Token = "Bearer " + token
}
