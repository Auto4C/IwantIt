# tkinter的函数与实例



## 1、使用tkinter.Tk() 生成主窗口（root=tkinter.Tk()）

root.title('标题名')    　　 　　修改框体的名字,也可在创建时使用className参数来命名；
root.resizable(0,0)   　　 　　框体大小可调性，分别表示x,y方向的可变性；
root.geometry('250x150')　　指定主框体大小；
root.quit()        　　　　 　　 退出；
root.update_idletasks()
root.update()        　　　　　刷新页面；

```
1
2
3
4
5
6
```

2、初级样例

1 import tkinter
2 root=tkinter.Tk() #生成root主窗口
3 label=tkinter.Label(root,text='Hello,GUI') #生成标签
4 label.pack()        #将标签添加到主窗口
5 button1=tkinter.Button(root,text='Button1') #生成button1
6 button1.pack(side=tkinter.LEFT)         #将button1添加到root主窗口
7 button2=tkinter.Button(root,text='Button2')
8 button2.pack(side=tkinter.RIGHT)
9 root.mainloop()             #进入消息循环（必需组件）

```
1
2
3
4
5
6
7
8
9
```

3、tkinter中的15种核心组件

```
Button        　　按钮；
Canvas        　　绘图形组件，可以在其中绘制图形；
Checkbutton      复选框；
Entry        　　 文本框（单行）；
Text             文本框（多行）；
Frame         　　框架，将几个组件组成一组
Label        　　 标签，可以显示文字或图片；
Listbox      　　 列表框；
Menu     　　     菜单；
Menubutton       它的功能完全可以使用Menu替代；
Message          与Label组件类似，但是可以根据自身大小将文本换行；
Radiobutton      单选框；
Scale      　　   滑块；允许通过滑块来设置一数字值
Scrollbar        滚动条；配合使用canvas, entry, listbox, and text窗口部件的标准滚动条；
Toplevel         用来创建子窗口窗口组件。
```
（在Tkinter中窗口部件类没有分级；所有的窗口部件类在树中都是兄弟。）

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
```

4、组件的放置和排版（pack,grid,place)

anchor的含义在这里插入图片描述
坐标
在这里插入图片描述

pack组件设置位置属性参数：
    after:    　　　 将组件置于其他组件之后；
    before:    　　　将组件置于其他组件之前；
    anchor:    　　  组件的对齐方式，顶对齐'n',底对齐's',左'w',右'e'
    side:    　　　　组件在主窗口的位置，可以为'top','bottom','left','right'（使用时tkinter.TOP,tkinter.E）；
    fill            填充方式 (Y,垂直，X，水平）
    expand          1可扩展，0不可扩展
grid组件使用行列的方法放置组件的位置，参数有：
    column:         组件所在的列起始位置；
    columnspam:     组件的列宽；
    row：      　　　组件所在的行起始位置；
    rowspam：    　　组件的行宽；
place组件可以直接使用坐标来放置组件，参数有：
    anchor:    　　　组件对齐方式；
    x:        　　　 组件左上角的x坐标；
    y:        　　   组件右上角的y坐标；
    relx:         　组件相对于窗口的x坐标，应为0-1之间的小数；
    rely:           组件相对于窗口的y坐标，应为0-1之间的小数；
    width:          组件的宽度；
    heitht:    　   组件的高度；
    relwidth:       组件相对于窗口的宽度，0-1；
    relheight:　    组件相对于窗口的高度，0-1；

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
```

5、使用tkinter.Button时控制按钮的参数

```
anchor:      　　　　  指定按钮上文本的位置；
background(bg)    　  指定按钮的背景色；
bitmap:       　　　　 指定按钮上显示的位图；
borderwidth(bd)　　　　指定按钮边框的宽度；
command:   　　　　　  指定按钮消息的回调函数；
cursor:        　　　　指定鼠标移动到按钮上的指针样式；
font:           　　  指定按钮上文本的字体；
foreground(fg)　　　　 指定按钮的前景色；
height:        　　　　指定按钮的高度；
image:        　　　　 指定按钮上显示的图片；
state:          　　　 指定按钮的状态（disabled）；
text:           　　　 指定按钮上显示的文本；
width:       　　　　  指定按钮的宽度
padx          　　　　 设置文本与按钮边框x的距离，还有pady;
activeforeground　　　 按下时前景色
textvariable    　　  可变文本，与StringVar等配合着用

1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
```

6、文本框tkinter.Entry,tkinter.Text控制参数

```
background(bg)   　　 文本框背景色；
foreground(fg)        前景色；
selectbackground　　  选定文本背景色；
selectforeground　　  选定文本前景色；
borderwidth(bd)    　 文本框边框宽度；
font                　字体；
show          　　    文本框显示的字符，若为*，表示文本框为密码框；
state            　　 状态；
width        　　　　  文本框宽度
textvariable    　　  可变文本，与StringVar等配合着用

1
2
3
4
5
6
7
8
9
10
```

7、标签tkinter.Label组件控制参数

```
Anchor        　　　　标签中文本的位置；
background(bg)　　　　背景色；
foreground(fg)　　    前景色；
borderwidth(bd)　　   边框宽度；
width        　　　　 标签宽度；
height        　　　　标签高度；
bitmap        　　　  标签中的位图；
font            　　　字体；
image        　　 　　标签中的图片；
justify        　　　 多行文本的对齐方式；
text    　　　　　　   标签中的文本，可以使用'\n'表示换行
textvariable  　　　  显示文本自动更新，与StringVar等配合着用

1
2
3
4
5
6
7
8
9
10
11
12
```

8、单选框和复选框Radiobutton,Checkbutton控制参数

