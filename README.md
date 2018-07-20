# dellog

## 使用

`dellog` 使用如下配置文件，默认从 `/etc/dellog.d/*.yml` 中获取

```yaml
configs:
    # 激活配置
  - enabled: true
    # 要查找的路径
    paths: 
      - /home/app/tomcat/logs/custom/**/*.log
    # 文件保存天数
    keep: 2
```

`dellog` 从指定的路径中，查找文件名中包含 `yyyy-MM-dd` 日期标志的文件，以当前时间为基准，删除过期的文件

## 选项

* `-dry-run` 使用该选项时，只会输出要删除的文件，不会执行删除操作

## 许可证

`dellog` 使用 MIT 许可证
