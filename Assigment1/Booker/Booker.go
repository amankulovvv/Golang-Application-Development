package Booker

type Booker struct {
	position string
	salary   float64
	address  string
}

func NewBooker(position string, salary float64, address string) *Booker {
	return &Booker{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (m *Booker) GetPosition() string {
	return m.position
}

func (m *Booker) SetPosition(position string) {
	m.position = position
}

func (m *Booker) GetSalary() float64 {
	return m.salary
}

func (m *Booker) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Booker) GetAddress() string {
	return m.address
}

func (m *Booker) SetAddress(address string) {
	m.address = address
}