```
anchor         　　文本位置；
background(bg) 　　背景色；
foreground(fg)    前景色；
borderwidth       边框宽度；
width        　　  组件的宽度；
height     　　    组件高度；
bitmap    　　     组件中的位图；
image    　　      组件中的图片；
font       　　    字体；
justify       　　 组件中多行文本的对齐方式；
text         　　  指定组件的文本；
value      　　    指定组件被选中中关联变量的值；
variable     　    指定组件所关联的变量；
indicatoron        特殊控制参数，当为0时，组件会被绘制成按钮形式;
textvariable       可变文本显示，与StringVar等配合着用

1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
```

9、组图组件Canvas控制参数

```
background(bg)  　　  背景色;
foreground(fg)       前景色;
borderwidth   　　　　组件边框宽度；
width     　　　　    组件宽度；
height     　　      高度;
bitmap 　　          位图;
image 　　　　        图片;
```
绘图的方法主要以下几种：
    create_arc          圆弧;
    create_bitmap  　　  绘制位图，支持XBM;
    create_image    　　 绘制图片，支持GIF(x,y,image,anchor);
    create_line         绘制支线；
    create_oval;        绘制椭圆；
    create_polygon   　　绘制多边形(坐标依次罗列，不用加括号，还有参数，fill,outline)；
    create_rectangle　　 绘制矩形((a,b,c,d),值为左上角和右下角的坐标)；
    create_text         绘制文字(字体参数font,)；
    create_window    　　绘制窗口；
    delete            　 删除绘制的图形；
    itemconfig          修改图形属性，第一个参数为图形的ID，后边为想修改的参数；
    move          　　   移动图像（1，4，0），1为图像对象，4为横移4像素，0为纵移像素，然后用root.update()刷新即可看到图像的移动，为了使多次移动变得可视，最好加上time.sleep()函数；
    只要用create_方法画了一个图形，就会自动返回一个ID,创建一个图形时将它赋值给一个变量，需要ID时就可以使用这个变量名。
    coords(ID)          返回对象的位置的两个坐标（4个数字元组）；
对于按钮组件、菜单组件等可以在创建组件时通过command参数指定其事件处理函数。方法为bind;或者用bind_class方法进行类绑定，bind_all方法将所有组件事件绑定到事件响应函数上。

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
```

10、菜单Menu

参数： 
    tearoff      　   分窗，0为在原窗，1为点击分为两个窗口
    bg,fg       　　  背景，前景
    borderwidth    　 边框宽度
    font              字体
    activebackgound   点击时背景，同样有activeforeground，activeborderwidth，disabledforeground
    cursor
    postcommand
    selectcolor    　 选中时背景
    takefocus
    title       
    type
    relief
方法：
    menu.add_cascade      添加子选项
    menu.add_command      添加命令（label参数为显示内容）
    menu.add_separator    添加分隔线
    menu.add_checkbutton  添加确认按钮
    delete                删除

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
```

11、事件关联

bind(sequence,func,add)——
bind_class(className,sequence,func,add)
bind_all(sequence,func,add)
事件参数：　　
sequence      　　　　　　　　所绑定的事件；
func            　　　　　　 所绑定的事件处理函数；
add             　　　　　　 可选参数，为空字符或‘+’；
className    　　　　　　　 　所绑定的类；
鼠标键盘事件
    <Button-1>        　  　鼠标左键按下，2表示中键，3表示右键；
    <ButtonPress-1>    　   同上；
    <ButtonRelease-1>　　　 鼠标左键释放；
    <B1-Motion>  　　       按住鼠标左键移动；
    <Double-Button-1>  　　 双击左键；
    <Enter>       　　      鼠标指针进入某一组件区域；
    <Leave>    　　         鼠标指针离开某一组件区域；
    <MouseWheel>  　   　　 滚动滚轮；
    <KeyPress-A> 　　  　　  按下A键，A可用其他键替代；
    <Alt-KeyPress-A>　　　   同时按下alt和A；alt可用ctrl和shift替代；
    <Double-KeyPress-A>　　  快速按两下A；
    <Lock-KeyPress-A>　　　  大写状态下按A；
窗口事件
    Activate        　　　　 当组件由不可用转为可用时触发；
    Configure      　　　　  当组件大小改变时触发；
    Deactivate    　　　　　 当组件由可用转变为不可用时触发；
    Destroy        　　　　  当组件被销毁时触发；
    Expose         　　　　　当组件从被遮挡状态中暴露出来时触发；
    Unmap        　　　　　　当组件由显示状态变为隐藏状态时触发；
    Map         　　　　     当组件由隐藏状态变为显示状态时触发；
    FocusIn       　　　 　  当组件获得焦点时触发；
    FocusOut      　　　　　 当组件失去焦点时触发；
    Property     　　　　    当窗体的属性被删除或改变时触发；
    Visibility       　　　　当组件变为可视状态时触发；
响应事件
event对象（def function(event)）：
    char        　　　　　　  按键字符，仅对键盘事件有效；
    keycode   　　　　　　  　按键名，仅对键盘事件有效；
    keysym     　　　　　　　 按键编码，仅对键盘事件有效；
    num          　　　　　　鼠标按键，仅对鼠标事件有效；
    type         　　　　    所触发的事件类型；
    widget      　　　　     引起事件的组件；
    width,heigh　　　　　　  组件改变后的大小，仅Configure有效；
    x,y       　  　　　　　　鼠标当前位置，相对于窗口；
    x_root,y_root　　　　　　 鼠标当前位置，相对于整个屏幕[![复制代码](https://common.cnblogs.com/images/copycode.gif)](javascript:void(0);

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
```

12、弹窗

messagebox._show函数的控制参数：
    default         指定消息框按钮；
    icon            指定消息框图标；
    message     　 　指定消息框所显示的消息；
    parent          指定消息框的父组件；
    title           标题；
    type            类型；
simpledialog模块参数：
    title           指定对话框的标题；
    prompt        　显示的文字；
    initialvalue    指定输入框的初始值；
　　filedialog　　　　模块参数：
    filetype   　　  指定文件类型；
    initialdir 　　  指定默认目录；
    initialfile 　　 指定默认文件；
    title    　　　  指定对话框标题
colorchooser模块参数：
    initialcolor  　 指定初始化颜色；
    title          　指定对话框标题；

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
```

13、字体（font)

一般格式：
（'Times -10 bold')
('Times',10,'bold','italic')    依次表示字体、字号、加粗、倾斜
补充：
config            重新配置
label.config(font='Arial -%d bold' % scale.get())
依次为字体，大小（大小可为字号大小），加粗
tkinter.StringVar    能自动刷新的字符串变量，可用set和get方法进行传值和取值，类似的还有IntVar,DoubleVar...
sys.stdout.flush()　　刷新输出

