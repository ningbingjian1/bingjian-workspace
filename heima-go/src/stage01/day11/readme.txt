/*
1、蛇结构体创建
	长度
	[]坐标
	方向
2、食物结构体创建
	坐标
3、蛇初始化
	设置蛇长度 2节
	设置蛇的坐标
	蛇头方向
4、界面初始化
	游戏界面
	食物初始化
	键盘输入
5、开始游戏
	程序流程控制
	程序更新周期 time.Sleep()

	蛇头和墙的碰撞
	蛇头和身体的碰撞
		游戏结束

	蛇头和食物碰撞
		长度增长
		分数增长
		随机新食物（不能出现在墙 障碍物 蛇身体）

	根据方向改变蛇的位置
		更新蛇的坐标
		身体坐标
		蛇头坐标

	绘制食物 蛇头和身体

6、计算分数
	排行榜 文件

 */