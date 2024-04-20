package employee

// EmployeeInterface is an interface for employee
type EmployeeInterface interface {
	SetName(name string)
	SetSalary(salary int)
	GetName() string
	GetSalary() int
	GetBonus() float64
}

// Employee is a struct for employee
type Employee struct {
	Name   string
	Salary int
	Bonus  float64
}

// SetName sets the name of the employee
func (e *Employee) SetName(name string) {
	e.Name = name // Set the name of the employee
}

// SetSalary sets the salary of the employee
func (e *Employee) SetSalary(salary int) {
	e.Salary = salary // Set the salary of the employee
}

// GetName gets the name of the employee
func (e *Employee) GetName() string {
	return e.Name // Get the name of the employee
}

// GetSalary gets the salary of the employee
func (e *Employee) GetSalary() int {
	return e.Salary // Get the salary of the employee
}

// GetBonus calculates and returns the bonus of the employee
func (e *Employee) GetBonus() float64 {
	return e.Bonus // Get the bonus of the employee
}
