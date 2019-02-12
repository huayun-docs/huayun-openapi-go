package chinacgatewayvpc

import (
)

type describeRouters struct {
    req        	map[string]string
}

func DescribeRouters() *describeRouters {
    c := &describeRouters {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeRouters"
    return c
}

func (c *describeRouters) GetRequest() map[string]string {
    return c.req
}

func (c *describeRouters) GetAction() string {
    return c.req["Action"]
}

func (c *describeRouters) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeRouters) GetRegion() string {
    return c.req["Region"]
}

func (c *describeRouters) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeRouters) GetOffset() string {
    return c.req["Offset"]
}

func (c *describeRouters) SetOffset(varOffset string) {
    c.req["Offset"] = varOffset
}

func (c *describeRouters) GetCount() string {
    return c.req["Count"]
}

func (c *describeRouters) SetCount(varCount string) {
    c.req["Count"] = varCount
}

func (c *describeRouters) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeRouters) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeRouters) GetName0() string {
    return c.req["Name.0"]
}

func (c *describeRouters) SetName0(varName0 string) {
    c.req["Name.0"] = varName0
}

func (c *describeRouters) GetEip() string {
    return c.req["Eip"]
}

func (c *describeRouters) SetEip(varEip string) {
    c.req["Eip"] = varEip
}

func (c *describeRouters) GetStartTime() string {
    return c.req["StartTime"]
}

func (c *describeRouters) SetStartTime(varStartTime string) {
    c.req["StartTime"] = varStartTime
}

func (c *describeRouters) GetEndTime() string {
    return c.req["EndTime"]
}

func (c *describeRouters) SetEndTime(varEndTime string) {
    c.req["EndTime"] = varEndTime
}

func (c *describeRouters) GetProductStatus() string {
    return c.req["ProductStatus"]
}

func (c *describeRouters) SetProductStatus(varProductStatus string) {
    c.req["ProductStatus"] = varProductStatus
}

func (c *describeRouters) GetPayType() string {
    return c.req["PayType"]
}

func (c *describeRouters) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}


