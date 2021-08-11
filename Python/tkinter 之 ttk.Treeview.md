# tkinter 之 ttk.Treeview

**树视图**（ttk.Treeview）小部件可以显示并允许浏览 Item 的层次结构，并可以将每个 Item 的一个或多个属性显示为树右侧的列。它允许您构建类似于在文件管理器（如 OS X 查找器或 Windows  资源管理器）中找到的树显示的用户界面。创建的方式是：

```
tree = ttk.Treeview(parent)
```

如果需要，可以按常规方式添加水平和垂直**滚动条**。

## 1 添加 Item 到 Tree

要对树视图执行任何有用的操作，您需要向其中添加一个或多个项。每个项目表示树中的单个节点，无论是叶节点还是内部节点。项目由唯一 ID 引用；这可以在首次创建项目时由程序员分配，或者小部件可以自动为项目选择 ID。

通过使用树视图的 "insert" 方法将 item 插入到树中来创建项。若要插入项，需要知道其在树中的位置，这意味着指定父项以及新项应在父项的子项列表中插入的位置。

Treeview 小部件会自动创建一个根节点（未被显示），其 ID 为 `"{}"`（即空字符串），可以用作添加的第一个级别项的父级。父项的子项列表中的位置由索引指定（0 是第一个，指定 `"end"` 索引表示在所有现有子级之后插入）。

通常，您还将指定显示在树中的每个项的名称。名称旁边还有其他选项可以添加图像，指定节点是打开还是关闭，等等。

方法 `insert` 的使用方法是：`tree.insert(parent, index, iid=None, **kw)`，可以看使用的例子：

![img](https://upload-images.jianshu.io/upload_images/1114626-7fc87ed67bc31c96.png?imageMogr2/auto-orient/strip|imageView2/2/w/1114)

图1 insert 函数的使用范例

从图1 可以看出如果指定 `iid`，则 `insert` 方法会返回 `iid` 名称，否则返回被自动分配的 `iid` 的值。

![img](https://upload-images.jianshu.io/upload_images/1114626-6779785a5996b3a8.png?imageMogr2/auto-orient/strip|imageView2/2/w/420)

图2 树视图的例子

## 2 重新排列项目

可以将节点（及其后代（如果有）移动到树中的不同位置；将节点及其后代（如果有）移动到树中的不同位置。唯一的限制是节点不能在其子级下移动。目标位置通过父位置指定，并索引到子项列表中，就像使用 "insert" 所做的那样。比如：

```
 # 移动 widgets 到 gallery 之下
tree.move('widgets', 'gallery', 'end') 
```

可以将项从树中**分离**（detached ），该树从层次结构中删除项及其后代，但不会销毁项，从而允许您稍后使用"move"重新插入它们。

```
tree.detach('widgets')
```

也可以使用 `delete` 删除项目：

```
tree.delete('widgets')
```

如果要导航层次结构，有一些方法允许您查找项（`"parent"`）的父项，查找项的下一个或以前的同级项（`"next"` 和 `"prev"`），并返回项（"children"）的子项列表。

您可以通过修改 `"open"` 项配置选项来控制项目是否处于打开状态并显示其子项。

```
tree.item('widgets', open=True)
isopen = tree.item('widgets', 'open')
```

## 3 显示每个项目的信息

树视图还可以显示有关每个项目的一个或多个附加信息，这些信息显示为主树显示右侧的列。

同样，每列都由您分配的符号名称引用。无论是在首次创建小部件时，还是以后创建小部件,您均可以使用树视图小部件的 `"columns"` 配置选项指定 columns 的列表。

```
tree = ttk.Treeview(root, columns=('size', 'modified', 'owner'))
# 等价于
tree['columns'] = ('size', 'modified', 'owner')
```

您可以指定列的宽度、列中项信息的显示方式对齐等。还可以提供有关列标题的信息，例如要显示的文本、可选图像、对齐方式以及单击项时要调用的脚本（例如，对树进行排序）。

```
tree.column('size', width=100, anchor='center')
tree.heading('size', text='Size')
```

可以单独指定要在每个项的每一列中放置的值，也可以通过提供项的值列表来指定。在后一种情况下，这是使用 "values"  项配置选项（因此可以在首次插入项时或更高版本使用），该选项采用值列表;列表的顺序必须与小部件配置选项 "columns" 中的顺序相同。

```
tree.set('widgets', 'size', '12KB')
size = tree.set('widgets', 'size')
tree.insert('', 'end', text='Listbox', values=('15KB Yesterday mark'))
```

## 4 项目外观和事件

与文本和画布小部件一样，树视图小部件使用标记来帮助修改树中项的外观。您可以使用 "tags" 项配置选项（在创建项时或以后创建项时）为每个项分配标记列表。

 然后可以指定标记配置选项，然后应用于具有该标记的所有项目。有效的标记选项包括 "foreground"（文本颜色）, "background", "font" 和 "image"（如果项目指定自己的图像）。

您还可以在标记上创建事件绑定，从而捕获鼠标单击、键盘事件等。

```
tree.insert('', 'end', text='button', tags=('ttk', 'simple'))
tree.tag_configure('ttk', background='yellow')
 # the item clicked can be found via tree.focus()
tree.tag_bind('ttk', '<1>', itemClicked) 
```

树视图将生成虚拟事件"<TreeviewSelect>", "<TreeviewOpen>" 和  "<TreeviewClose>" 允许您监视对用户所做的小部件的更改。您可以使用 "selection"  方法来确定当前选择（也可以从程序更改所选内容）。

## 5 自定义显示

treeview 小部件的显示方式有很多，您可以对其进行自定义。其中一些我们已经看到，如项目的文本，字体和颜色，列标题的名称，以及更多。下面是一些额外的。

- 使用 "height" 小部件配置选项指定要显示的所需行数。

- 使用列的 "width" 或者 "minwidth"  选项控制每列的宽度。可以使用符号名称 "#0" 访问持有树的列。小部件请求的总宽度基于列宽度的总和。

- 使用小部件配置选项 "displaycolumns" 选择要显示的列和显示顺序。

- 您可以选择使用小部件配置选项 "show"（默认为 "tree headings"  以同时显示两个标题）隐藏一个或两个列标题或树本身（仅保留列）。

   您可以指定用户是否可以通过小部件配置选项 "selectmode" 选择单个项目或多个项，传递 "browse"（单个项目）、"extended"（多个项目、默认值）或 "none"。