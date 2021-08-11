# python tkinter 窗口最大化





我是在linux下折腾的，python版本是python3.4.3+，linux用的是ubuntu kylin 发行版本，是一个中文汉化版，做的不错。

##### 方法1

root = Tk()

w, h = root.maxsize()

root.geometry("{}x{}".format(w, h)) #看好了，中间的是小写字母x

##### 方法2

root = Tk()

w = root.winfo_screenwidth()

h = root.winfo_screenheight()

root.geometry("%dx%d" %(w, h))

上面两个是比较常见的方法，应该是在windows和linux下都可以使用的。

##### 方法3

root = Tk()

screen = os.popen("xrandr | grep current")

cur = screen.read().split(',')[1].split(' ')

root.geometry(cur[2]+cur[3]+cur[4])

这个方法就只能在linux下使用了，它是利用了xrandr命令，这个命令可以查看和设置当前的屏幕分辨率。

##### 方法4

root = Tk()

root.state("zoomed")

这个方法只能在windows下使用，原因是”zoomed”这个参数只能在windowns下使用。

##### 方法5

root = Tk()

root.attributes("-fullscreen", true)





这个方法是设置root窗口的全屏属性为真，可以实现最大化，但是没有标题栏…

我找到的方法就这么多，希望可以帮到大家，也是给自己做个总结免得忘了…