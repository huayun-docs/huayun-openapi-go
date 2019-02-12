package chinacgatewayvpc

import (
)

type startRouter struct {
    req        	map[string]string
}

func StartRouter() *startRouter {
    c := &startRouter {
        req:			make(map[string]string),
    }

    c.req["Action"] = "StartRouter"
    return c
}

func (c *startRouter) GetRequest() map[string]string {
    return c.req
}

func (c *startRouter) GetAction() string {
    return c.req["Action"]
}

func (c *startRouter) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *startRouter) GetRegion() string {
    return c.req["Region"]
}

func (c *startRouter) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *startRouter) GetId() string {
    return c.req["Id"]
}

func (c *startRouter) SetId(varId string) {
    c.req["Id"] = varId
}


