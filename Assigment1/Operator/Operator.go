package Operator

type Operator struct {
	position string
	salary   float64
	address  string
}

func NewOperator(position string, salary float64, address string) *Operator {
	return &Operator{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (d *Operator) GetPosition() string {
	return d.position
}

func (d *Operator) SetPosition(position string) {
	d.position = position
}

func (d *Operator) GetSalary() float64 {
	return d.salary
}

func (d *Operator) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Operator) GetAddress() string {
	return d.address
}

func (d *Operator) SetAddress(address string) {
	d.address = address
}
