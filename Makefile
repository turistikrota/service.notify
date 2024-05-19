build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.notify . 

run:
	docker service create --name notify-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --secret firebase_service_account --env-file .env --publish 6017:6017 github.com/turistikrota/service.notify:latest

remove:
	docker service rm notify-api-turistikrota-com

stop:
	docker service scale notify-api-turistikrota-com=0

start:
	docker service scale notify-api-turistikrota-com=1

restart: remove build run
	