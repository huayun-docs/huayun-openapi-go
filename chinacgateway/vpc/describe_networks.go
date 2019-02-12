package chinacgatewayvpc

import (
    "strconv"
)

type describeNetworks struct {
    req        	map[string]string
}

func DescribeNetworks() *describeNetworks {
    c := &describeNetworks {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeNetworks"
    return c
}

func (c *describeNetworks) GetRequest() map[string]string {
    return c.req
}

func (c *describeNetworks) GetAction() string {
    return c.req["Action"]
}

func (c *describeNetworks) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeNetworks) GetRegion() string {
    return c.req["Region"]
}

func (c *describeNetworks) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeNetworks) GetOffset() int {
    nOffset, _ := strconv.Atoi(c.req["Offset"])
    return nOffset
}

func (c *describeNetworks) SetOffset(varOffset int) {
    c.req["Offset"] = strconv.Itoa(varOffset)
}

func (c *describeNetworks) GetCount() int {
    nCount, _ := strconv.Atoi(c.req["Count"])
    return nCount
}

func (c *describeNetworks) SetCount(varCount int) {
    c.req["Count"] = strconv.Itoa(varCount)
}

func (c *describeNetworks) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeNetworks) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeNetworks) GetName() string {
    return c.req["Name"]
}

func (c *describeNetworks) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *describeNetworks) GetStartTime() string {
    return c.req["StartTime"]
}

func (c *describeNetworks) SetStartTime(varStartTime string) {
    c.req["StartTime"] = varStartTime
}

func (c *describeNetworks) GetEndTime() string {
    return c.req["EndTime"]
}

func (c *describeNetworks) SetEndTime(varEndTime string) {
    c.req["EndTime"] = varEndTime
}


