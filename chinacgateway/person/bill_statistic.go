package cgperson

type billStatistic struct {
	req map[string]string
}

func BillStatistic() *billStatistic {
	b := &billStatistic{req: make(map[string]string)}
	b.req["Action"] = "BillStatistic"
	return b
}

func (b *billStatistic) GetRequest() map[string]string {
	return b.req
}

func (b *billStatistic) GetAction() string {
	return b.req["Action"]
}

func (b *billStatistic) SetDates(d string) {
	b.req["Dates"] = d
}

func (b *billStatistic) GetDates() string {
	return b.req["Dates"]
}
