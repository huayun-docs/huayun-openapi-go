package chinacgatewayvpc

import (
)

type stopRouter struct {
    req        	map[string]string
}

func StopRouter() *stopRouter {
    c := &stopRouter {
        req:			make(map[string]string),
    }

    c.req["Action"] = "StopRouter"
    return c
}

func (c *stopRouter) GetRequest() map[string]string {
    return c.req
}

func (c *stopRouter) GetAction() string {
    return c.req["Action"]
}

func (c *stopRouter) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *stopRouter) GetRegion() string {
    return c.req["Region"]
}

func (c *stopRouter) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *stopRouter) GetId() string {
    return c.req["Id"]
}

func (c *stopRouter) SetId(varId string) {
    c.req["Id"] = varId
}


