package employee

// Manager is a struct for manager
// It embeds Employee
type Manager struct {
	Employee
}

// NewManager creates a new manager
func NewManager() *Manager {
	manager := &Manager{
		Employee: Employee{
			Name:   "Manager",
			Salary: 1000,
			Bonus:  0, // Bonus will be calculated separately
		},
	}
	manager.calculateBonus()
	return manager
}

// calculateBonus calculates bonus for manager
func (m *Manager) calculateBonus() {
	m.Bonus = float64(m.Salary) * 0.20 // 20% bonus for manager
}
