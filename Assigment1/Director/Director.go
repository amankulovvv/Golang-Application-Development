package Director

type Director struct {
	position string
	salary   float64
	address  string
}

func NewDirector(position string, salary float64, address string) *Director {
	return &Director{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (m *Director) GetPosition() string {
	return m.position
}

func (m *Director) SetPosition(position string) {
	m.position = position
}

func (m *Director) GetSalary() float64 {
	return m.salary
}

func (m *Director) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Director) GetAddress() string {
	return m.address
}

func (m *Director) SetAddress(address string) {
	m.address = address
}
