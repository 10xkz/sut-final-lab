package entity_test

import(
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"github.com/10xkz/sut-final-lab/entity"
	

)

func TestPositiveEmployees(t *testing.T){
	t.Run("Positive Employees",func(t *testing.T){
		g := NewGomegaWithT(t)
		employess := entity.Employees{
			Name: "tar",
			Salary: 16000.0,
			EmployeeCode: "HR-1024",
		}
		ok,err := govalidator.ValidateStruct(employess)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}