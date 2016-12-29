# yphp

php useful function implement in golang

将php中的一些公用的方法用golang实现，主要目标用于提高golang的开发效率

# 提供的方法有

```
func Implode(glue string, pieces []string) string

func Explode(delimiter string, str string) []string

func In_array(needle string, haystack []string) bool

func Mysql_escape_string(escapestr string) string

func Ucwords(str string) string

func Lcfirst(str string) string
```

# Requirements

Go > 1.4

# License

WAVE TEAM MIT license.
