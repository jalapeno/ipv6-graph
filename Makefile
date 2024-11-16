REGISTRY_NAME?=docker.io/iejalapeno
IMAGE_VERSION?=latest

.PHONY: all bgpv6-graph container push clean test

ifdef V
TESTARGS = -v -args -alsologtostderr -v 5
else
TESTARGS =
endif

all: bgpv6-graph

bgpv6-graph:
	mkdir -p bin
	$(MAKE) -C ./cmd compile-bgpv6-graph

bgpv6-graph-container: bgpv6-graph
	docker build -t $(REGISTRY_NAME)/bgpv6-graph:$(IMAGE_VERSION) -f ./build/Dockerfile.bgpv6-graph .

push: bgpv6-graph-container
	docker push $(REGISTRY_NAME)/bgpv6-graph:$(IMAGE_VERSION)

clean:
	rm -rf bin

test:
	GO111MODULE=on go test `go list ./... | grep -v 'vendor'` $(TESTARGS)
	GO111MODULE=on go vet `go list ./... | grep -v vendor`
