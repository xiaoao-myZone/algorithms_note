## C tips

### basic
1. ! 的优先级比 > < 大  验证式 !1>-1
2. (int) 强制转换
3. char and int are same when compare and operate
4. squence's start index and number of it's elements (len) are beautiful.
5. 字符串之间也可以比较大小，从左至右单个字符先大于对方的较大

### skills
1. 善于利用指针, 不要再用 char * s; 将s[i]换成s++
2. 有些中间量需要问自己到底有没有必要？

### array
1. 多维数组如果要初始化貌似只能用`a[][5] = {{...}..}`的方式,必须指定子数列的长度



### learning
1. `printf("int: %d", *p++);`
会先返回int *p的值给printf,再将p这个地址自增,给*p加括号则是另一种熟悉的行为
这意味着传入参数也是一种看不见的运算符
2. `char * s="asdaf";`不能使用`s[2]='P'`修改值,有这个功能的是`s[20]="asdaf";`
3. 然而s[20]不可以使用s++,即不可以修改头指针,但是如果将`char *p; p=s; p++;`是没有问题的
4. 另外`char s[20]; s="asfag";`这种表达会出错
5. TODO找到一个底层原因解释这种现象
6. `char *s = "asfa";`像这样带双引号的的赋值,系统自动加上了一个"\0",而被赋值的双引号,其实返回的是一个首地址
>> TODO
7. 6_zigzagTransform_.c这个文件里,s会溢出,打印出来结果是`|A||`会多一个"|",但是将最后一行%s前后的"""删掉,结果是`A`,有点测不准的感觉了 
8. 本来用作char的变量去干int的溢出

>> TODO
9. 各种比较运算符的使用有没有什么讲究,速度有没有先后
