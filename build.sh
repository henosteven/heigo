git push origin master
glide cache-clear
glide update
go build -o mweb ./*.go
