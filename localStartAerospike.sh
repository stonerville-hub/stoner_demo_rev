result=$(docker container ls -aq -f "name=aerospike")
if [ $? -eq 0 ]; then
    # mkdir data
    docker run -d -v data:/opt/aerospike/data/ --name aerospike -p 3000-3002:3000-3002 aerospike/aerospike-server
fi
echo "Checking database connection..."