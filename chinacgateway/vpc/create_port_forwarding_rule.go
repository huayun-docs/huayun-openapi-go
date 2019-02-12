package chinacgatewayvpc

import (
    "strconv"
)

type createPortForwardingRule struct {
    req        	map[string]string
}

func CreatePortForwardingRule() *createPortForwardingRule {
    c := &createPortForwardingRule {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreatePortForwardingRule"
    return c
}

func (c *createPortForwardingRule) GetRequest() map[string]string {
    return c.req
}

func (c *createPortForwardingRule) GetAction() string {
    return c.req["Action"]
}

func (c *createPortForwardingRule) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createPortForwardingRule) GetRegion() string {
    return c.req["Region"]
}

func (c *createPortForwardingRule) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createPortForwardingRule) GetRouterId() string {
    return c.req["RouterId"]
}

func (c *createPortForwardingRule) SetRouterId(varRouterId string) {
    c.req["RouterId"] = varRouterId
}

func (c *createPortForwardingRule) GetName() string {
    return c.req["Name"]
}

func (c *createPortForwardingRule) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createPortForwardingRule) GetProtocol() string {
    return c.req["Protocol"]
}

func (c *createPortForwardingRule) SetProtocol(varProtocol string) {
    c.req["Protocol"] = varProtocol
}

func (c *createPortForwardingRule) GetExternalStartPort() int {
    nExternalStartPort, _ := strconv.Atoi(c.req["ExternalStartPort"])
    return nExternalStartPort
}

func (c *createPortForwardingRule) SetExternalStartPort(varExternalStartPort int) {
    c.req["ExternalStartPort"] = strconv.Itoa(varExternalStartPort)
}

func (c *createPortForwardingRule) GetExternalEndPort() int {
    nExternalEndPort, _ := strconv.Atoi(c.req["ExternalEndPort"])
    return nExternalEndPort
}

func (c *createPortForwardingRule) SetExternalEndPort(varExternalEndPort int) {
    c.req["ExternalEndPort"] = strconv.Itoa(varExternalEndPort)
}

func (c *createPortForwardingRule) GetInternalStartPort() int {
    nInternalStartPort, _ := strconv.Atoi(c.req["InternalStartPort"])
    return nInternalStartPort
}

func (c *createPortForwardingRule) SetInternalStartPort(varInternalStartPort int) {
    c.req["InternalStartPort"] = strconv.Itoa(varInternalStartPort)
}

func (c *createPortForwardingRule) GetInternalEndPort() int {
    nInternalEndPort, _ := strconv.Atoi(c.req["InternalEndPort"])
    return nInternalEndPort
}

func (c *createPortForwardingRule) SetInternalEndPort(varInternalEndPort int) {
    c.req["InternalEndPort"] = strconv.Itoa(varInternalEndPort)
}

func (c *createPortForwardingRule) GetLocalIpAddress() string {
    return c.req["LocalIpAddress"]
}

func (c *createPortForwardingRule) SetLocalIpAddress(varLocalIpAddress string) {
    c.req["LocalIpAddress"] = varLocalIpAddress
}



