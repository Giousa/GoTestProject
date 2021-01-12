package com.giousa.wx;


/**
 * Description:
 * Author:张蒙蒙
 * Date:2021-01-12
 */
@RestController
@RequestMapping("educ-info")
public class EducInfoController {

	@Autowired
    private EducInfoService educInfoService;

    @PostMapping("addEducInfo")
    public ResultVO addEducInfo(@RequestBody EducInfo educInfo) {
        return educInfoService.addEducInfo(educInfo);
    }

	@PostMapping("updateEducInfo")
    public ResultVO updateEducInfo(@RequestBody EducInfo educInfo) {
        return educInfoService.updateEducInfo(educInfo);
    }

    @GetMapping("findEducInfoById")
    public ResultVO findEducInfoById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return educInfoService.findEducInfoById(id);
    }

    @GetMapping("deleteEducInfoById")
    public ResultVO deleteEducInfoById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return educInfoService.deleteEducInfoById(id);
    }

 	@GetMapping("findEducInfoListByPage")
    public ResultVO findEducInfoListByPage(
                                    @RequestParam(value = "page",required = false,defaultValue = "1") int page,
                                    @RequestParam(value = "size",required = false,defaultValue = "10") int size) {
        return educInfoService.findEducInfoListByPage(page,size);
    }

}
