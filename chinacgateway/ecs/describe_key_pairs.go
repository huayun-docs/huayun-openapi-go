package chinacgatewayecs

import (
    "strconv"
)

type describeKeyPairs struct {
    req        	map[string]string
}

func DescribeKeyPairs() *describeKeyPairs {
    c := &describeKeyPairs {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeKeyPairs"
    return c
}

func (c *describeKeyPairs) GetRequest() map[string]string {
    return c.req
}

func (c *describeKeyPairs) GetAction() string {
    return c.req["Action"]
}

func (c *describeKeyPairs) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeKeyPairs) GetRegion() string {
    return c.req["Region"]
}

func (c *describeKeyPairs) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeKeyPairs) GetOffset() int {
    nOffset, _ := strconv.Atoi(c.req["Offset"])
    return nOffset
}

func (c *describeKeyPairs) SetOffset(varOffset int) {
    c.req["Offset"] = strconv.Itoa(varOffset)
}

func (c *describeKeyPairs) GetCount() string {
    return c.req["Count"]
}

func (c *describeKeyPairs) SetCount(varCount string) {
    c.req["Count"] = varCount
}

func (c *describeKeyPairs) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeKeyPairs) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeKeyPairs) GetName0() string {
    return c.req["Name.0"]
}

func (c *describeKeyPairs) SetName0(varName0 string) {
    c.req["Name.0"] = varName0
}


