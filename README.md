# mapinfo-kartrider-go 
A web app written in go. Check tier table with ease. 

# how to run
```sh
git clone https://github.com/alexlai97/mapinfo-kartrider-go
cd ./mapinfo-kartrider-go

make build
./mapinfo-kartrider load_default_maps
./mapinfo-kartrider serve

# Open your browser at localhost:8080, or using curl
curl localhost:8080/maps
curl localhost:8080/api/v1/maps/
curl localhost:8080/maps/2
curl localhost:8080/api/v1/maps/3
```

# What can you do so far?
1. You can register, and login, then logout. :)
2. You can check the tier table at "localhost:8080/maps" and "localhost:8080/api/v1/maps"

# TODO
## small
- flash error
- more testing
- user can store their record

## big
- deploy the app somewhere
- multiple users at the same time
- support filtering at "/maps" like "Easy", 
- cache maps on client
- live search for mapname
