package balancer

type LoadBalancer interface {
	NextServer([]string) string
}
