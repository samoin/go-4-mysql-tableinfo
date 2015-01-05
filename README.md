go-4-mysql-tableinfo
====================

i was trying to use golang to transform sql(export from db) to get create table info, this is can be used to transform mysql to one object ,with all table info.

====================

Example
====================
<code>
package main
</code>
<code>
import (
	"flag"
	"fmt"
	"transform"
)
</code>

func main() {
	flag.Parse()
	file := flag.Arg(0)  
	result := transform.ReadTableInfoFromSql(file)
	fmt.Println(result)
}
</code>
