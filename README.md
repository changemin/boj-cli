# BaekJoon CLI

## 설치하기

```
git clone https://github.com/changemin/BaekJoon-CLI

cd BaekJoon-CLI

go install bj
```

## 활용예제

```
$ bj get 1000 1001 1002 1003 1004
```

```
📦WORKSPACE
 ┣ 📂1000-A+B
 ┃ ┗ 📜1000.c
 ┣ 📂1001-A-B
 ┃ ┗ 📜1001.c
 ┣ 📂1002-터렛
 ┃ ┗ 📜1002.c
 ┣ 📂1003-피보나치 함수
 ┃ ┗ 📜1003.c
 ┗ 📂1004-어린 왕자
 ┃ ┗ 📜1004.c
 ┗ 📜config.bj.yaml
```

## ~~Config 파일~~
```yaml
username: {username}
language: {language}
```

~~지원가능언어 : `c`, `c++` ,`swift`, `Java`~~

```
/*
2021-03-31

Created By {username}

1000번 : A+B
https://www.acmicpc.net/problem/1000

* 문제

두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오.

* 입력

1 2

* 출력

3

*/
```

## 명령어

|명령어|설명|
|:---:|:---:|
|`bj init`|백준 WorkSpace를 생성합니다|
|`bj get [문제번호]`|문제를 Parse하여 WorkSpace안에 파일을 생성합니다|
|`bj cp [문제번호]`|작성된 소스코드를 클립보드에 복사합니다.|
