package employee_test

import (
	"go-factorypattern-company-case/employee"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employee", func() {

	Context("Manager Object", func() {

		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("manager")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Manager"))
			Expect(empl.GetSalary()).To(Equal(1000))
		})

		It("should return the correct bonus", func() {
			empl := employee.NewManager()

			bonus := empl.GetBonus()
			Expect(bonus).To(Equal(200.0))
		})

	})

	Context("Staff Object", func() {

		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("staff")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Staff"))
			Expect(empl.GetSalary()).To(Equal(500))
		})

		It("should return the correct bonus", func() {
			empl := employee.NewStaff()

			bonus := empl.GetBonus()
			Expect(bonus).To(Equal(50.0))
		})

	})

	Context("Intern Object", func() {

		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("intern")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Intern"))
			Expect(empl.GetSalary()).To(Equal(100))
		})

		It("should return the correct bonus", func() {
			empl := employee.NewIntern()

			bonus := empl.GetBonus()
			Expect(bonus).To(Equal(0.0))
		})
	})

	Context("Director Object", func() {
		It("should return the correct name and salary", func() {
			empl := employee.NewDirector()

			Expect(empl.GetName()).To(Equal("Director"))
			Expect(empl.GetSalary()).To(Equal(5000))
		})

		It("should return the correct bonus", func() {
			empl := employee.NewDirector()

			bonus := empl.GetBonus()
			Expect(bonus).To(Equal(1500.0))
		})
	})

	Context("Empty Employee", func() {

		It("should return an error", func() {
			_, err := employee.GetEmployeeFactory("non-existing")
			Expect(err).To(HaveOccurred())
		})

	})

})
