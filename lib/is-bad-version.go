package lib

type BadVersionAPI struct {
	bad int
}

func NewBadVersionAPI(bad int) *BadVersionAPI {
	return &BadVersionAPI{
		bad: bad,
	}
}
