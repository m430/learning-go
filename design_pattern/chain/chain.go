package chain

import "fmt"

// 处理者抽象接口
type Handler interface {
	Handle(*Patient)
	Next(Handler, *Patient)
}

type BaseHandler struct {
	handler Handler
}

func (h *BaseHandler) Next(next Handler, p *Patient) {
	if next != nil {
		h.handler.Handle(p)
	}
}

// 前台处理者
type ReceptionHandler struct {
	BaseHandler
}

func (h *ReceptionHandler) Handle(p *Patient) {
	fmt.Printf("The patient %s ask for help done\n", p.name)
	h.Next(h.handler, p)
}

// 医生处理者
type DoctorHandler struct {
	BaseHandler
}

func (h *DoctorHandler) Handle(p *Patient) {
	fmt.Printf("Doctor checkup the patient %s already done.\n", p.name)
	h.Next(h.handler, p)
}

// 去药部门
type MedicalHandler struct {
	BaseHandler
}

func (h *MedicalHandler) Handle(p *Patient) {
	fmt.Printf("Medicine already given to patient %s.\n", p.name)
	h.Next(h.handler, p)
}

// 结账
type CashHandler struct {
	BaseHandler
}

func (h *CashHandler) Handle(p *Patient) {
	fmt.Printf("The patient %s payment done\n", p.name)
	h.Next(h.handler, p)
}

type Patient struct {
	name string
}
