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



// DasNovelInfo ...
//type:  武侠古典
type DasNovelInfo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Type string `json:"type"`
	CreateTime time.Time `json:"create_time"`
}

// DasNovelDetl ...
type DasNovelDetl struct {
	Id int `json:"id"`
	NovelId int `json:"novel_id"`
	SubTitle string `json:"sub_title"`
	SubUrl string `json:"sub_url"`
	Content string `json:"content"`
	CreateTime time.Time `json:"create_time"`
}
