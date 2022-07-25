## Dns-server

dnsf is DNS forwarder which help foward queries queries from internal DNS servers to specified upstream servers. When a query is received by the forwarder, it forward requests checks it local cache and if still unresolved, it then forwards to alternative upstream servers.

## Features
- Improve the efficiency of the resolving dns queries by keeping cache of name server that has been queried frequently using priority queues
- Secures internal DNS server by not exposing the network directly to attackers
- It also reduces traffic on external internet DNS servers.

## How to run
clone the repository - git clone https://github.com/acentior/dnsf
go install you can have the dnsf command installed directly to you $GOPATH/bin folder
run dnsf to confirm if installed properly

## Commands and Flags
Here are some of the custom commands and flag that you can use when trying to model you dns query

## License
It is free and can be deployed anywhere if found safe. Please feel free to report and issue or bug when found and also any enhancements to the repository welcome.