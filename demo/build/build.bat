@REM cd ../geoindex
go build -o ../server.exe ../geoindex

set GOOS=linux
go build -o ../server ../geoindex
set GOOS=windows