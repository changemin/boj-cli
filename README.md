# BaekJoon CLI

## 설치하기

```
git clone https://github.com/changemin/BaekJoon-CLI

cd BaekJoon-CLI

go install bj
```

## 파일 구조

```
📦WORKSPACE
 ┣ 📂1000
 ┃  ┗ 📂1100
 ┃    ┣ 📂1110
 ┃    ┃ ┣ 📜1110.c
 ┃    ┃ ┣ 📜1111.c
 ┃    ┃ ┣ 📜1112.c
 ┃    ┃ ┣ 📜1113.c
 ┃    ┃ ┣ 📜1114.c
 ┃    ┃ ┗ 📜1115.c
 ┃    ┗ 📂1120
 ┃      ┗ 📜1120.c
 ┗ 📜config.bj.yaml
```
> 10개 단위로 묶음

## Config
```yaml
username: {username}
language: {language}
```

지원가능언어 : `c`, `c++` ,`swift`, `Java`

## Concept
```
// 
// 2020-03-30
// Created by {username}
//
// 1000.swift
// https://njssikq.com~~
//
// 문제 : A+B이 있다 합을 구하라
// 입력 : 어떻게 저렇게 온다
// 출력 : 이렇게 이렇게 해라
//
```

## 명령어

|명령어|설명|
|:---:|:---:|
|`bj init`|백준 WorkSpace를 생성합니다|
|`bj get [문제번호]`|문제를 Parse하여 WorkSpace안에 파일을 생성합니다|
|`bj home`|WorkSpace의 홈으로 이동합니다.|
