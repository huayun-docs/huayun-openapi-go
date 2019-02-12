package chinacgatewayvpc

import (
)

type modifyFirewallAttributes struct {
    req        	map[string]string
}

func ModifyFirewallAttributes() *modifyFirewallAttributes {
    c := &modifyFirewallAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifyFirewallAttributes"
    return c
}

func (c *modifyFirewallAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifyFirewallAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifyFirewallAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifyFirewallAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifyFirewallAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifyFirewallAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifyFirewallAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifyFirewallAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifyFirewallAttributes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *modifyFirewallAttributes) GetDescription() string {
    return c.req["Description"]
}

func (c *modifyFirewallAttributes) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


