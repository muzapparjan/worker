@echo off
set CGO_ENABLED=1
set CGO_CFLAGS=-D CL_TARGET_OPENCL_VERSION=220 -I %CD%\inc
set CGO_LDFLAGS=-L %CD%\lib\x64
@echo Start building worker...
@echo off
go build -o .\build\worker.exe .\cmd\worker\
echo Start building worker demo...
@echo off
go build -o .\build\demo.exe .\cmd\demo\