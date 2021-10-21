@REM start call ***.bat 是并发，适合同时启动多个服务
@REM call ***.bat 是顺序，适合流程执行
call build.bat
call copy.bat

call map_zip.bat

:: 删除文件,如果需要留着测试，则去掉
rmdir /s /q ..\map
