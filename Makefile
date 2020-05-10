.PHONY: build

DOCKERTAG?=latest

all: build push deploy update

build:
	sh ./scripts/docker_build.sh $(DOCKERTAG)

push:
	sh ./scripts/docker_push.sh $(DOCKERTAG)

deploy:
	sh ./scripts/k8s_create.sh

update:
	sh ./scripts/k8s_update.sh $(DOCKERTAG)

