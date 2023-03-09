
build_all:
	go install github.com/mitchellh/gox@latest
	cd bin/ && gox -os="linux darwin" -arch="amd64 arm64" ../
