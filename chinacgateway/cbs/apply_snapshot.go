package chinacgatewaycbs

import (
)

type applySnapshot struct {
    req        	map[string]string
}

func ApplySnapshot() *applySnapshot {
    c := &applySnapshot {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ApplySnapshot"
    return c
}

func (c *applySnapshot) GetRequest() map[string]string {
    return c.req
}

func (c *applySnapshot) GetAction() string {
    return c.req["Action"]
}

func (c *applySnapshot) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *applySnapshot) GetRegion() string {
    return c.req["Region"]
}

func (c *applySnapshot) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *applySnapshot) GetId() string {
    return c.req["Id"]
}

func (c *applySnapshot) SetId(varId string) {
    c.req["Id"] = varId
}


