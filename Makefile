PREFIX = /usr/local
all: jmp

jmp: cmd/jmp/main.go jumpdb/jumpdb.pb.go jumpdb/DB.go
	go build -o $@ $<

jumpdb/jumpdb.pb.go: jumpdb/jumpdb.proto
	protoc -I=./jumpdb --go_out=./jumpdb $<

.PHONY: completion
completion:
	mkdir -p ~/.jmp/functions
	cp -f ./completions/jmp.* ~/.jmp
	cp -f ./completions/_j ~/.jmp/functions
	@echo "Completions installed to ~/.jmp/"
	@echo "You want to source the according file"

install: jmp
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f jmp $(DESTDIR)$(PREFIX)/bin
	@echo "For completions, you want to run 'make completion' as regular user"
	@echo "Feel free to mv the locally installed completions to a system path (e.g. /etc/bash_completion.d/)"

.PHONY: test
test:
	cd jumpdb && go test
