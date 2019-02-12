package chinacgatewayecs

import (
    "strconv"
)

type runInstance struct {
    req        	map[string]string
}

func RunInstance() *runInstance {
    c := &runInstance {
        req:			make(map[string]string),
    }

    c.req["Action"] = "RunInstance"
    return c
}

func (c *runInstance) GetRequest() map[string]string {
    return c.req
}

func (c *runInstance) GetAction() string {
    return c.req["Action"]
}

func (c *runInstance) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *runInstance) GetRegion() string {
    return c.req["Region"]
}

func (c *runInstance) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *runInstance) GetPayType() string {
    return c.req["PayType"]
}

func (c *runInstance) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *runInstance) GetPeriod() int {
    nPeriod, _ := strconv.Atoi(c.req["Period"])
    return nPeriod
}

func (c *runInstance) SetPeriod(varPeriod int) {
    c.req["Period"] = strconv.Itoa(varPeriod)
}

func (c *runInstance) GetInstanceSeries() string {
    return c.req["InstanceSeries"]
}

func (c *runInstance) SetInstanceSeries(varInstanceSeries string) {
    c.req["InstanceSeries"] = varInstanceSeries
}

func (c *runInstance) GetImageId() string {
    return c.req["ImageId"]
}

func (c *runInstance) SetImageId(varImageId string) {
    c.req["ImageId"] = varImageId
}

func (c *runInstance) GetInstanceType() string {
    return c.req["InstanceType"]
}

func (c *runInstance) SetInstanceType(varInstanceType string) {
    c.req["InstanceType"] = varInstanceType
}

func (c *runInstance) GetInterface0NetworkId() string {
    return c.req["Interface.0.NetworkId"]
}

func (c *runInstance) SetInterface0NetworkId(varInterface0NetworkId string) {
    c.req["Interface.0.NetworkId"] = varInterface0NetworkId
}

func (c *runInstance) GetVolumes0Type() string {
    return c.req["Volumes.0.Type"]
}

func (c *runInstance) SetVolumes0Type(varVolumes0Type string) {
    c.req["Volumes.0.Type"] = varVolumes0Type
}

func (c *runInstance) GetVolumes0Size() int {
    nVolumes0Size, _ := strconv.Atoi(c.req["Volumes.0.Size"])
    return nVolumes0Size
}

func (c *runInstance) SetVolumes0Size(varVolumes0Size int) {
    c.req["Volumes.0.Size"] = strconv.Itoa(varVolumes0Size)
}

func (c *runInstance) GetVolumes1Type() string {
    return c.req["Volumes.1.Type"]
}

func (c *runInstance) SetVolumes1Type(varVolumes1Type string) {
    c.req["Volumes.1.Type"] = varVolumes1Type
}

func (c *runInstance) GetVolumes1Size() int {
    nVolumes1Size, _ := strconv.Atoi(c.req["Volumes.1.Size"])
    return nVolumes1Size
}

func (c *runInstance) SetVolumes1Size(varVolumes1Size int) {
    c.req["Volumes.1.Size"] = strconv.Itoa(varVolumes1Size)
}

func (c *runInstance) GetVolumes2Type() string {
    return c.req["Volumes.2.Type"]
}

func (c *runInstance) SetVolumes2Type(varVolumes2Type string) {
    c.req["Volumes.2.Type"] = varVolumes2Type
}

func (c *runInstance) GetVolumes2Size() int {
    nVolumes2Size, _ := strconv.Atoi(c.req["Volumes.2.Size"])
    return nVolumes2Size
}

func (c *runInstance) SetVolumes2Size(varVolumes2Size int) {
    c.req["Volumes.2.Size"] = strconv.Itoa(varVolumes2Size)
}

func (c *runInstance) GetVolumes3Type() string {
    return c.req["Volumes.3.Type"]
}

func (c *runInstance) SetVolumes3Type(varVolumes3Type string) {
    c.req["Volumes.3.Type"] = varVolumes3Type
}

func (c *runInstance) GetVolumes3Size() int {
    nVolumes3Size, _ := strconv.Atoi(c.req["Volumes.3.Size"])
    return nVolumes3Size
}

func (c *runInstance) SetVolumes3Size(varVolumes3Size int) {
    c.req["Volumes.3.Size"] = strconv.Itoa(varVolumes3Size)
}

func (c *runInstance) GetVolumes4Type() string {
    return c.req["Volumes.4.Type"]
}

func (c *runInstance) SetVolumes4Type(varVolumes4Type string) {
    c.req["Volumes.4.Type"] = varVolumes4Type
}

func (c *runInstance) GetVolumes4Size() int {
    nVolumes4Size, _ := strconv.Atoi(c.req["Volumes.4.Size"])
    return nVolumes4Size
}

func (c *runInstance) SetVolumes4Size(varVolumes4Size int) {
    c.req["Volumes.4.Size"] = strconv.Itoa(varVolumes4Size)
}

func (c *runInstance) GetInternetBandwidth() string {
    return c.req["Internet.Bandwidth"]
}

func (c *runInstance) SetInternetBandwidth(varInternetBandwidth string) {
    c.req["Internet.Bandwidth"] = varInternetBandwidth
}

func (c *runInstance) GetName() string {
    return c.req["Name"]
}

func (c *runInstance) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *runInstance) GetPassword() string {
    return c.req["Password"]
}

func (c *runInstance) SetPassword(varPassword string) {
    c.req["Password"] = varPassword
}

func (c *runInstance) GetKeyPair() string {
    return c.req["KeyPair"]
}

func (c *runInstance) SetKeyPair(varKeyPair string) {
    c.req["KeyPair"] = varKeyPair
}

func (c *runInstance) GetFirewallId() string {
    return c.req["FirewallId"]
}

func (c *runInstance) SetFirewallId(varFirewallId string) {
    c.req["FirewallId"] = varFirewallId
}

func (c *runInstance) GetInternetRouterId() string {
    return c.req["Internet.RouterId"]
}

func (c *runInstance) SetInternetRouterId(varInternetRouterId string) {
    c.req["Internet.RouterId"] = varInternetRouterId
}

func (c *runInstance) GetVolumeId() string {
    return c.req["VolumeId"]
}

func (c *runInstance) SetVolumeId(varVolumeId string) {
    c.req["VolumeId"] = varVolumeId
}


