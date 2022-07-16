echo GIN_MODE=release > .env
docker-compose -f deploy/docker-compose.yml -f deploy/prod/docker-compose.yml up --build