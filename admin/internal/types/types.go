// Code generated by goctl. DO NOT EDIT.
package types

type DeviceListRequest struct {
	Page int    `josn:"page,optional"`
	Size int    `json:"size,optional"`
	Name string `json:"name,optional"`
}

type DeviceListBaisc struct {
	Identity    string `json:"identity"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Secret      string `json:"secret"`
	ProductName string `json:"product_name"`
}

type DeviceListResponse struct {
	List  []*DeviceListBaisc `json:"list"`
	Count int64              `json:"count"`
}
