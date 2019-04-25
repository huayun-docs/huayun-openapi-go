package cgperson

type account struct {
	req map[string]string
}

func Account() *account {
	a := &account{req: make(map[string]string)}
	a.req["Action"] = "Account"
	return a
}

func (a *account) GetRequest() map[string]string {
	return a.req
}

func (a *account) GetAction() string {
	return a.req["Action"]
}
