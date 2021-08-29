# mapinfo-kartrider-go 
A web app written in go. Check tier table with ease. 

# how to run
```sh
git clone https://github.com/alexlai97/mapinfo-kartrider-go
cd ./mapinfo-kartrider-go

make build
./mapinfo-kartrider load_default_maps
./mapinfo-kartrider serve

# Open your browser at localhost:8080
```

# TODO
## small
- flash error
- more testing
- user can store their record

## big
- multiple users at the same time
- support filtering at "/maps" like "Easy", 
- cache maps on client
- live search for mapname
