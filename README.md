# terarkdb wrapper for Go

CGO_CFLAGS="-I/usr/local/include" CGO_LDFLAGS="-L/usr/local/lib -lterarkdb -lterark-zip-r -lboost_fiber -lboost_context -lstdc++ -lm -lz -lsnappy -llz4 -lzstd -lbz2 -ljemalloc -pthread -lm -lgomp -lrt -ldl -laio" go build 