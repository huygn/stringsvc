IMAGE=gnhuy91/stringsvc
DCR=docker-compose run --rm

.PHONY: clean test build release docker-build docker-push run

all: release

clean:
	rm -f bin/*

test:
	$(DCR) go-test

build:
	$(DCR) go-build

_run:
	bin/stringsvc

run: build _run

release: test build docker-build docker-push

docker-build:
	docker build --rm -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)
