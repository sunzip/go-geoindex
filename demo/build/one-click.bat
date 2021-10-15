@REM start call ***.bat 是并发，适合同时启动多个服务
@REM call ***.bat 是顺序，适合流程执行
call build.bat
call build4linux.bat
call copy.bat
set GOOS=windows