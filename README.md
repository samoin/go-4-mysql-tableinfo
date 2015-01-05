go-4-mysql-tableinfo
====================

i was trying to use golang to transform sql(export from db) to get create table info, this is can be used to transform mysql to one object ,with all table info.

====================

Example
====================
<p>trans.go</p>
<p>package main</p>

<p>import (</p>
<p>&nbsp;&nbsp;&nbsp;&nbsp;"flag"</p>
<p>&nbsp;&nbsp;&nbsp;&nbsp;"fmt"</p>
<p>&nbsp;&nbsp;&nbsp;&nbsp;"github.com/go-4-mysql-tableinfo-master"</p>
<p>)</p>


<p>func main() {</p>
<p>&nbsp;&nbsp;&nbsp;&nbsp;flag.Parse()</p>
<p>&nbsp;&nbsp;&nbsp;&nbsp;file := flag.Arg(0)  </p>
<p>&nbsp;&nbsp;&nbsp;&nbsp;result := transform.ReadTableInfoFromSql(file)</p>
<p>&nbsp;&nbsp;&nbsp;&nbsp;fmt.Println(result)</p>
<p>}</p>
====================
Run
====================
go run trans.go /home/samoin/test.sql
