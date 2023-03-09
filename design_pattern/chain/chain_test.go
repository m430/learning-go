package chain

import "testing"

func TestChain(t *testing.T) {
	cashier := &CashHandler{}

	medical := &MedicalHandler{}
	medical.handler = cashier

	doctor := &DoctorHandler{}
	doctor.handler = medical

	reception := &ReceptionHandler{}
	reception.handler = doctor

	p := &Patient{name: "andy"}

	reception.Handle(p)
}
