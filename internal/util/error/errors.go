package error

import "github.com/pkg/errors"

var ErrDB = errors.New("database error")
var ErrJWT = errors.New("validating jwt error")
