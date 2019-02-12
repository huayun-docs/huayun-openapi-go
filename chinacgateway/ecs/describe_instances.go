package chinacgatewayecs

import (
    "strconv"
)

type describeInstances struct {
    req        	map[string]string
}

func DescribeInstances() *describeInstances {
    c := &describeInstances {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeInstances"
    return c
}

func (c *describeInstances) GetRequest() map[string]string {
    return c.req
}

func (c *describeInstances) GetRegion() string {
    return c.req["Region"]
}

func (c *describeInstances) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeInstances) GetAction() string {
    return c.req["Action"]
}

func (c *describeInstances) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeInstances) GetEndTime() string {
    return c.req["EndTime"]
}

func (c *describeInstances) SetEndTime(varEndTime string) {
    c.req["EndTime"] = varEndTime
}

func (c *describeInstances) GetDueStartTime() string {
    return c.req["DueStartTime"]
}

func (c *describeInstances) SetDueStartTime(varDueStartTime string) {
    c.req["DueStartTime"] = varDueStartTime
}

func (c *describeInstances) GetDueEndTime() string {
    return c.req["DueEndTime"]
}

func (c *describeInstances) SetDueEndTime(varDueEndTime string) {
    c.req["DueEndTime"] = varDueEndTime
}

func (c *describeInstances) GetProductStatus() string {
    return c.req["ProductStatus"]
}

func (c *describeInstances) SetProductStatus(varProductStatus string) {
    c.req["ProductStatus"] = varProductStatus
}

func (c *describeInstances) GetPayType() string {
    return c.req["PayType"]
}

func (c *describeInstances) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *describeInstances) GetImageId() string {
    return c.req["ImageId"]
}

func (c *describeInstances) SetImageId(varImageId string) {
    c.req["ImageId"] = varImageId
}

func (c *describeInstances) GetOffset() int {
    nOffset, _ := strconv.Atoi(c.req["Offset"])
    return nOffset
}

func (c *describeInstances) SetOffset(varOffset int) {
    c.req["Offset"] = strconv.Itoa(varOffset)
}

func (c *describeInstances) GetCount() int {
    nCount, _ := strconv.Atoi(c.req["Count"])
    return nCount
}

func (c *describeInstances) SetCount(varCount int) {
    c.req["Count"] = strconv.Itoa(varCount)
}

func (c *describeInstances) GetInstanceSeries() string {
    return c.req["InstanceSeries"]
}

func (c *describeInstances) SetInstanceSeries(varInstanceSeries string) {
    c.req["InstanceSeries"] = varInstanceSeries
}

func (c *describeInstances) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeInstances) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeInstances) GetName() string {
    return c.req["Name"]
}

func (c *describeInstances) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *describeInstances) GetStatus() string {
    return c.req["Status"]
}

func (c *describeInstances) SetStatus(varStatus string) {
    c.req["Status"] = varStatus
}

func (c *describeInstances) GetFirewallId0() string {
    return c.req["FirewallId.0"]
}

func (c *describeInstances) SetFirewallId0(varFirewallId0 string) {
    c.req["FirewallId.0"] = varFirewallId0
}

func (c *describeInstances) GetStartTime() string {
    return c.req["StartTime"]
}

func (c *describeInstances) SetStartTime(varStartTime string) {
    c.req["StartTime"] = varStartTime
}


