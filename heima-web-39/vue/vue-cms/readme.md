

# 首页头部制作

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

```
<style lang="scss" scoped>
.app-container {
  padding-top: 40px;
  overflow-x: hidden;
}
</style>

```


 ![](./img/02.png)


 # 制作底部的 Tabbar 区域，使用的是 MUI 的 Tabbar.html


## 引入html

原始的html

```
		<nav class="mui-bar mui-bar-tab">
			<a class="mui-tab-item mui-active" href="#tabbar">
				<span class="mui-icon mui-icon-home"></span>
				<span class="mui-tab-label">首页</span>
			</a>
			<a class="mui-tab-item" href="#tabbar-with-chat">
				<span class="mui-icon mui-icon-email"><span class="mui-badge">9</span></span>
				<span class="mui-tab-label">消息</span>
			</a>
			<a class="mui-tab-item" href="#tabbar-with-contact">
				<span class="mui-icon mui-icon-contact"></span>
				<span class="mui-tab-label">通讯录</span>
			</a>
			<a class="mui-tab-item" href="#tabbar-with-map">
				<span class="mui-icon mui-icon-gear"></span>
				<span class="mui-tab-label">设置</span>
			</a>
		</nav>
```

经过改造后的html


 ```
     <!-- 中间的 路由 router-view 区域 -->
		<transition>
			<router-view></router-view>
		</transition>


    <!-- 底部 Tabbar 区域 -->
    <nav class="mui-bar mui-bar-tab">
			<router-link class="mui-tab-item" to="/home">
				<span class="mui-icon mui-icon-home"></span>
				<span class="mui-tab-label">首页</span>
			</router-link>
			<router-link class="mui-tab-item" to="/member">
				<span class="mui-icon mui-icon-contact"></span>
				<span class="mui-tab-label">会员</span>
			</router-link>
			<router-link class="mui-tab-item" to="/shopcar">
				<span class="mui-icon mui-icon-extra mui-icon-extra-cart">
					<span class="mui-badge">0</span>
				</span>
				<span class="mui-tab-label">购物车</span>
			</router-link>
			<router-link class="mui-tab-item" to="/search">
				<span class="mui-icon mui-icon-search"></span>
				<span class="mui-tab-label">搜索</span>
			</router-link>
		</nav>
 ```

 4个跳转:

 <font color=red>
 /home
 /member
 /shopcar
 /search
</font>



##  在app.vue中间引入router-view

 ```
  <!-- 中间的 路由 router-view 区域 -->
		<transition>
			<router-view></router-view>
		</transition>
 ```

## 引入mui的扩展css
需要用到
./lib/mui/css/icons-extra.css和./lib/mui/font/mui-icons-extra.ttf
其中icons-extra.css引用了ttf

确保在main.js引入
 ```
// 导入扩展图标样式
import './lib/mui/css/icons-extra.css'
 ```


注意:
购物车小图标的样式:

```
mui-icon-extra mui-icon-extra-cart
```

 ##  创建tabar上对应的4个组件

```
src\components\tabbar\HomeContainer.vue
src\components\tabbar\MemberContainer.vue
src\components\tabbar\SearchContainer.vue
src\components\tabbar\ShopcarContainer.vue
```

HomeContainer.vue

```
<template>
  <div>
   <h1> HomeContainer.Vue </h1>
  </div>
</template>
```


MemberContainer.vue

```
<template>
  <div>
    <h1>MemberContainer.Vue</h1>
  </div>
</template>
```

 

SearchContainer.vue

```
<template>
  <div>
    <h1>SearchContainer.Vue</h1>
  </div>
</template>
```


ShopcarContainer.vue

```
<template>
  <div>
    <h1>ShopcarContainer.Vue</h1>
  </div>
</template>
```

## 路由高亮

```
linkActiveClass: 'mui-active' // 覆盖默认的路由高亮的类，默认的类叫做 router-link-active
```

 ##  添加路由跳转功能

 在router.js添加路由跳转

 ```
 import VueRouter from 'vue-router'

// 导入对应的路由组件
import HomeContainer from './components/tabbar/HomeContainer.vue'
import MemberContainer from './components/tabbar/MemberContainer.vue'
import ShopcarContainer from './components/tabbar/ShopcarContainer.vue'
import SearchContainer from './components/tabbar/SearchContainer.vue'

// 3. 创建路由对象
var router = new VueRouter({
  routes: [ // 配置路由规则
    { path: '/', redirect: '/home' },
    { path: '/home', component: HomeContainer },
    { path: '/member', component: MemberContainer },
    { path: '/shopcar', component: ShopcarContainer },
    { path: '/search', component: SearchContainer }
  ],
  linkActiveClass: 'mui-active' // 覆盖默认的路由高亮的类，默认的类叫做 router-link-active
})

// 把路由对象暴露出去
export default router
 ```


 ## 

 ## 添加切换tabbar的动画跳转功能


 ```
 <style lang="scss" scoped>
.app-container {
  padding-top: 40px;
  overflow-x: hidden;
}

.v-enter {
  opacity: 0;
  transform: translateX(100%);
}

.v-leave-to {
  opacity: 0;
  transform: translateX(-100%);
  position: absolute;
}

.v-enter-active,
.v-leave-active {
  transition: all 0.5s ease;
}
</style>
 ```

<font color=red>
 注意: 这里的.v-enter .v-leave-to分开配置动画
 .v-enter：在X的方向最右边以外部作为起点开始进入
 .v-leave-to:在X的方向的最左边离开

 </font>


 # 制作首页轮播图布局

 ## 引入轮播组件mt-swipe

 ```
     <mt-swipe :auto="4000">
      <mt-swipe-item>1</mt-swipe-item>
      <mt-swipe-item>2</mt-swipe-item>
      <mt-swipe-item>3</mt-swipe-item>
    </mt-swipe>
 ```

 ## 轮播图样式设置

 轮播图直接引入的时候不能显示 ，其实是高度设置问题，只要把高度设置到正常值句可以了

 ```
 .mint-swipe {
  height: 200px;

  .mint-swipe-item {
    &:nth-child(1) {
      background-color: red;
    }
    &:nth-child(2) {
      background-color: blue;
    }
    &:nth-child(3) {
      background-color: cyan;
    }

    img {
      width: 100%;
      height: 100%;
    }
  }
}
 ```

 注意 这里轮播图的图片是从服务器读取的 img样式设置为占满整个div

 ## 从服务器获取轮播图数据

首先main.js必须引入vue-resource

 ```
// 2.1 导入 vue-resource
import VueResource from 'vue-resource'
// 2.2 安装 vue-resource
Vue.use(VueResource)

 ```

 配置vue-resource的全局访问域名,放在引入vue-resource的代码后面

 ```
 // 设置请求的根路径
Vue.http.options.root = 'http://www.escook.cn:3000/';
 ```


 通过vue-resource访问服务器端数据


 ```
 import { Toast } from "mint-ui";

export default {
  data() {
    return {
      lunbotuList: [] // 保存轮播图的数组
    };
  },
  created() {
    this.getLunbotu();
  },
  methods: {
    getLunbotu() {
      // 获取轮播图数据的方法
      this.$http.get("api/getlunbo").then(result => {
        // console.log(result.body);
        if (result.body.status === 0) {
          // 成功了
          this.lunbotuList = result.body.message;
        } else {
          // 失败的
          Toast("加载轮播图失败。。。");
        }
      });
    }
  }
};
 ```

 <font color=red>注意：由于已经配置了vue-resource的根域名，所以这里只需要写后面的路径就可以了 api/getlunbo</font>


# 首页六宫格

## 引入 mui的grid-default.html

```
   <ul class="mui-table-view mui-grid-view mui-grid-9">
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1557419977279&di=b93eecb0565a5297414701953a5846ed&imgtype=0&src=http%3A%2F%2Fpic1.16pic.com%2F00%2F25%2F00%2F16pic_2500836_b.jpg" alt="">
              <div class="mui-media-body">新闻资讯</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1557419977279&di=b93eecb0565a5297414701953a5846ed&imgtype=0&src=http%3A%2F%2Fpic1.16pic.com%2F00%2F25%2F00%2F16pic_2500836_b.jpg" alt="">
              <div class="mui-media-body">图片分享</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1557419977279&di=b93eecb0565a5297414701953a5846ed&imgtype=0&src=http%3A%2F%2Fpic1.16pic.com%2F00%2F25%2F00%2F16pic_2500836_b.jpg" alt="">
              <div class="mui-media-body">商品购买</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1557419977279&di=b93eecb0565a5297414701953a5846ed&imgtype=0&src=http%3A%2F%2Fpic1.16pic.com%2F00%2F25%2F00%2F16pic_2500836_b.jpg" alt="">
              <div class="mui-media-body">留言反馈</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1557419977279&di=b93eecb0565a5297414701953a5846ed&imgtype=0&src=http%3A%2F%2Fpic1.16pic.com%2F00%2F25%2F00%2F16pic_2500836_b.jpg" alt="">
              <div class="mui-media-body">视频专区</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1557419977279&di=b93eecb0565a5297414701953a5846ed&imgtype=0&src=http%3A%2F%2Fpic1.16pic.com%2F00%2F25%2F00%2F16pic_2500836_b.jpg" alt="">
              <div class="mui-media-body">联系我们</div></a></li>
  </ul>  
```

## 修改六宫格的样式

```
.mui-grid-view.mui-grid-9 {
  background-color: #fff;
  border: none;
  img {
    width: 60px;
    height: 60px;
  }

  .mui-media-body{
    font-size: 13px;
  }
}

.mui-grid-view.mui-grid-9 .mui-table-view-cell {
  border: 0;
}
```

第一： 背景色
第二：边框
第三： 图片，默认情况下美工的图片都会做大一倍，所以设置图片的长宽减少一倍
第四： 字体大小


# 最后六宫格的的html
引入本地image
```
  <!-- 九宫格 到 6宫格 的改造工程 -->
     <ul class="mui-table-view mui-grid-view mui-grid-9">
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="../../images/menu1.png" alt="">
              <div class="mui-media-body">新闻资讯</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="../../images/menu2.png" alt="">
              <div class="mui-media-body">图片分享</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="../../images/menu3.png" alt="">
              <div class="mui-media-body">商品购买</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="../../images/menu4.png" alt="">
              <div class="mui-media-body">留言反馈</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="../../images/menu5.png" alt="">
              <div class="mui-media-body">视频专区</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="../../images/menu6.png" alt="">
              <div class="mui-media-body">联系我们</div></a></li>
  </ul>  
```