```
1
2
3
4
5
6
7
8
9
```

14、tkinter中的颜色

[外链图片转存中...(img-WgnzWNBo-1563334229740)]
1、简单实例

下面的代码是创建出一个窗口，其他的操作就在这个平台上进行。执行之后会在桌面弹出一个窗口，窗口的标题就是代码中设置的win.title。这里说一下，我使用的版本是python3.6。后面的内容尽量按顺序看，后面的控件也许用到前面写到的东西。

#!/usr/bin/env python
# -*- coding:utf-8 -*-
import tkinter
# 创建主窗口
win = tkinter.Tk()
# 设置标题
win.title("yudanqu")
# 设置大小和位置
win.geometry("400x400+200+50")
# 进入消息循环，可以写控件
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
```

在这里插入图片描述
2、Label控件

!/usr/bin/env python
# -*- coding:utf-8 -*-
import tkinter
win = tkinter.Tk()
win.title("yudanqu")
win.geometry("400x400+200+50")
'''
Label:标签控件,可以显示文本
'''
# win：父窗体
# text：显示的文本内容
# bg：背景色
# fg：字体颜色
# font：字体
# wraplength：指定text文本中多宽之后换行
# justify：设置换行后的对齐方式
# anchor：位置 n北，e东，w西，s南，center居中；还可以写在一起：ne东北方向
label = tkinter.Label(win,
                      text="this is a word",
                      bg="pink", 
                      fg="red",
                      font=("黑体", 20),
                      width=20,
                      height=10,
                      wraplength=100,
                      justify="left",
                      anchor="ne")
# 显示出来
label.pack()
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
```

在这里插入图片描述
3、Button控件

#!/usr/bin/env python
# -*- coding:utf-8 -*-
import tkinter
def func():
    print("aaaaaaaaaaaaaaaaaaaaaaa")
win = tkinter.Tk()
win.title("yudanqu")
win.geometry("400x400+200+50")
# 创建按钮
button1 = tkinter.Button(win, text="按钮", command=func, width=10, height=10)
button1.pack()
button2 = tkinter.Button(win, text="按钮", command=lambda: print("bbbbbbbbbbbb"))
button2.pack()
button3 = tkinter.Button(win, text="退出", command=win.quit)
button3.pack()
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
```

在这里插入图片描述
4、Entry控件

#!/usr/bin/env python
# -*- coding:utf-8 -*-
import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
'''
 Entry：输入控件，用于显示简单的文本内容
'''
# 密文显示
entry1 = tkinter.Entry(win, show="*") # show="*" 可以表示输入密码
entry1.pack()
# 绑定变量
e = tkinter.Variable()
entry2 = tkinter.Entry(win, textvariable=e)
entry2.pack()
# e就代表输入框这个对象
# 设置值
e.set("set")
# 取值
print(e.get())
print(entry2.get())
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
```

在这里插入图片描述
5、点击按钮输出输入框中的内容

!/usr/bin/env python
# -*- coding:utf-8 -*-
import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
def showinfo():
    # 获取输入的内容
    print(entry.get())
entry = tkinter.Entry(win)
entry.pack()
button = tkinter.Button(win, text="按钮", command=showinfo)
button.pack()
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
```

在这里插入图片描述
6、Text控件

#!/usr/bin/env python
# -*- coding:utf-8 -*-
import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
'''
 文本控件：用于显示多行文本
'''
# height表示的是显示的行数
text = tkinter.Text(win, width=30, height=10)
text.pack()
str = '''文本内容：盖闻天地之数，有十二万九千六百岁为一元。将一元分为十二会，乃子、丑、寅、卯、辰、巳、午、未、申、酉、戌、亥之十二支也。每会该一万八百岁。且就一日而论：子时得阳气，而丑则鸡鸣；寅不通光，而卯则日出；辰时食后，而巳则挨排；日午天中，而未则西蹉；申时晡而日落酉；戌黄昏而人定亥。譬于大数，若到戌会之终，则天地昏蒙而万物否矣。再去五千四百岁，交亥会之初，则当黑暗，而两间人物俱无矣，故曰混沌。又五千四百岁，亥会将终，贞下起元，近子之会，而复逐渐开明。邵康节曰：“冬至子之半，天心无改移。一阳初动处，万物未生时。”到此，天始有根。再五千四百岁，正当子会，轻清上腾，有日，有月，有星，有辰。日、月、星、辰，谓之四象。故曰，天开于子。又经五千四百岁，子会将终，近丑之会，而逐渐坚实。易曰：“大哉乾元！至哉坤元！万物资生，乃顺承天。”至此，地始凝结。再五千四百岁，正当丑会，重浊下凝，有水，有火，有山，有石，有土。水、火、山、石、土谓之五形。故曰，地辟于丑。又经五千四百岁，丑会终而寅会之初，发生万物。'''
text.insert(tkinter.INSERT, str)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
```

在这里插入图片描述
7、带滚动条的Text

