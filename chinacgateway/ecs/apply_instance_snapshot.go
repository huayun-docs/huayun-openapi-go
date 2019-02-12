package chinacgatewayecs

import (
)

type applyInstanceSnapshot struct {
    req        	map[string]string
}

func ApplyInstanceSnapshot() *applyInstanceSnapshot {
    c := &applyInstanceSnapshot {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ApplyInstanceSnapshot"
    return c
}

func (c *applyInstanceSnapshot) GetRequest() map[string]string {
    return c.req
}

func (c *applyInstanceSnapshot) GetAction() string {
    return c.req["Action"]
}

func (c *applyInstanceSnapshot) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *applyInstanceSnapshot) GetRegion() string {
    return c.req["Region"]
}

func (c *applyInstanceSnapshot) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *applyInstanceSnapshot) GetId() string {
    return c.req["Id"]
}

func (c *applyInstanceSnapshot) SetId(varId string) {
    c.req["Id"] = varId
}


