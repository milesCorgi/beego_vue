## 后端搭建
1. 开个项目文件夹，GoPATH设定为该文件夹路径<br>
2. 文件夹路径下拿beego<br>

```
go get github.com/astaxie/beego
```
3. 进src路径，再开一个项目名文件夹，再进这个文件夹的路径，new个beego项目
实例：new前的文件树长这样：

```
├─bin
│      bee.exe
│
└─src
    ├─beego_vue

```

输入以下指令
```
E:\gitProject\beego_vue\src\beego_vue>bee new backend
```
4. 在新建立的main.go下，如果用goLand，现在可以下beego源码了
5. 下好以后在backend文件夹执行

```
bee run
```
应该现在可以看到beego的HelloWorld了（小蜜蜂好贱）
![image](14F7628DA23A4BA1B0EEF91632639A2E)

## 前端搭建，用Vue

- npm安装vue-cli脚手架工具(npm事先电脑装好)

```
npm install -g vue-cli
```

- 在project项目根目录下，新建一个前端工程目录

```
vue-init webpack [前端文件夹名]  //安装中把vue-router选上，我们须要它来做前端路由
```

- 进入appfront目录，安装vue所须要的node依赖

```
  npm install
```

。。。。。。。

- 写好前端后打包，不打包的话Django是收不到Vue这边的网页的

```
npm run build
```

调试过程可以走dev来看界面，不过涉及数据库的话还是有打包一下走Django


```
npm run dev
```

## 前后端结合
在vue的conf文件夹的index.js文件以下段修改如下

```
  build: {
    // Template for index.html
    // 这里改成beego那边放index.html的位置
    index: path.resolve(__dirname, '../../backend/views/index.html'),

    // Paths
    // 这里改成beego文件夹的位置
    assetsRoot: path.resolve(__dirname, '../../backend'),
    assetsSubDirectory: 'static',
    assetsPublicPath: '/',
```

修改后vue构建build出的js和html就都会直接放到beego的文件夹里了