#!/usr/bin/env python
# -*- coding:utf-8 -*-
import tkinter
win = tkinter.Tk()
win.title("title")
# win.geometry("400x400+200+50")
'''
 文本控件：用于显示多行文本
'''
# 创建滚动条
scroll = tkinter.Scrollbar()
text = tkinter.Text(win, width=30, height=10)
# side放到窗体的哪一侧,  fill填充
scroll.pack(side=tkinter.RIGHT, fill=tkinter.Y)
text.pack(side=tkinter.LEFT, fill=tkinter.Y)
# 关联
scroll.config(command=text.yview)
text.config(yscrollcommand=scroll.set)
text.pack()
str = '''text:盖闻天地之数，有十二万九千六百岁为一元。将一元分为十二会，乃子、丑、寅、卯、辰、巳、午、未、申、酉、戌、亥之十二支也。每会该一万八百岁。且就一日而论：子时得阳气，而丑则鸡鸣；寅不通光，而卯则日出；辰时食后，而巳则挨排；日午天中，而未则西蹉；申时晡而日落酉；戌黄昏而人定亥。譬于大数，若到戌会之终，则天地昏蒙而万物否矣。再去五千四百岁，交亥会之初，则当黑暗，而两间人物俱无矣，故曰混沌。又五千四百岁，亥会将终，贞下起元，近子之会，而复逐渐开明。邵康节曰：“冬至子之半，天心无改移。一阳初动处，万物未生时。”到此，天始有根。再五千四百岁，正当子会，轻清上腾，有日，有月，有星，有辰。日、月、星、辰，谓之四象。故曰，天开于子。又经五千四百岁，子会将终，近丑之会，而逐渐坚实。易曰：“大哉乾元！至哉坤元！万物资生，乃顺承天。”至此，地始凝结。再五千四百岁，正当丑会，重浊下凝，有水，有火，有山，有石，有土。水、火、山、石、土谓之五形。故曰，地辟于丑。又经五千四百岁，丑会终而寅会之初，发生万物。'''
text.insert(tkinter.INSERT, str) 
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
```

在这里插入图片描述
8、Checkbutton多选框控件

#!/usr/bin/env python
# -*- coding:utf-8 -*-
import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
def updata():
    message = ""
    if hobby1.get() == True:
        message += "variable1\n"
    if hobby2.get() == True:
        message += "variable2\n"
    if hobby3.get() == True:
        message += "variable3\n"
    # 清空text中所有内容
    text.delete(0.0, tkinter.END)
    text.insert(tkinter.INSERT, message)
# 要绑定的变量
hobby1 = tkinter.BooleanVar()
# 多选框
check1 = tkinter.Checkbutton(win, text="Checkbutton1", variable=hobby1, command=updata)
check1.pack()
hobby2 = tkinter.BooleanVar()
check2 = tkinter.Checkbutton(win, text="Checkbutton2", variable=hobby2, command=updata)
check2.pack()
hobby3 = tkinter.BooleanVar()
check3 = tkinter.Checkbutton(win, text="Checkbutton3", variable=hobby3, command=updata)
check3.pack()
text = tkinter.Text(win, width=50, height=5)
text.pack()
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
```

在这里插入图片描述
9、Radiobutton单选框

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
def updata():
    print(r.get())
# 绑定变量，一组单选框要绑定同一个变量，就能区分出单选框了
r = tkinter.IntVar()
radio1 = tkinter.Radiobutton(win, text="Radiobutton1", value=1, variable=r, command=updata)
radio1.pack()
radio2 = tkinter.Radiobutton(win, text="Radiobutton2", value=2, variable=r, command=updata)
radio2.pack()
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
```

在这里插入图片描述
10、Listbox控件一

#!/usr/bin/env python
# -*- coding:utf-8 -*-
import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
'''
 列表框控件：可以包含一个或多个文本框
 作用：在listbox控件的小窗口显示一个字符串
'''
# 创建一个listbox，添加几个元素
lb = tkinter.Listbox(win, selectmode=tkinter.BROWSE)
lb.pack()
for item in ["tkinter.END1", "tkinter.END2", "tkinter.END3", "tkinter.END4", "tkinter.END5", "tkinter.END6", "tkinter.END7"]:
    # 按顺序添加
    lb.insert(tkinter.END, item)
# 在开始添加
lb.insert(tkinter.ACTIVE, "tkinter.ACTIVE")
# 将列表当做一个元素添加
# lb.insert(tkinter.END, ["very good", "very nice"])
# 删除：参数1为开始的索引，参数2为结束的索引，如果不指定参数2，只删除第一个索引处的内容
lb.delete(1, 3)
# 选中：参数1为开始的索引，参数2为结束的索引，如果不指定参数2，只选中第一个索引处的内容
lb.select_set(2, 3)
# 取消选中：参数1为开始的索引，参数2为结束的索引，如果不指定参数2，只取消第一个索引处的内容
lb.select_clear(2)
# 获取到列表中的元素个数
print(lb.size())
# 获取值
print(lb.get(2, 3))
# 返回当前的索引项，不是item元素
print(lb.curselection())
# 判断：一个选项是否被选中
print(lb.selection_includes(3))
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
```

[外链图片转存失败(img-k47yyTan-1563417179863)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190717-160745-011.png)]
11、Listbox控件二

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# 绑定变量
lbv = tkinter.StringVar()
# 与BORWSE相似，但是不支持鼠标按下后移动选中位置
lb = tkinter.Listbox(win, selectmode=tkinter.SINGLE, listvariable=lbv)
lb.pack()
for item in ["tkinter.END1", "tkinter.END2", "tkinter.END3", "tkinter.END4", "tkinter.END5", "tkinter.END6", "tkinter.END7"]:
     # 按顺序添加
     lb.insert(tkinter.END, item)
