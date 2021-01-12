package com.giousa.sys.service;



/**
 * Description:
 * Author:
 * Date:
 */
public interface LoginInfoService {

    ResultVO addLoginInfo(LoginInfo loginInfo);

    ResultVO updateLoginInfo(LoginInfo loginInfo);

    ResultVO findLoginInfoById(String id);

    ResultVO deleteLoginInfoById(String id);

    ResultVO findLoginInfoListByPage(int page,int size);
}

