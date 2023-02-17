# sina_luck_draw
新浪微博的转发抽奖，写得比较随意，包括评论，转发，点赞，因为抽得太频繁被举报了，使用的时候建议不要太频繁

# 使用方法：
## 修改config文件夹下的db.go和sina.go

### db.go
  修改数据库地址、数据库名、数据库用户名和密码

### sina.go
  修改SUB：微博的sub，登录微博账号后，打开主页，按F12，在cookie中查找sub
  修改UID：自己微博的uid（浏览器打开自己微博主页，地址栏中一串数字就是uid）
  
## VSCode编译生成exe
  执行命令：go build .
