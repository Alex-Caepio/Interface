package main

type user struct {
	FIO, Address, Phone string
}

func (u *user) ChangeFio(newFio string) {
	u.FIO = newFio
}

func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress
}

type User interface {
	ChangeFio(newFio string)
	ChangeAddress(newAddress string)
}

func main() {

}
