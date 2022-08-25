FROM ubuntu:latest

RUN apt update && apt install -y git wget

RUN git clone https://github.com/vmware-tanzu/k-bench.git

RUN cd k-bench && sh install.sh