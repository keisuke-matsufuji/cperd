FROM ubuntu:20.04

RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

# Goのインストール
RUN apt-get update \
    && apt-get install -y golang-go \
    && apt-get update \
    && apt-get install -y git

RUN apt install -y -qq curl python3 jq

# Install reviewdog
RUN curl -sfL https://raw.githubusercontent.com/reviewdog/reviewdog/master/install.sh | sh -s

COPY main.go main.go
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
