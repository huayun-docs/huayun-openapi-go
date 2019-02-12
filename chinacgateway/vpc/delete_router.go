package chinacgatewayvpc

import (
)

type deleteRouter struct {
    req        	map[string]string
}

func DeleteRouter() *deleteRouter {
    c := &deleteRouter {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteRouter"
    return c
}

func (c *deleteRouter) GetRequest() map[string]string {
    return c.req
}

func (c *deleteRouter) GetAction() string {
    return c.req["Action"]
}

func (c *deleteRouter) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteRouter) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteRouter) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteRouter) GetId() string {
    return c.req["Id"]
}

func (c *deleteRouter) SetId(varId string) {
    c.req["Id"] = varId
}


