package chinacgatewaycbs

import (
)

type describeVolumeType struct {
    req        	map[string]string
}

func DescribeVolumeType() *describeVolumeType {
    c := &describeVolumeType {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeVolumeType"
    return c
}

func (c *describeVolumeType) GetRequest() map[string]string {
    return c.req
}

func (c *describeVolumeType) GetAction() string {
    return c.req["Action"]
}

func (c *describeVolumeType) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeVolumeType) GetRegion() string {
    return c.req["Region"]
}

func (c *describeVolumeType) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}


