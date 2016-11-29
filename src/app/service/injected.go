package service

type Injected struct{}

func (Injected) Call(arg string) string {
	return "injected!"
}