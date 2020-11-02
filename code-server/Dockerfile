FROM codercom/code-server:latest

RUN curl -fsSL https://golang.org/dl/go1.15.2.linux-amd64.tar.gz | sudo tar -C /usr/local -xzf -
# RUN curl -fsSL https://golang.org/dl/go1.15.2.linux-amd64.tar.gz > go.tar.gz \
# 	&& tar -xvf go.tar.gz \
# 	&& sudo mv go /usr/local

ENV GOROOT=/usr/local/go
ENV GOPATH=/home/coder/gocode
ENV PATH="${GOPATH}/bin:${GOROOT}/bin:${PATH}"

