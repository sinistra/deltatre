### Deltatre test

These instructions assume that GoLang & Docker are installed on the target computer.

#### Deployment instructions

Execute 'deploy.sh' from the deltatre folder.
This will;
* create a docker network for use by the containers
* build, test & deploy the backend container. It accepts requests on port 8001.
* build, test & deploy the frontend container. It accepts requests on port 8000.
* the frontend container uses the docker network to communicate with the backend.

The following URLs are available on frontend;

* http://localhost:8000/api/v1/search/{name} where {name} is the word to search for
* http://localhost:8000/api/v1/add/{name} where name is the word to add to the list
* http://localhost:8000/api/v1/top

The following URLs are available on backend;

* http://localhost:8001/api/v1/words - the current word list
* http://localhost:8001/api/v1/search/{name} where {name} is the word to search for
* http://localhost:8001/api/v1/add/{name} where name is the word to add to the list
* http://localhost:8001/api/v1/top

The frontend communicates with the backend using a Docker hostname that is stored in an environment variable in the frontend container.

The communication between the containers is via REST (not gRPC). 
I am familiar with the theory of gRPC but have not used it sufficiently to justify use in this project.
