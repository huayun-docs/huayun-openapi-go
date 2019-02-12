package chinacgatewayvpc

import (
    "strconv"
)

type describeFirewalls struct {
    req        	map[string]string
}

func DescribeFirewalls() *describeFirewalls {
    c := &describeFirewalls {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeFirewalls"
    return c
}

func (c *describeFirewalls) GetRequest() map[string]string {
    return c.req
}

func (c *describeFirewalls) GetAction() string {
    return c.req["Action"]
}

func (c *describeFirewalls) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeFirewalls) GetRegion() string {
    return c.req["Region"]
}

func (c *describeFirewalls) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeFirewalls) GetOffset() int {
    nOffset, _ := strconv.Atoi(c.req["Offset"])
    return nOffset
}

func (c *describeFirewalls) SetOffset(varOffset int) {
    c.req["Offset"] = strconv.Itoa(varOffset)
}

func (c *describeFirewalls) GetCount() int {
    nCount, _ := strconv.Atoi(c.req["Count"])
    return nCount
}

func (c *describeFirewalls) SetCount(varCount int) {
    c.req["Count"] = strconv.Itoa(varCount)
}

func (c *describeFirewalls) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeFirewalls) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeFirewalls) GetName() string {
    return c.req["Name"]
}

func (c *describeFirewalls) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *describeFirewalls) GetStartTime() string {
    return c.req["StartTime"]
}

func (c *describeFirewalls) SetStartTime(varStartTime string) {
    c.req["StartTime"] = varStartTime
}

func (c *describeFirewalls) GetEndTime() string {
    return c.req["EndTime"]
}

func (c *describeFirewalls) SetEndTime(varEndTime string) {
    c.req["EndTime"] = varEndTime
}


