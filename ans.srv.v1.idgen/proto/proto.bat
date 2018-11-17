@echo off
echo ----------------------------------------------------------
echo                      Welcome Use     
echo ----------------------------------------------------------
echo working .....
@echo on
protoc --proto_path=. --micro_out=. --go_out=. idgen.proto
protoc --go_out=plugins=micro:. idgen.proto
@echo off
echo.