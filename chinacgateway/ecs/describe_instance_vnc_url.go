package chinacgatewayecs

import (
)

type describeInstanceVncUrl struct {
    req        	map[string]string
}

func DescribeInstanceVncUrl() *describeInstanceVncUrl {
    c := &describeInstanceVncUrl {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeInstanceVncUrl"
    return c
}

func (c *describeInstanceVncUrl) GetRequest() map[string]string {
    return c.req
}

func (c *describeInstanceVncUrl) GetRegion() string {
    return c.req["Region"]
}

func (c *describeInstanceVncUrl) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeInstanceVncUrl) GetAction() string {
    return c.req["Action"]
}

func (c *describeInstanceVncUrl) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeInstanceVncUrl) GetId() string {
    return c.req["Id"]
}

func (c *describeInstanceVncUrl) SetId(varId string) {
    c.req["Id"] = varId
}


