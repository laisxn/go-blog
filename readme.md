# gin-blog

###开始
```$xslt
创建数据库 blog
复制 config_test.ini -> config.ini && 编辑对应配置

go run main.go
```

### 首次执行可选择导入 sql, 将清空数据，谨慎操作
```
go run main.go -init_sql=1
```