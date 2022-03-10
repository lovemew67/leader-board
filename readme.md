# leader-board

## binaries

```
- go:                       1.16.4
- protoc: 					libprotoc 3.14.0
- protoc-gen-go:			v1.25.0
- protoc-gen-go-grpc:		1.0.1
- protoc-gen-openapiv2:     Version 2.8.0, commit bd0d991030a391d41b30a895a33cc92b16735c46, built at 2022-03-02T17:58:16Z
- protoc-go-inject-tag:		v1.3.0
```

## how to

- 單元測試
```
make test
```

- 本地起Redis (optional)
```
docker-compose up -d
```

- 本地起API server (single server)
```
make leaderboard
```
預設讀取local.toml作為config, API server listen 8999 port, sqlite檔案放在data資料夾下, redis連線到localhost:6379, 使用DB 0

- 本地起API server (multiple server)
```
make install
```
go install到$GOPATH/bin底下, 使用leader-board binary起API server. leader-board會預設查找$HOME/.leaderboard.toml, 當想要起多個API server, 有兩個建議做法

1. 準備多份config file, 使用--config flag overwrite預設路徑
```
leader-board leaderboard --config somewhere.toml
```

或是

2. 先touch一份空的config放在$HOME/.leaderboard.toml, 再用環境變數做設定
```
REST_PORT=7999 DATABASE_REDIS_HOST=localhost:6379 DATABASE_SQLITE_FOLDER=/Users/zzh/Downloads/sqlite leader-board leaderboard
```

- 起本地API文件viewer
```
make swaggerui
```

## library usages

- github.com/lovemew67/public-misc/cornerstone: 存放靜態的build time info, 重新包裝原生context, 使用sync.Map作為Getter/Setter的storage, 包裝logrus作為logger, 預設debug level, 會把sync.Map內容也一併寫到log去
- github.com/lovemew67/public-misc/connection-poller: 簡單包裝redigo, expose Ping, Set, GetBytes, Delete, Exist methods
- github.com/gin-gonic/gin: HTTP server framework
- github.com/gin-contrib/cors: 為了SwaggerUI能夠使用, 需要cors middleware
- github.com/jinzhu/gorm: sqlite ORM library
- github.com/spf13/cobra: CLI interface, 讓API server作為一個sub command啟動, 方便未來擴充
- github.com/spf13/viper: Configuration solution, 跟cobra同樣的大神作者, 支持多層config來源
- github.com/mitchellh/go-homedir: 搭配viper使用, 抓取home folder位置
- github.com/ory/dockertest/v3: 在單元測試的時候幫忙起Docker container做測試
- github.com/robfig/cron/v3: cron library
- github.com/stretchr/testify: test assertion in go

## design principles

- 使用protobuf定義DAO是為了搭配grpc-gateway protoc-gen-openapiv2幫忙產生swagger文件, 搭配container就能就地使用. 不過目前公司是使用blueprint, 寫markdown
- 使用cobra寫成command line pattern, 這樣同一份code base可以實作不同API server, 方便打包成不同docker container. 如果有需要也可以release成CLI tool
- 使用viper讀取config, 能夠輕鬆在不同層級overwrite
- database跟cache部分使用sqlite跟redis, 因為relational database跟cache目前只使用過這兩種. 平台其他大部分應用使用mongoDB, 記帳使用OracleDB
- 使用clean architectire寫法, 切割controller, service, repositoy跟domain. 層層包法讓各層能專注在各自的功能, 不用在意下層的實作, 單元測試也能用gomock簡單mock掉. 不過取捨就是當實作開始多了, 一旦有interface change會很痛苦
- 因為multi server設定, API server裡面不存放狀態 (local cache, etc.), 盡量從remote cache跟database存取狀態, 寫入資料的時候清除cache, 讀取時先從cache讀, 如果miss在從database讀資料再塞入cache
- 每隔十分鐘清除的設計因為來不及思考是從什麼時候開始算十分鐘, 所以實作cronjob每10th分鐘清理database table跟cache
- API response加上transit id, 如果server有接ELK的話, 方便在上面查找

## TODO
- rate limitation
- singleflight
- more clear spec about leader board reset

## FIXME
- upsert in sqlite
- protoc-gen-openapiv2 does not support snake case