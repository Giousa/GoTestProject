package com.chdspine.system;


/**
 * Description:
 * Author:张蒙蒙
 * Date:2021-01-12
 */
@RestController
@RequestMapping("sys-clause")
public class SysClauseController {

	@Autowired
    private SysClauseService sysClauseService;

    @PostMapping("addSysClause")
    public ResultVO addSysClause(@RequestBody SysClause sysClause) {
        return sysClauseService.addSysClause(sysClause);
    }

	@PostMapping("updateSysClause")
    public ResultVO updateSysClause(@RequestBody SysClause sysClause) {
        return sysClauseService.updateSysClause(sysClause);
    }

    @GetMapping("findSysClauseById")
    public ResultVO findSysClauseById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return sysClauseService.findSysClauseById(id);
    }

    @GetMapping("deleteSysClauseById")
    public ResultVO deleteSysClauseById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return sysClauseService.deleteSysClauseById(id);
    }

 	@GetMapping("findSysClauseListByPage")
    public ResultVO findSysClauseListByPage(
                                    @RequestParam(value = "page",required = false,defaultValue = "1") int page,
                                    @RequestParam(value = "size",required = false,defaultValue = "10") int size) {
        return sysClauseService.findSysClauseListByPage(page,size);
    }

}
