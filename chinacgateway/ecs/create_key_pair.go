package chinacgatewayecs

import (
)

type createKeyPair struct {
    req        	map[string]string
}

func CreateKeyPair() *createKeyPair {
    c := &createKeyPair {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateKeyPair"
    return c
}

func (c *createKeyPair) GetRequest() map[string]string {
    return c.req
}

func (c *createKeyPair) GetAction() string {
    return c.req["Action"]
}

func (c *createKeyPair) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createKeyPair) GetRegion() string {
    return c.req["Region"]
}

func (c *createKeyPair) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createKeyPair) GetName() string {
    return c.req["Name"]
}

func (c *createKeyPair) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createKeyPair) GetEncryption() string {
    return c.req["Encryption"]
}

func (c *createKeyPair) SetEncryption(varEncryption string) {
    c.req["Encryption"] = varEncryption
}

func (c *createKeyPair) GetDescription() string {
    return c.req["Description"]
}

func (c *createKeyPair) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


