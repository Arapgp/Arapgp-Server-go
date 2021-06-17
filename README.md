# :broken_heart: Arapgp-Server-go

原后端仓库已 `private archive`。

## :heart: 使用

### 1. 使用 docker

#### i. 安装 docker 与 docker-compose

关于 `docker`：

* [For CentOS](https://docs.docker.com/engine/install/centos/)
* [For Debian](https://docs.docker.com/engine/install/debian/)

```bash
# to check docker
docker version
```

关于 `docker-compose`：

[Read this doc to install docker-compose](https://docs.docker.com/compose/install/)

```bash
# to check docker-compose
sudo docker-compose --version
```

#### ii. 修改配置文件 `arapgp.server.json`

在默认的 `arapgp.server.json` 中，`db.mongo.host` 为 `127.0.0.1`，为方便本地调试而设置。这里应该改成数据库所在的 `host`；`port` 等字段同理。

```json
{
  // some other options...

  "db": {
    "mongo": { "host": "127.0.0.1", "port": 27017, "username": "ljg", "password": "ljg", "database": "ljgtest" }
  }

  // some other options...
}
```

#### iii. 直接执行

由于本仓库使用 `go` 作为主要语言，在 `build` / `up` 的过程中，指令 `go build` 可能由于 `go get xxxx` 耗费时间过长而被迫中断。这个时候需要想想办法。

```bash
# at repo root to build
sudo chmod +x ./script/build/docker-init.sh
./script/build/docker-init.sh
```

### 2. 直接 go run

#### i. 安装 go

[Read these doc to install go](https://golang.org/doc/install)

```bash
# check whether go installed well
go version
```

#### ii. 安装 MongoDB

[Read these doc to install MongoDB](https://docs.mongodb.com/manual/administration/install-community/)

[Here to get MongoDB-Community](https://www.mongodb.com/try/download/community)

```bash
# MongoDB-Path is where you install MongoDB
cd <MongoDB-Path>

# to enter MongoDB shell
./bin/mongo
```

还请注意 `arapgp.server.json`，现在使用的是 `ljgtest` 数据库中的 `ljg` 用户。如果没有创建的话，可能得手动创建一下：

```bash
# in MongoDB shell
# to check existed users
> show users

# create a new database (if not existed, MongoDB shell will help to create one)
> use ljgtest

# create a new user
> db.createUser({user: "ljg", pwd: "ljg", roles: [{role: "userAdminAnyDatabase", "db": "admin"}]})
```

#### iii. clone this repo && run it

```bash
# clone this repo
git clone https://github.com/Arapgp/Arapgp-Server-go

# run arapgp.go
cd Arapgp-Server-go
go run arapgp.go
```

## :triangular_flag_on_post: TODO LIST

目前进度：

* [x] `POST     /api/v1/signup`
* [x] `POST     /api/v1/login`
* [x] `POST     /api/v1/logout`
* [x] `GET      /api/v1/user`
* [x] `GET      /api/v1/pubKey`
* [x] `POST     /api/v1/pubKey`
* [x] `PUT      /api/v1/pubKey`
* [x] `DELETE   /api/v1/pubKey`
* [x] `GET      /api/v1/user/:username/file`
* [x] `POST     /api/v1/user/:username/file`
* [x] `PUT      /api/v1/user/:username/file`
* [x] `DELETE   /api/v1/user/:username/file`
* [ ] `GET      /api/v1/ping`
