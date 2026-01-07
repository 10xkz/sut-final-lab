package entity_test

import(
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"github.com/10xkz/sut-final-lab/entity"
	

)

func TestEmployeeCodeEmployees(t *testing.T){
	t.Run("Negative EmployeeCode Employees",func(t *testing.T){
		g := NewGomegaWithT(t)
		employess := entity.Employees{
			Name: "tar",
			Salary: 16000.0,
			EmployeeCode: "HR1024",
		}
		ok,err := govalidator.ValidateStruct(employess)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))
	})
}