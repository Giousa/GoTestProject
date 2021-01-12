package com.giousa.wx.service;



/**
 * Description:
 * Author:
 * Date:
 */
public interface LoginService {

    ResultVO addLogin(Login login);

    ResultVO updateLogin(Login login);

    ResultVO findLoginById(String id);

    ResultVO deleteLoginById(String id);

    ResultVO findLoginListByPage(int page,int size);
}

