:: 解决bat utf-8编码，显示乱码问题
CHCP 65001 
:: 解决将新的内容压缩到旧zip文件里
del map.zip 

mkdir ..\map\static
mkdir ..\map\static-gaode
mkdir ..\map\static-gaode\效果
:: copy /y ..\server  +  ..\index.html ..\map
copy /y ..\static ..\map\static
copy /y ..\static-gaode ..\map\static-gaode
:: copy /y ..\static-gaode\效果 ..\map\static-gaode\效果
XCOPY ..\static-gaode\效果\  ..\map\static-gaode\效果\  /E /Y
:: 可执行文件不需要保留到原始目录
move /y ..\server ..\map 
move /y ..\server.exe ..\map
copy /y ..\index.html ..\map
copy /y ..\README.md ..\map
