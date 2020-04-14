APPNAME=fan-control
TARGETS=$(shell find . -type f -name '*.go')

usage:
	@echo "USAGE:"
	@echo "  make run"
	@echo "  make build"
	@echo "  make install"
	@echo "  make uninstall"

build: $(TARGETS)
	go build -o bin/$(APPNAME)

run:
	DEBUG=1 go run . -config ./etc/$(APPNAME).yaml

install:
	cp bin/$(APPNAME)      /usr/local/sbin/$(APPNAME)
	cp etc/$(APPNAME).yaml /usr/local/etc/$(APPNAME).yaml
	chmod 755 /usr/local/sbin/$(APPNAME)
	cp dist/systemd/$(APPNAME).service /etc/systemd/system/$(APPNAME).service
	systemctl daemon-reload
	systemctl enable $(APPNAME)
	systemctl start $(APPNAME)

uninstall:
	-systemctl stop $(APPNAME)
	-systemctl disable $(APPNAME)
	-rm -f /etc/systemd/system/$(APPNAME).service
	-rm -f /usr/local/sbin/$(APPNAME)
	-rm -f /usr/local/etc/$(APPNAME).yaml
	systemctl daemon-reload

