
test:
	cd marshal/ && go test

gentestpb:
	cd marshal/marshal_test && \
		protoc -I=$$(pwd) \
		--go_out=$$(pwd) \
		$$(pwd)/test.proto
