README.md

# 使用

1. windows

    直接运行server.exe

    浏览 http://127.0.0.1:8090/index
    
2. linux

    sudo chmod +x server (ubuntu)
    
    直接运行 ./server

    浏览 http://127.0.0.1:8090/index

# 高德
1. 搜索最近的点 

   点击地图，显示最近的5公里范围内的7个点

   中心点大概在上海，随机生成的300个点

2. 点聚合

    支持后端聚合后再聚合，每次查询当前窗口内的点，如果明细数据量大，可以后端聚合后返回前端

    优化高德地图最顶层点，样式控，高德默认明细的点是固定的样式，插件外不能控制

    中心点大概在上海，随机生成的300个点

3. 热力图

    默认随机生成3000个点，权重随机，这样才能看到热力图的效果，如果权重都一样，看到就是孤立的点的聚合效果

    中心点大概在上海，随机生成的300个点

    web页面，支持windows和ubuntu linux

# google 
需要科学上网
1. K-Nearest (click on a map to show the k-nearest points)

2. Cluster

3. Heatmap