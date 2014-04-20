package main

import (
	"fmt"
)

type Template struct {
	ReqId   string  `json:"reqid"`
	Name    string  `json:"templatename"`
	Content Content `json:"content"`
}

type Content struct {
	VmList      VmList      `json:"vmlist"`
	NetworkList NetworkList `json:"networks"`
	Flavor      Flavor      `json:"flavors"`
}

type VmList []struct {
	VmName   string `json:"vmname"`
	VmType   string `json:"vmtype"`
	MemCap   string `json:"memcap"`
	MaxCount int    `json:"maxcount"`
	MinCount int    `json:"mincount"`
}

type NetworkList []struct {
	NetName string `json:"netname"`
	NetType string `json:"nettype"`
	NetMask string `json:"netmask"`
	SubNet  string `json: "subnet"`
}

type Flavor []struct {
	FlavorName string `json:"flavorname"`
	FlavorType string `json:"flavortype"`
	FlavorRef  string `json:"flavorref"`
	FlavorMem  int    `json:"flavormem"`
}

func GeneratVmList() VmList {
	return VmList{{
		"cmp",
		"small",
		"64M",
		1,
		5},
		{
			"dmp",
			"small",
			"64M",
			1,
			5}}
}

func GeneratNetworkList() NetworkList {
	return NetworkList{{
		"controlplane",
		"lcp",
		"255.255.255.0",
		"192.168.1.1"},
		{
			"dataplane",
			"dcp",
			"255.255.255.0",
			"192.168.2.1"}}
}

func GeneratFlavor() Flavor {
	return Flavor{{
		"sirors",
		"simple",
		"/usri/idf/qqq",
		64},
		{
			"sirors",
			"simple",
			"/usri/idf/fdsfdsfdsfdsfds",
			128}}
}

func CreateSampleTemplate(name string) (t Template) {
	t.ReqId = "4"
	t.Name = name
	t.Content.VmList = GeneratVmList()
	t.Content.NetworkList = GeneratNetworkList()
	t.Content.Flavor = GeneratFlavor()
	return t
}

func CreateVappFromTemplate(t Template) {

}

// The Album data structure, serializable in JSON, XML and text using the Stringer interface.
type TemplateResource struct {
	Id          string `json:"id" `
	Name        string `json:"name" `
	Content     string `json:"content" `
	TemplateObj *Template
	Index       int
}

func (ts *TemplateResource) String() string {
	return fmt.Sprintf("%s - %s (%d)", ts.Id, ts.Name, ts.Content)
}

func CreateTemplateResource(reqContent string, t *Template) *TemplateResource {
	return &TemplateResource{t.ReqId, t.Name, reqContent, t, 0}
}
