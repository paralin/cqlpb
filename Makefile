
gentestpb:
	cd pkg/marshal/marshal_test && \
		protoc -I=$$(pwd) \
		--go_out=$$(pwd) \
		$$(pwd)/test.proto
