package chinacgatewayecs

import (
    "strconv"
)

type describeImages struct {
    req        	map[string]string
}

func DescribeImages() *describeImages {
    c := &describeImages {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeImages"
    return c
}

func (c *describeImages) GetRequest() map[string]string {
    return c.req
}

func (c *describeImages) GetAction() string {
    return c.req["Action"]
}

func (c *describeImages) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeImages) GetRegion() string {
    return c.req["Region"]
}

func (c *describeImages) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeImages) GetOffset() int {
    nOffset, _ := strconv.Atoi(c.req["Offset"])
    return nOffset
}

func (c *describeImages) SetOffset(varOffset int) {
    c.req["Offset"] = strconv.Itoa(varOffset)
}

func (c *describeImages) GetCount() string {
    return c.req["Count"]
}

func (c *describeImages) SetCount(varCount string) {
    c.req["Count"] = varCount
}

func (c *describeImages) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeImages) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeImages) GetName() string {
    return c.req["Name"]
}

func (c *describeImages) SetName(varName string) {
    c.req["Name"] = varName
}


