package employee

// Intern is a struct for intern
// It embeds Employee
// This means that Intern has all the fields and methods of Employee
type Intern struct {
	Employee // Embedding Employee struct
}

// NewIntern creates a new intern
// It returns a pointer to the intern
// Creational method
func NewIntern() *Intern {
	// Create a new intern
	intern := &Intern{
		Employee: Employee{
			Name:   "Intern", // Set the name to "Intern"
			Salary: 100,      // Set the salary to 100
		},
	}
	return intern
}

// GetBonus calculates and returns the bonus of the intern
func (i *Intern) GetBonus() float64 {
	// Interns do not receive any bonus
	return 0
}
