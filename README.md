<div align="center">

<h1 align="center">go-bind-mp</h1>

[![Auth](https://img.shields.io/badge/Auth-eryajf-ff69b4)](https://github.com/eryajf)
[![Go Version](https://img.shields.io/github/go-mod/go-version/eryajf/xirang)](https://github.com/eryajf/xirang)
[![Gin Version](https://img.shields.io/badge/Gin-1.6.3-brightgreen)](https://github.com/eryajf/xirang)
[![Gorm Version](https://img.shields.io/badge/Gorm-1.20.12-brightgreen)](https://github.com/eryajf/xirang)
[![GitHub Issues](https://img.shields.io/github/issues/eryajf/xirang.svg)](https://github.com/eryajf/xirang/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/eryajf/xirang)](https://github.com/eryajf/xirang/pulls)
[![GitHub Pull Requests](https://img.shields.io/github/stars/eryajf/xirang)](https://github.com/eryajf/xirang/stargazers)
[![HitCount](https://views.whatilearened.today/views/github/eryajf/xirang.svg)](https://github.com/eryajf/xirang)
[![GitHub license](https://img.shields.io/github/license/eryajf/xirang)](https://github.com/eryajf/xirang/blob/main/LICENSE)

<p> ð ç®åå¥½ç¨ï¼ä¸ç¼ ä¸ç»ï¼ç´æ¥ä¸æçgo-backendæ¡æ¶ </p>

![logo](cmd/docs/logo.png)
![api](cmd/docs/postman.png)


## ð¥¸ é¡¹ç®ä»ç»

`go-bind-mp` æ¯ä¸ä¸ªéå¸¸ç®åç `gin+gorm` æ¡æ¶çåºç¡æ¶æï¼ä½ åªéè¦ä¿®æ¹ç®åçä»£ç ï¼å³å¯å¼å§ä¸æç¼åä½ çæ¥å£ã

åªéè¦æ ¹æ®æåµä¿®æ¹éç½®`config.yml`ï¼ç¶åéç½®éè¾¹çæ°æ®åºéç½®ä¿¡æ¯ï¼å³å¯å¼å§å¼åã

æ°æ®è¡¨ä¼èªå¨åå»ºã

## ð¨âð» é¡¹ç®å°å

| åç±» |                        GitHub                       
| :--: | :--------------------------------------------------
| åç«¯ |  https://github.com/yin-zt/go-bind-mp
|
## ð ç®å½ç»æ

```
go-bind-mp
âââ cmd ----------------ç¨åºå¯å¨èæ¬
    âââ doce---ä»ç»ææ¡£æå¡ãå¾å®æã
    âââ server---------------ä¸»æå¡
        âââ config.yml   ä¸»éç½®æä»¶
        âââ main.go  ä¸»æå¡å¯å¨èæ¬
âââ pkg---------------- é»è¾ä»£ç ç®å½
    âââ config-----------éç½®æä»¶è¯»å
    âââ controller------------æ§å¶å±
    âââ middleware------------ä¸­é´ä»¶
    âââ model---------------å¯¹è±¡å®ä¹
    âââ routers-----------------è·¯ç±
    âââ service---------------æå¡å±
        âââ logic --------------é»è¾å±
        âââ isql --------- æ°æ®åºäº¤äºå±
    âââ util--------ä¸äºå¬å±ç»ä»¶ä¸å·¥å·
```


## ð å¿«éå¼å§

go-bind-mpé¡¹ç®çåºç¡ä¾èµé¡¹åªæMySQLï¼æ¬å°åå¤å¥½mysqlæå¡ä¹åï¼å°±å¯ä»¥å¯å¨é¡¹ç®ï¼è¿è¡è°è¯ã

### æåä»£ç 

```sh
# åç«¯ä»£ç 
$ git clone https://github.com/yin-zt/go-bind-mp.git

### æ´æ¹éç½®

```sh
# ä¿®æ¹åç«¯éç½®
$ cd go-bind-mp
# æä»¶è·¯å¾ config.yml, æ ¹æ®èªå·±æ¬å°çæåµï¼è°æ´æ°æ®åºç­éç½®ä¿¡æ¯ã
$ vim config.yml
```

### å¯å¨æå¡

```sh
# å¯å¨åç«¯
$ cd go-bind-mp
$ go mod tidy
$ go run main.go

```

