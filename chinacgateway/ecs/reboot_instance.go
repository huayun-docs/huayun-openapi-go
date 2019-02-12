package chinacgatewayecs

import (
)

type rebootInstance struct {
    req        	map[string]string
}

func RebootInstance() *rebootInstance {
    c := &rebootInstance {
        req:			make(map[string]string),
    }

    c.req["Action"] = "RebootInstance"
    return c
}

func (c *rebootInstance) GetRequest() map[string]string {
    return c.req
}

func (c *rebootInstance) GetRegion() string {
    return c.req["Region"]
}

func (c *rebootInstance) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *rebootInstance) GetAction() string {
    return c.req["Action"]
}

func (c *rebootInstance) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *rebootInstance) GetId() string {
    return c.req["Id"]
}

func (c *rebootInstance) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *rebootInstance) GetForce() string {
    return c.req["Force"]
}

func (c *rebootInstance) SetForce(varForce string) {
    c.req["Force"] = varForce
}


