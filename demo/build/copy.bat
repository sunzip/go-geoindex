:: 解决bat utf-8编码，显示乱码问题
CHCP 65001 

mkdir ..\map\static
mkdir ..\map\static-gaode
:: copy /y ..\server  +  ..\index.html ..\map
copy /y ..\static ..\map\static
copy /y ..\static-gaode ..\map\static-gaode
:: 可执行文件不需要保留到原始目录
move /y ..\server ..\map 
move /y ..\server.exe ..\map
copy /y ..\index.html ..\map
copy /y ..\README.md ..\map

@REM 移除每一项
del ..\map\static-gaode\nearest.html
del ..\map\static-gaode\heatmap.html
del ..\map\static-gaode\cluster.html
