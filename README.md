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
docker build --build-arg GITHUB_USER=<USER_NAME> --build-arg GITHUB_TOKEN=<ACCESS_TOKEN> -t api.turistikrota.com/notify .  
```

## 4. run container

```bash
docker service create --name notify-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --secret firebase_service_account --env-file .env --publish 6017:6017 api.turistikrota.com/notify:latest
```
