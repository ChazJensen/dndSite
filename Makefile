export GOBIN=/home/chimpgimp/git/dndSite/

cycle: run clean
	@echo Done

run: install ./server
	@./server

install:
	@go install

clean: ./server
	@rm server
