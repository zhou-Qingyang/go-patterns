package proxy

type server interface {
	handlerHttpRequest(url string, ip string) (int, string)
}
type HttpServer struct {
}

func (h *HttpServer) handlerHttpRequest(url string, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Ok"
}

type Nginx struct {
	application       server
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginx(maxAllowedRequest int) *Nginx {
	return &Nginx{
		application:       &HttpServer{},
		maxAllowedRequest: maxAllowedRequest,
		rateLimiter:       make(map[string]int),
	}
}
func (n *Nginx) handlerHttpRequest(url string, ip string) (int, string) {
	if n.rateLimiter[ip] >= n.maxAllowedRequest {
		return 403, "Not Allowed"
	}
	n.rateLimiter[ip]++
	return n.application.handlerHttpRequest(url, ip)
}
