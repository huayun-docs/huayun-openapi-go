package chinacgatewayecs

import (
)

type stopInstance struct {
    req        	map[string]string
}

func StopInstance() *stopInstance {
    c := &stopInstance {
        req:			make(map[string]string),
    }

    c.req["Action"] = "StopInstance"
    return c
}

func (c *stopInstance) GetRequest() map[string]string {
    return c.req
}

func (c *stopInstance) GetRegion() string {
    return c.req["Region"]
}

func (c *stopInstance) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *stopInstance) GetAction() string {
    return c.req["Action"]
}

func (c *stopInstance) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *stopInstance) GetId() string {
    return c.req["Id"]
}

func (c *stopInstance) SetId(varId string) {
    c.req["Id"] = varId
}


