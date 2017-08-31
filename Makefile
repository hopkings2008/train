.PHONY: all clean bin

TEST_TARGET="test"

GOFLAGS=-gcflags "-N -l"

all:
	go test -check.vv

bin:
	export CGO_ENABLED=0 && go test -c $(GOFLAGS) -o $(TEST_TARGET)

clean:
	-@rm -f $(TEST_TARGET)
