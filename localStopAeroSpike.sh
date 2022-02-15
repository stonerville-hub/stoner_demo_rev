result=$(docker container ls -aq -f "name=aerospike")
if [[ -n "$result" ]]; then
  echo Container $(docker rm -f $result) was removed
  rm -r data
else
  echo "No container to stop and remove"
fi