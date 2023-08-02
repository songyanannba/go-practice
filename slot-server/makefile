all:init-doc

# 初始化并运行后台
doc-run:init-doc run

# 初始化文档
init-doc:
	@swag init -t Seamless-Provider,Seamless-Operator
	
# 运行后台
run:
	@go run main.go -type=backend

# 运行api
run-api:
	@go run main.go -type=api

# git排除配置文件
skip-conf:
	@git update-index --skip-worktree ./config.yaml

# git取消排除配置文件
un-skip-conf:
	@git update-index --no-skip-worktree ./config.yaml