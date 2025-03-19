package main

// Interface renomeada para evitar conflito
type Service interface {
	Request()
}

// Implementação concreta do serviço
type RealService struct{}

func (r *RealService) Request() {
	println("Real request")
}

// Proxy com nova nomenclatura
type ServiceProxy struct {
	service *RealService
}

func (p *ServiceProxy) Request() {
	if p.service == nil {
		p.service = &RealService{}
	}
	println("Proxy: Verificando acesso antes de delegar...")
	p.service.Request()
}

// Exemplo de uso
func main() {
	proxy := &ServiceProxy{}
	proxy.Request() // Cria e delega para o RealService
}
