package chinacgatewayvpc

import (
)

type enablePortForwardingRules struct {
    req        	map[string]string
}

func EnablePortForwardingRules() *enablePortForwardingRules {
    c := &enablePortForwardingRules {
        req:			make(map[string]string),
    }

    c.req["Action"] = "EnablePortForwardingRules"
    return c
}

func (c *enablePortForwardingRules) GetRequest() map[string]string {
    return c.req
}

func (c *enablePortForwardingRules) GetAction() string {
    return c.req["Action"]
}

func (c *enablePortForwardingRules) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *enablePortForwardingRules) GetRegion() string {
    return c.req["Region"]
}

func (c *enablePortForwardingRules) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *enablePortForwardingRules) GetId0() string {
    return c.req["Id.0"]
}

func (c *enablePortForwardingRules) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}


