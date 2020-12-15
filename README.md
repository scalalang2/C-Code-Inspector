## C Code 검사기
본 프로젝트는 40명이 넘는 수강생의 C 프로그래밍 수업 제출물을 채점해야 하는 조교의 귀차니즘을 해결하기 위해 제작되었습니다.

#### Features
* C언어 코드 자동 컴파일
* 입출력 자동/검사 후 틀린 코드 검출하기

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