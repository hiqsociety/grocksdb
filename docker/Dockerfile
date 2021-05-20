FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive

RUN apt update -y && apt upgrade -y --no-install-recommends && \
    apt install -y --no-install-recommends ca-certificates gcc g++ git pkg-config autoconf make cmake automake libssl-dev libaio-dev wget

# setup cxx standard
ENV CFLAGS='-fPIC -O3 -pipe'
ENV CXXFLAGS='-fPIC -O3 -pipe'

RUN cd /tmp && git clone https://github.com/bytedance/terarkdb.git && cd terarkdb && git submodule update --init --recursive

RUN set -e; cd /tmp/terarkdb && \
    ./build.sh && ./build.sh && \
    mv output/lib/* /usr/local/lib/ && \
    mv output/include/* /usr/local/include/ && \
    mv output/third-party/terark-zip/3rdparty/boost-include/lib/* /usr/local/lib/ && \
    cp -R include/rocksdb /usr/local/include/ && \
    cd /tmp && rm -rf * && ldconfig

# install go
ENV GOLANG_PACKAGE go1.16.4.linux-amd64.tar.gz

RUN wget https://dl.google.com/go/${GOLANG_PACKAGE} && \
    tar -C /usr/local -xzf ${GOLANG_PACKAGE} && rm ${GOLANG_PACKAGE}

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# setup pkg-config path
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig:/usr/local/lib64/pkgconfig
