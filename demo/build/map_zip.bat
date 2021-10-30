:: 解决bat utf-8编码，显示乱码问题
CHCP 65001

del ..\map-nearest.zip
del ..\map-cluster.zip
del ..\map-heatmap.zip

mkdir ..\map\static-gaode\效果

copy /y ..\static-gaode\nearest.html ..\map\static-gaode\nearest.html
@REM mkdir ..\map\static-gaode\效果\nearest
XCOPY ..\static-gaode\效果\nearest  ..\map\static-gaode\效果\  /E /Y
:: 需要安装7z，并且windows
D:\"Program Files"\7-Zip\7z.exe a ..\map-nearest.zip ..\map
del ..\map\static-gaode\nearest.html
@REM todo:check
rmdir /s /q ..\map\static-gaode\效果\*

@REM mkdir ..\map\static-gaode\效果\cluster
XCOPY ..\static-gaode\效果\cluster  ..\map\static-gaode\效果\  /E /Y
copy /y ..\static-gaode\cluster.html ..\map\static-gaode\cluster.html
:: 需要安装7z，并且windows
D:\"Program Files"\7-Zip\7z.exe a ..\map-cluster.zip ..\map
del ..\map\static-gaode\cluster.html
@REM todo:check
rmdir /s /q ..\map\static-gaode\效果\*

@REM mkdir ..\map\static-gaode\效果\heatmap
XCOPY ..\static-gaode\效果\heatmap  ..\map\static-gaode\效果\  /E /Y
copy /y ..\static-gaode\heatmap.html ..\map\static-gaode\heatmap.html
:: 需要安装7z，并且windows
D:\"Program Files"\7-Zip\7z.exe a ..\map-heatmap.zip ..\map
del ..\map\static-gaode\heatmap.html
@REM rmdir /s /q ..\map\static-gaode\效果\*

rmdir /s /q ..\map\static-gaode\效果
