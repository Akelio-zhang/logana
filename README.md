# Logana 日志分析工具

主要用于日志统计分析，包括各个类输出日志的级别统计、异常的统计。

## Example
### usage
logana -dir=logs/

### output
CLASS-INFO-STAT|2022-01-21
+--------------------------------------+-------+------+------+-------+-------+
|                CLASS                 | ERROR | WARN | INFO | DEBUG | COUNT |
+--------------------------------------+-------+------+------+-------+-------+
| com.akelio.risk.impl.ClmsServiceImpl |     0 |    0 |    1 |     0 |     1 |
| com.akelio.risk.impl.WarnServiceImpl |     1 |    0 |    0 |     0 |     1 |
+--------------------------------------+-------+------+------+-------+-------+
EXCEPTION-INFO-STAT|2022-01-21
+--------------------------------+-------+
|           EXCEPTION            | COUNT |
+--------------------------------+-------+
| java.lang.NullPointerException |     1 |
+--------------------------------+-------+