package transform

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

/**
* 读取文件(read file by append)
*/
func ReadByAppend(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}
/**
* 读取文件(read file by io)
*/
func ReadByBufio(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}
/**
* 读取文件(read file by ioutil)
*/
func ReadByIoutil(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

/**
* 表对应的类(one table info)
 */
type TableInfo struct {
	/**
	*表名(tablename)
	 */
	name string `json:"name"`
	/**
	*表的中文说明(comment for table)
	 */
	comment string `json:"comment"`
	/**
	* 列对应的字段有name（名称），types（类型），constraint（约束），comment（说明）<br/>
	* (all columns in this table)
	*/
	column []map[string]string `json:"column"`
}
/**
* 返回表信息的表名(return the filed "name" of this TableInfo)
*/
func GetTableInfoName(t TableInfo) string{
	return t.name
}
/**
* 返回表信息的表的中文说明(return the filed "comment" of this TableInfo)
*/
func GetTableInfoComment(t TableInfo) string{
	return t.comment
}
/**
* 返回表信息的列对应的字段(return the filed "column" of this TableInfo)
*/
func GetTableInfoColumn(t TableInfo) []map[string]string{
	return t.column
}

/**
* 进行初始化的过滤，主要是：<br/>
* 1.替换特殊字符"`"<br/>
* 2.根据特殊关键词"CREATE TABLE"，进行打散到数组<br/>
* 3.根据";"打散到数组，并返回对应数组的第一个元素<br/>
* <font color=red>注意：如果通过正常导出的sql文件，在执行后，返回数组的第一个元素不是实际的库表信息内容</font><br/>
*/
func GetFristFitleArr(sql string) (index []string) {
	specialReg, err := regexp.Compile("`")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sql = specialReg.ReplaceAllString(sql, "'")
	tablearr := strings.Split(sql, "CREATE TABLE")
	for i, v := range tablearr {
		if strings.Contains(v, "ENGINE") {
			tablearr2 := strings.Split(v, ";\r\n")
			tablearr[i] = tablearr2[0]
		} else {

			tablearr[i] = ""
		}
	}
	return tablearr
}

/**
* 通过正则获取对应匹配的参数(read matched params from string by a regex)
*/
func GetByReg(regStr string, regIndex int, str string) string {
	tableReg, err := regexp.Compile(regStr)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	matches := tableReg.FindAllStringSubmatch(str, 1)
	for _, item := range matches {
		return item[regIndex]
	}
	return ""
}

/**
* 设置数据库表的信息(set table info by a string array)
*/
func SetTableInfo(arr []string) []TableInfo {
	resultTableInfo := make([]TableInfo, len(arr)-1)
	for i, v := range arr {
		if strings.Contains(v, "ENGINE") {
			tbInfo := TableInfo{}
			tbInfo.name = GetByReg("'([^']+)'", 1, v)
			tbInfo.comment = GetByReg(`CHARSET=[\s\S]+ COMMENT='([^']+)'`, 1, v)
			tmpArr := strings.Split(v, "\n")
			columnArr := make([]map[string]string, len(tmpArr))
			tbInfo.column = columnArr
			var primarykeyName string
			for i2, v2 := range tmpArr {
				v2 = strings.TrimLeft(v2, " ")
				if strings.HasPrefix(v2, "'") && i2 > 0 {
					//解析字段的信息'tid' int(11) NOT NULL DEFAULT '0' COMMENT '其他平台编号',
					columnMap := make(map[string]string)
					columnMap["name"] = GetByReg("'([^']+)'", 1, v2)
					columnMap["types"] = GetByReg("'[^']+' ([^ ]+)", 1, v2)
					columnMap["comment"] = GetByReg(`[\s\S]+COMMENT '([^']+)'`, 1, v2)
					var constraint string = GetByReg(`[\s\S]+\)([\s\S]+) COMMENT`, 1, v2)
					if constraint == "" {
						constraint = GetByReg(`[\s\S]+\) ([\s\S]+)`, 1, v2)
					}
					columnMap["constraint"] = constraint
					columnArr[i2] = columnMap
				} else if strings.HasPrefix(v2, "PRIMARY KEY") {
					//解析主键idPRIMARY KEY (`ad_id`),
					primarykeyName = GetByReg("PRIMARY KEY \\('([^']+)'\\)", 1, v2)
				}
			}
			for i2, v2 := range columnArr {
				if v2["name"] == primarykeyName {
					if columnArr[i2] == nil {
						break
					}
					columnArr[i2]["constraint"] = "PRIMARY KEY ; " + columnArr[i2]["constraint"]
				}
			}
			resultTableInfo[i-1] = tbInfo
		}
	}
	return resultTableInfo
}

/**
* 读取sql文件数中的据库表的信息(read all table info from a sql file, exported from db, like mysql)
*/
func ReadTableInfoFromSql(file string) []TableInfo {
	str := ReadByIoutil(file) 
	resultTableInfo := SetTableInfo(GetFristFitleArr(str))
	return resultTableInfo
}
/**
* 读取sql文本的据库表的信息(read all table info from a sql text, exported from db, like mysql)
*/
func ReadTableInfoFromSqlInfo(str string) []TableInfo {
	resultTableInfo := SetTableInfo(GetFristFitleArr(str))
	return resultTableInfo
}
