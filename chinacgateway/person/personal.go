package cgperson

type personal struct {
	req map[string]string
}

func Personal() *personal {
	p := &personal{req: make(map[string]string)}
	p.req["Action"] = "Personal"
	return p
}

func (p *personal) GetRequest() map[string]string {
	return p.req
}

func (p *personal) GetAction() string {
	return p.req["Action"]
}
