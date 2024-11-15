docker_start:
	docker pull jordan/icinga2
	docker run -d --name icinga2 -p 8080:80 -p 8443:443 -p 5665:5665 -it jordan/icinga2:latest
	sleep 20

docker_get_root_password:
	$(eval password:=$(shell docker exec icinga2 bash -c 'grep password /etc/icinga2/conf.d/api-users.conf' | awk -F'"' '{ print $$2}' ))
	echo $(password)

docker_reset:
	make docker_clean
	make docker_start
	make set_env

docker_clean:
	docker stop icinga2
	docker rm icinga2

set_env:
	$(eval password:=$(shell docker exec icinga2 bash -c 'grep password /etc/icinga2/conf.d/api-users.conf' | awk -F'"' '{ print $$2}' ))
	echo "ICINGA2_API_PASSWORD=$(password)" > .env

test:
	$(eval password:=$(shell docker exec icinga2 bash -c 'grep password /etc/icinga2/conf.d/api-users.conf' | awk -F'"' '{ print $$2}' ))
	( export ICINGA2_API_PASSWORD="$(password)" && go test -p 1 ./... -v  )
