.PHONY: proto data build

proto:
	for d in api srv; do \
		for f in $$d/**/proto/**/*.proto; do \
			protoc --proto_path=.:${GOPATH}/src --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done

lint:
	./scripts/lint.sh

build:
	./scripts/build.sh

data:
	go-bindata -o data/bindata.go -pkg data data/*.json

run:
	docker-compose build
	docker-compose up
