abi:
	solc --abi ./business/contract.sol -o ./business/build
bin:
	solc --bin ./business/contract.sol -o ./business/build
abigen:
	abigen --abi ./business/build/NonVotedCapped.abi --bin ./business/build/NonVotedCapped.bin --pkg business --out ./business/NonVotedCapped.go
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
evans:
	evans --host localhost --port 9090 -r repl
.PHONY: abi bin abigen proto