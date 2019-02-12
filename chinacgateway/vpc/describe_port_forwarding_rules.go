package chinacgatewayvpc

import (
)

type describePortForwardingRules struct {
    req        	map[string]string
}

func DescribePortForwardingRules() *describePortForwardingRules {
    c := &describePortForwardingRules {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribePortForwardingRules"
    return c
}

func (c *describePortForwardingRules) GetRequest() map[string]string {
    return c.req
}

func (c *describePortForwardingRules) GetAction() string {
    return c.req["Action"]
}

func (c *describePortForwardingRules) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describePortForwardingRules) GetRegion() string {
    return c.req["Region"]
}

func (c *describePortForwardingRules) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describePortForwardingRules) GetOffset() string {
    return c.req["Offset"]
}

func (c *describePortForwardingRules) SetOffset(varOffset string) {
    c.req["Offset"] = varOffset
}

func (c *describePortForwardingRules) GetCount() string {
    return c.req["Count"]
}

func (c *describePortForwardingRules) SetCount(varCount string) {
    c.req["Count"] = varCount
}

func (c *describePortForwardingRules) GetId0() string {
    return c.req["Id.0"]
}

func (c *describePortForwardingRules) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describePortForwardingRules) GetRouterId0() string {
    return c.req["RouterId.0"]
}

func (c *describePortForwardingRules) SetRouterId0(varRouterId0 string) {
    c.req["RouterId.0"] = varRouterId0
}


