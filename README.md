go-4-mysql-tableinfo
====================

i was trying to use golang to transform sql(export from db) to get create table info, this is can be used to transform mysql to one object ,with all table info.

====================

Example
====================

<p>package main</p>

<p>import (</p>
<p>&nbsp;&nbsp;"flag"</p>
<p>&nbsp;&nbsp;"fmt"</p>
<p>&nbsp;&nbsp;"transform"</p>
<p>)</p>


<p>func main() {</p>
<p>&nbsp;&nbsp;flag.Parse()</p>
<p>&nbsp;&nbsp;file := flag.Arg(0)  </p>
<p>&nbsp;&nbsp;result := transform.ReadTableInfoFromSql(file)</p>
<p>&nbsp;&nbsp;fmt.Println(result)</p>
<p>}</p>

