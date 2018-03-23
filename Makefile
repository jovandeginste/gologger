BINARY=gologger
PACKAGE=$(BINARY)
VENDOR="KU Leuven ICTS Linux team <linuxicts@kuleuven.be>"

$(BINARY).%.amd64: GOARCH=amd64
$(BINARY).%.amd64: ARCH=amd64
$(BINARY).%.i686: GOARCH=386
$(BINARY).%.i686: ARCH=i686

$(BINARY).bin.%:
	GOARCH=$(GOARCH) go build -ldflags="-s -w" -o $(BINARY).bin.$(ARCH)

$(BINARY).rpm.%: $(BINARY).bin.%
	fpm --vendor $(VENDOR) -s dir -t rpm -a $(ARCH) -n $(PACKAGE) ./$(BINARY).bin.$(ARCH)=/usr/bin/$(BINARY)

clean: clean-rpm clean-bin

clean-rpm:
	rm -vf *.rpm

clean-bin:
	rm -vf $(BINARY).bin.*

copy:
	rsync -av *.rpm kul-repo@yum-repo-upload.icts.kuleuven.be:repo/

all: bin rpm
bin: $(BINARY).bin.amd64 $(BINARY).bin.i686
rpm: $(BINARY).rpm.amd64 $(BINARY).rpm.i686
