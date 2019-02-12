package chinacgatewayvpc

import (
)

type deleteFirewallRules struct {
    req        	map[string]string
}

func DeleteFirewallRules() *deleteFirewallRules {
    c := &deleteFirewallRules {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteFirewallRules"
    return c
}

func (c *deleteFirewallRules) GetRequest() map[string]string {
    return c.req
}

func (c *deleteFirewallRules) GetAction() string {
    return c.req["Action"]
}

func (c *deleteFirewallRules) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteFirewallRules) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteFirewallRules) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteFirewallRules) GetId0() string {
    return c.req["Id.0"]
}

func (c *deleteFirewallRules) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}


