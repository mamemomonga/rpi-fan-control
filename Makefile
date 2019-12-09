APPNAME=fan-control

usage:
	@echo "USAGE:"
	@echo "  make install"
	@echo "  make uninstall"

install:
	cp $(APPNAME).sh /usr/local/sbin/$(APPNAME).sh
	chmod 755 /usr/local/sbin/$(APPNAME).sh
	cp dist/systemd/$(APPNAME).service /etc/systemd/system/$(APPNAME).service
	systemctl daemon-reload
	systemctl enable $(APPNAME)
	systemctl start $(APPNAME)

uninstall:
	-systemctl stop $(APPNAME)
	-systemctl disable $(APPNAME)
	-rm -f /etc/systemd/system/$(APPNAME).service
	-rm -f /usr/local/sbin/$(APPNAME).sh
	systemctl daemon-reload

