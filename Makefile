.PHONY: zap
zap:
	@go build -o zapmain ./zap/main.go && \
	du -h zapmain && \
	go version -m ./zapmain

.PHONY: zerolog
zerolog:
	@go build -o zerologmain ./zerolog/main.go && \
	du -h zerologmain && \
	go version -m ./zerologmain
