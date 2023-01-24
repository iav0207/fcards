bin_name=fcards

play:
	make build
	bin/$(bin_name) play ~/.fcards/*

build:
	go build -o bin/$(bin_name)

