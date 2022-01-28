package sap_api_patch_header_setup

import (
	"fmt"
	"testing"
)

type option struct {
}

func (o option) User() string {
	return "XXXXXXXXX"
}

func (o option) Pass() string {
	return "XXXXXXXXX"
}

func (o option) RefreshTokenURL() string {
	return "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

func (o option) RetryMax() int {
	return 1
}

func (o option) RetryInterval() int {
	return 1000
}

func Test_a(t *testing.T) {
	c := NewSAPpatchClientWithOption(option{})
	res, err := c.patch(
		"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		map[string]string{
			"sap-client": "XXX",
		},
		`
{
	"Product": "",
	"IndustrySector": "M",
	"ProductType": "FERT",
	"BaseUnit": "PC",
	"to_Description": {
		"results": [
			{
				"Product": "",
				"Language": "EN",
				"ProductDescription": "Test Material"
			}
		]
	}
}`,
	)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", res.Body)
}
