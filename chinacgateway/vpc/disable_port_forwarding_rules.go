package chinacgatewayvpc

import (
)

type disablePortForwardingRules struct {
    req        	map[string]string
}

func DisablePortForwardingRules() *disablePortForwardingRules {
    c := &disablePortForwardingRules {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DisablePortForwardingRules"
    return c
}

func (c *disablePortForwardingRules) GetRequest() map[string]string {
    return c.req
}

func (c *disablePortForwardingRules) GetAction() string {
    return c.req["Action"]
}

func (c *disablePortForwardingRules) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *disablePortForwardingRules) GetRegion() string {
    return c.req["Region"]
}

func (c *disablePortForwardingRules) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *disablePortForwardingRules) GetId0() string {
    return c.req["Id.0"]
}

func (c *disablePortForwardingRules) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}


