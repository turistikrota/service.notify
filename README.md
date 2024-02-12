# Run Project

How to run project

## 1. add swarm network

```bash
docker network create --driver overlay --attachable turistikrota

```

## 2. add secrets

```bash
docker secret create jwt_private_key ./jwtRS256.key
docker secret create jwt_public_key ./jwtRS256.key.pub
docker secret create firebase_service_account ./firebase_service_account.json

```

## 3. build image

```bash
docker build --build-arg GITHUB_USER=<GITHUB_USER> --build-arg GITHUB_TOKEN=<GITHUB_TOKEN> -t github.com/turistikrota/service.notify .  
```

## 4. run container

```bash
docker service create --name notify-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --secret firebase_service_account --env-file .env --publish 6017:6017 github.com/turistikrota/service.notify:latest
```
