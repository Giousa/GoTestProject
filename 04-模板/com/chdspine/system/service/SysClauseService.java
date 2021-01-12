package com.chdspine.system.service;



/**
 * Description:
 * Author:
 * Date:
 */
public interface SysClauseService {

    ResultVO addSysClause(SysClause sysClause);

    ResultVO updateSysClause(SysClause sysClause);

    ResultVO findSysClauseById(String id);

    ResultVO deleteSysClauseById(String id);

    ResultVO findSysClauseListByPage(int page,int size);
}

