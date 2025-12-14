package types

type Student struct {
	ID    int
	Name  string `validate:"required"`
	Age   int
	Email string `validate:"required,email"`
}
