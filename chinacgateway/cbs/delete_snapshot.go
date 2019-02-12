package chinacgatewaycbs

import (
)

type deleteSnapshot struct {
    req        	map[string]string
}

func DeleteSnapshot() *deleteSnapshot {
    c := &deleteSnapshot {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteSnapshot"
    return c
}

func (c *deleteSnapshot) GetRequest() map[string]string {
    return c.req
}

func (c *deleteSnapshot) GetAction() string {
    return c.req["Action"]
}

func (c *deleteSnapshot) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteSnapshot) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteSnapshot) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteSnapshot) GetId() string {
    return c.req["Id"]
}

func (c *deleteSnapshot) SetId(varId string) {
    c.req["Id"] = varId
}


