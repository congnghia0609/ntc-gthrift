# Author:       nghiatc
# Email:        congnghia0609@gmail.com

.PHONY: deps
deps:
	@./deps.sh

.PHONY: gen
gen:
	@cd ./example/thrift; thrift -r --gen go tutorial.thrift; cd ../..;
	@echo thrift gen go complete.

.PHONY: server
server:
	@go run main.go

.PHONY: client
client:
	@go run cli.go

.PHONY: ssl
ssl:
	@cd ./ssl; ./gen_ssl.sh; cd ..;
