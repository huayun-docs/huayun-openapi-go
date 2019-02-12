package chinacgatewayvpc

import (
)

type deleteFirewall struct {
    req        	map[string]string
}

func DeleteFirewall() *deleteFirewall {
    c := &deleteFirewall {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteFirewall"
    return c
}

func (c *deleteFirewall) GetRequest() map[string]string {
    return c.req
}

func (c *deleteFirewall) GetAction() string {
    return c.req["Action"]
}

func (c *deleteFirewall) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteFirewall) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteFirewall) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteFirewall) GetId() string {
    return c.req["Id"]
}

func (c *deleteFirewall) SetId(varId string) {
    c.req["Id"] = varId
}


