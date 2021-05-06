# :broken_heart: Arapgp-Server-go

由于 `TypeScript` 我不会配环境，为了前端开发的顺利进行，临时搓了个 `golang` 的后端，以供前端开发时测试验证。

（与 `TypeScript` 同步开发中）

## :heart: 使用

### 1. 安装 go

[Read these doc to install go](https://golang.org/doc/install)

```bash
# check whether go installed well
go version
```

### 2. 安装 MongoDB

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
> db.createUser({user: "ljg", pwd: "ljg", roles: [{role: "readWrite", "db": "ljgtest"}]})
```

### 3. clone this repo && run it

```bash
# clone this repo
git clone https://github.com/Arapgp/Arapgp-Server-go

# run arapgp.go
cd Arapgp-Server-go
go run arapgp.go
```

## :triangular_flag_on_post: TODO LIST

目前进度：

* [x] `POST /api/v1/signup`
* [ ] `POST /api/v1/login`
* [ ] `POST /api/v1/logout`
* [ ] `GET /api/v1/user`
* [ ] `GET /api/v1/pubKey`
* [ ] `POST /api/v1/pubKey`
* [ ] `PUT /api/v1/pubKey`
* [ ] `DELETE /api/v1/pubKey`
