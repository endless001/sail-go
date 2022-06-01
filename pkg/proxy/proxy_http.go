package proxy

type ProxyHttp struct{}

func New() (*ProxyHttp, error) {
	p := &ProxyHttp{}
	return p, nil
}
