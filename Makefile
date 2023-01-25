bin_name=fcards
bin_path=bin/$(bin_name)
content_folder=~/.fcards
tsv_folder=$(content_folder)/tsv
card_files=$(content_folder)/tsv/*
project_url=github.com/iav0207/fcards
src_installation_path=$$GOPATH/src/$(project_url)
bin_installation_path=$$GOPATH/bin

build_and_play: build play

play:
	$(bin_path) play $(card_files)

install:
	mkdir -p $(tsv_folder)
	go install
	fcards --help

build:
	go build -o $(bin_path)

publish:
ifndef v
	$(error version argument `v` is undefined)
endif
	go fmt
	go test
	git tag $v
	git push origin $v

