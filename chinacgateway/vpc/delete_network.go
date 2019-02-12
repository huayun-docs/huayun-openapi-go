package chinacgatewayvpc

import (
)

type deleteNetwork struct {
    req        	map[string]string
}

func DeleteNetwork() *deleteNetwork {
    c := &deleteNetwork {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteNetwork"
    return c
}

func (c *deleteNetwork) GetRequest() map[string]string {
    return c.req
}

func (c *deleteNetwork) GetAction() string {
    return c.req["Action"]
}

func (c *deleteNetwork) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteNetwork) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteNetwork) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteNetwork) GetId() string {
    return c.req["Id"]
}

func (c *deleteNetwork) SetId(varId string) {
    c.req["Id"] = varId
}


