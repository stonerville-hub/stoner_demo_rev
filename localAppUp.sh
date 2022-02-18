unset HOST
unset PORT
unset NAMESPACE
unset SET

# . localStartAeroSpike.sh
echo Building App...
go build .
echo Starting App...
go run .