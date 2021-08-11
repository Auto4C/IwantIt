# pyftpdlib中文

\# from pyftpdlib.filesystems import AbstractedFS

from pyftpdlib.handlers import FTPHandler

from pyftpdlib.servers import FTPServer

from pyftpdlib.authorizers import DummyAuthorizer













class quickftp:

  '''

  为支持中文，修改了pyftplib的编码，改动如下：

  from pyftpdlib.filesystems import AbstractedFS

  1》 filesystem.py在类 AbstractedFS 中增加属性：encoding

  def __init__(self, root, cmd_channel,encoding):

  。。

  。。

  \###CHJ begin

​    self.__encoding = encoding



  @property

  def encoding(self):

​    return self.__encoding



  @encoding.setter

  def encoding(self,value):

​    self.__encoding=value

  \###CHJ end

  2》将filesystem.py文件中所有'utf8'替换为：self.encoding（共4个）

  3》在类 FTPHandler 中同样增加属性encoding

  def __init__(self, conn, server, ioloop=None):

  。。

  。。

​    self.__encoding = 'utf8'

  @property

  def encoding(self):

​    return self.__encodingg

  @encoding.setter

  def encoding(self,value):

​    self.__encoding=value

​    self.fs.encoding=value

  4》在方法中 def handle_auth_success（。。。）修改一下（因为改了AbstractFS构造函数）：

​    self.fs = self.abstracted_fs(home, self,self.encoding)

  5》将handlers.py文件中所有'utf8'替换为：self.encoding（共7个）

  5》使用：定义好handler之后，将这赋给ftpserver前：

​    handler.encoding='gbk'

  '''

  def __init__(self,port,user="admin",pwd="admin"):

​    reportdir = getnewdir("./report/")

​    authorizer = DummyAuthorizer()

​    authorizer.add_user(user, pwd, reportdir, perm='elradfmwM')

​    handler = FTPHandler

​    handler.encoding = 'gbk' # 修改pyftplib增加的编码支持

​    handler.authorizer = authorizer

​    server = FTPServer(('0.0.0.0', port), handler)

​    _thread.start_new_thread(server.serve_forever,())

​    \# server.serve_forever()