FROM ubuntu:22.04
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone
RUN apt-get update -y && \
    apt-get install -y wget build-essential libssl-dev gdb python3-pip git && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
RUN wget https://github.com/Kitware/CMake/releases/download/v3.28.2/cmake-3.28.2.tar.gz && \
    tar -xzvf cmake-3.28.2.tar.gz && \
    cd cmake-3.28.2 && \
    ./bootstrap && \
    make -j$(nproc) && \
    make install && \
    cd .. && \
    rm -rf cmake-3.28.2 && \
    rm cmake-3.28.2.tar.gz
RUN mkdir -p /home/workdir
WORKDIR /home/workdir
COPY ./requirements.txt /home/workdir/requirements.txt
RUN pip3 install -r /home/workdir/requirements.txt
WORKDIR /home/workdir