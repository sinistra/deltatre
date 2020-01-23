docker image build -t frontend:latest .
docker container run -p 8000:8000 -d --rm --name frontend frontend:latest