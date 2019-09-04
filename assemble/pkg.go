package assemble

/*
use `go tool compile -S pkg.go`
to see the assemble info ,like :
`
"".Id SNOPTRDATA size=8
        0x0000 01 00 00 00 00 00 00 00                          ........

`
"".Id 对应Id变量符号
变量的内存大小为8个字节
初始内容为：0x0000 1f 00 00 00 00 00 00 00 对应16进制的1f，对应10进制的31
SNOPTRDATA 是相关的标志，NOPTR表示数据中不包含指针数据
*/
var Id = 31
