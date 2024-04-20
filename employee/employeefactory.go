// employee.go
package employee

import "errors"

// GetEmployeeFactory is a factory for employee
// It returns an employee based on the input
// If the employee is not found, it returns an error
func GetEmployeeFactory(empl string) (EmployeeInterface, error) {
	if empl == "manager" {
		return NewManager(), nil
	}

	if empl == "staff" {
		return NewStaff(), nil
	}

	if empl == "intern" {
		return NewIntern(), nil
	}

	if empl == "director" {
		return NewDirector(), nil
	}

	return nil, errors.New("Employee not found")
}
