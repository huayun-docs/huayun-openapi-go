package chinacgatewayecs

import (
)

type changeInstanceInterface struct {
    req        	map[string]string
}

func ChangeInstanceInterface() *changeInstanceInterface {
    c := &changeInstanceInterface {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ChangeInstanceInterface"
    return c
}

func (c *changeInstanceInterface) GetRequest() map[string]string {
    return c.req
}

func (c *changeInstanceInterface) GetRegion() string {
    return c.req["Region"]
}

func (c *changeInstanceInterface) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *changeInstanceInterface) GetAction() string {
    return c.req["Action"]
}

func (c *changeInstanceInterface) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *changeInstanceInterface) GetInstanceId() string {
    return c.req["InstanceId"]
}

func (c *changeInstanceInterface) SetInstanceId(varInstanceId string) {
    c.req["InstanceId"] = varInstanceId
}

func (c *changeInstanceInterface) GetOldNetworkId() string {
    return c.req["OldNetworkId"]
}

func (c *changeInstanceInterface) SetOldNetworkId(varOldNetworkId string) {
    c.req["OldNetworkId"] = varOldNetworkId
}

func (c *changeInstanceInterface) GetNewNetworkId() string {
    return c.req["NewNetworkId"]
}

func (c *changeInstanceInterface) SetNewNetworkId(varNewNetworkId string) {
    c.req["NewNetworkId"] = varNewNetworkId
}

func (c *changeInstanceInterface) GetNetworkIpAddress() string {
    return c.req["NetworkIpAddress"]
}

func (c *changeInstanceInterface) SetNetworkIpAddress(varNetworkIpAddress string) {
    c.req["NetworkIpAddress"] = varNetworkIpAddress
}


