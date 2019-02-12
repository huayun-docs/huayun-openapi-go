package chinacgatewayecs

import (
)

type listFlavor struct {
    req        	map[string]string
}

func ListFlavor() *listFlavor {
    c := &listFlavor {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ListFlavor"
    return c
}

func (c *listFlavor) GetRequest() map[string]string {
    return c.req
}

func (c *listFlavor) GetAction() string {
    return c.req["Action"]
}

func (c *listFlavor) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *listFlavor) GetRegion() string {
    return c.req["Region"]
}

func (c *listFlavor) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}


