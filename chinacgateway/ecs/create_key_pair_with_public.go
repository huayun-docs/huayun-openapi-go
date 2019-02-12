package chinacgatewayecs

import (
)

type createKeyPairWithPublic struct {
    req        	map[string]string
}

func CreateKeyPairWithPublic() *createKeyPairWithPublic {
    c := &createKeyPairWithPublic {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateKeyPairWithPublic"
    return c
}

func (c *createKeyPairWithPublic) GetRequest() map[string]string {
    return c.req
}

func (c *createKeyPairWithPublic) GetAction() string {
    return c.req["Action"]
}

func (c *createKeyPairWithPublic) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createKeyPairWithPublic) GetRegion() string {
    return c.req["Region"]
}

func (c *createKeyPairWithPublic) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createKeyPairWithPublic) GetName() string {
    return c.req["Name"]
}

func (c *createKeyPairWithPublic) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createKeyPairWithPublic) GetPublicKey() string {
    return c.req["PublicKey"]
}

func (c *createKeyPairWithPublic) SetPublicKey(varPublicKey string) {
    c.req["PublicKey"] = varPublicKey
}

func (c *createKeyPairWithPublic) GetDescription() string {
    return c.req["Description"]
}

func (c *createKeyPairWithPublic) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


