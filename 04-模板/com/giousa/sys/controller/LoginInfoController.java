package com.giousa.sys;


/**
 * Description:
 * Author:张蒙蒙
 * Date:2021-01-12
 */
@RestController
@RequestMapping("login-info")
public class LoginInfoController {

	@Autowired
    private LoginInfoService loginInfoService;

    @PostMapping("addLoginInfo")
    public ResultVO addLoginInfo(@RequestBody LoginInfo loginInfo) {
        return loginInfoService.addLoginInfo(loginInfo);
    }

	@PostMapping("updateLoginInfo")
    public ResultVO updateLoginInfo(@RequestBody LoginInfo loginInfo) {
        return loginInfoService.updateLoginInfo(loginInfo);
    }

    @GetMapping("findLoginInfoById")
    public ResultVO findLoginInfoById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return loginInfoService.findLoginInfoById(id);
    }

    @GetMapping("deleteLoginInfoById")
    public ResultVO deleteLoginInfoById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return loginInfoService.deleteLoginInfoById(id);
    }

 	@GetMapping("findLoginInfoListByPage")
    public ResultVO findLoginInfoListByPage(
                                    @RequestParam(value = "page",required = false,defaultValue = "1") int page,
                                    @RequestParam(value = "size",required = false,defaultValue = "10") int size) {
        return loginInfoService.findLoginInfoListByPage(page,size);
    }

}
