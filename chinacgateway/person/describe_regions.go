package cgperson

type describeRegions struct {
	req map[string]string
}

func DescribeRegions() *describeRegions {
	d := &describeRegions{req: make(map[string]string)}
	d.req["Action"] = "DescribeRegions"
	return d
}

func (d *describeRegions) GetRequest() map[string]string {
	return d.req
}

func (d *describeRegions) GetAction() string {
	return d.req["Action"]
}
