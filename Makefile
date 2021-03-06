mweb:
	mkdir ./log
	git push origin master
	glide cache-clear
	glide update
	go build -o mweb ./*.go

build:
	go build -o mweb ./*.go

update:
	glide update

push:
	git push origin master

clean:
	rm mweb
	rm -rf log
	glide cache-clear
