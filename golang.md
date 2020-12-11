#### golang小细节 
- string和strconv.Itoa()

两个方法都可以将一个int类型的数字转为字符串，不同的是string(int)是根据ascii转的。
而strconv.Itoa()是将数字字面量转成字符串字面量

```go
strconv.Itoa(123) // ->"123"
string(123) // "{"  123的ascii是{
```

- fmt.Sprintf 
"%b" 代表以二进制的形式输出   
```go
//打印的过程中会清除 前面的0
num:= 3
bitStr:=  fmt.Sprintf("%0.32b",num) 
// 00000000000000000000000000000011
// 11
bitStr:=  fmt.Sprintf("%b",num)
```


- 字符串

```go
func main(){
	var s  string= "1223李"

	length :=  len(s)
	fmt.Printf("len=%v\n",length)

	for i:= 0;i<length;i++{
		fmt.Printf("utf8 char = %b\n", s[i])
	}

	for _,val := range s{
		fmt.Printf("unicode char =%b\n",val)
	}

	//utf8 char = 110001
	//utf8 char = 110010
	//utf8 char = 110010
	//utf8 char = 110011
	//utf8 char = 11100110
	//utf8 char = 10011101
	//utf8 char = 10001110
	//unicode char =110001
	//unicode char =110010
	//unicode char =110010
	//unicode char =110011
	//unicode char =110011101001110
}
```
```go
type stringStruct struct{
    str unsafe.Pointer
    len int
}
```
go中的字符串本质上[]byte,长度为字符串中包含的字符个数。由于golang采用的utf-8编码。所以汉字占用
三个字节。

```
UTF-8 最大的一个特点，就是它是一种变长的编码方式。它可以使用1~4个字节表示一个符号，根据不同的符号而变化字节长度。

UTF-8 的编码规则很简单，只有二条：

1）对于单字节的符号，字节的第一位设为0，后面7位为这个符号的 Unicode 码。因此对于英语字母，UTF-8 编码和 ASCII 码是相同的。

2）对于n字节的符号（n > 1），第一个字节的前n位都设为1，第n + 1位设为0，后面字节的前两位一律设为10。剩下的没有提及的二进制位，全部为这个符号的 Unicode 码。

下表总结了编码规则，字母x表示可用编码的位。

Unicode符号范围     |        UTF-8编码方式
(十六进制)        |              （二进制）
----------------------+---------------------------------------------
0000 0000-0000 007F | 0xxxxxxx
0000 0080-0000 07FF | 110xxxxx 10xxxxxx
0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx

下面，还是以汉字严为例，演示如何实现 UTF-8 编码。

严的 Unicode 是4E25（100111000100101），根据上表，可以发现4E25处在第三行的范围内（0000 0800 - 0000 FFFF），因此严的 UTF-8 编码需要三个字节，
即格式是1110xxxx 10xxxxxx 10xxxxxx。然后，从严的最后一个二进制位开始，
依次从后向前填入格式中的x，多出的位补0。这样就得到了，
严的 UTF-8 编码是11100100 10111000 10100101，转换成十六进制就是E4B8A5。
```
[字符编码](http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html)

根据上面的结果可以知道 "李" 对应的utf-8编码为
```
//utf8 char = 11100110
//utf8 char = 10011101
//utf8 char = 10001110
```

符合上面的utf8编码规则，去掉里面的固定编码[...]，得到的是unicode 编码。

```
//utf8 char = [111]00110
//utf8 char = [10]011101
//utf8 char = [10]001110
```

而go中的rune(utf32)类型就是unicode编码字符串。而byte类型（uint8）是以utf-8编码显示的。

有上面我们可以看出两种for循环的不同。











