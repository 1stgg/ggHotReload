
只测试了macOS，其他系统自行修改代码build


使用方法
```sh
# 首次设置
在`~/.bash_profile`中添加
alias ggHR=ggHotReload可执行文件的路径

# 使用
ggHR [sh="sh ./ggHR.sh"] [watch="./"] 
# path watch 默认当前路径
```

参数
|参数	|必填	  |默认值			|说明						|
|-    |-    |-          |-              |
|sh 	|否		|sh ./ggHR.sh	|执行bash文件路径	|
|watch|否		|./					|监控路径				|
