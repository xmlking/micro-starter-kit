VERSION:=$(shell cat ./VERSION)

.PHONY: proto data build

proto:
	for d in api srv; do \
		for f in $$d/**/proto/**/*.proto; do \
			protoc --proto_path=.:${GOPATH}/src --micro_out=. --go_out=. --validate_out=lang=go:. $$f; \
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

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	# goreleaser --rm-dist

clean:
	for d in api srv; do \
		for sd in $$d/*; do \
			$(MAKE) -C $$sd $(MAKECMDGOALS); \
			echo cleaned: $$sd; \
		done \
	done
