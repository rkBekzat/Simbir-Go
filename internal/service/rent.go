package service

type rent struct {
}

func NewRent() Renting {
	return &rent{}
}

func (r *rent) AccessTransport() {

}

func (r *rent) GetById() {

}

func (r *rent) History() {

}

func (r *rent) TransportHistory() {

}

func (r *rent) StartRenting() {

}

func (r *rent) EndRenting() {

}
