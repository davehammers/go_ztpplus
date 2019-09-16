export GOOS=linux
mips: GOARCH=mips
onie: GOARCH=386

all: mips onie

mips onie:
	go fmt
	export GOOS=$(GOOS) GOARCH=$(GOARCH) && go build -o clissrver_$@
include ../../../../rules.mk
