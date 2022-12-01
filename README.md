# guBasic
Go Util on Basic Functions

# go mod

```shell
# go mod init
goUtils/guBasic> go mod init github.com/moondevgo/guBasic
# go mod tidy
goUtils/guBasic> go mod tidy
```

# env 설정
## windows10
시스템 환경 변수 편집 > 고급 > 환경변수
user에 대한 사용자 변수 > 새로만들기
변수: "CONFIG_ROOT"	값: "C:/Dev/inGo/with_go/_config/"

## linux

```bash
$ sudo vi ~/.profile
```

> 편집/수정
```bash
export CONFIG_ROOT=/home/ubuntu/dev/inGo/go-mods/_config/ 추가
```

```bash
# 설정 적용
$ source ~/.profile

# 확인
$ env | grep CONFIG_ROOT
```