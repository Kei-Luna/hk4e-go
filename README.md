# hk4e-go

[English](README-EN.md)
***

#### [hkrpg-go](https://github.com/gucooing/hkrpg-go) 朋友的项目，欢迎支持

## 简介

#### 『原神』 Game Server But Golang Ver.

#### 本项目的目标为构建一个高性能高可用的ARPG游戏服务端，并非以完整还原游戏内原本功能点为目的

#### 项目的客户端协议、配置表主要基于3.2版本修改而来，因此请尽量使用3.2版本的客户端，但不是必须的

#### 客户端需要打破解补丁才能正常使用，详情请参考目前主流私服连接方法，如[Grasscutter](https://github.com/Grasscutters/Grasscutter)

## 特性

* 原生的高可用集群架构，任意节点宕机不会影响到整个系统，可大量水平扩展
* 玩家级无状态游戏服务器，无锁单线程模型，开发省时省力，完善的玩家数据交换机制(内存-缓存-数据库)，拒绝同步阻塞的数据库访问
* 新颖的玩家在线跨服无缝迁移功能
* 独创的网关服务器侧客户端协议代理转换功能，拒绝因协议号消息号混淆而带来代码改动的烦恼
* 完整的密钥交换机制实现，安全性++，拒绝一个写死的随机数种子和XOR密钥文件用到天荒地老

## 编译和运行环境

* Go >= 1.18
* Protoc >= 3.21
* Protoc Gen Go >= 1.28
* Docker >= 20.10
* Docker Compose >= 1.29

## 快速启动

* 首次需要安装工具

```shell
make dev_tool
```

* 生成协议

```shell
make gen_natsrpc      # 生成natsrpc协议
make gen_proto        # 生成客户端协议
make gen_client_proto # 生成客户端协议代理(非必要 详见gate/client_proto/README.md)
```

* 构建

```shell
make build         # 构建服务器二进制文件
make docker_config # 复制配置模板等文件
make docker_build  # 构建镜像
```

* 启动

```shell
cd docker
# 启动前请先确保各服务器的配置文件正确(如docker/node/bin/application.toml)
docker-compose up -d # 启动服务器
```

#### 第三方组件

* mongodb
* nats-server
* redis

#### 服务器组件

* node 节点服务器 (仅单节点 有状态)
* dispatch 登录服务器 (可多节点 无状态)
* gate 网关服务器 (可多节点 有状态)
* multi 多功能服务器 (可多节点 有状态 尚不完善非必要启动)
* gs 游戏服务器 (可多节点 有状态)
* gm 游戏管理服务器 (仅单节点 无状态)

#### 其它

* 部分服务器组件以本地原生进程方式启动需要添加以下环境变量

```shell
GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore
```

## 代码提交规范

#### 欢迎提交PR

* 提交前**必须**格式化你的代码，如运行`go fmt`
* 进行全局格式化时，请跳过`gdconf/game_data_config`目录，这是配置表数据，包含大量的`json`、`lua`、`txt`等文件
