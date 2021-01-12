// das_ktdm_detl...
type DasKtdmDetl struct {
    Id int `json:id`
    KtdmId int `json:ktdm_id`
    PicUrl string `json:pic_url`
    Page int `json:page`
    CreateTime time.Time `json:create_time`
}



// das_ktdm_info...
type DasKtdmInfo struct {
    Id int `json:id`
    Url string `json:url`
    Title string `json:title`
    CreateTime time.Time `json:create_time`
}



// das_novel_detl...
type DasNovelDetl struct {
    Id int `json:id`
    NovelId int `json:novel_id`
    SubTitle string `json:sub_title`
    Content string `json:content`
    SubUrl string `json:sub_url`
    CreateTime time.Time `json:create_time`
}



// das_novel_info...
type DasNovelInfo struct {
    Id int `json:id`
    Title string `json:title`
    Url string `json:url`
    Type string `json:type`
    CreateTime time.Time `json:create_time`
}



