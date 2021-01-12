package com.giousa.wx.service;



/**
 * Description:
 * Author:
 * Date:
 */
public interface EducInfoService {

    ResultVO addEducInfo(EducInfo educInfo);

    ResultVO updateEducInfo(EducInfo educInfo);

    ResultVO findEducInfoById(String id);

    ResultVO deleteEducInfoById(String id);

    ResultVO findEducInfoListByPage(int page,int size);
}

