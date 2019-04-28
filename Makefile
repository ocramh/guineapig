.PHONY: docker_build_app
docker_build_app:
	docker build --tag ocramh/guineapig:v0.0.1 --file=docker/Dockerfile.app .

.PHONY: docker_build_db
docker_build_db:
	docker build --tag ocramh/guineapig_db:v0.0.1 --file=docker/Dockerfile.db .

.PHONY: docker_run_app
docker_run_app:
	docker run -d --rm --env-file .env --name guineapig -p 8080:8080 ocramh/guineapig:v0.0.1

.PHONY: docker_run_db
docker_run_db:
	docker run --rm -d --name guineapig_db -p 5432:5432 ocramh/guineapig_db:v0.0.1

.PHONY: docker_psql
docker_psql:
	docker run -it --link guineapig_db:postgres ocramh/guineapig_db:v1 psql -h postgres -d postgres -U user

.PHONY: psql_ip
psql_ip:
	docker inspect guineapig_db | grep IPAddr