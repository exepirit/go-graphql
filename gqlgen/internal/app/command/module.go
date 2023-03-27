package command

import "go.uber.org/fx"

type Commands struct {
	fx.In

	CreateUser Command[CreateUserArgs]
}

var Module = fx.Options(
	fx.Provide(fx.Annotate(NewCreateUserCommand, fx.As(new(Command[CreateUserArgs])))),
)
