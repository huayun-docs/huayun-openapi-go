package chinacgatewayvpc

import (
)

type modifyFirewallRuleAttributes struct {
    req        	map[string]string
}

func ModifyFirewallRuleAttributes() *modifyFirewallRuleAttributes {
    c := &modifyFirewallRuleAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifyFirewallRuleAttributes"
    return c
}

func (c *modifyFirewallRuleAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifyFirewallRuleAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifyFirewallRuleAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifyFirewallRuleAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifyFirewallRuleAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifyFirewallRuleAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifyFirewallRuleAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifyFirewallRuleAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifyFirewallRuleAttributes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *modifyFirewallRuleAttributes) GetDirection() string {
    return c.req["Direction"]
}

func (c *modifyFirewallRuleAttributes) SetDirection(varDirection string) {
    c.req["Direction"] = varDirection
}

func (c *modifyFirewallRuleAttributes) GetPortStart() string {
    return c.req["PortStart"]
}

func (c *modifyFirewallRuleAttributes) SetPortStart(varPortStart string) {
    c.req["PortStart"] = varPortStart
}

func (c *modifyFirewallRuleAttributes) GetPortEnd() string {
    return c.req["PortEnd"]
}

func (c *modifyFirewallRuleAttributes) SetPortEnd(varPortEnd string) {
    c.req["PortEnd"] = varPortEnd
}

func (c *modifyFirewallRuleAttributes) GetProtocol() string {
    return c.req["Protocol"]
}

func (c *modifyFirewallRuleAttributes) SetProtocol(varProtocol string) {
    c.req["Protocol"] = varProtocol
}

func (c *modifyFirewallRuleAttributes) GetPriority() string {
    return c.req["Priority"]
}

func (c *modifyFirewallRuleAttributes) SetPriority(varPriority string) {
    c.req["Priority"] = varPriority
}

func (c *modifyFirewallRuleAttributes) GetRemoteIpPrefix() string {
    return c.req["RemoteIpPrefix"]
}

func (c *modifyFirewallRuleAttributes) SetRemoteIpPrefix(varRemoteIpPrefix string) {
    c.req["RemoteIpPrefix"] = varRemoteIpPrefix
}



