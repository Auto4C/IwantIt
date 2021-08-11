# UCI指令

1.UCI命令



一个众所周知的原因，在Linux下各种软件包有各种不同的配置脚本，每个配置脚本的语法格式和操作方式不同，

这样的设计虽然可以体现出各软件包自身的优势，同时也增加了学习曲线。在这一点上OpenWrt的UCI无疑处理的更

胜一筹。UCI是集中式配置信息管理接口(Unified Configuration Interface)的缩写，他是OpenWrt引进的一套配置

参数管理系统。UCI管理了OpenWrt下最主要的系统配置参数并且提供了简单、容易、标准化的人机交互接口。UCI中

已经包含了网络配置、无线配置、系统信息配置等作为基本路由器所需的主要配置参数。同时UCI也可以帮助开发人

员快速的建立一套基于OpenWrt的智能路由产品控制界面。



2.UCI的文件和流程



UCI的配置文件全部存储在/etc/config目录下。

1. root@OpenWrt:/# ls /etc/config/
2. dhcp        dropbear  firewall  network    system     wireless





日前已有大量软件包支持UCI模式管理，但不是所有的软件包，支持的软件包是这样来完成

启动的(以samba举例):

1.启动脚本/etc/init.d/samba

2.启动脚本通过UCI分析库从/etc/config/samba获得启动参数

3.启动脚本完成正常启动



  由于UCI的数据文件较为简单，并且具备了很nice的直接观感，所以配置文件既可以使用UCI

命令进行修改，也可以使用VI编辑器直接修改文件。但如果两种方式都是用时需要注意UCI命

令修改会产生缓存，每次修改好要尽快确认保存避免出现冲突。

 最常见的几个UCI配置作用说明

| 文件                 | 作用                              |
| -------------------- | --------------------------------- |
| /etc/config/dhcp     | 面向LAN口提供的IP地址分配服务配置 |
| /etc/config/dropbear | SSH服务配置                       |
| /etc/config/firewall | 路由转发，端口转发，防火墙规则    |
| /etc/config/network  | 自身网络接口配置                  |
| /etc/config/system   | 时间服务器时区配置                |
| /etc/config/wireless | 无线网络配置                      |





3.UCI的文件语法

UCI文件语法举例

1. config 'section-type' 'section'
2. ​     option 'key'    'value'
3. ​     list  'list_key' 'list_value'
4. config 'example' 'test'
5. ​     option 'string'     'some value'
6. ​     option 'boolean'    '1'
7. ​     list  'collection'  'first item'
8. ​     list  'collection'  'second item'







config 节点 以关键字 config 开始的一行用来代表当前节点

​        section-type 节点类型

​        section 节点名称

option 选项 表示节点中的一个元素

​        key 键

​        value 值

list 列表选项 表示列表形式的一组参数。

​       list_key 列表键

​       list_value 列表值



config 节点语法格式

1. config 'section-type' 'section'





config 节点(后文统一称为节点)原则

​       UCI 允许只有节点类型的匿名节点存在

​       节点类型和名字建议使用单引号包含以免引起歧义

​       节点中可以包含多个 option 选项或 list 列表选项。

​       节点遇到文件结束或遇到下一个节点代表完成。

option 选项语法格式

1. option 'key' 'value'





option 选项(后文统一称为选项)原则

​       选项的键与值建议使用单引号包含

​       避免相同的选项键存在于同一个节点,否则只有一个生效

list 列表选项语法格式

1. list 'list_key' 'list_value'





list 列表选项(后文统一称为列表)原则

​    选项的键与值建议使用单引号包含

​    列表键的名字如果相同,则相同键的值将会被当作数组传递给相应软件

UCI 的语法容错

1. option example  value
2. option 'example'  value
3. option example  "value"
4. option "example" 'value'
5. option 'example'  "value"





UCI 无法容忍的语法

1. option 'example" "value'
2. option example some value with space





尽量使用常规字符去处理器 UCI,特殊字符有可能会破坏数据结构的完整性。



4.UCI 命令读写配置

语法格式

1. uci [<options>] <command> [<arguments>]





读写规则

​    UCI 读取总是先读取内存中的缓存,然后再读取文件中的

​    进行过增加,修改,删除操作后要执行生效指令,否则所做修改只存留在缓存中

读取类语法

取得节点类型

1. uci get <config>.<section>





取得一个值

1. uci get <config>.<section>.<option>





显示全部 UCI 配置

1. uci show





显示指定文件配置

1. uci show <config>





显示指定节点名字配置

1. uci show <config>.<section>





显示指定选项配置

1. uci show <config>.<section>.<option>





显示尚未生效的修改记录

1. uci changes <config>





匿名节点显示(如果所显示内容有匿名节点,使用-X 参数可以显示出匿名节点的 ID)

1. uci show -X <config>.<section>.<option>





写入类语法

增加一个匿名节点到文件

1. uci add <config> <section-type>





增加一个节点到文件中

1. uci set <config>.<section>=<section-type>





增加一个选项和值到节点中

1. uci set <config>.<section>.<option>=<value>





增加一个值到列表中

1. uci add_list <config>.<section>.<option>=<value>





修改一个节点的类型

1. uci set <config>.<section>=<section-type>





修改一个选项的值

1. uci set <config>.<section>.<option>=<value>





删除指定名字的节点

1. uci delete <config>.<section>





删除指定选项

1. uci delete <config>.<section>.<option>





删除列表

1. uci delete <config>.<section>.<list>





删除列表中一个值

1. uci del_list <config>.<section>.<option>=<string>





生效修改(任何写入类的语法,最终都要执行生效修改,否则所做修改只在缓存中,切记!)

1. uci commit <config>







5.综合实例