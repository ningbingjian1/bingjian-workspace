
<!-- TOC -->

- [hello world](#hello-world)
- [namespace 命名空间](#namespace-命名空间)
    - [定义命名空间](#定义命名空间)
    - [使用命名空间](#使用命名空间)
- [C++对C的增强](#c对c的增强)
    - [C++ 实用性增强](#c-实用性增强)
    - [对全局变量的检测能力加强](#对全局变量的检测能力加强)
    - [struct类型增强](#struct类型增强)
    - [C++中所有变量和函数都必须有类型](#c中所有变量和函数都必须有类型)
    - [新增bool类型关键字](#新增bool类型关键字)
    - [三目运算符功能增强](#三目运算符功能增强)
    - [const增强](#const增强)
        - [const基础](#const基础)
        - [const 和 #define 的相同](#const-和-define-的相同)
    - [真正的枚举](#真正的枚举)
- [C++对C语言的拓展](#c对c语言的拓展)
    - [引用](#引用)
    - [指针引用](#指针引用)
    - [const 引用](#const-引用)
    - [const 引用的原理](#const-引用的原理)
- [类和对象](#类和对象)

<!-- /TOC -->
# hello world

C语言的写法:

```
#define _CRT_SECURE_NO_WARNINGS
#include<stdio.h>
int main(void)
{
	int a = 0;
	printf("hello world!! \n");//输出屏幕
	printf("请输入一个数字");
	scanf("%d", &a);
	return 0;
}
```


C++的写法
```
#include<iostream>
using namespace std;
int main(void) {
	int a = 0;
	cout << "hello world !!" << endl;
	cout << "请输入一个数字" << endl;
	cin >> a ;
	return 0;
}
```


# namespace 命名空间


```
#include <iostream>


// 命名空间的引入方式  第一种 
using std::cout; //using 关键字 不是引入整个命名空间 而是引入命名空间一个变量
using std::endl; 



//第二种
using namespace  std; //using namespace 是引入整个命名空间



```

## 定义命名空间

```
//定义一个明明空间

namespace namespaceA//定义一个命名空间 namespace是命名空间关键字类型， namespaceA 是命名空间的名字
{
	// namespace A 空间领域
	int a = 10;

	int b = 20;
}

namespace namespaceB
{
	int a = 20;

	namespace namespaceC
	{
		struct teacher
		{
			int id;
			char name[64];
		};
	}

	namespace namespaceD
	{
		struct teacher
		{

		};
	}
}
```

## 使用命名空间

```

//使用自定义的命名空间
void test()
{
	//using namespaceA::a;//真个test（）函数领域中所有的a 默认都是 namespaceA中的a
	//using namespace namespaceA; //引入整个namespaceA空间
	cout << "A::a = " << namespaceA::a << endl;
	cout << "B::a = " << namespaceB::a << endl;
	//cout << "a = " << a << endl;


	//创建一个struct teacher的变量
	//using namespace namespaceB::namespaceC;//把namepaceC中的所有定义的遍历都引入
	//using namespaceB::namespaceC::teacher;

	//namespaceB::namespaceC::teacher t;
	using namespace namespaceB::namespaceC;
	teacher t;

}
```

# C++对C的增强

## C++ 实用性增强

```
using namespace	std;	
//C语⾔中的变量都必须在作⽤域开始的位置定义！！
//C++中更强调语⾔的“实⽤性”，所有的变量都可以在需要使⽤时再定义。
int	main(void)	
{	
				int	i	=	0;	
				cout	<<	"i	=	"	<<i	<<endl;	
				int	k;	
				k	=	4;	
				cout	<<	"k	=	"	<<k	<<endl;	
				return 0;	
}	
```

## 对全局变量的检测能力加强

对全局变量的检测能力加强， 一个变量不管是声明，还是定义，只能出现一次

```
int g_val ; //全局变量
//int g_val = 10;// 右一个全局变量
```

## struct类型增强

```
/*
				C语⾔的struct定义了⼀组变量的集合，C编译器并不认为这是⼀种新的类型
				C++中的struct是⼀个新类型的定义声明
*/
#include	<iostream>
struct	Student	
9
{	
				char	name[100];	
				int	age;	
};	
int	main(int	argc,	char	*argv[])	
{	
				Student	s1	=	{"wang",	1};	
				Student	s2	=	{"wang2",	2};	
				return 0;	
}
```

##  C++中所有变量和函数都必须有类型 

##  新增bool类型关键字

```
/*
			C++中的布尔类型
			C++在C语⾔的基本类型系统之上增加了bool
			C++中的bool可取的值只有true和false
			理论上bool只占⽤⼀个字节，
			如果多个bool变量定义在⼀起，可能会各占⼀个bit，这取决于编译器的实现
			true代表真值，编译器内部⽤1来表⽰
			false代表⾮真值，编译器内部⽤0来表⽰
			bool类型只有true（⾮0）和false（0）两个值
			C++编译器会在赋值时将⾮0值转换为true，0值转换为false
	*/
	
```

##  三目运算符功能增强

```
	int	a	=	10;	
	int	b	=	20;	
	//返回⼀个最⼩数 并且给最⼩数赋值成30
	//三⺫运算符是⼀个表达式 ，表达式不可能做左值
	(a	<	b	?	a	:	b	)	=	30;	
	printf("a	=	%d,	b	=	%d\n",	a,	b);	
```

## const增强 


### const基础

```
#include	<iostream>
int	main(void)	
{	
				//const	定义常量--->	const	意味只读
				const int	a;	
				int const	b;	
				//第1个第2个意思一样 代表一个常整形数
				const int	*c;	
				//第三个	c是一个指向常整形数的指针(所指向的内存数据不能被修改，但是本身可以修改)
				int	*	const	d;	
				//第四个	d	常指针（指针变量不能被修改，但是它所指向内存空间可以被修改）
				const int	*	const	e	;	
				//第五个	e 一个指向常整形的常指针（指针和它所指向的内存空间，均不能被修改）
				return 0;	
}

```

### const 和 #define 的相同 

C++中的const修饰的，是一个真正的常量，而不是C中变量（只读）。在
const修饰的常量编译期间，就已经确定下来了

C++中的const常量类似于宏定义
const int c = 5; ≈ #define c 5
C++中的const常量与宏定义不同
const常量是由编译器处理的，提供类型检查和作用域检查
宏定义由预处理器处理，单纯的文本替换


## 真正的枚举

```
#include<iostream>
using namespace std;
int main(void) {
	enum color {
		RED,GREEN,BLUE,YELLOW
	};
	enum color c1 = RED;
	cout << "c1 = " << c1 << endl;
	c1 = YELLOW; 
	cout << "c1 = " << c1 << endl;


	return 0;
}
```



#  C++对C语言的拓展 

## 引用

* 变量引用
	*  1 引用没有定义,是一种关系型声明。声明它和原有某一变量(实体)的关
系。故 而类型与原类型保持一致,且不分配内存。与被引用的变量有相同的地
址。 
 	*  2 声明的时候必须初始化,一经声明,不可变更。
 	*  3 可对引用,再次引用。多次引用的结果,是某一变量具有多个别名。 
 	*  4 &符号前有数据类型时,是引用。其它皆为取地址。 
* 引用作为函数参数 
	普通引用在声明时必须用其它的变量进行初始化，引用作为函数参数声
明时不进行初始化。
* 引用的意义 
	* 1）引用作为其它变量的别名而存在，因此在一些场合可以代替指针
	* 2）引用相对于指针来说具有更好的可读性和实用性
* 引用的本质 

```
引用做参数传递时编译器会替我们将实参取地址给引用
//int &a = main：：&b；//a是b的引用
当对引用进行操作赋值时，编译器帮我们隐藏*操作
// cout<<a其实是cout<<*a;*被编译器隐去
```

```
引用是常指针
```


```
引用在实现上，只不过是把：间接赋值成立的三个条件的后两步和二为一.
 当实参传给形参引用的时候，只不过是c++编译器帮我们程序员手工取了
一个实参地址，传给了形参引用（常量指针）。
```
* 引用作为函数的返回值（引用当左值） 

```
I. 当函数返回值为引用时,
若返回栈变量:
 不能成为其它引用的初始值（不能作为左值使用）
```

```
 当函数返回值为引用时, 
 若返回静态变量或全局变量
 可以成为其他引用的初始值（可作为右值使用，也可作为左值使用）
```

```
引用作为函数返回值， 
 如果返回值为引用可以当左值， 
22
 如果返回值为普通变量不可以当左值。 
```

## 指针引用 


##  const 引用 

```
const 引用有较多使用。它可以防止对象的值被随意修改。因些特性。 
  (1)const 对象的引用必须是 const 的,将普通引用绑定到 con合法的。这个原因比较简单。既然对象是 const 的,表示不能被修改也不 能修改,必须使用 const 引用。实际上,
  const int a=1; 
  int &b=a;
这种写法是不合法 的,编译不过。 
  (2)const 引用可使用相关类型的对象(常量,非同类型的变量或始化。这个是 const 引用与普通引用最大的区别。
  const int &a=2;
是合法的。
  double x=3.14; 
  const int &b=a;
也是合法的。 
```
##  const 引用的原理


```
double	val	=	3.14;	
const int	&ref	=	val;	
double	&	ref2	=	val;	
cout<<ref<<"	"<<ref2<<endl;	
val	=	4.14;	
cout<<ref<<"	"<<ref2<<endl;	
```

上述输出结果为 3 3.14 和 3 4.14。因为 ref 是 const 的,在初始化的过
程中已经给定值,不允许修改。而被引用的对象是 val,是非 const 的,所以 val 
的修改并未 影响 ref 的值,而 ref2 的值发生了相应的改变。 
 那么,为什么非 const 的引用不能使用相关类型初始化呢?实际上,const 
引用 使用相关类型对象初始化时发生了如下过程: 

```
int	temp	=	val;	
const int	&ref	=	temp;
```

如果 ref 不是 const 的,那么改变 ref 值,修改的是 temp,而不是 val。期
望对 ref 的赋值会修改 val 的程序员会发现 val 实际并未修改。


# 类和对象

