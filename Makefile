.PHONY: all clean

GOFLAGS=-gcflags "-N -l"

all:
	#export CGO_ENABLED=0 && go build $(GOFLAGS) -o gts

bin:
	export CGO_ENABLED=0 && go test -c $(GOFLAGS)

clean:
	-@rm -f gts
