# BaekJoon Online Judge File Management Tool
![백준 로고](imgs/logo.png)

## 만들게된 계기 

적지 않은 사람들이 자신이 해결한 백준 문제를 깃허브와 같은 저장소에 올리는 것을 보았다.
하지만 체계적으로 파일을 업로드하기 위해서는 매번 폴더, 파일을 생성하고 이름을 바꾸는 등의 반복적인 프로세스를 간단하게 해결해보고 싶었다. 
> [https://github.com/Changemin/boj-solving](https://github.com/Changemin/boj-solving) 에서 위 프로젝트를 적용하여 문풀을 하고 있다!

## 설치하기

```
git clone https://github.com/Changemin/boj-cli

cd boj-cli

go install bj
```

## 명령어

|명령어|설명|
|:---:|:---:| 
|`bj init`|백준 WorkSpace를 생성합니다|
|`bj get [문제번호]`|문제를 Parse하여 WorkSpace안에 파일을 생성합니다|
|`bj cp [문제번호]`|작성된 소스코드를 클립보드에 복사합니다.|
|`bj open [문제번호]`|문제가 저장되어 있는 폴더를 엽니다.|

## 활용예제

```
$ bj get 1000~1005 2000~2005
```

```
📦 WORKSPACE
 ┣ 📂1000번~1099번
 ┃ ┣ 📂1000번-A+B
 ┃ ┃ ┗ 📜1000.c
 ┃ ┣ 📂1001번-A-B
 ┃ ┃ ┗ 📜1001.c
 ┃ ┣ 📂1002번-터렛
 ┃ ┃ ┗ 📜1002.c
 ┃ ┣ 📂1003번-피보나치 함수
 ┃ ┃ ┗ 📜1003.c
 ┃ ┣ 📂1004번-어린 왕자
 ┃ ┃ ┗ 📜1004.c
 ┃ ┗ 📂1005번-ACM Craft
 ┃ ┃ ┗ 📜1005.c
 ┣ 📂2000번~2099번
 ┃ ┣ 📂2000번-책장제작
 ┃ ┃ ┗ 📜2000.c
 ┃ ┣ 📂2001번-보석 줍기
 ┃ ┃ ┗ 📜2001.c
 ┃ ┣ 📂2002번-추월
 ┃ ┃ ┗ 📜2002.c
 ┃ ┣ 📂2003번-수들의 합 2
 ┃ ┃ ┗ 📜2003.c
 ┃ ┣ 📂2004번-조합 0의 개수
 ┃ ┃ ┗ 📜2004.c
 ┃ ┗ 📂2005번-사발
 ┃ ┃ ┗ 📜2005.c
 ┗ 📜.BaekJoon.yml
```

## BjConfig
```yaml
username: {username}
file-extension: {languageExtension}
comment-style: {commentStyle}
```

> `extension`을 통해 어떠한 언어로도 커스텀 가능

```
//
// 2021-04-08
//
// Created By 변경민
//
// 1000번 : A+B
// https://www.acmicpc.net/problem/1000
//
// * 문제
//
// 두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오.
//
// * 입력
//
// 1 2
//
// * 출력
//
// 3
//
```
