# 参与贡献

感谢你愿意参与 cosh 的开发！在提交代码前请阅读以下约定。

## 开发环境

- Go 1.26+
- 腾讯云 COS 账号（用于本地功能验证，非必须）

## 工作流

1. Fork 本仓库并克隆到本地
2. 创建功能分支：`git checkout -b feat/xxx`
3. 提交代码，commit message 遵循 [Conventional Commits](https://www.conventionalcommits.org/)：
   - `feat: 新增 xxx 功能`
   - `fix: 修复 xxx 问题`
   - `docs: 更新文档`
   - `refactor: 重构 xxx`
4. 本地验证通过后推送到你的仓库，并提交 Pull Request

## 代码规范

- 代码必须通过 `gofmt` 格式化
- 提交前运行 `go vet ./...` 确认无静态检查问题
- 改动保持最小化，不做与需求无关的重构
- 新命令请遵循现有命令风格（cobra 子命令，注册到 `cmd/` 下的独立文件）

## 提交 PR 前检查

- [ ] `gofmt -l .` 无输出
- [ ] `go vet ./...` 通过
- [ ] `go build ./...` 通过
- [ ] 涉及配置文件路径/格式变更时同步更新 README
