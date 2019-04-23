.PHONY: docker_build_app
docker_build_app:
	docker build --tag popcore/guineapig:v2.1 --file=docker/Dockerfile.app .

.PHONY: docker_build_db
docker_build_db:
	docker build --tag popcore/guineapig_db:v1 --file=docker/Dockerfile.db .

.PHONY: docker_run_app
docker_run_app:
	docker run --rm -d --name guineapig -p 8080:8080 popcore/guineapig v2.1

.PHONY: docker_run_db
docker_run_db:
	docker run --rm -d --name guineapig_db -p 5432:5432 popcore/guineapig_db:v1

.PHONY: docker_psql
docker_psql:
	docker run -it --link guineapig_db:postgres popcore/guineapig_db:v1 psql -h postgres -d postgres -U user
