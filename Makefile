.PHONY: build

DOCKERTAG?=latest

all: build push

build:
	sh ./scripts/build_docker.sh $(DOCKERTAG)

push:
	sh ./scripts/push_docker.sh $(DOCKERTAG)
