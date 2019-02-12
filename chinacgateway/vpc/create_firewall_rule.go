package chinacgatewayvpc

import (
)

type createFirewallRule struct {
    req        	map[string]string
}

func CreateFirewallRule() *createFirewallRule {
    c := &createFirewallRule {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateFirewallRule"
    return c
}

func (c *createFirewallRule) GetRequest() map[string]string {
    return c.req
}

func (c *createFirewallRule) GetAction() string {
    return c.req["Action"]
}

func (c *createFirewallRule) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createFirewallRule) GetRegion() string {
    return c.req["Region"]
}

func (c *createFirewallRule) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createFirewallRule) GetFirewallId() string {
    return c.req["FirewallId"]
}

func (c *createFirewallRule) SetFirewallId(varFirewallId string) {
    c.req["FirewallId"] = varFirewallId
}

func (c *createFirewallRule) GetName() string {
    return c.req["Name"]
}

func (c *createFirewallRule) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createFirewallRule) GetDirection() string {
    return c.req["Direction"]
}

func (c *createFirewallRule) SetDirection(varDirection string) {
    c.req["Direction"] = varDirection
}

func (c *createFirewallRule) GetPortStart() string {
    return c.req["PortStart"]
}

func (c *createFirewallRule) SetPortStart(varPortStart string) {
    c.req["PortStart"] = varPortStart
}

func (c *createFirewallRule) GetPortEnd() string {
    return c.req["PortEnd"]
}

func (c *createFirewallRule) SetPortEnd(varPortEnd string) {
    c.req["PortEnd"] = varPortEnd
}

func (c *createFirewallRule) GetProtocol() string {
    return c.req["Protocol"]
}

func (c *createFirewallRule) SetProtocol(varProtocol string) {
    c.req["Protocol"] = varProtocol
}

func (c *createFirewallRule) GetPriority() string {
    return c.req["Priority"]
}

func (c *createFirewallRule) SetPriority(varPriority string) {
    c.req["Priority"] = varPriority
}

func (c *createFirewallRule) GetRemoteIpPrefix() string {
    return c.req["RemoteIpPrefix"]
}

func (c *createFirewallRule) SetRemoteIpPrefix(varRemoteIpPrefix string) {
    c.req["RemoteIpPrefix"] = varRemoteIpPrefix
}



