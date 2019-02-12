package chinacgatewayvpc

import (
    "strconv"
)

type describeFirewallRules struct {
    req        	map[string]string
}

func DescribeFirewallRules() *describeFirewallRules {
    c := &describeFirewallRules {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeFirewallRules"
    return c
}

func (c *describeFirewallRules) GetRequest() map[string]string {
    return c.req
}

func (c *describeFirewallRules) GetAction() string {
    return c.req["Action"]
}

func (c *describeFirewallRules) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeFirewallRules) GetRegion() string {
    return c.req["Region"]
}

func (c *describeFirewallRules) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeFirewallRules) GetOffset() int {
    nOffset, _ := strconv.Atoi(c.req["Offset"])
    return nOffset
}

func (c *describeFirewallRules) SetOffset(varOffset int) {
    c.req["Offset"] = strconv.Itoa(varOffset)
}

func (c *describeFirewallRules) GetCount() int {
    nCount, _ := strconv.Atoi(c.req["Count"])
    return nCount
}

func (c *describeFirewallRules) SetCount(varCount int) {
    c.req["Count"] = strconv.Itoa(varCount)
}

func (c *describeFirewallRules) GetFirewallRule() string {
    return c.req["FirewallRule"]
}

func (c *describeFirewallRules) SetFirewallRule(varFirewallRule string) {
    c.req["FirewallRule"] = varFirewallRule
}

func (c *describeFirewallRules) GetFirewall0() string {
    return c.req["Firewall.0"]
}

func (c *describeFirewallRules) SetFirewall0(varFirewall0 string) {
    c.req["Firewall.0"] = varFirewall0
}


