package chinacgatewayecs

import (
)

type deleteInstanceSnapshot struct {
    req        	map[string]string
}

func DeleteInstanceSnapshot() *deleteInstanceSnapshot {
    c := &deleteInstanceSnapshot {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteInstanceSnapshot"
    return c
}

func (c *deleteInstanceSnapshot) GetRequest() map[string]string {
    return c.req
}

func (c *deleteInstanceSnapshot) GetAction() string {
    return c.req["Action"]
}

func (c *deleteInstanceSnapshot) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteInstanceSnapshot) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteInstanceSnapshot) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteInstanceSnapshot) GetId() string {
    return c.req["Id"]
}

func (c *deleteInstanceSnapshot) SetId(varId string) {
    c.req["Id"] = varId
}


