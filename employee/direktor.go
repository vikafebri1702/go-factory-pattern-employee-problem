// director.go
package employee

// Director is a struct for director
// It embeds Employee
type Director struct {
	Employee // Director embeds Employee
}

// NewDirector creates a new director
func NewDirector() *Director {
	director := &Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
			Bonus:  0, // Bonus will be calculated separately
		},
	}
	director.calculateBonus()
	return director
}

// calculateBonus calculates bonus for director
func (d *Director) calculateBonus() {
	d.Bonus = float64(d.Salary) * 0.30 // 30% bonus for director
}
