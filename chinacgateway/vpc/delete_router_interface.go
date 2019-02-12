package chinacgatewayvpc

import (
)

type deleteRouterInterface struct {
    req        	map[string]string
}

func DeleteRouterInterface() *deleteRouterInterface {
    c := &deleteRouterInterface {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteRouterInterface"
    return c
}

func (c *deleteRouterInterface) GetRequest() map[string]string {
    return c.req
}

func (c *deleteRouterInterface) GetAction() string {
    return c.req["Action"]
}

func (c *deleteRouterInterface) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteRouterInterface) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteRouterInterface) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteRouterInterface) GetId() string {
    return c.req["Id"]
}

func (c *deleteRouterInterface) SetId(varId string) {
    c.req["Id"] = varId
}


