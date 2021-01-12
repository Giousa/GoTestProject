package com.giousa.sys;


/**
 * Description:
 * Author:张蒙蒙
 * Date:2021-01-12
 */
@RestController
@RequestMapping("register")
public class RegisterController {

	@Autowired
    private RegisterService registerService;

    @PostMapping("addRegister")
    public ResultVO addRegister(@RequestBody Register register) {
        return registerService.addRegister(register);
    }

	@PostMapping("updateRegister")
    public ResultVO updateRegister(@RequestBody Register register) {
        return registerService.updateRegister(register);
    }

    @GetMapping("findRegisterById")
    public ResultVO findRegisterById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return registerService.findRegisterById(id);
    }

    @GetMapping("deleteRegisterById")
    public ResultVO deleteRegisterById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return registerService.deleteRegisterById(id);
    }

 	@GetMapping("findRegisterListByPage")
    public ResultVO findRegisterListByPage(
                                    @RequestParam(value = "page",required = false,defaultValue = "1") int page,
                                    @RequestParam(value = "size",required = false,defaultValue = "10") int size) {
        return registerService.findRegisterListByPage(page,size);
    }

}
