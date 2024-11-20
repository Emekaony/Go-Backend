package structbasics

/*
in Go, u want to strive to alwaya ccept interfaces but return concrete types.

Dependency Injection: This is when an object gets its dependencies from an
external source as opposed to cretaing them itself. Example:
*/

type Logger interface {
	Log(message string)
}

type Service struct {
	logger Logger
}

func NewService(logger Logger) *Service {
	// we create and return a new Service object but the logger dependency
	// is given to use from an extenral source, we do not create it.
	// this is a design pattern I need to look into!
	return &Service{logger: logger}
}
