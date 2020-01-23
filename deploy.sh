docker network create deltatre
docker image build -t backend:latest ./backend
docker container run -p 8001:8001 -d --rm --net deltatre --name backend backend:latest
docker image build -t frontend:latest ./frontend
docker container run -p 8000:8000 -d --rm --net deltatre --name frontend frontend:latest
