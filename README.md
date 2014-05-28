GNVM: Node.exe version manager for Windows by GO
================================
`gnvm` 是个Windows下面的多Node.exe版本管理工具，类似 [nvm](https://github.com/creationix/nvm) [nvmw](https://github.com/hakobera/nvmw)

下载
---
<http://pan.baidu.com/s/1jGgdjiI>

配置
---

#### 本机已有node.exe
* 方式1: 将下载的`gnvm.exe`放到`node.exe`目录下。（推荐方式）
* 方式2: 将下载的`gnvm.exe`放到任意文件夹下。（确保此文件夹在Path环境下，或者手动添加此文件夹到Path）

#### 本机没有node.exe
* 将下载的`gnvm.exe`放到任意文件夹下。（确保此文件夹在Path环境下，或者手动添加此文件夹到Path）

验证
---
在cmd下（如是Win7/8系统，确保获取管理员权限）输入：`gnvm version`，如有`Current version x.x.x`则说明配置成功。（注：`x.xx.xx`以下载的版本为准。）

使用
---

    Usage:
      gnvm
      gnvm [command]

    Available Commands:
      version                   :: Print the version number of gnvm.exe
      install                   :: Install any node.exe version
      uninstall                 :: Uninstall local node.exe version
      use                       :: Use any version of the local already exists
      update                    :: Update latest node.exe
      ls                        :: List show all <local> <remote> node.exe version
      node-version              :: Show <global> <latest> node.exe version
      config                    :: Setter and getter registry
      help [command]            :: Help about any command

最佳实践
---
* 在cmd中运行`gnvm.exe`需要管理员权限。
* 第一次运行`gnvm.exe`时，使用`gnvm config INIT`
* 虽然`gnvm.exe`支持任意文件夹，但建议将`node.exe`与`gnvm.exe`放在同一目录下。
* 使用`gnvm config registry xxx`更换库，默认库：<http://nodejs.org/dist/>，只要xxx的结构与默认库一致就可以使用。
* `gnvm.exe`的使用依赖与`.gnvmrc`，请不要手动修改此文件。
* `gnvm install npm`支持安装最新版的npm，但`0.1.0`版本只支持安装最新版本到`node.exe`所在文件夹，不可自定义`npm`的文件夹。（npm的最新版本取决于`gnvm config registry`对应的最新版本。）

使用场景之一（本机已有node.exe）
---
    gnvm config INIT （第一次使用时，推荐做法）
    gnvm config registry dist.u.qiniudn.com （更换库）
    gnvm update latest （如果本机的latest过低，可以使用此方式升级。或者使用gnvm install latest）
    gnvm install 0.11.1 0.11.2 0.11.3 （下载任意版本的node.exe）
    gnvm use 0.11.1 （切换本机已安装的任意版本node.exe）
    gnvm ls （查看当前共有多少个node.exe）
    gnvm uninstall 0.11.1 （删除0.11.1）

使用场景之二（本机没有node.exe）
---
    gnvm config INIT （第一次使用时，推荐做法）
    gnvm config registry dist.u.qiniudn.com （更换库）
    gnvm install latest -g （下载最新版本的latest并设置为全局node.exe）
    gnvm ls （查看当前共有多少个node.exe）
    gnvm install npm （安装最新版本的npm到node.exe所在目录）


使用的第三方lib
---
* <https://github.com/spf13/cobra>
* <https://github.com/tsuru/config>
* <https://github.com/pierrre/archivefile/zip>
* <https://github.com/daviddengcn/go-colortext>
* icon <http://www.easyicon.net/1143807-update_icon.html>

FAQ
---

#### Q. 在某些Win 7环境下，如果安装了XXX卫士，那么在使用`gnvm`的时候，如`gnvm use x.xx.xx`会弹出警告。
A. 建议将`gnvm.exe`加入白名单。

LICENSE
---
(The MIT License)

Copyright (c) 2014 Kenshin Wang <kenshin@ksria.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
