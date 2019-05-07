# 制作项目首页的header和Tabar

1. 使用mint-ui的header组件


1.1 首先使用mint-ui之前需要引入mint-ui的组件

```
// 按需导入 Mint-UI 中的组件   
import { Header, Swipe, SwipeItem } from 'mint-ui'
Vue.component(Header.name, Header)
```

然后在App.vue中编写header组件

```
<div class="app-container">
        <!-- 顶部 Header 区域 -->
    <mt-header fixed title="程序员·Vue项目"></mt-header>
    <h1>123</h1>
  </div>
```

观察上面的123显示,发现没有123 ![](./img/01.png)

可以发现123是直接顶到顶部了，可以设置padding来迫使123距离顶部一段距离，高度是header组件的高度


 ![](./img/02.png)


 # 