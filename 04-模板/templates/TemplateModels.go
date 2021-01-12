/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/12
 */
package templates

var TextController =
`package {{.Package}};


/**
 * Description:
 * Author:{{.Author}}
 * Date:{{.DateTime}}
 */
@RestController
@RequestMapping("{{.NameSeparate}}")
public class {{.Name}}Controller {

	@Autowired
    private {{.Name}}Service {{.NameHumpLower}}Service;

    @PostMapping("add{{.Name}}")
    public ResultVO add{{.Name}}(@RequestBody {{.Name}} {{.NameHumpLower}}) {
        return {{.NameHumpLower}}Service.add{{.Name}}({{.NameHumpLower}});
    }

	@PostMapping("update{{.Name}}")
    public ResultVO update{{.Name}}(@RequestBody {{.Name}} {{.NameHumpLower}}) {
        return {{.NameHumpLower}}Service.update{{.Name}}({{.NameHumpLower}});
    }

    @GetMapping("find{{.Name}}ById")
    public ResultVO find{{.Name}}ById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return {{.NameHumpLower}}Service.find{{.Name}}ById(id);
    }

    @GetMapping("delete{{.Name}}ById")
    public ResultVO delete{{.Name}}ById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return {{.NameHumpLower}}Service.delete{{.Name}}ById(id);
    }

 	@GetMapping("find{{.Name}}ListByPage")
    public ResultVO find{{.Name}}ListByPage(
                                    @RequestParam(value = "page",required = false,defaultValue = "1") int page,
                                    @RequestParam(value = "size",required = false,defaultValue = "10") int size) {
        return {{.NameHumpLower}}Service.find{{.Name}}ListByPage(page,size);
    }

}
`

var TextService =
	`package {{.Package}}.service;



/**
 * Description:
 * Author:{{.Author}}
 * Date:{{.DateTime}}
 */
public interface {{.Name}}Service {

    ResultVO add{{.Name}}({{.Name}} {{.NameHumpLower}});

    ResultVO update{{.Name}}({{.Name}} {{.NameHumpLower}});

    ResultVO find{{.Name}}ById(String id);

    ResultVO delete{{.Name}}ById(String id);

    ResultVO find{{.Name}}ListByPage(int page,int size);
}

`

var TextServiceImpl =
	`package {{.Package}}.service.impl;



/**
 * Description:
 * Author:{{.Author}}
 * Date:{{.DateTime}}
 */
@Service
public class {{.Name}}ServiceImpl implements {{.Name}}Service {

    @Autowired
    private {{.Name}}Mapper {{.NameHumpLower}}Mapper;

    @Override
    public ResultVO add{{.Name}}({{.Name}} {{.NameHumpLower}}) {

		return null;
    }

	@Override
    public ResultVO update{{.Name}}({{.Name}} {{.NameHumpLower}}) {
		
		return null;
    }

    @Override
    public ResultVO find{{.Name}}ById(String id) {

		return null;
    }

    @Override
    public ResultVO delete{{.Name}}ById(String id) {
        
		return null;
    }

    @Override
    public ResultVO find{{.Name}}ListByPage(int page,int size) {

		return null;
    }
}

`