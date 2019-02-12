package chinacgatewayecs

import (
)

type startInstance struct {
    req        	map[string]string
}

func StartInstance() *startInstance {
    c := &startInstance {
        req:			make(map[string]string),
    }

    c.req["Action"] = "StartInstance"
    return c
}

func (c *startInstance) GetRequest() map[string]string {
    return c.req
}

func (c *startInstance) GetRegion() string {
    return c.req["Region"]
}

func (c *startInstance) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *startInstance) GetAction() string {
    return c.req["Action"]
}

func (c *startInstance) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *startInstance) GetId() string {
    return c.req["Id"]
}

func (c *startInstance) SetId(varId string) {
    c.req["Id"] = varId
}


