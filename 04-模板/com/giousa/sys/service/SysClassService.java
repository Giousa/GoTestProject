package com.giousa.sys.service;



/**
 * Description:
 * Author:
 * Date:
 */
public interface SysClassService {

    ResultVO addSysClass(SysClass sysClass);

    ResultVO updateSysClass(SysClass sysClass);

    ResultVO findSysClassById(String id);

    ResultVO deleteSysClassById(String id);

    ResultVO findSysClassListByPage(int page,int size);
}

