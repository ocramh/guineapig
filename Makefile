.PHONY: docker_build
docker_build:
	docker build --tag ocramh/guineapig .

.PHONY: docker_run
docker_run:
	docker run --rm -d --name guineapig -p 8080:8080 ocramh/guineapig lateste
