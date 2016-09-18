package gosupervisor

type SupervisorRpc struct {
	Url string
}

func New(url string) *SupervisorRpc {
	return &SupervisorRpc{Url: url}
}

