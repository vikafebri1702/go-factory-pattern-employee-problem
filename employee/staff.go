package employee

// Staff is a struct for staff
// It embeds Employee
type Staff struct {
	Employee
}

// NewStaff creates a new staff
func NewStaff() *Staff {
	staff := &Staff{
		Employee: Employee{
			Name:   "Staff",
			Salary: 500,
			Bonus:  0, // Bonus will be calculated separately
		},
	}
	staff.calculateBonus()
	return staff
}

// calculateBonus calculates bonus for staff
func (s *Staff) calculateBonus() {
	s.Bonus = float64(s.Salary) * 0.10 // 10% bonus for staff
}
