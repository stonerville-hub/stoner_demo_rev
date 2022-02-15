unset HOST
unset PORT
unset NAMESPACE
unset SET

docker run -d --name aerospike -p 3000-3002:3000-3002 aerospike/aerospike-server