# 改造 新闻资讯 路由链接

添加新闻组件,使用mui的media-list.html

下面的图片链接必须存在，不然会报错显示不出来，可以先设置为某个互联网图片的链接

```
<template>
    <div>
        <ul class="mui-table-view">
				<li class="mui-table-view-cell mui-media">
					<a href="javascript:;">
						<img class="mui-media-object mui-pull-left" src="../images/shuijiao.jpg">
						<div class="mui-media-body">
							幸福
							<p class='mui-ellipsis'>能和心爱的人一起睡觉，是件幸福的事情；可是，打呼噜怎么办？</p>
						</div>
					</a>
				</li>
				<li class="mui-table-view-cell mui-media">
					<a href="javascript:;">
						<img class="mui-media-object mui-pull-left" src="../images/muwu.jpg">
						<div class="mui-media-body">
							木屋
							<p class='mui-ellipsis'>想要这样一间小木屋，夏天挫冰吃瓜，冬天围炉取暖.</p>
						</div>
					</a>
				</li>
				<li class="mui-table-view-cell mui-media">
					<a href="javascript:;">
						<img class="mui-media-object mui-pull-left" src="../images/cbd.jpg">
						<div class="mui-media-body">
							CBD
							<p class='mui-ellipsis'>烤炉模式的城，到黄昏，如同打翻的调色盘一般.</p>
						</div>
					</a>
				</li>

			</ul>
    </div>
</template>

<script>
export default {
    
}
</script>

<style lang="sass" scoped>

</style>



```

## router.js添加路由

```
import NewsList from './components/news/NewsList.vue'

// 3. 创建路由对象
var router = new VueRouter({
  routes: [ // 配置路由规则
    { path: '/', redirect: '/home' },
    { path: '/home', component: HomeContainer },
    { path: '/member', component: MemberContainer },
    { path: '/shopcar', component: ShopcarContainer },
    { path: '/search', component: SearchContainer },
    { path: '/home/newslist', component: NewsList },
    { path: '/home/newsinfo/:id', component: NewsInfo }
  ],
  linkActiveClass: 'mui-active' // 覆盖默认的路由高亮的类，默认的类叫做 router-link-active
})
```
新闻组件的路由就是下面这个
```
   { path: '/home/newslist', component: NewsList },
```

在hHomeContainer.vue添加新闻链接

```
   <router-link to="/home/newslist">
              <img src="../../images/menu1.png" alt="">
              <div class="mui-media-body">新闻资讯</div></router-link></li>
```


## 修改新闻资讯的样式




```
<ul class="mui-table-view">
				<li class="mui-table-view-cell mui-media">
					<a href="javascript:;">
						<img class="mui-media-object mui-pull-left" src="https://pg-ad-b1.ws.126.net/yixiao/17665/bd264df5640015c7599291132504d2c1.jpg">
						<div class="mui-media-body">
							<h1>幸福</h1>
							<p class='mui-ellipsis'>
								<span>发表时间:2019-01-01 12:12:32</span>
								<span>点击:0次</span>
							</p>
						</div>
					</a>
				</li>
				<li class="mui-table-view-cell mui-media">
					<a href="javascript:;">
						<img class="mui-media-object mui-pull-left" src="https://pg-ad-b1.ws.126.net/yixiao/17665/bd264df5640015c7599291132504d2c1.jpg">
						<div class="mui-media-body">
							<h1>幸福</h1>
							<p class='mui-ellipsis'>
								<span>发表时间:2019-01-01 12:12:32</span>
								<span>点击:0次</span>
							</p>
						</div>
					</a>
				</li>
				<li class="mui-table-view-cell mui-media">
					<a href="javascript:;">
						<img class="mui-media-object mui-pull-left" src="https://pg-ad-b1.ws.126.net/yixiao/17665/bd264df5640015c7599291132504d2c1.jpg">
						<div class="mui-media-body">
							<h1>幸福</h1>
							<p class='mui-ellipsis'>
								<span>发表时间:2019-01-01 12:12:32</span>
								<span>点击:0次</span>
							</p>
						</div>
					</a>
				</li>

			</ul>
```




NewList.vue

```
<style lang="scss" scoped>
.mui-table-view {
  li {
    h1 {
      font-size: 14px;
    }
    .mui-ellipsis {
      font-size: 12px;
      color: #226aff;
      display: flex;
      justify-content: space-between;
    }
  }
}
</style>
```

