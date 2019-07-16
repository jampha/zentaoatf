<?php
<<<TC
caseId:   200
title:    登录失败账号锁定策略
steps:    @开头的为含验证点的步骤
   step2000         连续输入3次错误的密码
   @step2010        第4次尝试登录
   group2100        不连续输入3次错误的密码
      step2101      输入2次错误的密码
      step2102      输入1次正确的密码
      step2103      再输入1次错误的密码
      @step2104     再输入1次正确的密码

expects:
# 
aaa
bbb

# 
111
222

readme:
- 脚本输出日志和expects章节中，#号标注的验证点需保持一致对应
- 脚本中/* */标注的需用代码替换，//注解的为说明文字
- 参考样例https://github.com/easysoft/zentaoatf/tree/master/xdoc/sample

TC;

/* 此处编写操作步骤代码 */

echo "#\n";  // @step2010: 系统提示账号被锁定
echo "aaa\n";
echo "bbb\n";

echo "#\n";  // @step2104: 登录成功，账号未被锁定
echo "111\n";
echo "222\n";

?>