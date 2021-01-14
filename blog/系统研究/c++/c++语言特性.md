# C++语言特性

## 类

### 类成员

类包括两类成员变量，static和no static成员变量，所有的类对象均可以访问static成员，每个类对象只能访问自己的no static成员。使用类可以访问static成员，不能访问no static成员。

类包括三类成员函数，static、no static、virtual成员函数。类只能访问static成员函数，static成员函数只能访问static成员变量。所有的类对象均可以访问static成员函数、no static成员函数，no static成员函数只能访问no static成员变量而不能访问static成员变量。

虚函数与纯虚函数，虚函数用来实现多态，纯虚函数用来实现接口。

### 类的实现

三种实现机制：

1. 类对象直接存放所有的成员变量，函数调用直接指向类成员函数。问题：无法实现static成员变量、无法实现多态。
2. 类对象存放两个表的指针，成员变量表和成员函数表。可以实现多态，但效率比较低。
3. C++对象模型，static成员变量、static成员函数均不放在类对象中，no static成员变量放在类对象中，对于虚函数，则使用虚函数表来实现。每个对象有一个虚函数表指针vptr，指向一个虚函数表，虚函数表按照声明的顺序放置虚函数的指针。

### 重载、覆盖、隐藏

重载：函数签名不同，函数是相同的函数名，但参数不一样，视为重载。
覆盖：函数签名相同，有相同的函数名、参数也一样，且为虚函数，视为覆盖。
隐藏：函数签名一样，但不是虚函数，视为隐藏。

### 继承、多态

子类覆盖父类的虚函数，父类指针的函数调用表现为多态。

## static的作用

### static的引入

作用域的static：只能在作用域被访问，具有隐藏性，同时具有全局变量的生命期。
类的static：为整个类服务而不是具体的一个对象，同时被隐藏在类的内部。

### static的存储

static存储在程序的BSS段中，在程序加载后被系统自动清0。所以不能在函数内分配static成员和初始化。

### static的作用

1. 修饰变量：static修饰的局部变量具有了全局变量一样的生命周期。在程序加载后被初始化，在程序退出后释放。
2. 修饰全局变量：static修饰的全局变量只能在本文件中被访问，不能被其他文件访问。
3. 修饰函数：static修饰的函数，只能在本文件中被调用。
4. 修饰类成员函数：static修饰的类成员函数只能访问static的类成员变量，类可以直接调用该成员函数。

