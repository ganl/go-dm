# go-dm
达梦官方 Go 驱动源码 

DM 8.0+

Go 1.13+

源码来源于官方安装包，驱动包的路径： `dmdbms/drivers/go/dm-go-driver.zip`  , 详见说明： [Go 数据库接口](https://eco.dameng.com/document/dm/zh-cn/app-dev/go-go.html) 

## DataSourceName

```
	dm://username:password@host:port?schema=schemaName[&logLevel=debug&sslFilesPath&=pathValue&timeout=timeoutValue&param1=value1&...&paramN=valueN]
```

`dm://SYSDBA:SYSDBA@localhost:5236`

## Example

`go get github.com/ganl/go-dm@v1.2.193`

PS: 8.1.2.192 / 8.2.2.192 官方未改驱动构建32位，编译报错： constant 4294967295 overflows int

```golang
package main

import (
	"database/sql"
	_ "github.com/ganl/go-dm"
	"log"
)

func main() {
	//url := "dm://" + os.Getenv("dm_username") + ":" + os.Getenv("dm_password") + "@" + os.Getenv("dm_host") + "?noConvertToHex=true"
	conn, err := sql.Open("dm", "dm://SYSDBA:SYSDBA@192.168.100.168:7777")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	for _, sqlstr := range []string{
		"DROP TABLE IF EXISTS tmp_testabc",
		`CREATE TABLE  tmp_testabc (
			id bigint primary key,
			name varchar(12)
		)`,
	} {
		_, err := conn.Exec(sqlstr)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	var id int64
	sqlstr := "insert into tmp_testabc(id, name) values(1, 'a') returning id into :id"
	_, err = conn.Exec(sqlstr, sql.Out{Dest: &id})
	if err != nil {
		log.Fatal(err)
		return
	}
	if id != 1{
		log.Fatal("Error cmp")
	}
	sqlstr = "insert into tmp_testabc(id, name) values(2, 'b') returning id into :id"
	_, err = conn.Exec(sqlstr, sql.Named("id123", sql.Out{Dest: &id}))
	if err != nil {
		log.Fatal(err)
		return
	}
	if id != 2{
		log.Fatal("Error cmp")
	}
}

```
