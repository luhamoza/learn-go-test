build:
	@go build -o happy-fun-book
run:build
	@./happy-fun-book
test:
	@go test -v ./... -count=1
git:
	@rm happy-fun-book  && git add . && git commit -m "feat: building the app" && git push