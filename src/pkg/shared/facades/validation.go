package facades

type Validation struct {
}

type ValidationManager interface {
	SetField(field string) *Validation
}

func NewValidation() *Validation {
	return &Validation{}
}

func (v *Validation) SetField(field string) {

}
