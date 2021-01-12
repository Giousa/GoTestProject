/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/12
 */
package templates

var TextMysqlTemplate =
`
{{- range .}}
{{.TableName}}:
	echo "from {{.ColumnMap}}"
{{end}}
`
