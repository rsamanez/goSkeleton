package model

type Partner struct {
	idpartner int
	code string
}

func (p Partner) GetId() int {
	return p.idpartner
}

func (p Partner) GetCode() string {
	return p.code
}

func (p *Partner) SetId(v int) {
	p.idpartner = v
}

func (p *Partner) SetCode(v string) {
	p.code = v
}
