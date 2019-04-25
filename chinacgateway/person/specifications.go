package cgperson

type specifications struct {
	req map[string]string
}

func Specifications() *specifications {
	s := &specifications{req: make(map[string]string)}
	s.req["Action"] = "Specifications"
	return s
}

func (s *specifications) GetRequest() map[string]string {
	return s.req
}

func (s *specifications) GetAction() string {
	return s.req["Action"]
}
