package chinacgatewayecs

import (
)

type resizeInstance struct {
    req        	map[string]string
}

func ResizeInstance() *resizeInstance {
    c := &resizeInstance {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ResizeInstance"
    return c
}

func (c *resizeInstance) GetRequest() map[string]string {
    return c.req
}

func (c *resizeInstance) GetRegion() string {
    return c.req["Region"]
}

func (c *resizeInstance) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *resizeInstance) GetAction() string {
    return c.req["Action"]
}

func (c *resizeInstance) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *resizeInstance) GetId() string {
    return c.req["Id"]
}

func (c *resizeInstance) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *resizeInstance) GetInstanceType() string {
    return c.req["InstanceType"]
}

func (c *resizeInstance) SetInstanceType(varInstanceType string) {
    c.req["InstanceType"] = varInstanceType
}


