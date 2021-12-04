package depend_inject

type Hello interface {
	Say(name string) string
}

type HelloImpl struct {}

func (h *HelloImpl) Say(name string) string {
	return "Hello " + name
}

func NewHelloImpl() *HelloImpl {
	return &HelloImpl{}
}



type HelloService struct {
	Hello
}

func NewHelloService(hello Hello) *HelloService {
	return &HelloService{Hello: hello}
}

