# Author:       nghiatc
# Email:        congnghia0609@gmail.com

.PHONY: deps
deps:
	@./deps.sh

.PHONY: gen
gen:
	@cd ./thrift; thrift -r --gen go tutorial.thrift; cd ..;

.PHONY: server
server:
	@go run server/server.go

.PHONY: client
client:
	@go run client/client.go

.PHONY: ssl
ssl:
	@cd ./ssl; ./gen_ssl.sh; cd ..;
