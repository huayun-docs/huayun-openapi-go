package chinacgatewayvpc

import (
)

type deletePortForwardingRule struct {
    req        	map[string]string
}

func DeletePortForwardingRule() *deletePortForwardingRule {
    c := &deletePortForwardingRule {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeletePortForwardingRule"
    return c
}

func (c *deletePortForwardingRule) GetRequest() map[string]string {
    return c.req
}

func (c *deletePortForwardingRule) GetAction() string {
    return c.req["Action"]
}

func (c *deletePortForwardingRule) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deletePortForwardingRule) GetRegion() string {
    return c.req["Region"]
}

func (c *deletePortForwardingRule) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deletePortForwardingRule) GetId() string {
    return c.req["Id"]
}

func (c *deletePortForwardingRule) SetId(varId string) {
    c.req["Id"] = varId
}


