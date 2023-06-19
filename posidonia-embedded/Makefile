CC=clang
CFLAGS=-g -Wall

BIN=posidonia

all: $(BIN)

%: %.c
	$(CC) $(CFLAGS) $< -o $@ -lcurl

clean:
	rm -rf $(BIN) *.dSYM
