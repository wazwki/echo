package v1

type Server struct {
	service service.ServiceInterface
}

func NewControllers(s service.ServiceInterface) *Server {
	return &Server{service: s}
}
