bin_name=fcards
bin_path=bin/$(bin_name)
content_folder=~/.fcards
tsv_folder=$(content_folder)/tsv
tsv_folder_backup=$(content_folder)/tsv.bak
card_files=$(content_folder)/tsv/*
project_url=github.com/iav0207/fcards
src_installation_path=$$GOPATH/src/$(project_url)
bin_installation_path=$$GOPATH/bin

install: fmt
	mkdir -p $(tsv_folder)
	go install
	fcards --help

play:
	fcards play --direc random $(card_files)

build: fmt
	go build -o $(bin_path)

fmt:
	gofmt -s -w .

publish:
ifndef v
	$(error version argument `v` is undefined)
endif
	go fmt
	go test
	git tag $v
	git push origin $v

use_cards_from: # this should be a command of the tool itself
ifndef path
	$(error path argument `path` is undefined)
endif
	rm -r $(tsv_folder_backup) || true
	mv $(tsv_folder) $(tsv_folder_backup)
	ln -sfF $(path) $(tsv_folder)

