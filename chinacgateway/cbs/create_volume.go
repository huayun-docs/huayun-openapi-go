package chinacgatewaycbs

import (
    "strconv"
)

type createVolume struct {
    req        	map[string]string
}

func CreateVolume() *createVolume {
    c := &createVolume {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateVolume"
    return c
}

func (c *createVolume) GetRequest() map[string]string {
    return c.req
}

func (c *createVolume) GetAction() string {
    return c.req["Action"]
}

func (c *createVolume) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createVolume) GetRegion() string {
    return c.req["Region"]
}

func (c *createVolume) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createVolume) GetPayType() string {
    return c.req["PayType"]
}

func (c *createVolume) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *createVolume) GetPeriod() int {
    nPeriod, _ := strconv.Atoi(c.req["Period"])
    return nPeriod
}

func (c *createVolume) SetPeriod(varPeriod int) {
    c.req["Period"] = strconv.Itoa(varPeriod)
}

func (c *createVolume) GetSize() int {
    nSize, _ := strconv.Atoi(c.req["Size"])
    return nSize
}

func (c *createVolume) SetSize(varSize int) {
    c.req["Size"] = strconv.Itoa(varSize)
}

func (c *createVolume) GetType() string {
    return c.req["Type"]
}

func (c *createVolume) SetType(varType string) {
    c.req["Type"] = varType
}

func (c *createVolume) GetSnapshotId() string {
    return c.req["SnapshotId"]
}

func (c *createVolume) SetSnapshotId(varSnapshotId string) {
    c.req["SnapshotId"] = varSnapshotId
}

func (c *createVolume) GetSourceVolumeId() string {
    return c.req["SourceVolumeId"]
}

func (c *createVolume) SetSourceVolumeId(varSourceVolumeId string) {
    c.req["SourceVolumeId"] = varSourceVolumeId
}

func (c *createVolume) GetCount() int {
    nCount, _ := strconv.Atoi(c.req["Count"])
    return nCount
}

func (c *createVolume) SetCount(varCount int) {
    c.req["Count"] = strconv.Itoa(varCount)
}

func (c *createVolume) GetName() string {
    return c.req["Name"]
}

func (c *createVolume) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createVolume) GetDescription() string {
    return c.req["Description"]
}

func (c *createVolume) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


