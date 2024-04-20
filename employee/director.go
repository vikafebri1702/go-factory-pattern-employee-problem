// director.go
package employee

type Director struct {
	Employee
}

func NewDirector() *Director {
	director := &Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
			Bonus:  0,
		},
	}
	director.calculateBonus()
	return director
}

func (d *Director) calculateBonus() {
	d.Bonus = float64(d.Salary) * 0.30
}
