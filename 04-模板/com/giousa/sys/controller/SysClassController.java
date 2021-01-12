package com.giousa.sys;


/**
 * Description:
 * Author:张蒙蒙
 * Date:2021-01-12
 */
@RestController
@RequestMapping("sys-class")
public class SysClassController {

	@Autowired
    private SysClassService sysClassService;

    @PostMapping("addSysClass")
    public ResultVO addSysClass(@RequestBody SysClass sysClass) {
        return sysClassService.addSysClass(sysClass);
    }

	@PostMapping("updateSysClass")
    public ResultVO updateSysClass(@RequestBody SysClass sysClass) {
        return sysClassService.updateSysClass(sysClass);
    }

    @GetMapping("findSysClassById")
    public ResultVO findSysClassById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return sysClassService.findSysClassById(id);
    }

    @GetMapping("deleteSysClassById")
    public ResultVO deleteSysClassById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return sysClassService.deleteSysClassById(id);
    }

 	@GetMapping("findSysClassListByPage")
    public ResultVO findSysClassListByPage(
                                    @RequestParam(value = "page",required = false,defaultValue = "1") int page,
                                    @RequestParam(value = "size",required = false,defaultValue = "10") int size) {
        return sysClassService.findSysClassListByPage(page,size);
    }

}
