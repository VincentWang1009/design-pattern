package chainofresponsibility

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor already checkup")
	d.next.execute(p)
}

func (d *Doctor) setNext(department Department) {
	d.next = department
}
