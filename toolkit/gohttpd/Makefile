TARGET := gohttpd

.PHONY  : all clean clean rebuild install

all     : $(TARGET)

build   : all

rebuild : clean all

test    :

clean   :
	rm -fv $(TARGET)

install :
	install -d -m0755 /usr/local/gohttpd/html
	install -m0755 ./gohttpd /usr/local/gohttpd/gohttpd
	install -m0600 ./gohttpd.conf /usr/local/gohttpd/gohttpd.conf
	install -m0755 ./gohttpd.service /usr/lib/systemd/system/gohttpd.service
	install -m0644 ./html/favicon.ico /usr/local/gohttpd/html/favicon.ico
	install -m0644 ./html/index.html /usr/local/gohttpd/html/index.html

$(TARGET) :
	go build -o ./$(TARGET) -ldflags "-w -s" $(TARGET)
