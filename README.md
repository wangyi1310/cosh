# ☁ COSH — Tencent Cloud COS CLI

[![Go Version](https://img.shields.io/badge/Go-1.26-blue)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![CI](https://github.com/wangyi1310/cosh/actions/workflows/ci.yml/badge.svg)](https://github.com/wangyi1310/cosh/actions/workflows/ci.yml)

一个「FTP 式」的腾讯云 COS 命令行工具：像操作本地文件一样浏览、上传、下载 COS 上的对象。

```
╔══════════════════════════════════════════╗
║  ☁  COSH - Tencent Cloud COS CLI         ║
║  Browse, upload, download files easily   ║
╚══════════════════════════════════════════╝
```

## ✨ 特性

- 📁 浏览 bucket / 对象：`ls`、`buckets`、`mkdir`
- ⬆️ 上传本地文件到 COS：`put`
- ⬇️ 下载 COS 对象到本地：`get`（带进度条）
- 🗑 删除对象：`rm`
- ⚙️ 交互式或命令行参数配置：`config init`
- 🎨 彩色输出，下载/上传进度一目了然

## 安装

### 从源码构建

需要 Go 1.26+。

```bash
git clone git@github.com:wangyi1310/cosh.git
cd cosh
go build -o cosh .
```

把生成的 `cosh` 二进制放到 `$PATH` 下即可：

```bash
sudo mv cosh /usr/local/bin/
```

### Go install

```bash
go install github.com/wangyi1310/cosh@latest
```

## 快速开始

### 1. 初始化配置

```bash
cosh config init
```

按提示输入：

- **Secret ID** / **Secret Key**：腾讯云 API 密钥（[访问管理控制台](https://console.cloud.tencent.com/cam/capi)获取）
- **Region**：COS 地域，默认 `ap-guangzhou`（如 `ap-beijing`、`na-siliconvalley`）
- **Default Bucket**：默认桶名，如 `mybucket-1250000000`（可留空，之后用 `--bucket` 指定）

也支持一次性传参（非交互）：

```bash
cosh config init \
  --secret_id AKIDxxxx \
  --secret_key xxxx \
  --region ap-guangzhou \
  --bucket mybucket-1250000000
```

配置文件保存在 `~/.cosh.toml`（权限敏感，请勿提交到仓库）。

### 2. 常用命令

```bash
# 列出所有 bucket
cosh buckets

# 列出桶内对象（支持前缀过滤）
cosh ls
cosh ls pics/

# 上传本地文件到 COS
cosh put ./photo.jpg photos/photo.jpg
cosh put ./photo.jpg .          # 远端键名取文件名

# 下载 COS 对象到本地
cosh get photos/photo.jpg
cosh get photos/photo.jpg ./downloads/

# 创建目录
cosh mkdir archive/

# 删除对象
cosh rm photos/photo.jpg
```

### 指定桶 / 地域

每个命令都支持 `--bucket` 和 `--region` 覆盖配置：

```bash
cosh ls --bucket mybucket-1250000000 --region ap-beijing
```

## 命令一览

| 命令      | 说明                     | 用法                              |
|-----------|--------------------------|-----------------------------------|
| `buckets` | 列出所有 COS bucket      | `cosh buckets`                    |
| `ls`      | 列出桶内对象             | `cosh ls [prefix]`                |
| `put`     | 上传本地文件             | `cosh put <local-file> <remote-key>` |
| `get`     | 下载对象                 | `cosh get <remote-key> [local-path]` |
| `mkdir`   | 创建文件夹               | `cosh mkdir <remote-path>`        |
| `rm`      | 删除对象                 | `cosh rm <remote-key>`            |
| `config`  | 管理配置（`init` 初始化）| `cosh config init`                |
| `completion` | 生成 shell 补全脚本   | `cosh completion bash` 等         |

## 配置说明

配置通过 `cosh config init` 生成，存放在 `~/.cosh.toml`：

```toml
secret_id  = "AKIDxxxx"
secret_key = "xxxx"
region     = "ap-guangzhou"
bucket     = "mybucket-1250000000"
```

优先级：命令行参数 `--bucket` / `--region` > 配置文件。

> ⚠️ 请确保 `~/.cosh.toml` 权限为 `600`：`chmod 600 ~/.cosh.toml`

## 开发

```bash
go build ./...      # 编译
go vet ./...        # 静态检查
```

## 许可证

[MIT](LICENSE)
