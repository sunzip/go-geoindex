:: 解决bat utf-8编码，显示乱码问题
CHCP 65001

del ..\map-nearest.zip
del ..\map-cluster.zip
del ..\map-heatmap.zip

mkdir ..\map\static-gaode\效果

copy /y ..\static-gaode\nearest.html ..\map\static-gaode\nearest.html
mkdir ..\map\static-gaode\效果\nearest
XCOPY ..\static-gaode\效果\nearest  ..\map\static-gaode\效果\nearest  /E /Y
:: 需要安装7z，并且windows
D:\"Program Files"\7-Zip\7z.exe a ..\map-nearest.zip ..\map
del ..\map\static-gaode\nearest.html
rmdir /s /q ..\map\static-gaode\效果\nearest

mkdir ..\map\static-gaode\效果\cluster
XCOPY ..\static-gaode\效果\cluster  ..\map\static-gaode\效果\cluster  /E /Y
copy /y ..\static-gaode\cluster.html ..\map\static-gaode\cluster.html
:: 需要安装7z，并且windows
D:\"Program Files"\7-Zip\7z.exe a ..\map-cluster.zip ..\map
del ..\map\static-gaode\cluster.html
rmdir /s /q ..\map\static-gaode\效果\cluster

@REM mkdir ..\map\static-gaode\效果\heatmap
@REM XCOPY ..\static-gaode\效果\heatmap  ..\map\static-gaode\效果\heatmap  /E /Y
copy /y ..\static-gaode\heatmap.html ..\map\static-gaode\heatmap.html
:: 需要安装7z，并且windows
D:\"Program Files"\7-Zip\7z.exe a ..\map-heatmap.zip ..\map
del ..\map\static-gaode\heatmap.html
@REM rmdir /s /q ..\map\static-gaode\效果\heatmap

rmdir /s /q ..\map\static-gaode\效果
