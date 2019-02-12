package chinacgatewayecs

import (
)

type detachKeyPair struct {
    req        	map[string]string
}

func DetachKeyPair() *detachKeyPair {
    c := &detachKeyPair {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DetachKeyPair"
    return c
}

func (c *detachKeyPair) GetRequest() map[string]string {
    return c.req
}

func (c *detachKeyPair) GetAction() string {
    return c.req["Action"]
}

func (c *detachKeyPair) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *detachKeyPair) GetRegion() string {
    return c.req["Region"]
}

func (c *detachKeyPair) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *detachKeyPair) GetKeyPairId() string {
    return c.req["KeyPairId"]
}

func (c *detachKeyPair) SetKeyPairId(varKeyPairId string) {
    c.req["KeyPairId"] = varKeyPairId
}

func (c *detachKeyPair) GetInstanceId() string {
    return c.req["InstanceId"]
}

func (c *detachKeyPair) SetInstanceId(varInstanceId string) {
    c.req["InstanceId"] = varInstanceId
}