# 打印当前列表中的选型
print(lbv.get())
# 设置选项
# lbv.set(("1","2","3"))
# 绑定事件
def myprint(event):
    # print(lb.curselection()) # 返回下标
     print(lb.get(lb.curselection()))  # 返回值
     lb.bind("<Double-Button-1>", myprint)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
```

[外链图片转存失败(img-ArR40J29-1563417179863)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190717-161412-012.png)]
12、Listbox控件三（EXTENDED：可以使listbox支持shift和Ctrl）

import tkinter
win = tkinter.Tk()
win.title("title")
# win.geometry("400x400+200+50")
# EXTENDED：可以使listbox支持shift和Ctrl
lb = tkinter.Listbox(win,
                     selectmode=tkinter.EXTENDED)
lb.pack()
for item in ["tkinter.END1", "tkinter.END2", "tkinter.END3", "tkinter.END4", "tkinter.END5", "tkinter.END6", "tkinter.END7", "tkinter.END8", "tkinter.END9", "tkinter.END10", "tkinter.END11", "tkinter.END12", "tkinter.END13", "tkinter.END14", "tkinter.END15", "tkinter.END16", "tkinter.END17", "tkinter.END18", "tkinter.END18", "tkinter.END20", "tkinter.END21", "tkinter.END22", "tkinter.END23", "tkinter.END24", "tkinter.END25", "tkinter.END26", "tkinter.END27", "tkinter.END28", "tkinter.END29", "tkinter.END30", "tkinter.END31", "tkinter.END32", "tkinter.END33", "tkinter.END34", "tkinter.END35"]:
     # 按顺序添加
     lb.insert(tkinter.END, item)
# 滚动条
sc = tkinter.Scrollbar(win)
sc.pack(side=tkinter.RIGHT,
        fill=tkinter.Y)
# 配置
lb.configure(yscrollcommand=sc.set)
lb.pack(side=tkinter.LEFT,
        fill=tkinter.BOTH)
# 额外给属性赋值
sc["command"] = lb.yview
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
```

[外链图片转存失败(img-pRyPeZNP-1563417179863)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190717-162125-013.png)]
13、Listbox四(多选：selectmode=tkinter.MULTIPLE)

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# MULTIPLE支持多选
lb = tkinter.Listbox(win, selectmode=tkinter.MULTIPLE)
lb.pack()
for item in ["tkinter.END1", "tkinter.END2", "tkinter.END3", "tkinter.END4", "tkinter.END5", "tkinter.END6", "tkinter.END7"]:
     # 按顺序添加
     lb.insert(tkinter.END, item)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
```

[外链图片转存失败(img-8k0VpPFx-1563417179863)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190717-163411-014.png)]
14、Scale控件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
'''
 供用户通过拖拽指示器来改变变量的值，可以水平，也可以竖直
'''
# tkinter.HORIZONTAL水平
# tkinter.VERTICAL 竖直（默认）
# length:水平时表示宽度，竖直时表示高度
# tickintervar ：选择值将会为该值得倍数
scale1 = tkinter.Scale(win, from_=0, to=100, orient=tkinter.HORIZONTAL, tickinterval=10, length=200)
scale1.pack()
# 设置值
scale1.set(20)
# 取值
# print(scale1.get())
def showNum():
     print(scale1.get())
tkinter.Button(win, text="Button", command=showNum).pack()
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
```

[外链图片转存失败(img-K5MA7fp0-1563417179864)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190717-165217-015.png)]
15、Spinbox控件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
'''
 数值范围控件
'''
# 绑定变量
v = tkinter.StringVar()
def updata():
    print(v.get())
# increment：步长，默认为1
# values要输入一个元组 最好不要和from和to同时使用，而且步长也没用
# command 只要值改变就会执行updata方法
sp = tkinter.Spinbox(win,
                     from_=0,
                     to=100,
                     increment=5,
                     textvariable=v,
                     command=updata)
# sp = tkinter.Spinbox(win, values=(0,2,4,6,8))
sp.pack()
# 赋值
v.set(20)
# 取值
print(v.get())
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
```

[外链图片转存失败(img-8Fx1F6YN-1563417179864)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190717-165557-016.png)]
16、Menu顶层菜单

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# 菜单条
menubar = tkinter.Menu(win)
win.config(menu=menubar)
def func():
     print("command")
# 创建一个菜单选项
menu1 = tkinter.Menu(menubar, tearoff=False)
# 给菜单选项添加内容
for item in ['add_command1','add_command2','add_command3','add_command4', 'add_command5','add_command6','add_command7','add_command8退出']:
     if item == '退出':
         # 添加分割线
         menu1.add_separator()
         menu1.add_command(label=item, command=win.quit)
     else:
         menu1.add_command(label=item, command=func)
# 向菜单条上添加菜单选项
menubar.add_cascade(label='add_cascade', menu=menu1)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
```

[外链图片转存失败(img-feR0mwEc-1563417179864)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\1.png)]
17、Menu鼠标右键菜单

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# 菜单条
menubar = tkinter.Menu(win)
def func():
     print("func")
# 菜单
menu = tkinter.Menu(menubar, tearoff=False)
# 给菜单选项添加内容
for item in ['add_command1','add_command2','add_command3','add_command4', 'add_command5','add_command6','add_command7','add_command退出']:
     if item == 'add_command退出':
         # 添加分割线
         menu.add_separator()
         menu.add_command(label=item, command=win.quit)
     else:
         menu.add_command(label=item, command=func)
menubar.add_cascade(label="add_cascade", menu=menu)
def showMenu(event):
     menubar.post(event.x_root, event.y_root)
win.bind("<Button-3>", showMenu)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
```

[外链图片转存失败(img-LxtjFa5a-1563417275772)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\2.png)]
18、Combobox下拉控件

import tkinter
from tkinter import ttk
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# 绑定变量
cv = tkinter.StringVar()
com = ttk.Combobox(win, textvariable=cv)
com.pack()
# 设置下拉数据
com["value"] = ("Combobox1", "Combobox2", "Combobox3")
# 设置默认值
com.current(0)
# 绑定事件
def func(event):
    print(com.get())
    print(cv.get())
com.bind("<<ComboboxSelected>>", func)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
```

