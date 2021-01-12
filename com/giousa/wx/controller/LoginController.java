package com.giousa.wx;


/**
 * Description:
 * Author:张蒙蒙
 * Date:2021-01-12
 */
@RestController
@RequestMapping("login")
public class LoginController {

	@Autowired
    private LoginService loginService;

    @PostMapping("addLogin")
    public ResultVO addLogin(@RequestBody Login login) {
        return loginService.addLogin(login);
    }

	@PostMapping("updateLogin")
    public ResultVO updateLogin(@RequestBody Login login) {
        return loginService.updateLogin(login);
    }

    @GetMapping("findLoginById")
    public ResultVO findLoginById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return loginService.findLoginById(id);
    }

    @GetMapping("deleteLoginById")
    public ResultVO deleteLoginById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return loginService.deleteLoginById(id);
    }

 	@GetMapping("findLoginListByPage")
    public ResultVO findLoginListByPage(
                                    @RequestParam(value = "page",required = false,defaultValue = "1") int page,
                                    @RequestParam(value = "size",required = false,defaultValue = "10") int size) {
        return loginService.findLoginListByPage(page,size);
    }

}
