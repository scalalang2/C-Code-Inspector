## C Code 검사기
본 프로젝트는 40명이 넘는 수강생의 C 프로그래밍 수업 제출물을 채점해야 하는 조교의 귀차니즘을 해결하기 위해 제작되었습니다.

#### Features
* C언어 코드 자동 컴파일 및 실행하기
* 입력값 부여하기

#### Requirements
```shell script
$ go version
go version go1.13.4 darwin/amd64
```

#### Usage
```shell script
$ git clone https://github.com/scalalang2/C-Code-Inspector
$ cd ./C-Code-Inspector && make build
$ ./bin/inspector

Usage: 
  -d string
    	directory
  -i string
    	input file
  -s string
    	student size (default "50")
  -t string
    	target files format
```

#### Execution
아래 결과와 같이 수강생의 모든 C파일을 컴파일하고 실행한 뒤 실행결과를 출력해줍니다.

```shell script
$ ./bin/inspector -d ./ -i input.txt -t p1.c
dir: "./"
students: "50"
input: "input.txt"
target: "p1.c"

input:
7 2

filename: cp2_20200001_p1.c, compile: success, runtime: success
filename: cp2_20160002_p1.c, compile: success, runtime: success
filename: cp2_20180003_p1.c, compile: success, runtime: success
total student: 3

---- output ----
[ filename: cp1_20200001_p1.c ]
Input number: Result: 7

[ filename: cp2_20160002_p1.c ]
Input number: Result: 7

[ filename: cp2_20180003_p1.c ]
Input number:Result: 7
```
