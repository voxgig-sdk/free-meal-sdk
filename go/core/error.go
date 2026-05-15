package core

type FreeMealError struct {
	IsFreeMealError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewFreeMealError(code string, msg string, ctx *Context) *FreeMealError {
	return &FreeMealError{
		IsFreeMealError: true,
		Sdk:              "FreeMeal",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *FreeMealError) Error() string {
	return e.Msg
}
