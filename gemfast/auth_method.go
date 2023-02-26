package gemfast

import (
	"fmt"
)

type AuthMethod interface {
	fmt.Stringer
}

type LocalAuth struct {
	JWTToken string
}

func (l LocalAuth) String() (string) {
	return fmt.Sprintf("Bearer %s", l.JWTToken)
}