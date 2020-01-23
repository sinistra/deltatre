docker image build -t backend:latest .
docker container run -p 8001:8001 -d --rm --name backend backend:latest