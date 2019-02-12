package chinacgatewayecs

import (
    "strconv"
)

type describeInstanceSnapshots struct {
    req        	map[string]string
}

func DescribeInstanceSnapshots() *describeInstanceSnapshots {
    c := &describeInstanceSnapshots {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeInstanceSnapshots"
    return c
}

func (c *describeInstanceSnapshots) GetRequest() map[string]string {
    return c.req
}

func (c *describeInstanceSnapshots) GetAction() string {
    return c.req["Action"]
}

func (c *describeInstanceSnapshots) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeInstanceSnapshots) GetRegion() string {
    return c.req["Region"]
}

func (c *describeInstanceSnapshots) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeInstanceSnapshots) GetInstanceId0() string {
    return c.req["InstanceId.0"]
}

func (c *describeInstanceSnapshots) SetInstanceId0(varInstanceId0 string) {
    c.req["InstanceId.0"] = varInstanceId0
}

func (c *describeInstanceSnapshots) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeInstanceSnapshots) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeInstanceSnapshots) GetName() string {
    return c.req["Name"]
}

func (c *describeInstanceSnapshots) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *describeInstanceSnapshots) GetStartTime() string {
    return c.req["StartTime"]
}

func (c *describeInstanceSnapshots) SetStartTime(varStartTime string) {
    c.req["StartTime"] = varStartTime
}

func (c *describeInstanceSnapshots) GetEndTime() string {
    return c.req["EndTime"]
}

func (c *describeInstanceSnapshots) SetEndTime(varEndTime string) {
    c.req["EndTime"] = varEndTime
}

func (c *describeInstanceSnapshots) GetOffset() int {
    nOffset, _ := strconv.Atoi(c.req["Offset"])
    return nOffset
}

func (c *describeInstanceSnapshots) SetOffset(varOffset int) {
    c.req["Offset"] = strconv.Itoa(varOffset)
}

func (c *describeInstanceSnapshots) GetCount() int {
    nCount, _ := strconv.Atoi(c.req["Count"])
    return nCount
}

func (c *describeInstanceSnapshots) SetCount(varCount int) {
    c.req["Count"] = strconv.Itoa(varCount)
}

func (c *describeInstanceSnapshots) GetStatus() string {
    return c.req["Status"]
}

func (c *describeInstanceSnapshots) SetStatus(varStatus string) {
    c.req["Status"] = varStatus
}


