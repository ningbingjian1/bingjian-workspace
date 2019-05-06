# 安装mint-ui

```
npm install mint-ui
```

![](./img/01.png)

# main.js引入mint-ui

```
// 按需导入 Mint-UI组件
import { Button } from 'mint-ui'
// 使用 Vue.component 注册 按钮组件
Vue.component(Button.name, Button)
```

# app.vue使用mint-ui的组件



## 按钮

首先引入mint-ui按钮组件

main.js

```
// 按需导入 Mint-UI组件
import { Button } from 'mint-ui'
// 使用 Vue.component 注册 按钮组件
Vue.component(Button.name, Button)

```
在根组件app.vue中创建几个按钮组件


```

    <!-- 为什么这里叫做 mt-button 的 button 直接就能用 -->
    <mt-button type="danger" icon="more" @click="show">default</mt-button>

    <mt-button type="danger" size="large" plain>default</mt-button>

    <mt-button type="danger" size="small" disabled>default</mt-button>
```

效果:
![](./img/02.png)


## Toast组件

app.vue中引入Toast组件

```
// 按需导入 Toast 组件
import { Toast } from "mint-ui";
```

使用toast组件

```
// 按需导入 Toast 组件
import { Toast } from "mint-ui";

export default {
  data() {
    return {
      toastInstanse: null
    };
  },
  created() {
    this.getList();
  },
  methods: {
    getList() {
      // 模拟获取列表的 一个 AJax 方法
      // 在获取数据之前，立即 弹出 Toast 提示用户，正在加载数据
      this.show();
      setTimeout(() => {
        //  当 3 秒过后，数据获取回来了，要把 Toast 移除
        this.toastInstanse.close();
      }, 3000);
    },
    show() {
      // Toast("提示信息");
      this.toastInstanse = Toast({
        message: "这是消息",
        duration: -1, // 如果是 -1 则弹出之后不消失
        position: "top",
        iconClass: "glyphicon glyphicon-heart", // 设置 图标的类
        className: "mytoast" // 自定义Toast的样式，需要自己提供一个类名
      });
    }
  }
};
```

toast组件的使用在show方法中

效果:
![](./img/03.png)


## 修改toast的样式:

首先在main.js引入样式
```
// 导入bootstrap样式
import 'bootstrap/dist/css/bootstrap.css'
import './css/app.css'

```

app.css
```
.mytoast i{
  color:blue !important;
}
```

使用自己定义的样式:

```
  // Toast("提示信息");
      this.toastInstanse = Toast({
        message: "这是消息",
        duration: -1, // 如果是 -1 则弹出之后不消失
        position: "top",
        iconClass: "glyphicon glyphicon-heart", // 设置 图标的类
        className: "mytoast" // 自定义Toast的样式，需要自己提供一个类名
      });
```

这样显示出来的toast，颜色就变了

![](./img/04.png)







