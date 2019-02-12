package chinacgatewayecs

import (
)

type changeInstanceFirewall struct {
    req        	map[string]string
}

func ChangeInstanceFirewall() *changeInstanceFirewall {
    c := &changeInstanceFirewall {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ChangeInstanceFirewall"
    return c
}

func (c *changeInstanceFirewall) GetRequest() map[string]string {
    return c.req
}

func (c *changeInstanceFirewall) GetRegion() string {
    return c.req["Region"]
}

func (c *changeInstanceFirewall) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *changeInstanceFirewall) GetAction() string {
    return c.req["Action"]
}

func (c *changeInstanceFirewall) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *changeInstanceFirewall) GetId() string {
    return c.req["Id"]
}

func (c *changeInstanceFirewall) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *changeInstanceFirewall) GetFirewallId() string {
    return c.req["FirewallId"]
}

func (c *changeInstanceFirewall) SetFirewallId(varFirewallId string) {
    c.req["FirewallId"] = varFirewallId
}


