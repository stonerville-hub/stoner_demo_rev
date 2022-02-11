 #!/bin/bash
echo "DOCKER: Verifying installation..."

# Verify that we can at least get version output 
##  TODO: Install if missing
if ! docker --version; then
	echo "ERROR: Did Docker get installed?"
	exit 1
fi

echo "DOCKER: Successfully verified installation!"

ENV=dev
HOST=127.0.0.1
PORT=8080
NAMESPACE=test
SET=users

if [ "$#" != "0" ]
   then
        ENV=$1
        NAMESPACE=$2
        SET=$3
 fi

echo "---------------------------"
echo "Connection Info"
echo "Environment: " $ENV
echo "---------------------------"
echo "Host:      "$HOST
echo "Port:      "$PORT
echo "Namespace: "$NAMESPACE
echo "Set:       "$SET
echo ""
 # Install - Docker must be installed
docker build --pull --rm -f "Dockerfile" -t demorevcontent:1.0 "." --build-arg environment=$ENV --build-arg aero_host=$HOST --build-arg aero_port=$PORT --build-arg aero_namespace=$NAMESPACE --build-arg aero_set=$SET
# Install - Docker compose must be installed
docker-compose up -d

exit 1
