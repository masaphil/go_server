package tester

type TesterID string
type GoogleMail string
type Tester struct {
	Id   TesterID
	Mail GoogleMail
}

func New(id TesterID, gm GoogleMail) *Tester {
	return &Tester{
		Id:   id,
		Mail: gm,
	}
}
