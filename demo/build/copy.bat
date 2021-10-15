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
