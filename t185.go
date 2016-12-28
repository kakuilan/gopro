// mysql 选择字段
package main
import (
    "database/sql"
    //_ "database/sql/convert"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "errors"
    "reflect"
)

type MyRows struct {
    *sql.Rows
    values map[string]interface{}
    columnName []string
}

func (this *MyRows) Next() bool {
    bResult := this.Rows.Next()
    if bResult {
	//如果成功,取所有列的数据到values里
	if this.columnName == nil || len(this.columnName) ==0 {
	    this.columnName,_ = this.Rows.Columns()
	}
	if this.values == nil {
	    this.values = make(map[string]interface{})
	}
	var arr []interface{}
	for i:=0;i<len(this.columnName);i++{
	    var inf interface{}
	    arr = append(arr, &inf)
	}
	//将数据接收到interface{}变量里
	this.Rows.Scan(arr...)

	for i:=0;i<len(this.columnName);i++{
	    this.values[this.columnName[i]] = reflect.ValueOf(arr[i]).Elem().Interface()
	}
    }
    return bResult
}

func (this *MyRows) GetValue(name string,value interface{}) error {
    if this.values == nil || len(this.values)==0 {
	return errors.New("没有调用Next,或没有可用的行数据")
    }
    i,ok := this.values[name]
    if ok {
	err := ConvertAssign(value, 1)
	if err != nil {
	    return err
	}
	return nil
    }
    return errors.New("字段不存在,请注意大小写")
}

func (this *MyRows) Scan(dest ...interface{}) error {
    if this.values == nil || len(this.values)==0 {
	return errors.New("没有调用Next,或没有可用的行数据")
    }

    for i:=0;i<len(dest);i++{
	err := ConvertAssign(dest[i], this.values[this.columnName[i]])
	if err != nil {
	    return err
	}
    }
    return nil
}
func main(){
    db,err := sql.Open("mysql","test:123456@tcp(192.168.128.1:3306)/test?charset=utf8")
    if err != nil {
	fmt.Println(err)
	return
    }

    defer db.Close()
    var row *sql.Rows
    row,err = db.Query("select * from person limit 1")
    if row.Next(){
	//返回所有的列名,列的顺序跟SCan一致
	fmt.Println(row.Columns())
	var id,name,age,isBoy interface{}
	err = row.Scan(&id,&name,&age,&isBoy)
	if err != nil {
	    fmt.Println(err)
	}
	fmt.Println("id", id, string(id.([]byte)))
	fmt.Println("name", name, string(name.([]byte)))
    }



}

func ConvertAssign(dest, src interface{}) error {
        // Common cases, without reflect.
        switch s := src.(type) {
        case string:
                switch d := dest.(type) {
                case *string:
                        if d == nil {
                                return errNilPtr
                        }
                        *d = s
                        return nil
                case *[]byte:
                        if d == nil {
                                return errNilPtr
                        }
                        *d = []byte(s)
                        return nil
                }
        case []byte:
                switch d := dest.(type) {
                case *string:
                        if d == nil {
                                return errNilPtr
                        }
                        *d = string(s)
                        return nil
                case *interface{}:
                        if d == nil {
                                return errNilPtr
                        }
                        *d = cloneBytes(s)
                        return nil
                case *[]byte:
                        if d == nil {
                                return errNilPtr
                        }
                        *d = cloneBytes(s)
                        return nil
                case *RawBytes:
                        if d == nil {
                                return errNilPtr
                        }
                        *d = s
                        return nil
                }
        case nil:
                switch d := dest.(type) {
                case *interface{}:
                        if d == nil {
                                return errNilPtr
                        }
                        *d = nil
                        return nil
                case *[]byte:
                        if d == nil {
                                return errNilPtr
                        }
                        *d = nil
                        return nil
                case *RawBytes:
                        if d == nil {
                                return errNilPtr
                        }
                        *d = nil
                        return nil
                }
        }

        var sv reflect.Value

        switch d := dest.(type) {
        case *string:
                sv = reflect.ValueOf(src)
                switch sv.Kind() {
                case reflect.Bool,
                        reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
                        reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
                        reflect.Float32, reflect.Float64:
                        *d = fmt.Sprintf("%v", src)
                        return nil
                }
        case *[]byte:
                sv = reflect.ValueOf(src)
                switch sv.Kind() {
                case reflect.Bool,
                        reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
                        reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
                        reflect.Float32, reflect.Float64:
                        *d = []byte(fmt.Sprintf("%v", src))
                        return nil
                }
        case *RawBytes:
                sv = reflect.ValueOf(src)
                switch sv.Kind() {
                case reflect.Bool,
                        reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
                        reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
                        reflect.Float32, reflect.Float64:
                        *d = RawBytes(fmt.Sprintf("%v", src))
                        return nil
                }
        case *bool:
                bv, err := driver.Bool.ConvertValue(src)
                if err == nil {
                        *d = bv.(bool)
                }
                return err
        case *interface{}:
                *d = src
                return nil
        }

        if scanner, ok := dest.(Scanner); ok {
                return scanner.Scan(src)
        }

        dpv := reflect.ValueOf(dest)
        if dpv.Kind() != reflect.Ptr {
                return errors.New("destination not a pointer")
        }
        if dpv.IsNil() {
                return errNilPtr
        }

        if !sv.IsValid() {
                sv = reflect.ValueOf(src)
        }

        dv := reflect.Indirect(dpv)
        if dv.Kind() == sv.Kind() {
                dv.Set(sv)
                return nil
        }

        switch dv.Kind() {
        case reflect.Ptr:
                if src == nil {
                        dv.Set(reflect.Zero(dv.Type()))
                        return nil
                } else {
                        dv.Set(reflect.New(dv.Type().Elem()))
                        return convertAssign(dv.Interface(), src)
                }
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
                s := asString(src)
                i64, err := strconv.ParseInt(s, 10, dv.Type().Bits())
                if err != nil {
                        return fmt.Errorf("converting string %q to a %s: %v", s, dv.Kind(), err)
                }
                dv.SetInt(i64)
                return nil
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
                s := asString(src)
                u64, err := strconv.ParseUint(s, 10, dv.Type().Bits())
                if err != nil {
                        return fmt.Errorf("converting string %q to a %s: %v", s, dv.Kind(), err)
                }
                dv.SetUint(u64)
                return nil
        case reflect.Float32, reflect.Float64:
                s := asString(src)
                f64, err := strconv.ParseFloat(s, dv.Type().Bits())
                if err != nil {
                        return fmt.Errorf("converting string %q to a %s: %v", s, dv.Kind(), err)
                }
                dv.SetFloat(f64)
                return nil
        }

        return fmt.Errorf("unsupported driver -> Scan pair: %T -> %T", src, dest)
}