[C/C++ 中 static 的用法全局变量与局部变量](https://www.runoob.com/w3cnote/cpp-static-usage.html)

## const的使用

### 修饰普通类型的变量
```
const int  a = 7; 
int  b = a;  // 正确
a = 8;       // 错误，不能改变
```

```
#include<iostream>
using namespace std;
int main(void)
{
    const int  a = 7;
    int  *p = (int*)&a;
    *p = 8;
    cout<<a;
    system("pause");
    return 0;
}
```
可以看到上面的输出为7，但a确实被修改为8，这是因为编译器根据const修饰做了编译优化。

### 修饰指针变量

1. 修饰指针指向的内容为不可变
```
const int *p = 8;
```
2. 修饰指针本身不可变
```
int a = 8;
int* const p = &a;
*p = 9; // 正确
int  b = 7;
p = &b; // 错误
```
3. 修饰指针和指针指向的内容均不可变
```
int a = 8;
const int * const  p = &a;
```

### 修饰函数参数
1. 值传递的const修饰传递，一般这种情况不需要const修饰，因为函数会自动产生临时变量复制实参值。
```
#include<iostream>
using namespace std;
void Cpf(const int a)
{
    cout<<a;
    // ++a;  是错误的，a 不能被改变
}
```
2. 当const参数为指针时，可以防止指针被意外篡改。
```
#include<iostream>
using namespace std;
void Cpf(int *const a)
{
    cout<<*a<<" ";
    *a = 9;
}

int main(void)
{
    int a = 8;
    Cpf(&a);
    cout<<a; // a 为 9
    system("pause");
    return 0;
}
```
3. 自定义类型的参数传递，需要临时对象复制参数，对于临时对象的构造，需要调用构造函数，比较浪费时间，因此我们采取const外加引用传递的方法。
```
#include<iostream>
using namespace std;
class Test
{
public:
    Test(){}
    Test(int _m):_cm(_m){}
    int get_cm()const
    {
       return _cm;
    }
private:
    int _cm;
};
void Cmf(const Test& _tt)
{
    cout<<_tt.get_cm();
}

int main(void)
{
    Test t(8);
    Cmf(t);
    system("pause");
    return 0;
}
```

### 修饰函数返回值
1. const修饰普通类型的返回值，修饰与不修饰返回值作用一样。
2. const修饰自定义类型的返回值，此时返回的值不能作为左值使用，既不能被赋值，也不能被修改。
```
const rational operator*(const rational& lhs, const rational& rhs);
rational a, b, c;
(a * b) = c;       // 对a*b的结果赋值
```
3. const修饰返回的指针或者引用，是否返回一个指向const的指针，取决于我们想让用户干什么。

### 修饰类成员函数
const修饰类成员函数，其目的是防止成员函数修改被调用对象的值，如果我们不想修改一个调用对象的值，所有的成员函数都应当声明为const成员函数。
const不能和static同时使用，因为static成员函数不含this指针。
```
#include<iostream>
using namespace std;
class Test
{
public:
    Test(){}
    Test(int _m):_cm(_m){}
    int get_cm()const
    {
       return _cm;
    }
 
private:
    int _cm;
};
void Cmf(const Test& _tt)
{
    cout<<_tt.get_cm();
}
int main(void)
{
    Test t(8);
    Cmf(t);
    system("pause");
    return 0;
}
```

[C++ const 关键字](https://www.runoob.com/w3cnote/cpp-const-keyword.html)

## 指针和引用

### 区别

指针是一个对象的内存地址，而引用是一个对象的别名。指针本质上就是存放变量地址的一个变量，和变量独立，可以被修改以指向新的对象。

而引用时一个变量的别名，逻辑上和变量不独立，引用在一开始就被初始化且引用的变量在整个生命周期都固定。

指针和引用经常用于函数的参数传递，然而，指针传递参数和引用传递参数是有本质上的不同的：

1. 指针传递参数本质上是值传递的方式，它所传递的是一个地址值。值传递过程中，被调函数的形式参数作为被调函数的局部变量处理，即在栈中开辟了内存空间以存放由主调函数放进来的实参的值，从而成为了实参的一个副本。值传递的特点是被调函数对形式参数的任何操作都是作为局部变量进行，不会影响主调函数的实参变量的值。

2. 引用传递过程中，被调函数的形式参数虽然也作为局部变量在栈中开辟了内存空间，但是这时存放的是由主调函数放进来的实参变量的地址。被调函数对形参的任何操作都被处理成间接寻址，即通过栈中存放的地址访问主调函数中的实参变量。正因为如此，被调函数对形参做的任何操作都影响了主调函数中的实参变量。

引用传递和指针传递是不同的，虽然它们都是在被调函数栈空间上的一个局部变量，但是任何对于引用参数的处理都会通过一个间接寻址的方式操作到主调函数中的相关变量。

对于指针传递的参数，如果改变被调函数中的指针地址，它将影响不到主调函数的相关变量。如果想通过指针参数传递来改变主调函数中的相关变量，那就得使用指向指针的指针，或者指针引用。

### 引用的三种用法

1. 独立引用
```
#include<iostream>
using namespace std; 
int main()
{
    int a=3;
    int& b=a;     //b与a绑定在一起，同生共死
    cout<<b<<" "<<a<<endl;
    b=5;
    cout<<b<<" "<<a<<endl;
    return 0;
}
```

2. 函数参数
```
#include<iostream>
using namespace std;
void f(int& b)  //b在这里与实参a无异
{
    b++;
}

int main()
{
    int a=3;
    f(a);        //a受函数体内部影响
    cout<<a<<endl;
    return 0;
}
```

3. 返回值
```
#include<iostream> 
using namespace std;
// f函数返回一个（*p）的引用，即a的引用。此引用可作为左值进行运算。
int& f(int* p)
{
    (*p)++;
    return *p;
}

int main()
{
    int a=3,b;
    b=f(&a)*5;
    f(&a)+=10;
    cout<<b<<" "<<a<<endl;  //输出20与15
    
    return 0;
}
```

[指针与引用的区别以及引用的三种用法](https://blog.csdn.net/u011857683/article/details/78348212)

