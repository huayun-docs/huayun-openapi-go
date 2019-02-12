package chinacgatewayecs

import (
)

type deleteKeyPair struct {
    req        	map[string]string
}

func DeleteKeyPair() *deleteKeyPair {
    c := &deleteKeyPair {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteKeyPair"
    return c
}

func (c *deleteKeyPair) GetRequest() map[string]string {
    return c.req
}

func (c *deleteKeyPair) GetAction() string {
    return c.req["Action"]
}

func (c *deleteKeyPair) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteKeyPair) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteKeyPair) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteKeyPair) GetId() string {
    return c.req["Id"]
}

func (c *deleteKeyPair) SetId(varId string) {
    c.req["Id"] = varId
}