[外链图片转存中...(img-iFYmEnX3-1563417275773)]
19、Frame控件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
'''
 框架控件：在桌面上显示一个矩形区域，多作为一个容器控件
'''
frm = tkinter.Frame(win)
frm.pack()
# left
frm_l = tkinter.Frame(frm)
tkinter.Label(frm_l, text="左上", bg="pink").pack(side=tkinter.TOP)
tkinter.Label(frm_l, text="左下", bg="blue").pack(side=tkinter.TOP)
frm_l.pack(side=tkinter.LEFT)
# right
frm_r = tkinter.Frame(frm)
tkinter.Label(frm_r, text="右上", bg="green").pack(side=tkinter.TOP)
tkinter.Label(frm_r, text="右下", bg="red").pack(side=tkinter.TOP)
frm_r.pack(side=tkinter.RIGHT)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
```

[外链图片转存中...(img-wlWJxmRh-1563417275773)]
20、表格数据

import tkinter
from tkinter import ttk
win = tkinter.Tk()
win.title("title")
win.geometry("600x400+200+50")
# 表格
tree = ttk.Treeview(win)
tree.pack()
# 定义列
tree["columns"] = ("columns1","columns2","columns3","columns4")
# 设置列，列还不显示
tree.column("columns1", width=100)
tree.column("columns2", width=100)
tree.column("columns3", width=100)
tree.column("columns4", width=100)
# 设置表头
tree.heading("columns1", text="heading1")
tree.heading("columns2", text="heading2")
tree.heading("columns3", text="heading3")
tree.heading("columns4", text="heading4")
# 添加数据
tree.insert("", 0, text="insert.text1", values=("insert.values11","insert.values12","insert.values13","insert.values14"))
tree.insert("", 1, text="insert.text2", values=("insert.values21","insert.values22","insert.values23","insert.values24"))
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
```

[外链图片转存中...(img-DHZdBSdg-1563417275773)]
21、树状数据

import tkinter
from tkinter import ttk
win = tkinter.Tk()
win.title("yudanqu")
win.geometry("400x400+200+50")
tree = ttk.Treeview(win)
tree.pack()
# 添加一级树枝
treeF1 = tree.insert("",0,"中国",text= "insert.text1", values=("F1"))
treeF2 = tree.insert("",1,"美国",text= "insert.text2", values=("F1"))
treeF3 = tree.insert("",2,"英国",text= "insert.text3", values=("F1"))
# 二级树枝
treeF1_1 = tree.insert(treeF1,0,"黑龙江",text="insert.text11",values=("F1_1"))
treeF1_2 = tree.insert(treeF1,1,"吉林",text="insert.text12",values=("F1_2"))
treeF1_3 = tree.insert(treeF1,2,"辽宁",text="insert.text13",values=("F1_3"))
treeF2_1 = tree.insert(treeF2,0,"aaa",text="insert.text21",values=("F2_1"))
treeF2_2 = tree.insert(treeF2,1,"bbb",text="insert.text22",values=("F2_2"))
treeF2_3 = tree.insert(treeF2,2,"ccc",text="insert.text23",values=("F2_3"))
# 三级树枝
treeF1_1_1 = tree.insert(treeF1_1,0,"哈尔滨",text="insert.text111")
treeF1_1_2 = tree.insert(treeF1_1,1,"五常",text="insert.text112")
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
```

[外链图片转存中...(img-7yaKP8nO-1563417275773)]
22、绝对布局

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
label1 = tkinter.Label(win, text="label1", bg="blue")
label2 = tkinter.Label(win, text="label2", bg="red")
label3 = tkinter.Label(win, text="label3", bg="green")
# 绝对布局，窗口的变化对位置没有影响
label1.place(x=10,y=10)
label2.place(x=50,y=50)
label3.place(x=100,y=100)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
```

[外链图片转存中...(img-TzQOJjHF-1563417275774)]
23、相对布局

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
label1 = tkinter.Label(win, text="label1", bg="yellow")
label2 = tkinter.Label(win, text="label2", bg="red")
label3 = tkinter.Label(win, text="label3", bg="green")
# 相对布局，窗体改变对控件有影响
label1.pack(fill=tkinter.Y, side=tkinter.LEFT)
label2.pack(fill=tkinter.X, side=tkinter.TOP)
label3.pack()
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
```

[外链图片转存中...(img-HSopweQI-1563417275774)]
24、表格布局

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
label1 = tkinter.Label(win, text="label1", bg="blue")
label2 = tkinter.Label(win, text="label2", bg="red")
label3 = tkinter.Label(win, text="label3", bg="green")
label4 = tkinter.Label(win, text="label4", bg="yellow")
# 表格布局
label1.grid(row=0,column=0)
label2.grid(row=0,column=1)
label3.grid(row=1,column=0)
label4.grid(row=1,column=1)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
```

[外链图片转存失败(img-uyR7QOdC-1563417275774)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190717-193440-023.png)]
25、鼠标点击事件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
def func(event):
     print(event.x, event.y)
# <Button-1>  鼠标左键
# <Button-2>  鼠标滚轮
# <Button-1>  鼠标右键
# <Double-Button-1>  鼠标双击左键
# <Triple-Button-1>  鼠标三击左键
button1 = tkinter.Button(win, text="leftmouse button")
# bind 给控件绑定数据（参数一是绑定的事件，参数二是触发事件的函数）
button1.bind("<Button-1>", func)
button1.pack()
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
```

[外链图片转存失败(img-ummtgAzO-1563417275774)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190717-193839-024.png)]多次点击的坐标：[外链图片转存失败(img-QdenON5S-1563417275775)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190717-193849-025.png)]
多次点击多次点击
26、鼠标移动事件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# <B1-Motion>  左键移动
# <B2-Motion>  中键移动
# <B3-Motion>  右键移动
label = tkinter.Label(win, text="Label")
label.pack()
def func(event):
     print(event.x, event.y)
label.bind("<B1-Motion>", func)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
```

