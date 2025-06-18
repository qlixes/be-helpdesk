package facades

type Signature struct {
	path string
}

type SignatureManager interface {
}

func NewSignature() *Signature {
	return &Signature{}
}
