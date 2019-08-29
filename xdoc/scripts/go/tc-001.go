///usr/local/go/bin/go run $0 $@; exit $?

package main

import "fmt"

/*
<<<TC
caseId:   100
title:    用户登录
steps:    @开头的为含验证点的步骤
   1000         打开登录页面
   1010         输入正确的用户名和密码
   @1020        点击'登录'按钮

expects:
#
@step1020期望结果, 可以有多行

readme:
- 脚本输出日志和expects章节中，#号标注的验证点需保持一致对应
- 脚本中CODE注解的需用代码替换，双斜标注的为说明文字
- 参考样例https://github.com/easysoft/zentaoatf/tree/master/xdoc/sample
TC
*/

func main() {
	fmt.Println("hello world.")
}