[外链图片转存失败(img-vdaK8DrY-1563417275775)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190718-095625-001.png)]
27、鼠标释放事件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# <ButtonRelease-1> 释放鼠标左键
# <ButtonRelease-2> 释放鼠标中键
# <ButtonRelease-3> 释放鼠标右键
label = tkinter.Label(win, text="*********", bg="red")
label.pack()
def func(event):
    print(event.x, event.y)
label.bind("<ButtonRelease-1>", func)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
```

[外链图片转存失败(img-PYGPARxB-1563417275775)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190718-100004-002.png)]
28、进入和离开事件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# <Enter>  当鼠标进入控件时触发事件
# <Leave>  当鼠标离开控件时触发事件
label = tkinter.Label(win, text="label", bg="red")
label.pack()
def func(event):
    print(event.x, event.y)
label.bind("<Enter>", func)
label.bind("<Leave>", func)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
```

[外链图片转存失败(img-ciwcaPj8-1563417275775)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190718-100409-003.png)]
29、响应所有按键的事件

import tkinter
win = tkinter.Tk()
win.title("label")
win.geometry("400x400+200+50")
# <Key>  响应所有的按键（要有焦点）
label = tkinter.Label(win, text="label", bg="red")
# 设置焦点
label.focus_set()
label.pack()
def func(event):
     print("event.char=", event.char)
     print("event.keycode=", event.keycode)
label.bind("<Key>", func)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
```

[外链图片转存失败(img-tjZGbe3c-1563417275775)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190718-101510-004.png)]
30、响应特殊按键事件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# <Shift_L>  只响应左侧的shift键
# <Shift_R>
# <F5>
# <Return>  也就是回车键
# <BackSpace>  返回,也就是退格键
label = tkinter.Label(win, text="label", bg="red")
# 设置焦点
label.focus_set()
label.pack()
def func(event):
     print("event.char=", event.char)
     print("event.keycode=", event.keycode)
label.bind("<Shift_L>", func)
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
```

[外链图片转存失败(img-aAfpdLUA-1563417275776)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190718-102019-005.png)]
31、指定按键事件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
label = tkinter.Label(win, text="label", bg="red")
# 设置焦点
label.focus_set()
label.pack()
def func(event):
     print("event.char=", event.char)
     print("event.keycode=", event.keycode)
win.bind("a", func) # 注意前面改成了win，只需要写出按键名即可
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
```

[外链图片转存失败(img-0DxC70k4-1563417275776)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190718-102319-006.png)]
32、组合按键事件

import tkinter
win = tkinter.Tk()
win.title("title")
win.geometry("400x400+200+50")
# <Control-Alt-a>
# <Shift-Up>
# 只是control+alt不行
label = tkinter.Label(win, text="label", bg="red")
# 设置焦点
label.focus_set()
label.pack()
def func(event):
     print("event.char=", event.char)
     print("event.keycode=", event.keycode)
win.bind("<Control-Alt-a>", func) # 注意前面改成了win，只需要写出按键名即可
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
```

[外链图片转存失败(img-5z4lX15h-1563417275777)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190718-102641-007.png)]
33、更改标题‘羽毛‘图标

import tkinter
win = tkinter.Tk()
#更改标题‘羽毛’图标
win.iconbitmap('./图片路径.ico')
win.mainloop()

```
1
2
3
4
5
```

打开favicon.ico在线制作做.ico文件图片
在这里插入图片描述
34、增加背景图片

import tkinter
win = tkinter.Tk()
# 增加背景图片，注意背景图片大小应与窗口大小一致，否则填不满窗口
background_image = tkinter.PhotoImage(file="背景图片.png")
theLabel = tkinter.Label(win,
                         justify = tkinter.LEFT,# 对齐方式
                          image = background_image,# 加入背景图片
                          compound = tkinter.CENTER,# 关键:设置为背景图片
                          font = ("华文行楷", 20),# 字体和字号
                          fg = "white")  # 前景色
