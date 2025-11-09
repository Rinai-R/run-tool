# run-tool

自用，如需使用可能需要修改代码

用于快速启动 qq，飞书，vscode 而无需指定环境变量的小工具

## 环境

go 1.25.4

makefile

## 用法

```bash
make do

run-code [目录]
```

对于 `run-code` 只需要第一次唤起 vscode 时使用 `run-code` 即可，
后续直接在 vscode 内部终端使用 `code` 唤起其他窗口可以正常工作。