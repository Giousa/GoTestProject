package com.giousa.sys.service;



/**
 * Description:
 * Author:
 * Date:
 */
public interface RegisterService {

    ResultVO addRegister(Register register);

    ResultVO updateRegister(Register register);

    ResultVO findRegisterById(String id);

    ResultVO deleteRegisterById(String id);

    ResultVO findRegisterListByPage(int page,int size);
}