样式说明:
1. 标题字体大小
2. 底部字体大小 12
3. 颜色
4. 弹性布局：
```
display: flex;
````
5. 元素排列

```
justify-content: space-between;  /* 均匀排列每个元素
```

## 获取数据渲染新闻列表

NewsList.vue获取数据

```


<script>
import { Toast } from "mint-ui";

export default {
  data() {
    return {
      newslist: [] // 新闻列表
    };
  },
  created() {
    this.getNewsList();
  },
  methods: {
    getNewsList() {
      // 获取新闻列表
      this.$http.get("api/getnewslist").then(result => {
        if (result.body.status === 0) {
          // 如果没有失败，应该把数据保存到 data 上
          this.newslist = result.body.message;
        } else {
          Toast("获取新闻列表失败！");
        }
      });
    }
  }
};
</script>


```


修改html,遍历新闻数据

```
 <ul class="mui-table-view">
      <li class="mui-table-view-cell mui-media" v-for="item in newslist" :key="item.id">
        <router-link :to="'/home/newsinfo/' + item.id">
          <img class="mui-media-object mui-pull-left" :src="item.img_url">
          <div class="mui-media-body">
            <h1>{{ item.title }}</h1>
            <p class='mui-ellipsis'>
              <span>发表时间：{{ item.add_time  }}</span>
              <span>点击：{{item.click}}次</span>
            </p>
          </div>
        </router-link>
      </li>

    </ul>

```

效果:

![](./img/03.png)

7. 关于底部贴太近的问题，在app.vue组件中添加

```
.app-container {
  padding-top: 40px;
	padding-bottom: 50px;
  overflow-x: hidden;
}
```
新增的样式是:
<font color=red>padding-bottom: 50px;</font>


8. 通过全局日志过滤器格式化时间
在main.js添加过滤器

```
// 导入格式化时间的插件
import moment from 'moment'
// 定义全局的过滤器
Vue.filter('dateFormat', function (dataStr, pattern = "YYYY-MM-DD HH:mm:ss") {
  return moment(dataStr).format(pattern)
})
```

在新闻列表组件NewsList.vue添加格式化

```
  <span>发表时间：{{ item.add_time |dateFormat  }}</span>
```

最终效果

![](./img/05.png)


## 实现 新闻资讯列表 点击跳转到新闻详情

1. 创建新闻详情组件NewInfo.vue

新闻详情组件的初始样子

```
<template>
  <div class="newsinfo-container">
      <h1>NewInfo.vue</h1>
  </div>
</template>

<script>
// 1. 导入 评论子组件
export default {
 
};
</script>

<style lang="scss">

</style>

```

2. 在router.js中引入新闻详情组件NewInfo.vue,创建对应的路由

```
    { path: '/home/newsinfo/:id', component: NewsInfo }

```

<font color=red>
注意：这里的:id是一个路径参数,id是参数名称
</font>




```

```

3. 把列表中的每一项改造为 router-link,同时，在跳转的时候应该提供唯一的Id标识符

```

       <ul class="mui-table-view">
      <li class="mui-table-view-cell mui-media" v-for="item in newslist" :key="item.id">
        <router-link :to="'/home/newsinfo/' + item.id">
          <img class="mui-media-object mui-pull-left" :src="item.img_url">
          <div class="mui-media-body">
            <h1>{{ item.title }}</h1>
            <p class='mui-ellipsis'>
              <span>发表时间：{{ item.add_time  }}</span>
              <span>点击：{{item.click}}次</span>
            </p>
          </div>
        </router-link>
      </li>

    </ul>
    </div>

```

注意链接:

<font color=red>
  <router-link :to="'/home/newsinfo/' + item.id">
</font>


4. 查看效果

![](./img/04.png)


## 新闻详情页样式绘制


```
<div class="newsinfo-container">
    <!-- 大标题 -->
    <h3 class="title"这是标题</h3>
    <!-- 子标题 -->
    <p class="subtitle">
      <span>发表时间：2019-01-02 01:12:12</span>
      <span>点击：100次</span>
    </p>

    <hr>

    <!-- 内容区域 -->
    <div class="content" v-html="newsinfo.content">
      这是内容区域
    </div>

  </div>

```
现在的效果是:

![](./img/06.png)

需要修改样式满足需求

添加对应的样式

```
<style lang="scss">
.newsinfo-container {
  padding: 0 4px;
  .title {
    font-size: 16px;
    text-align: center;
    margin: 15px 0;
    color: red;
  }
  .subtitle {
    font-size: 13px;
    color: #226aff;
    display: flex;
    justify-content: space-between;
  }
  .content {
    img {
      width: 100%;
    }
  }
}
</style>
```
样式说明:
* title
  * 字体大小
  * 文本居中
  * 上下相隔
  * 红色字体

* subtitle
  * 字体大小
  * 颜色
  * 两端分布 ``` display: flex;  ````
  * justify-content: space-between;  均匀排列每个元素
* content 内容图片  100%显示




<font color=red>
注意 这里没有scoped  其实没有scoped也没有问题，scope本来就是为了避免和其他样式冲突的，在我们项目里面这个样式一般不太可能冲突

</font>

## 查询新闻详情页的数据并且渲染


1. 查询新闻详情

```
// 1. 导入 评论子组件
export default {
 data() {
    return {
      id: this.$route.params.id, // 将 URL 地址中传递过来的 Id值，挂载到 data上，方便以后调用
      newsinfo: {} // 新闻对象
    };
  },
  created() {
    this.getNewsInfo();
  },
   methods: {
    getNewsInfo() {
      // 获取新闻详情
      this.$http.get("api/getnew/" + this.id).then(result => {
        if (result.body.status === 0) {
          this.newsinfo = result.body.message[0];
        } else {
          Toast("获取新闻失败！");
        }
      });
    }
  },
};
```

this.$route.params.id:是从新闻列表的链接传递过来的id参数

2. 用数据渲染html页面

```
  <div class="newsinfo-container">
    <!-- 大标题 -->
    <h3 class="title">{{ newsinfo.title }}</h3>
    <!-- 子标题 -->
    <p class="subtitle">
      <span>发表时间：{{ newsinfo.add_time | dateFormat }}</span>
      <span>点击：{{ newsinfo.click }}次</span>
    </p>
    <hr>

    <!-- 内容区域 -->
    <div class="content" v-html="newsinfo.content"></div>

  </div>
```

3. 效果

![](./img/07.png)


## 新闻详情页封装一个评论组件

*  先创建一个 单独的 comment.vue 组件模板

html 内容:

```
  <div class="cmt-container">
    <h3>发表评论</h3>
    <hr>
    <textarea placeholder="请输入要BB的内容（做多吐槽120字）" maxlength="120"></textarea>

    <mt-button type="primary" size="large">发表评论</mt-button>

    <div class="cmt-list">
      <div class="cmt-item" v-for="(item, i) in comments" :key="item.add_time">
        <div class="cmt-title">
          第{{ i+1 }}楼&nbsp;&nbsp;用户：{{ item.user_name }}&nbsp;&nbsp;发表时间：{{ item.add_time | dateFormat }}
        </div>
        <div class="cmt-body">
          {{ item.content === 'undefined' ? '此用户很懒，嘛都没说': item.content }}
        </div>
      </div>

    </div>

    <mt-button type="danger" size="large" plain @click="getMore">加载更多</mt-button>
  </div>

```

样式:

```
<style lang="scss" scoped>
.cmt-container {
  h3 {
    font-size: 18px;
  }
  textarea {
    font-size: 14px;
    height: 85px;
    margin: 0;
  }

  .cmt-list {
    margin: 5px 0;
    .cmt-item {
      font-size: 13px;
      .cmt-title {
        line-height: 30px;
        background-color: #ccc;
      }
      .cmt-body {
        line-height: 35px;
        text-indent: 2em;
      }
    }
  }
}
</style>
```
样式说明:
* h3 字体大小
* textarea
  * 字体大小
  * 高度
  * 元素间隔
* 评论列表 cmt-list
  * 元素间隔
  * 字体大小
  * 文本高度
  * 背景颜色
  * 文本间隔


获取评论内容:

```

<script>
import { Toast } from "mint-ui";
export default {
  data() {
    return {
      pageIndex: 1, // 默认展示第一页数据
      comments: [] // 所有的评论数据
    };
  },
  created() {
    this.getComments();
  },
  methods: {
    getComments() {
      // 获取评论
      this.$http
        .get("api/getcomments/" + this.id + "?pageindex=" + this.pageIndex)
        .then(result => {
          if (result.body.status === 0) {
            // this.comments = result.body.message;
            // 每当获取新评论数据的时候，不要把老数据清空覆盖，而是应该以老数据，拼接上新数据
            this.comments = this.comments.concat(result.body.message);
          } else {
            Toast("获取评论失败！");
          }
        });
    },
    getMore() {
      // 加载更多
      this.pageIndex++;
      this.getComments();
    }
  },
  props: ["id"]
};
</script>
```

<font color=red>
注意：
这里props的属性是和父组件绑定的，父组件传递过来的id
</font>

记得在main.js引入Button组件

```
// 按需导入 Mint-UI 中的组件   
import { Header, Swipe, SwipeItem, Button } from 'mint-ui'
Vue.component(Header.name, Header)
Vue.component(Swipe.name, Swipe)
Vue.component(SwipeItem.name, SwipeItem)
Vue.component(Button.name, Button)
```



在新闻详情NewsInfo.vue组件中使用comment组件，新闻详情组件作为comment的父组件，需要绑定id属性，这样评论组件才能通过ID查询对应的评论


引入comment组件

```
// 1. 导入 评论子组件
import comment from "../subcomponents/comment.vue";

```

创建组件:

```
  methods: {
    getNewsInfo() {
      // 获取新闻详情
      this.$http.get("api/getnew/" + this.id).then(result => {
        if (result.body.status === 0) {
          this.newsinfo = result.body.message[0];
        } else {
          Toast("获取新闻失败！");
        }
      });
    }
  },
  components: {
    // 用来注册子组件的节点
    "comment-box": comment
  }
};
```

<font color=red>  "comment-box": comment</font>



在NewsInfo.vue组件页面上添加评论组件代码

```
 <!-- 评论子组件区域 -->
    <comment-box :id="this.id"></comment-box>
```



<font color=red>
注意这里的绑定属性id是传递个comment组件的参数
相当于子组件绑定父组件的属性,同样在comment子组件中要设置子组件绑定的父组件属性是哪个
代码如下，在comment.vue中
```
 props: ["id"]
```
</font>

最后的效果如图:

![](./img/08.png)

## 实现点击加载更多评论的功能
1. 为加载更多按钮，绑定点击事件，在事件中，请求 下一页数据
2. 点击加载更多，让 pageIndex++ , 然后重新调用 this.getComments() 方法重新获取最新一页的数据
3. 为了防止 新数据 覆盖老数据的情况，我们在 点击加载更多的时候，每当获取到新数据，应该让 老数据 调用 数组的 concat 方法，拼接上新数组


具体的代码已经在comment.vue中了


## 发表评论

1. 把文本框做双向数据绑定
comment.vue的评论输入框添加v-model
```
 <textarea placeholder="请输入要BB的内容（做多吐槽120字）" maxlength="120"></textarea>
```

2. 为发表按钮绑定一个事件

comment.vue
```

    <mt-button type="primary" size="large"  @click="postComment">发表评论</mt-button>
```


<font color=red>
@click="postComment"
</font>

3. 校验评论内容是否为空，如果为空，则Toast提示用户 评论内容不能为空
4. 通过 vue-resource 发送一个请求，把评论内容提交给 服务器
5. 当发表评论OK后，重新刷新列表，以查看最新的评论
 + 如果调用 getComments 方法重新刷新评论列表的话，可能只能得到 最后一页的评论，前几页的评论获取不到
 + 换一种思路： 当评论成功后，在客户端，手动拼接出一个 最新的评论对象，然后 调用 数组的 unshift 方法， 把最新的评论，追加到  data 中 comments 的开头；这样，就能 完美实现刷新评论列表的需求；

comment.vue

methods中添加方法
```
    postComment() {
      // 校验是否为空内容
      if (this.msg.trim().length === 0) {
        return Toast("评论内容不能为空！");
      }

      // 发表评论
      // 参数1： 请求的URL地址
      // 参数2： 提交给服务器的数据对象 { content: this.msg }
      // 参数3： 定义提交时候，表单中数据的格式  { emulateJSON:true }
      this.$http
        .post("api/postcomment/" + this.$route.params.id, {
          content: this.msg.trim()
        })
        .then(function(result) {
          if (result.body.status === 0) {
            // 1. 拼接出一个评论对象
            var cmt = {
              user_name: "匿名用户",
              add_time: Date.now(),
              content: this.msg.trim()
            };
            this.comments.unshift(cmt);
            this.msg = "";
          }
        });
    }
```

效果:

![](./img/09.png)
![](./img/10.png)


# 改造图片分析 按钮为 路由的链接并显示对应的组件页面

1. 添加图片列表组件PhotoList.vue

```
<template>
    <div>
        <h1>PhotoList</h1>
    </div>
</template>

<script>
export default {
    
}
</script>

<style lang="scss" scoped>

</style>

```

2. router.js添加图片列表组件的路由

```
import PhotoList from './components/photos/PhotoList.vue'

```

路由

```
  { path: '/home/photolist', component: PhotoList },
```

3. 修改首页的图片链接为route-link

HomeContainer.vue

```
   <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><router-link to="/home/photolist">
              <img src="../../images/menu2.png" alt="">
              <div class="mui-media-body">图片分享</div></router-link></li>
```

4. 效果如图:

![](./img/11.png)


# 绘制 图片列表 组件页面结构并美化样式

## 制作 顶部的滑动条

使用mui的tab-top-webview-main.html

```
			<div id="slider" class="mui-slider mui-fullscreen">
				<div id="sliderSegmentedControl" class="mui-scroll-wrapper mui-slider-indicator mui-segmented-control mui-segmented-control-inverted">
					<div class="mui-scroll">
						<a class="mui-control-item mui-active" href="#item1mobile" data-wid="tab-top-subpage-1.html">
							推荐
						</a>
						<a class="mui-control-item" href="#item2mobile" data-wid="tab-top-subpage-2.html">
							热点
						</a>
						<a class="mui-control-item" href="#item3mobile" data-wid="tab-top-subpage-3.html">
							北京
						</a>
						<a class="mui-control-item" href="#item4mobile" data-wid="tab-top-subpage-4.html">
							社会
						</a>
						<a class="mui-control-item" href="#item5mobile" data-wid="tab-top-subpage-5.html">
							娱乐
						</a>
						<!--<a class="mui-control-item" href="#item6mobile" data-wid="tab-top-subpage-6.html">
							科技
						</a>-->
					</div>
				</div>

			</div>
```

这时候还不能滑动，需要引入初始化scroll的代码

PhotoList.vue引入mui的 mui.js 文件:

```
// 1. 导入 mui 的js文件
import mui from "../../lib/mui/js/mui.min.js";

```


在PhotoList.vue中初始化滑动组件:

```
  mui('.mui-scroll-wrapper').scroll({
    deceleration: 0.0005 //flick 减速系数，系数越大，滚动速度越慢，滚动距离越小，默认值0.0006
  });
```


这时候出现了严格模式不能使用mui的错误:

这是mui.js文件的原因,错误如下

```
Uncaught TypeError: 'caller', 'callee', and 'arguments' properties may not be accessed on strict mode functions or the arguments objects for calls to them

```

解决办法:

解决方案：
 1. 把 mui.js 中的 非严格 模式的代码改掉；但是不现实； 
 2. 把 webpack 打包时候的严格模式禁用掉；
  + 最终，我们选择了 plan B  移除严格模式： 使用这个插件 

```
npm install babel-plugin-transform-remove-strict-mode
```

.babelrc 文件添加配置

```
{
    "plugins": ["transform-remove-strict-mode"]
}

```



警告处理

```
[Intervention] Unable to preventDefault inside passive event listener due to target being treated as passive. See <URL>
```

解决办法:

```
* {
  touch-action: pan-y;
}

```

刚进页面不能滑动的解决办法:
在mounted的时候初始化滑动组件
```
 mounted() {
    // 当 组件中的DOM结构被渲染好并放到页面中后，会执行这个 钩子函数
    // 如果要操作元素了，最好在 mounted 里面，因为，这里时候的 DOM 元素 是最新的
    // 2. 初始化滑动控件
    mui(".mui-scroll-wrapper").scroll({
      deceleration: 0.0005 //flick 减速系数，系数越大，滚动速度越慢，滚动距离越小，默认值0.0006
    });
  },
```

### 引入mui.js之后tabar不能切换的解决办法

当 滑动条 调试OK后，发现， tabbar 无法正常工作了，这时候，我们需要把 每个 tabbar 按钮的 样式中  `mui-tab-item` 重新改一下名字；

App.vue中的

```
 <!-- 底部 Tabbar 区域 -->
    <nav class="mui-bar mui-bar-tab">
			<router-link class="mui-tab-item" to="/home">
				<span class="mui-icon mui-icon-home"></span>
				<span class="mui-tab-label">首页</span>
			</router-link>
			<router-link class="mui-tab-item" to="/member">
				<span class="mui-icon mui-icon-contact"></span>
				<span class="mui-tab-label">会员</span>
			</router-link>
			<router-link class="mui-tab-item" to="/shopcar">
				<span class="mui-icon mui-icon-extra mui-icon-extra-cart">
					<span class="mui-badge">0</span>
				</span>
				<span class="mui-tab-label">购物车</span>
			</router-link>
			<router-link class="mui-tab-item" to="/search">
				<span class="mui-icon mui-icon-search"></span>
				<span class="mui-tab-label">搜索</span>
			</router-link>
		</nav>
```
这里的mui-tab-item样式冲突，我们需要修改它的名字


首先是新名字的样式:
app.vue

```
// 该类名，解决 tabbar 点击无法切换的问题
.mui-bar-tab .mui-tab-item-llb.mui-active {
    color: #007aff;
}

.mui-bar-tab .mui-tab-item-llb {
    display: table-cell;
    overflow: hidden;
    width: 1%;
    height: 50px;
    text-align: center;
    vertical-align: middle;
    white-space: nowrap;
    text-overflow: ellipsis;
    color: #929292;
}

.mui-bar-tab .mui-tab-item-llb .mui-icon {
    top: 3px;
    width: 24px;
    height: 24px;
    padding-top: 0;
    padding-bottom: 0;
}

.mui-bar-tab .mui-tab-item-llb .mui-icon~.mui-tab-label {
    font-size: 11px;
    display: block;
    overflow: hidden;
    text-overflow: ellipsis;
}
```

然后修改对应名字
```
 <!-- 底部 Tabbar 区域 -->
    <nav class="mui-bar mui-bar-tab">
			<router-link class="mui-tab-item-llb" to="/home">
				<span class="mui-icon mui-icon-home"></span>
				<span class="mui-tab-label">首页</span>
			</router-link>
			<router-link class="mui-tab-item-llb" to="/member">
				<span class="mui-icon mui-icon-contact"></span>
				<span class="mui-tab-label">会员</span>
			</router-link>
			<router-link class="mui-tab-item-llb" to="/shopcar">
				<span class="mui-icon mui-icon-extra mui-icon-extra-cart">
					<span class="mui-badge">0</span>
				</span>
				<span class="mui-tab-label">购物车</span>
			</router-link>
			<router-link class="mui-tab-item-llb" to="/search">
				<span class="mui-icon mui-icon-search"></span>
				<span class="mui-tab-label">搜索</span>
			</router-link>
		</nav>

```

这个时候的效果如图:

![](./img/12.png)

并且底部按钮也能切换了

 ### 获取所有分类，并渲染 分类列表；

 在HhotoList.vue中的script中


 添加data:

```
 data() {
    return {
      cates: [], // 所有分类的列表数组
      list: [] // 图片列表的数组
    };
  },
```

cates就是保存图片分类数据

 添加method:
```
 getAllCategory() {
      // 获取所有的图片分类
      this.$http.get("api/getimgcategory").then(result => {
        if (result.body.status === 0) {
          // 手动拼接出一个最完整的 分类列表
          result.body.message.unshift({ title: "全部", id: 0 });
          this.cates = result.body.message;
        }
      });
    },
```

created时候调用:

```
 created() {
    this.getAllCategory();
  },
```

PhotoList.vue 模板渲染返回的分类数据

```
	<div id="slider" class="mui-slider">
		<div id="sliderSegmentedControl" class="mui-scroll-wrapper mui-slider-indicator mui-segmented-control mui-segmented-control-inverted">
			<div class="mui-scroll">
				  <a :class="['mui-control-item', item.id == 0 ? 'mui-active' : '']" v-for="item in cates" :key="item.id" @tap="getPhotoListByCateId(item.id)">
           			 {{ item.title }}
         		 </a>
			</div>
		</div>
	</div>
```

<font color=red>
注意这里绑定的class属性:
:class="['mui-control-item', item.id == 0 ? 'mui-active' : '']"
默认情况刚进入该页面的时候 第一个全部的链接激活状态，其他分类链接点击的时候调用的是mui自己的激活功能来修改样式的
</font>




### 制作图片列表区域
1. 图片列表需要使用懒加载技术，我们可以使用 Mint-UI 提供的现成的 组件 `lazy-load`
2. 根据`lazy-load`的使用文档，尝试使用
3. 渲染图片列表数据


需要注意的是使用lazy-load需要全部引入mint-ui，像之前的按需引入的代码都需要注释掉

原来在main.js按需引入mint-ui
```
import { Header, Swipe, SwipeItem, Button } from 'mint-ui'
Vue.component(Header.name, Header)
Vue.component(Swipe.name, Swipe)
Vue.component(SwipeItem.name, SwipeItem)
Vue.component(Button.name, Button)
```

修改成全部引入,不然图片懒加载有问题:


```
import MintUI from 'mint-ui'
Vue.use(MintUI)
import 'mint-ui/lib/style.css'
```


PhotoList.vue中编写图片列表显示的html:

首先官网提供使用实例本来就很简单:


```
<ul>
  <li v-for="item in list">
    <img v-lazy="item">
  </li>
</ul>

image[lazy=loading] {
  width: 40px;
  height: 300px;
  margin: auto;
}
```


所以我们先调用API获取每个分类的图片列表数据



```
getPhotoListByCateId(cateId) {
      	// 根据 分类Id，获取图片列表
			this.$http.get("api/getimages/" + cateId).then(result => {
				if (result.body.status === 0) {
				this.list = result.body.message;
				}
			});
		}
```


```
created() {
    	this.getAllCategory();
    	// 默认进入页面，就主动请求 所有图片列表的数据
		this.getPhotoListByCateId(0);
		 // 默认进入页面，就主动请求 所有图片列表的数据
    	this.getPhotoListByCateId(0);
  	},
```

```
data() {
		return {
		cates: [], // 所有分类的列表数组
		list: [] // 图片列表的数组
		};
	  },
```


在分类中添加点击事件:


```
	  <a :class="['mui-control-item', item.id == 0 ? 'mui-active' : '']" v-for="item in cates" :key="item.id" @tap="getPhotoListByCateId(item.id)">
           			 {{ item.title }}
         		 </a>


```


@tap就是我们的事件



在html循坏遍历图片列表:

```
<!-- 图片列表区域 -->
    <ul class="photo-list">
      <router-link v-for="item in list" :key="item.id" :to="'/home/photoinfo/' + item.id" tag="li">
        <img v-lazy="item.img_url">
        <div class="info">
          <h1 class="info-title">{{ item.title }}</h1>
          <div class="info-body">{{ item.zhaiyao }}</div>
        </div>
      </router-link>
    </ul>
```

图片列表渲染出来之后需要调整样式，完整的样式如下:

```
<style lang="scss" scoped>
* {
  touch-action: pan-y;
}

.photo-list {
  list-style: none;
  margin: 0;
  padding: 10px;
  padding-bottom: 0;
  li {
    background-color: #ccc;
    text-align: center;
    margin-bottom: 10px;
    box-shadow: 0 0 9px #999;
    position: relative;
    img {
      width: 100%;
      vertical-align: middle;
    }
    img[lazy="loading"] {
      width: 40px;
      height: 300px;
      margin: auto;
    }

    .info {
      color: white;
      text-align: left;
      position: absolute;
      bottom: 0;
      background-color: rgba(0, 0, 0, 0.4);
      max-height: 84px;
      .info-title {
        font-size: 14px;
      }
      .info-body {
        font-size: 13px;
      }
    }
  }
}
</style>
```

样式说明:




最后还有一个问题：在把内容网上滑动的时候发现图片分类哪一栏的分类链接会浮动到头部前面

![](./img/14.png)

这个时候需要修改app.vue的zindex

app.vue

```
.mint-header{
  z-index: 99;
}
```







最终效果如图:

![](./img/13.png)



## 图片详情页

1. 在改造 li 成 router-link 的时候，需要使用 tag 属性指定要渲染为 哪种元素
创建PhotoInfo.vue组件


```
<template>
    <div><h1>PhotoInfo.vue</h1></div>
</template>
<script>
export default {
    
}
</script>

<style lang="scss" scoped>

</style>



```

router.js引入PhotoList.vue组件

```
import PhotoInfo from './components/photos/PhotoInfo.vue'
```

```
    { path: '/home/photoinfo/:id', component: PhotoInfo },

```

PhotoList.vue添加对应的路由跳转

```
<!-- 图片列表区域 -->
    <ul class="photo-list">
      <router-link v-for="item in list" :key="item.id" :to="'/home/photoinfo/' + item.id" tag="li">
        <img v-lazy="item.img_url">
        <div class="info">
          <h1 class="info-title">{{ item.title }}</h1>
          <div class="info-body">{{ item.zhaiyao }}</div>
        </div>
      </router-link>
    </ul>

```

<font color=red>:to="'/home/photoinfo/' + item.id"</font>就是我们需要的链接
这个时候的效果:
![](./img/15.png)


## 实现图片详情页面的布局和美化，同时获取数据渲染页面


1. 图片详情页html结构

```
<div class="photoinfo-container">
    <h3>图片标题</h3>
    <p class="subtitle">
      <span>发表时间:9-01-01 12:14:01</span>
      <span>点击：100次</span>
    </p>
    <hr>
     <!-- 图片内容区域 -->
    <div class="content">
      这是图片区域的内容
    </div>
  </div>   
   
```

2. 获取图片详情页的数据

```
 data() {
    return {
      id: this.$route.params.id, // 从路由中获取到的 图片Id
      photoinfo: {}, // 图片详情
      list: [] // 缩略图的数组
    };
  },
```

```
 getPhotoInfo() {
      // 获取图片的详情
      this.$http.get("api/getimageInfo/" + this.id).then(result => {
        if (result.body.status === 0) {
          this.photoinfo = result.body.message[0];
        }
      });
    },
```

```
created() {
    this.getPhotoInfo();
    this.getThumbs();
  },
```

html渲染获取回来的数据

```
<template>
  <div class="photoinfo-container">
    <h3>{{ photoinfo.title }}</h3>
    <p class="subtitle">
      <span>发表时间：{{ photoinfo.add_time | dateFormat }}</span>
      <span>点击：{{ photoinfo.click }}次</span>
    </p>

    <hr>

    <!-- 缩略图区域 -->
    <div class="thumbs">
      <img class="preview-img" v-for="(item, index) in list" :src="item.src" height="100" @click="$preview.open(index, list)" :key="item.src">
    </div>

    <!-- 图片内容区域 -->
    <div class="content" v-html="photoinfo.content"></div>

    <!-- 放置一个现成的 评论子组件 -->
    <cmt-box :id="id"></cmt-box>
  </div>
</template>
```

样式:

```
<style lang="scss" scoped>
.photoinfo-container {
  padding: 3px;
  h3 {
    color: #26a2ff;
    font-size: 15px;
    text-align: center;
    margin: 15px 0;
  }
  .subtitle {
    display: flex;
    justify-content: space-between;
    font-size: 13px;
  }

  .content {
    font-size: 13px;
    line-height: 30px;
  }

  .thumbs{
    img{
      margin: 10px;
      box-shadow: 0 0 8px #999;
    }
  }
}
</style>
```

样式就不先不细说了

### 图片详情页添加评论组件

首先引入评论组件
PhotoInfo.vue

```
// 1. 导入评论子组件
import comment from "../subcomponents/comment.vue";
```

创建评论组件

PhotoInfo.vue
```
 components: {
    // 注册 评论子组件
    "cmt-box": comment
  }
```

在PhotoInfo.vue的页面中引入评论组件

```
<!-- 放置一个现成的 评论子组件 -->
    <cmt-box :id="id"></cmt-box>
```

最终效果：


![](./img/16.png)


### 实现 图片详情中 缩略图的功能
1. 使用 插件 vue-preview 这个缩略图插件
2. 获取到所有的图片列表，然后使用 v-for 指令渲染数据
3. 注意： img标签上的class不能去掉
4. 注意： 每个 图片数据对象中，必须有 w 和 h 属性


main.js引入vue-preview组件

```
// 安装 图片预览插件
import VuePreview from 'vue-preview'
Vue.use(VuePreview)
```

PhotoInfo.vue中添加vue-preview图片预览组件
```
 <!-- 缩略图区域 -->
    <div class="thumbs">
      <!-- <img class="preview-img" v-for="(item, index) in list" :src="item.src" height="100" @click="$preview.open(index, list)" :key="item.src"> -->
      <vue-preview :slides="list" ></vue-preview>

    </div>
```

获取缩略图，因为预览组件需要w,h属性，所以需要遍历图片然后加上w h 的属性

```
 getThumbs() {
      // 获取缩略图
      this.$http.get("api/getthumimages/" + this.id).then(result => {
        if (result.body.status === 0) {
          // 循环每个图片数据，补全图片的宽和高
          result.body.message.forEach(item => {
            item.w = 600;
            item.h = 400;
            item.msrc = item.src;
          });
          // 把完整的数据保存到 list 中
          this.list = result.body.message;
        }
      });
    }
```
缩略图数组:
```
 data() {
    return {
      id: this.$route.params.id, // 从路由中获取到的 图片Id
      photoinfo: {}, // 图片详情
      list: [] // 缩略图的数组
    };
  },
```

最终效果如图:

![](./img/17.png)


# 商品列表页面


* 商品列表跳转制作
  * 创建商品列表组件GoodsList.vue

```
  <template>
    <div>
        <h1>商品列表页面</h1>
    </div>
</template>
<script>
export default {
    
}
</script>

<style lang="sass" scoped>

</style>

```

  * router.js配置路由导航

  ```
  import GoodsList from './components/goods/GoodsList.vue'
  ```

  ```
      { path: '/home/goodslist', component: GoodsList },
  ```

  * 首页HomeContainer.vue配置跳转到商品列表页面

  ```
     <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3">
        <router-link to="/home/goodslist">
          <img src="../../images/menu3.png" alt>
          <div class="mui-media-body">商品购买</div>
        </router-link>
      </li>
  ```
  * 商品列表页面初步跳转显示完成


![](./img/18.png)


* 绘制商品列表页面
  * html绘制

```
   <!-- 在网页中，有两种跳转方式： -->
    <!-- 方式1： 使用 a 标签 的形式叫做 标签跳转  -->
    <!-- 方式2： 使用 window.location.href 的形式，叫做 编程式导航 -->
    <div class="goods-item" >
      <img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1559236033398&di=32301904d20cd9fae50e4044fc036290&imgtype=0&src=http%3A%2F%2Fimg003.hc360.cn%2Fm1%2FM01%2FE9%2FD4%2FwKhQb1RJmFuEcjmkAAAAAExCHG4129.jpg" alt="">
      <h1 class="title">商品标题</h1>
      <div class="info">
        <p class="price">
          <span class="now">￥ 3999</span>
          <span class="old">￥ 2999</span>
        </p>
        <p class="sell">
          <span>热卖中</span>
          <span>剩 9999 件</span>
        </p>
      </div>
    </div>

    <div class="goods-item" >
      <img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1559236033398&di=32301904d20cd9fae50e4044fc036290&imgtype=0&src=http%3A%2F%2Fimg003.hc360.cn%2Fm1%2FM01%2FE9%2FD4%2FwKhQb1RJmFuEcjmkAAAAAExCHG4129.jpg" alt="">
      <h1 class="title">商品标题</h1>
      <div class="info">
        <p class="price">
          <span class="now">￥ 3999</span>
          <span class="old">￥ 2999</span>
        </p>
        <p class="sell">
          <span>热卖中</span>
          <span>剩 9999 件</span>
        </p>
      </div>
    </div>


    <div class="goods-item" >
      <img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1559236033398&di=32301904d20cd9fae50e4044fc036290&imgtype=0&src=http%3A%2F%2Fimg003.hc360.cn%2Fm1%2FM01%2FE9%2FD4%2FwKhQb1RJmFuEcjmkAAAAAExCHG4129.jpg" alt="">
      <h1 class="title">商品标题</h1>
      <div class="info">
        <p class="price">
          <span class="now">￥ 3999</span>
          <span class="old">￥ 2999</span>
        </p>
        <p class="sell">
          <span>热卖中</span>
          <span>剩 9999 件</span>
        </p>
      </div>
    </div>
    <mt-button type="danger" size="large">加载更多</mt-button>
  </div>


```
  * 效果图


  ![](./img/19.png)


  * css美化

1. 列表的整体布局 

```
.goods-list {
  display: flex;
  flex-wrap: wrap;
  padding: 7px;
  justify-content: space-between;

}
```

flex:弹性布局

flex-wrap: 子元素的换行模式  换行，第二行在第一行下面，从左到右

padding:内容距离外部间距 

justify-content:子元素的对齐方式





商品的css样式:



```

<style lang="scss" scoped>
.goods-list {
  display: flex;
  flex-wrap: wrap;
  padding: 7px;
  justify-content: space-between;

  .goods-item {
    width: 49%;
    border: 1px solid #ccc;
    box-shadow: 0 0 8px #ccc;
    margin: 4px 0;
    padding: 2px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    min-height: 293px;
    img {
      width: 100%;
    }
    .title {
      font-size: 14px;
    }

    .info {
      background-color: #eee;
      p {
        margin: 0;
        padding: 5px;
      }
      .price {
        .now {
          color: red;
          font-weight: bold;
          font-size: 16px;
        }
        .old {
          text-decoration: line-through;
          font-size: 12px;
          margin-left: 10px;
        }
      }
      .sell {
        display: flex;
        justify-content: space-between;
        font-size: 13px;
      }
    }
  }
}
```



需要关注的：



商品占用宽度的一半 这样一行可以显示两列

```
    width: 49%;

```



布局的方向是列方向弹性布局 这样上下对齐

```
flex-direction: column;
```

保证每个商品占用的高度一样

```
 min-height: 293px;
```

图片百分百填充

```
 width: 100%;
```





## 获取商品列表数据并渲染



```
 data() {
    // data 是往自己组件内部，挂载一些私有数据的
    return {
      pageindex: 1, // 分页的页数
      goodslist: [] // 存放商品列表的数组
    };
  },
  created() {
    this.getGoodsList();
  },
  methods: {
    getGoodsList() {
      // 获取商品列表
      this.$http
        .get("api/getgoods?pageindex=" + this.pageindex)
        .then(result => {
          if (result.body.status === 0) {
            // this.goodslist = result.body.message;
            this.goodslist = this.goodslist.concat(result.body.message);
          }
        });
    },
    getMore() {
      this.pageindex++;
      this.getGoodsList();
    },
    goDetail(id) {
      // 使用JS的形式进行路由导航

      // 注意： 一定要区分 this.$route 和 this.$router 这两个对象，
      // 其中： this.$route 是路由【参数对象】，所有路由中的参数， params, query 都属于它
      // 其中： this.$router 是一个路由【导航对象】，用它 可以方便的 使用 JS 代码，实现路由的 前进、后退、 跳转到新的 URL 地址

      console.log(this);
      // 1. 最简单的
      // this.$router.push("/home/goodsinfo/" + id);
      // 2. 传递对象
      // this.$router.push({ path: "/home/goodsinfo/" + id });
      // 3. 传递命名的路由
      this.$router.push({ name: "goodsinfo", params: { id } });
    }
  }
```



说明:

pageindex:保存的是页码

goodslist:保存的是商品列表

getGoodsList:方法是请求商品列表,在vue创建时期初始化商品列表，点击更多的时候请求商品列表，并追加到原来的商品列表

```
  getGoodsList() {
      // 获取商品列表
      this.$http
        .get("api/getgoods?pageindex=" + this.pageindex)
        .then(result => {
          if (result.body.status === 0) {
            // this.goodslist = result.body.message;
            this.goodslist = this.goodslist.concat(result.body.message);
          }
        });
    },
```



获取商品列表之后需要渲染:



```
<div class="goods-list">
    <div class="goods-item" v-for="item in goodslist" :key="item.id" @click="goDetail(item.id)">
      <img :src="item.img_url" alt="">
      <h1 class="title">{{ item.title}}</h1>
      <div class="info">
        <p class="price">
          <span class="now">{{ item.sell_price}}</span>
          <span class="old">{{ item.market_price}}</span>
        </p>
        <p class="sell">
          <span>热卖中</span>
          <span>剩 9999 件</span>
        </p>
      </div>
    </div>
    <mt-button type="danger" size="large" @click="getMore">加载更多</mt-button>
  </div>
```





点击获取更多

```
 getMore() {
      this.pageindex++;
      this.getGoodsList();
    },
```





按钮也需要添加点击事件

```
<mt-button type="danger" size="large" @click="getMore">加载更多</mt-button>
```





# 商品详情页面



## 创建商品详情页面的组件GoodsInfo.vue

```
<template>
    <div>
        <h1>商品详情页面</h1>
    </div>
</template>
<script>
export default {
    
}
</script>
<style lang="scss" scoped>

</style>


```

router.js引入

```
import GoodsInfo from './components/goods/GoodsInfo.vue'
```



```
  { path: '/home/goodsinfo/:id', component: GoodsInfo, name: 'goodsinfo' },
```

id是参数



## 商品列表跳转商品详情页面



在GoodsList.vue中定义的getDetails方法就是跳转商品详情页面的

```
  goDetail(id) {
      // 使用JS的形式进行路由导航

      // 注意： 一定要区分 this.$route 和 this.$router 这两个对象，
      // 其中： this.$route 是路由【参数对象】，所有路由中的参数， params, query 都属于它
      // 其中： this.$router 是一个路由【导航对象】，用它 可以方便的 使用 JS 代码，实现路由的 前进、后退、 跳转到新的 URL 地址

      console.log(this);
      // 1. 最简单的
      // this.$router.push("/home/goodsinfo/" + id);
      // 2. 传递对象
      // this.$router.push({ path: "/home/goodsinfo/" + id });
      // 3. 传递命名的路由
      this.$router.push({ name: "goodsinfo", params: { id } });
    }
```

注意这里的跳转使用的是编程式导航跳转

```
{ name: "goodsinfo", params: { id } }
```

name:是在router.js定义的name 

```
 { path: '/home/goodsinfo/:id', component: GoodsInfo, name: 'goodsinfo' },
```

点击跳转:

```
 <div class="goods-item" v-for="item in goodslist" :key="item.id" @click="goDetail(item.id)">
```



效果

![](C:\work\github\bingjian-workspace\heima-web-39\vue\vue-cms\img\20.PNG)



## 商品详情页面绘制



使用mui-card的代码片段来绘制商品详情页面

商品录播图区域

```
 
    <!-- 商品轮播图区域 -->
    <div class="mui-card">
      <div class="mui-card-content">
        <div class="mui-card-content-inner">
          <swiper :lunbotuList="lunbotu" :isfull="false"></swiper>
        </div>
      </div>
    </div>

```

商品购买区域

````
    <!-- 商品购买区域 -->
    <div class="mui-card">
      <div class="mui-card-header">{{ goodsinfo.title }}</div>
      <div class="mui-card-content">
        <div class="mui-card-content-inner">
          <p class="price">
            市场价：<del>￥{{ goodsinfo.market_price }}</del>&nbsp;&nbsp;销售价：<span class="now_price">￥{{ goodsinfo.sell_price }}</span>
          </p>
          <p>购买数量：<numbox @getcount="getSelectedCount" :max="goodsinfo.stock_quantity"></numbox></p>
          <p>
            <mt-button type="primary" size="small">立即购买</mt-button>
            <mt-button type="danger" size="small" @click="addToShopCar">加入购物车</mt-button>
            <!-- 分析： 如何实现加入购物车时候，拿到 选择的数量 -->
            <!-- 1. 经过分析发现： 按钮属于 goodsinfo 页面， 数字 属于 numberbox 组件 -->
            <!-- 2. 由于涉及到了父子组件的嵌套了，所以，无法直接在 goodsinfo 页面zhong 中获取到 选中的商品数量值-->
            <!-- 3. 怎么解决这个问题：涉及到了 子组件向父组件传值了（事件调用机制） -->
            <!-- 4. 事件调用的本质： 父向子传递方法，子调用这个方法， 同时把 数据当作参数 传递给这个方法 -->
          </p>
        </div>
      </div>
    </div>
````

商品参数区域

```
    <!-- 商品参数区域 -->
    <div class="mui-card">
      <div class="mui-card-header">商品参数</div>
      <div class="mui-card-content">
        <div class="mui-card-content-inner">
          <p>商品货号：{{ goodsinfo.goods_no }}</p>
          <p>库存情况：{{ goodsinfo.stock_quantity }}件</p>
          <p>上架时间：{{ goodsinfo.add_time | dateFormat }}</p>
        </div>
      </div>
      <div class="mui-card-footer">
        <mt-button type="primary" size="large" plain @click="goDesc(id)">图文介绍</mt-button>
        <mt-button type="danger" size="large" plain @click="goComment(id)">商品评论</mt-button>
      </div>
    </div>
```



### css美化



```
.goodsinfo-container {
  background-color: #eee;
  overflow: hidden;

  .now_price {
    color: red;
    font-size: 16px;
    font-weight: bold;
  }

  .mui-card-footer {
    display: block;
    button {
      margin: 15px 0;
    }
  }

  .ball {
    width: 15px;
    height: 15px;
    border-radius: 50%;
    background-color: red;
    position: absolute;
    z-index: 99;
    top: 390px;
    left: 146px;
  }
}
```



获取录播图数据渲染

```
   getLunbotu() {
      this.$http.get("api/getthumimages/" + this.id).then(result => {
        if (result.body.status === 0) {
          // 先循环轮播图数组的每一项，为 item 添加 img 属性，因为 轮播图 组件中，只认识 item.img， 不认识 item.src
          result.body.message.forEach(item => {
            item.img = item.src;
          });
          this.lunbotu = result.body.message;
        }
      });
    },
```

渲染轮播图



渲染轮播图需要先抽离出之前的轮播图组件，一个公共的轮播图组件

```
<template>
  <div>
    <mt-swipe :auto="4000">
      <!-- 在组件中，使用v-for循环的话，一定要使用 key -->
      <!-- 将来，谁使用此 轮播图组件，谁为我们传递 lunbotuList -->
      <!-- 此时，lunbotuList 应该是 父组件向子组件传值来设置 -->
      <mt-swipe-item v-for="item in lunbotuList" :key="item.url">
        <img :src="item.img" alt="" :class="{'full': isfull}">
      </mt-swipe-item>
    </mt-swipe>
  </div>

  <!-- 分析：为什么 商品评论中的 轮播图那么丑： -->
  <!-- 1. 首页中的图片，它的宽和高，都是 使用了 100% 的宽度 -->
  <!-- 2. 在商品详情页面中，轮播图的 图片，如果也使用 宽高 为 100%的话，页面不好看 -->
  <!-- 3. 商品详情页面中的轮播图，期望 高度是 100%， 但是 宽度为 自适应 -->
  <!-- 4. 经过分析，得到 问题的原因： 首页中的轮播图 和 详情中的轮播图，分歧点是 宽度到底是 100% 还是 自适应 -->
  <!-- 5. 既然这两个 轮播图，其它方面都是没有冲突的，只是 宽度有分歧， 那么，我们可以定义一个 属性，让 使用轮播图的 调用者，手动指定 是否为 100% 的宽度 -->

</template>

<script>
export default {
  props: ["lunbotuList", "isfull"]
};
</script>

<style lang="scss" scoped>
.mint-swipe {
  height: 200px;

  .mint-swipe-item {
    text-align: center;

    img {
      // width: 100%;
      height: 100%;
    }
  }
}

.full {
  width: 100%;
}
</style>

```

需要绑定父组件的两个属性

```
"lunbotuList", "isfull"
```

lunbotuList:是轮播图的图片列表数据

isfull:表示是否需要宽度百分百  首页需要百分百，但是商品详情页的宽度不需要百分百



### 调整首页轮播图组件

抽离轮播图组件之后我们重新调整首页的轮播图 HomeContainer.vue

```
import swiper from "../subcomponents/swiper.vue";
```

 注册组轮播图组件

```
components: {
    swiper
  }
```



渲染轮播图组件

```
  <!-- 轮播图区域 -->
    <swiper :lunbotuList="lunbotuList" :isfull="true"></swiper>
```





### 商品轮播图组件

```
 <!-- 商品轮播图区域 -->
    <div class="mui-card">
      <div class="mui-card-content">
        <div class="mui-card-content-inner">
          <swiper :lunbotuList="lunbotu" :isfull="false"></swiper>
        </div>
      </div>
    </div>
```



### 商品购买区域

```


    <!-- 商品购买区域 -->
    <div class="mui-card">
      <div class="mui-card-header">{{ goodsinfo.title }}</div>
      <div class="mui-card-content">
        <div class="mui-card-content-inner">
          <p class="price">
            市场价：<del>￥{{ goodsinfo.market_price }}</del>&nbsp;&nbsp;销售价：<span class="now_price">￥{{ goodsinfo.sell_price }}</span>
          </p>
          <p>购买数量：<numbox @getcount="getSelectedCount" :max="goodsinfo.stock_quantity"></numbox></p>
          <p>
            <mt-button type="primary" size="small">立即购买</mt-button>
            <mt-button type="danger" size="small" @click="addToShopCar">加入购物车</mt-button>
            <!-- 分析： 如何实现加入购物车时候，拿到 选择的数量 -->
            <!-- 1. 经过分析发现： 按钮属于 goodsinfo 页面， 数字 属于 numberbox 组件 -->
            <!-- 2. 由于涉及到了父子组件的嵌套了，所以，无法直接在 goodsinfo 页面zhong 中获取到 选中的商品数量值-->
            <!-- 3. 怎么解决这个问题：涉及到了 子组件向父组件传值了（事件调用机制） -->
            <!-- 4. 事件调用的本质： 父向子传递方法，子调用这个方法， 同时把 数据当作参数 传递给这个方法 -->
          </p>
        </div>
      </div>
    </div>

```



购买数量组件 goodsinfo_numbox.vue

```
<template>
<!-- 问题： 我们不知道什么时候能够拿到 max 值，但是，总归有一刻，会得到一个真正的 max 值 -->
<!-- 我们可以 使用 watch 属性监听，来 监听 父组件传递过来的 max 值，不管 watch 会被触发几次，但是，最后一次，肯定是一个 合法的 max 数值 -->
  <div class="mui-numbox" data-numbox-min='1'>
    <button class="mui-btn mui-btn-numbox-minus" type="button">-</button>
    <input id="test" class="mui-input-numbox" type="number" value="1" @change="countChanged" ref="numbox" />
    <button class="mui-btn mui-btn-numbox-plus" type="button">+</button>
  </div>

  <!-- 分析： 子组件什么时候把 数据传递给父组件 -->
  <!--  -->
</template>

<script>
import mui from "../../lib/mui/js/mui.min.js";

export default {
  mounted() {
    // 初始化数字选择框组件
    mui(".mui-numbox").numbox();
    console.log(this.max);
  },
  methods: {
    countChanged() {
      // 每当 文本框的数据被修改的时候，立即把 最新的数据，通过事件调用，传递给父组件
      // console.log(this.$refs.numbox.value);
      this.$emit("getcount", parseInt(this.$refs.numbox.value));
    }
  },
  props: ["max"],
  watch: {
    // 属性监听
    max: function(newVal, oldVal) {
      // 使用 JS API 设置 numbox 的最大值
      mui(".mui-numbox")
        .numbox()
        .setOption("max", newVal);
    }
  }
};
</script>

<style lang="scss" scoped>

</style>

```





初始化numbox组件

```
import mui from "../../lib/mui/js/mui.min.js";


  mounted() {
    // 初始化数字选择框组件
    mui(".mui-numbox").numbox();
    console.log(this.max);
  },
```



绑定父组件属性max

```
 props: ["max"],
```



监听max属性

```
 watch: {
    // 属性监听
    max: function(newVal, oldVal) {
      // 使用 JS API 设置 numbox 的最大值
      mui(".mui-numbox")
        .numbox()
        .setOption("max", newVal);
    }
  }
```

如果max属性监听到变化就重新设置numbox的最大值



购买数量输入框的变化事件

```

<input id="test" class="mui-input-numbox" type="number" value="1" @change="countChanged" ref="numbox" />

```

调用父组件的getCount方法

```
 methods: {
    countChanged() {
      // 每当 文本框的数据被修改的时候，立即把 最新的数据，通过事件调用，传递给父组件
      // console.log(this.$refs.numbox.value);
      this.$emit("getcount", parseInt(this.$refs.numbox.value));
    }
  },
```



父组件GoodsInfo.vue组件引入数量框组件

```
// 导入 数字选择框 组件
import numbox from "../subcomponents/goodsinfo_numbox.vue";
```



使用:

```
  <p>购买数量：<numbox @getcount="getSelectedCount" :max="goodsinfo.stock_quantity"></numbox></p>

```

### 商品参数区域



```
 <!-- 商品参数区域 -->
    <div class="mui-card">
      <div class="mui-card-header">商品参数</div>
      <div class="mui-card-content">
        <div class="mui-card-content-inner">
          <p>商品货号：{{ goodsinfo.goods_no }}</p>
          <p>库存情况：{{ goodsinfo.stock_quantity }}件</p>
          <p>上架时间：{{ goodsinfo.add_time | dateFormat }}</p>
        </div>
      </div>
```



### 图文介绍和商品评论

````
      <div class="mui-card-footer">
        <mt-button type="primary" size="large" plain @click="goDesc(id)">图文介绍</mt-button>
        <mt-button type="danger" size="large" plain @click="goComment(id)">商品评论</mt-button>
      </div>
````





#### 图文介绍

```
 goDesc(id) {
      // 点击使用编程式导航跳转到 图文介绍页面
      this.$router.push({ name: "goodsdesc", params: { id } });
    },
```

创建图文介绍组件GoodsDesc.vue

```
<template>
  <div class="goodsdesc-container">
    <h3>{{ info.title }}</h3>

    <hr>

    <div class="content" v-html="info.content"></div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      info: {} // 图文数据
    };
  },
  created() {
    this.getGoodsDesc();
  },
  methods: {
    getGoodsDesc() {
      this.$http
        .get("api/goods/getdesc/" + this.$route.params.id)
        .then(result => {
          if (result.body.status === 0) {
            this.info = result.body.message[0];
          }
        });
    }
  }
};
</script>

<style lang="scss">
.goodsdesc-container {
  padding: 4px;
  h3 {
    font-size: 16px;
    color: #226aff;
    text-align: center;
    margin: 15px 0;
  }
  .content{
    img{
      width: 100%;
    }
  }
}
</style>

```



#### 商品评论

```
goComment(id) {
      // 点击跳转到 评论页面
      this.$router.push({ name: "goodscomment", params: { id } });
    },
```

创建商品品论组件，之前已经抽离好商品评论组件直接使用就好了

GoodsComment.vue

```
<template>
  <div>
    <cmtbox :id="$route.params.id"></cmtbox>
  </div>
</template>

<script>
import cmtbox from "../subcomponents/comment.vue";

export default {
  components: {
    cmtbox
  }
};
</script>

<style lang="scss" scoped>

</style>

```



# 加入购物车功能

## 加入vuex组件

```
npm install vuex
```
main.js添加:

```
//注册vuex
import Vuex from 'vuex'
Vue.use(Vuex)
var store = new Vuex.Store({
  state:{
    car:[]
  }
})
```

```
var vm = new Vue({
  el: '#app',
  render: c => c(app),
  router, // 1.4 挂载路由对象到 VM 实例上
  store // 挂载 store 状态管理对象
})
```

store就是vuex用来管理对象的


## 点击加入购物车的功能

这里主要涉及几个要点：
1. 加入store的购物车
2. 本地存储持久化购物车的内容，方便在刷新之后重新渲染购物车的内容

需要在store中添加mutations:
```
  addToCar(state,goodsinfo){
      var flag = false;
      state.car.some(item =>{
        if(item.id == goodsinfo.id){
          item.count += parseInt(goodsinfo.count)
          flag = true ;
          return true ;

        }
      })
      if(!flag){
        state.car.push(goodsinfo)
      }
      localStorage.setItem('car',JSON.stringify(state.car))
    }
```
然后在goodsInfo.vue添加加入购物的点击事件
```
    addToShopCar() {
      // 添加到购物车
      this.ballFlag = !this.ballFlag;
      // { id:商品的id, count: 要购买的数量, price: 商品的单价，selected: false  }
      // 拼接出一个，要保存到 store 中 car 数组里的 商品信息对象
      var goodsinfo = {
        id: this.id,
        count: this.selectedCount,
        price: this.goodsinfo.sell_price,
        selected: true
      };
      // 调用 store 中的 mutations 来将商品加入购物车
      this.$store.commit("addToCar", goodsinfo);
    },
```

## 实现微标的数值自动更新

首先我们需要在store中定义个getter方法用来获取store中car的物品数量

main.js中store中添加getters方法:


```
  getters:{
     getAllCount(state){
       var c = 0 ;
       state.car.forEach(item =>{
         c += item.count
       })
       return c;
     }
   }
```

然后在App.vue中获取商品数量

```
		<router-link class="mui-tab-item-llb" to="/shopcar">
				<span class="mui-icon mui-icon-extra mui-icon-extra-cart">
					<span class="mui-badge">{{ $store.getters.getAllCount}}</span>
				</span>
				<span class="mui-tab-label" id="badge">购物车</span>
			</router-link>
```

## 利用本地持久化保持购物车数据

刷新的时候发现购物车的数据又重新变为0了，所以我们可以利用本地持久化来保存购物车的数据，启动项目的
时候先从本地存储获取购物车数据来渲染
main.js中启动读取本地存储

```
//注册vuex
import Vuex from 'vuex'
Vue.use(Vuex)
var car = JSON.parse(localStorage.getItem('car')||[])
```

在store中设置car为本地存储的car
```
Vue.use(Vuex)
var car = JSON.parse(localStorage.getItem('car')||[])
var store = new Vuex.Store({
  state:{
    car:car
  },

  ...其他代码
```

这样刷新就不会丢失为购物车的数据了




## 购物车图片商品列表页面的布局

分为商品列表区域和商品结算区域

![](./img/21.png)

ShopcarContainer.vue

html代码
```
<template>
  <div class="shopcar-container">

    <!-- 商品列表区域 -->
    <div class="goods-list">
      <div class="mui-card">
        <div class="mui-card-content">
          <div class="mui-card-content-inner">
              <mt-switch>

              </mt-switch>
              <img src="">
              <div class="info">
                <h1>商品标题</h1>
                <p>
                  <span class="price">￥10000</span>
                  <numbox :initcount="0"></numbox>
                  <!-- 问题：如何从购物车中获取商品的数量呢 -->
                  <!-- 1. 我们可以先创建一个 空对象，然后循环购物车中所有商品的数据， 把 当前循环这条商品的 Id， 作为 对象 的 属性名，count值作为 对象的 属性值，这样，当把所有的商品循环一遍，就会得到一个对象： { 88: 2, 89: 1, 90: 4 } -->
                  <a href="#" @click.prevent="remove(item.id, i)">删除</a>
                </p>
            </div>
          </div>
        </div>
      </div>
    </div>

     <div class="goods-list">
      <div class="mui-card">
        <div class="mui-card-content">
          <div class="mui-card-content-inner">
              <mt-switch>

              </mt-switch>
              <img src="">
              <div class="info">
                <h1>商品标题</h1>
                <p>
                  <span class="price">￥10000</span>
                  <numbox :initcount="0"></numbox>
                  <!-- 问题：如何从购物车中获取商品的数量呢 -->
                  <!-- 1. 我们可以先创建一个 空对象，然后循环购物车中所有商品的数据， 把 当前循环这条商品的 Id， 作为 对象 的 属性名，count值作为 对象的 属性值，这样，当把所有的商品循环一遍，就会得到一个对象： { 88: 2, 89: 1, 90: 4 } -->
                  <a href="#" @click.prevent="remove(item.id, i)">删除</a>
                </p>
            </div>
          </div>
        </div>
      </div>
    </div>

  <!-- 商品结算区域 -->

    <!-- 结算区域 -->
    <div class="mui-card">
      <div class="mui-card-content">
        <div class="mui-card-content-inner jiesuan">
          <div class="left">
            <p>总计（不含运费）</p>
            <p>
              已勾选商品
              <span class="red">000</span> 件， 总价
              <span class="red">111111</span>
            </p>
          </div>
          <mt-button type="danger">去结算</mt-button>
        </div>
      </div>
    </div>

  </div>
</template>

```
css代码:
```

<style lang="scss" scoped>
.shopcar-container {
  background-color: #eee;
  overflow: hidden;
  .goods-list {
    .mui-card-content-inner {
      display: flex;
      align-items: center;
    }
    img {
      width: 60px;
    }
    h1 {
      font-size: 13px;
    }
    .info {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      .price {
        color: red;
        font-weight: bold;
      }
    }
  }
  .jiesuan {
    display: flex;
    justify-content: space-between;
    align-items: center;
    .red {
      color: red;
      font-weight: bold;
      font-size: 16px;
    }
  }
}
</style>
```

用到了一个组件numbox

```
import numbox from "../subcomponents/shopcar_numbox.vue";

export default {
  data() {
    return {
      goodslist: [] // 购物车中所有商品的数据
    };
  },
  created() {
  },
  methods: {
   
  },
  components: {
    numbox
  }
};
```

numbox组件的代码subcompoonents/shopcar_numbox.vue
```
<template>
  <div class="mui-numbox" data-numbox-min='1' style="height:25px;">
    <button class="mui-btn mui-btn-numbox-minus" type="button">-</button>
    <input id="test" class="mui-input-numbox" type="number" :value="initcount" @change="countChanged" ref="numbox" readonly />
    <button class="mui-btn mui-btn-numbox-plus" type="button">+</button>
  </div>

  <!-- 分析： 子组件什么时候把 数据传递给父组件 -->
  <!--  -->
</template>

<script>
import mui from "../../lib/mui/js/mui.min.js";

export default {
  mounted() {
    // 初始化数字选择框组件
    mui(".mui-numbox").numbox();
  },
  methods: {
    countChanged() {
      // 数量改变了
      // console.log(this.$refs.numbox.value);
      // 每当数量值改变，则立即把最新的数量同步到 购物车的  store 中，覆盖之前的数量值
      this.$store.commit("updateGoodsInfo", {
        id: this.goodsid,
        count: this.$refs.numbox.value
      });
    }
  },
  props: ["initcount", "goodsid"]
};
</script>

<style lang="scss" scoped>

</style>

```

需要绑定父组件的两个属性，初始化的数量initcount,商品id goodsid

## 获取购物车所有的商品列表并显示

我们需要从store获取所有的商品ID，然后请求后台返回所有的商品列表

ShopcarContainer.vue
```
   getGoodsList() {
      // 1. 获取到 store 中所有的商品的Id，然后拼接出一个 用逗号分隔的 字符串
      var idArr = [];
      this.$store.state.car.forEach(item => idArr.push(item.id));
      // 如果 购物车中没有商品，则直接返回，不需要请求数据接口，否则会报错
      if (idArr.length <= 0) {
        return;
      }
      // 获取购物车商品列表
      this.$http
        .get("api/goods/getshopcarlist/" + idArr.join(","))
        .then(result => {
          if (result.body.status === 0) {
            this.goodslist = result.body.message;
          }
        });
    },
```
在组件创建完成之后调用
```
created() {
    this.getGoodsList();
  },
```

html中遍历商品列表
```
    <!-- 商品列表区域 -->
    <div class="goods-list">
      <div class="mui-card" v-for="(item,i) in goodslist" :key="item.id">
        <div class="mui-card-content">
          <div class="mui-card-content-inner">
              <mt-switch></mt-switch>
              <img src="">
              <div class="info">
                <h1>{{ item.title }}</h1>
                <p>
                  <span class="price">￥{{ item.sell_price }}</span>
                  <numbox :initcount="0"></numbox>
                  <!-- 问题：如何从购物车中获取商品的数量呢 -->
                  <!-- 1. 我们可以先创建一个 空对象，然后循环购物车中所有商品的数据， 把 当前循环这条商品的 Id， 作为 对象 的 属性名，count值作为 对象的 属性值，这样，当把所有的商品循环一遍，就会得到一个对象： { 88: 2, 89: 1, 90: 4 } -->
                  <a href="#" @click.prevent="remove(item.id, i)">删除</a>
                </p>
            </div>
          </div>
        </div>
      </div>
    </div>
```

## 循坏购物车列表初始化数量值
需要在store中获取商品的数量
main.js中的store添加getters
```
 getGoodsCount(state) {
      var o = {}
      state.car.forEach(item => {
        o[item.id] = item.count
      })
      return o
    },
```
返回的是一个map,对应商品ID和商品数量
在numbox中绑定initcout的属性

ShopcarContainer.vue
```
 <numbox :initcount="$store.getters.getGoodsCount[item.id]"></numbox>
```
## 购物车数量的变化同步到store中
首先我们需要在store添加updateGoodsInfo方法更新store中的购物车状态
main.js store的mutations

```
    updateGoodsInfo(state, goodsinfo) {
      // 修改购物车中商品的数量值
      // 分析： 
      state.car.some(item => {
        if (item.id == goodsinfo.id) {
          item.count = parseInt(goodsinfo.count)
          return true
        }
      })
      // 当修改完商品的数量，把最新的购物车数据，保存到 本地存储中
      localStorage.setItem('car', JSON.stringify(state.car))
    },
```
然后在ShopcarContainer.vue中添加更新购物车状态的事件,由于数值变化是在子组价numbox中，所以我们在shopcar_numbox.vue添加更新数值的更新事件

shopcar_numbox.vue

```
  countChanged() {
      // 数量改变了
      // console.log(this.$refs.numbox.value);
      // 每当数量值改变，则立即把最新的数量同步到 购物车的  store 中，覆盖之前的数量值
      this.$store.commit("updateGoodsInfo", {
        id: this.goodsid,
        count: this.$refs.numbox.value
      });
    }
  },
```
mui-numbox组件添加改变事件
```
 @change="countChanged"
```


注意这里的goodsid是从父组件绑定过来的属性,在ShopcarContainer.vue中
```
 <numbox :initcount="$store.getters.getGoodsCount[item.id]"  :goodsid="item.id"></numbox>

```

## 删除购物车商品功能

在store中添加删除的mutations
```
 removeFormCar(state, id) {
      // 根据Id，从store 中的购物车中删除对应的那条商品数据
      state.car.some((item, i) => {
        if (item.id == id) {
          state.car.splice(i, 1)
          return true;
        }
      })
      // 将删除完毕后的，最新的购物车数据，同步到 本地存储中
      localStorage.setItem('car', JSON.stringify(state.car))
    },
```
在ShopcarContainer.vue中添加删除事件
method
```
 remove(id, index) {
      // 点击删除，把商品从 store 中根据 传递的 Id 删除，同时，把 当前组件中的 goodslist 中，对应要删除的那个商品，使用 index 来删除
      this.goodslist.splice(index, 1);
      this.$store.commit("removeFormCar", id);
    },
```
事件:
```
 <a href="#" @click.prevent="remove(item.id, i)">删除</a>
```


## store选中状态同步到页面

store中添加查询选中状态的mutations

main.js

```
  getGoodsSelected(state) {
      var o = {}
      state.car.forEach(item => {
        o[item.id] = item.selected
      })
      return o
    },
```

然后需要在ShopcarContainer.vue中的mt-switch组件添加v-model

```
  <mt-switch
              v-model="$store.getters.getGoodsSelected[item.id]"
              @change="selectedChanged(item.id, $store.getters.getGoodsSelected[item.id])"
            ></mt-switch>
```

对于修改选中状态我们也需要更新到store中，所以需要给mt-switch添加change事件

```
  <mt-switch
              v-model="$store.getters.getGoodsSelected[item.id]"
              @change="selectedChanged(item.id, $store.getters.getGoodsSelected[item.id])"
            ></mt-switch>
```


在store中添加更新选中状态的mutations

```
 updateGoodsSelected(state, info) {
      state.car.some(item => {
        if (item.id == info.id) {
          item.selected = info.selected
        }
      })
      // 把最新的 所有购物车商品的状态保存到 store 中去
      localStorage.setItem('car', JSON.stringify(state.car))
    }
  },
```

然后在ShopcarContainer.vue添加method
```
selectedChanged(id, val) {
      // 每当点击开关，把最新的 快关状态，同步到 store 中
      // console.log(id + " --- " + val);
      this.$store.commit("updateGoodsSelected", { id, selected: val });
    }
```

## 结算区域数量和总价计算
store中添加getter方法
main.js


getGoodsCountAndAmount
```
  getGoodsCountAndAmount(state) {
      var o = {
        count: 0, // 勾选的数量
        amount: 0 // 勾选的总价
      }
      state.car.forEach(item => {
        if (item.selected) {
          o.count += item.count
          o.amount += item.price * item.count
        }
      })
      return o
    }
```

在ShopcarContainer.vue
```
   <!-- 结算区域 -->
    <div class="mui-card">
      <div class="mui-card-content">
        <div class="mui-card-content-inner jiesuan">
          <div class="left">
            <p>总计（不含运费）</p>
            <p>
              已勾选商品
              <span class="red">{{ $store.getters.getGoodsCountAndAmount.count }}</span> 件， 总价
              <span class="red">￥{{ $store.getters.getGoodsCountAndAmount.amount }}</span>
            </p>
          </div>
          <mt-button type="danger">去结算</mt-button>
        </div>
      </div>
    </div>
```

## 实现返回功能


在App.vue头部添加返回的按钮

```
<!-- 顶部 Header 区域 -->
    <mt-header fixed title="黑马程序员·Vue项目">
      <span slot="left" @click="goBack" v-show="flag">
        <mt-button icon="back">返回</mt-button>
      </span>
    </mt-header>
```

```
 <mt-button icon="back">返回</mt-button>
```

然后添加method:
```
 goBack() {
      // 点击后退
      this.$router.go(-1);
    }
```


其他需要关注的地方:
```
export default {
  data() {
    return {
      flag: false
    };
  },
  created() {
    this.flag = this.$route.path === "/home" ? false : true;
  },
  methods: {
    goBack() {
      // 点击后退
      this.$router.go(-1);
    }
  },
  watch: {
    "$route.path": function(newVal) {
      if (newVal === "/home") {
        this.flag = false;
      } else {
        this.flag = true;
      }
    }
  }
};
```
1. 在创建的时候判断当前的路径如果是家目录就不显示返回按钮
2. 监听当前的路径变化，如果不是/home路径就设置显示返回按钮






