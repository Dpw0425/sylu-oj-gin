package error

import "github.com/pkg/errors"

var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	Unwrap       = errors.Unwrap
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
	Cause        = errors.Cause
	Errorf       = errors.Errorf
	Is           = errors.Is
	As           = errors.As
)

func (r *ResponseError) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}
	return r.Message
}
