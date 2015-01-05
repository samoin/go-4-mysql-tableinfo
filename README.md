go-4-mysql-tableinfo
====================

i was trying to use golang to transform sql(export from db) to get create table info, this is can be used to transform mysql to one object ,with all table info.

====================

you can use this like this.
--------------------
package main

import (
	"flag"
	"fmt"
	"transform"
)


func main() {
	flag.Parse()
	file := flag.Arg(0)  
	result := transform.ReadTableInfoFromSql(file)
	fmt.Println(result)
}
