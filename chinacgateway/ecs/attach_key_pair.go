package chinacgatewayecs

import (
)

type attachKeyPair struct {
    req        	map[string]string
}

func AttachKeyPair() *attachKeyPair {
    c := &attachKeyPair {
        req:			make(map[string]string),
    }

    c.req["Action"] = "AttachKeyPair"
    return c
}

func (c *attachKeyPair) GetRequest() map[string]string {
    return c.req
}

func (c *attachKeyPair) GetAction() string {
    return c.req["Action"]
}

func (c *attachKeyPair) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *attachKeyPair) GetRegion() string {
    return c.req["Region"]
}

func (c *attachKeyPair) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *attachKeyPair) GetKeyPairId() string {
    return c.req["KeyPairId"]
}

func (c *attachKeyPair) SetKeyPairId(varKeyPairId string) {
    c.req["KeyPairId"] = varKeyPairId
}

func (c *attachKeyPair) GetInstanceId() string {
    return c.req["InstanceId"]
}

func (c *attachKeyPair) SetInstanceId(varInstanceId string) {
    c.req["InstanceId"] = varInstanceId
}


