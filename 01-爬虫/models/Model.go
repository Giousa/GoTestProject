/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/4
 */
package models

import "time"

// DasKtdmDetl ...
type DasKtdmDetl struct {
	Id int `json:"id"`
	KtdmId int `json:"ktdm_id"`
	PicUrl string `json:"pic_url"`
	Page int `json:"page"`
	CreateTime time.Time `json:"create_time"`
}

// DasKtdmInfo ...
type DasKtdmInfo struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Title string `json:"title"`
	CreateTime time.Time `json:"create_time"`
}


