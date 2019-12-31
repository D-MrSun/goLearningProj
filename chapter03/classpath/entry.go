/*
   The regexp package implements a simple library for
   regular expressions.

   The syntax of the regular expressions accepted is:

   regexp:
       concatenation { '|' concatenation }
   concatenation:
       { closure }
   closure:
       term [ '*' | '+' | '?' ]
   term:
       '^'
       '$'
       '.'
       character
       '[' [ '^' ] character-ranges ']'
       '(' regexp ')'
*/
package classpath

import (
	"os"
	"strings"
)

//定义文件分割符号
const pathListSeparator = string(os.PathSeparator)

type Entry interface {

	/**
	description:负责寻找和加载class文件;readClass（）方法的参数是class文件的相对路径，路径之间用斜线（/）分隔，文件名有.class后缀。
	example:比如要读取java.lang.Object类，传入的参数应该是java/lang/Object.class。
	返回值是读取到的字节数据、最终定位到class文件的Entry，以及错误信息。
	Go的函数或方法允许返回多个值，按照惯例，可以使用最后一个返回值作为错误信息。
	*/
	readClass(className string) ([]byte, Entry, error)

	//作用相当于Java中的toString（），用于返回变量的字符串表示。
	String() string
}

//composite pattern
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
