Setup
-
``make docker`` command will build all the required docker images which then can be run with 
``docker-compose up -d``

Technology Used
-
``redis`` - in memory key value db used to cache the offered challenges to validate on request.\
``grpc-gateway`` - plugin to run RESTful HTTP API server based on the proto file. Also used to 
generate swagger.

POW
-
**Hashcash** was chosen based on it's popularity and simplicity of implementation. Also in this case 
we do not require a validation of chain, previous answers, race or several step protection which 
other algorithms do offer. Simple seek of the valid nonce based on the offered challenge was 
enough. 