package entity_test

import(
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"github.com/10xkz/sut-final-lab/entity"
	

)

func TestSalaryEmployees(t *testing.T){
	t.Run("Negative Salary Employees",func(t *testing.T){
		g := NewGomegaWithT(t)
		employess := entity.Employees{
			Name: "tar",
			Salary: 160.0,
			EmployeeCode: "HR-1024",
		}
		ok,err := govalidator.ValidateStruct(employess)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})
}