theLabel.pack()
win.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
```

在这里插入图片描述
35、用canvas加载背景图片

from tkinter import *
window = Tk()
window.title("登录系统")
window.geometry('450x300')
# 用画布加载图片
canvas = Canvas(window, height=200, width=500)
image_file = PhotoImage(file='background.png')
image = canvas.create_image(0, 0, anchor='nw', image=image_file)
canvas.pack(side='top')
mainloop()

```
1
2
3
4
5
6
7
8
9
10
```

[外链图片转存失败(img-dOf4PQRa-1563500010688)(C:\Users\xiahuadong\Documents\Pointofix\Screenshots\20190719-093224-001.png)]
37、quit和destory的区别

command = win2.destroy

```
1
```

只关闭win2窗口

command = win2.quit

```
1
```

关闭win2窗口的同时，关闭主窗口等所有窗口
38、tkinter+opencv摄像头拍照功能

#!/usr/bin/python
# -*- coding: UTF-8 -*-
from tkinter import *
import cv2 as cv
from PIL import Image, ImageTk
import os
#拍照类
class photo_graph():
    # 去掉文件名，返回目录
    workpath = os.path.dirname(sys.argv[0])
    os.chdir(workpath)  # 指定py文件执行路径为当前工作路径
    # 界面相关
    window_width = 640
    window_height = 480
    image_width = int(window_width * 0.6)
    image_height = int(window_height * 0.6)
    imagepos_x = int(window_width * 0.2)
    imagepos_y = int(window_height * 0.1)
    butpos_x = 250
    butpos_y = 350
    # 摄像机设置
    # 0是代表摄像头编号，只有一个的话默认为0
    capture = cv.VideoCapture(0)
    top = Tk()
    canvas = Canvas(top, bg='white', width=image_width, height=image_height)  # 绘制画布
    # 得到图片
    def getframe(self):
        ref, frame = self.capture.read()
        cv.imwrite(self.image_path, frame)
    # 关闭摄像头
    def closecamera(self):
        self.capture.release()
    # 更改标题‘羽毛’图标
    def icon_logo(self,win):
        win.iconbitmap('./gui_image/face.ico')
        win.title("有限公司")
        win.geometry("500x500")
    def tkImage(self):
        ref, frame = self.capture.read()
        cvimage = cv.cvtColor(frame, cv.COLOR_BGR2RGBA)
        pilImage = Image.fromarray(cvimage)
        pilImage = pilImage.resize((self.image_width, self.image_height), Image.ANTIALIAS)
        tkImage = ImageTk.PhotoImage(image=pilImage)
        return tkImage
    # 拍照按钮，写入图片
    def Photograph(self):
        # 图片文件夹路径
        path = './data/123'
        # path文件夹下的文件列表
        files = os.listdir(path)
        # 列表长度+1
        l = len(files) + 1
        # 图片路径（为了可以多次拍照，否则只能拍一张照片）
        image_path = path + '/'+ '_' + str(l) + '.png'
        print('image_path:' + image_path)
        ref, frame = self.capture.read()
        cv.imwrite(image_path, frame)
    #创建窗口
    def create_win(self):
        self.icon_logo(win=self.top)
        self.top.wm_title("上海芯灵科技有限公司")
        self.top.geometry(str(self.window_width) + 'x' + str(self.window_height))
        # 画布定义
        b = Button(self.top, text='拍照', width=15, height=2, command = self.Photograph)
        # 画布位置设置
        self.canvas.place(x=self.imagepos_x, y=self.imagepos_y)
        b.place(x=self.butpos_x, y=self.butpos_y)
        while (True):
            picture = self.tkImage()
            # 绘制图片，支持GIF
            self.canvas.create_image(0, 0, anchor='nw', image=picture)
            # 刷新tkinter任务
            self.top.update()
            # 等待
            self.top.after(100)
        top.mainloop()
        closecamera()
p=photo_graph()
p.create_win()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64
65
66
67
68
69
70
71
72
73
74
75
76
77
78
79
```

在这里插入图片描述
tkinter上传照片
tkinter上传照片并显示

import tkinter
import tkinter.filedialog
top = tkinter.Tk()
top.title = 'new'
top.geometry('640x480')
def choose_fiel():
    selectFileName = tkinter.filedialog.askopenfilename(title='选择文件')  # 选择文件
    e.set(selectFileName)
e = tkinter.StringVar()
e_entry = tkinter.Entry(top, width=68, textvariable=e)
e_entry.pack()
def upload_func(a):
    '''
    要自己写个方法，ftp等方法，上传文件到服务器
    '''
    print(a)
    pass
submit_button = tkinter.Button(top, text ="选择文件", command = choose_fiel)
submit_button.pack()
submit_button = tkinter.Button(top, text ="上传", command = lambda:upload_func(e_entry.get()))
submit_button.pack()
from PIL import Image, ImageTk
def showImg(img1):
    load = Image.open(img1)
    render = ImageTk.PhotoImage(load)
    img = tkinter.Label(image=render)
    img.image = render
    img.place(x=200, y=100)
submit_button = tkinter.Button(top, text ="显示图片", command = lambda :showImg(showImg('服务器上的pic路径')))
submit_button.pack()
top.mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
```

tkinter上传照片到网上

# -*- coding: UTF-8 -*-
#上传到网上
from tkinter import *
import tkinter.filedialog
import requests
def Upload():
    print('upload')
    selectFileName = tkinter.filedialog.askopenfilename(title='选择文件')  # 选择文件
    #上传路径
    r = requests.post('http://127.0.0.1:8000/upload', files={'file': open(selectFileName, 'rb')})
    print(r.content.decode('utf-8'))
    setText = r.content.decode('utf-8')
    print(setText.__class__)
    e1.delete(0, END)
    e1.insert(0, setText)
def Download():
    link = e1.get()
    files = requests.get(link)
    files.raise_for_status()
    path = tkinter.filedialog.asksaveasfilename()
    print(files.content)
    with open(path, 'wb') as f:
        f.write(files.content)
root = Tk()
root.title('Download')
root.geometry('+500+300')
e1 = Entry(root, width=50)
e1.grid(row=0, column=0)
btn1 = Button(root, text=' 上传 ', command=Upload).grid(row=1, column=0, pady=5)
btn2 = Button(root, text=' 下载 ', command=Download).grid(row=2, column=0, pady=5)
btn3 = Button(root, text=' 复制 ', ).grid(row=3, column=0, pady=5)
mainloop()

```
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
```

tkinter上传（移动）照片到本地

# -*- coding: UTF-8 -*-
from tkinter import *
import tkinter.filedialog
import shutil
def Upload():
    print('upload')
    #源文件的路径
    selectFileName = tkinter.filedialog.askopenfilename(title='选择文件')  # 选择文件
    #目标文件夹的路径
    destination = r"C:\Users\xiahuadong\Pictures"
    #复制文件selectFileName到destination_path
    shutil.copy(selectFileName, destination_path)
root = Tk()
root.title('Download')
root.geometry('+500+300')
e1 = Entry(root, width=50)
e1.grid(row=0, column=0)
btn1 = Button(root, text=' 上传 ', command=Upload).grid(row=1, column=0, pady=5)
mainloop()
————————————————
版权声明：本文为CSDN博主「夏华东的博客」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_44493841/article/details/96147625