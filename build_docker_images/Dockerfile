FROM centos:7
MAINTAINER Yang Sen

COPY ./Centos-7.repo /etc/yum.repos.d/CentOS-Base.repo
COPY ./pip.conf /root/.pip/pip.conf

COPY ./profile /etc/profile
COPY ./PRC /etc/localtime
COPY ./resolv.conf /etc/resolv.conf
COPY ./simsun.ttf /simsun.ttf

RUN yum -y install epel-release && yum clean all && \
    yum -y install python-pip && yum clean all && \
    yum install -y subversion git curl wget net-tools.x86_64 mkfontscale fontconfig zlib zlib-devel pcre pcre-devel openssl openssl-devel openssh-clients openssh-server telnet-server telnet && \
    mkdir -p /usr/share/fonts/chinese/TrueType /root/default/webapps && \
    cd /usr/share/fonts/chinese/TrueType && \
    cp /simsun.ttf . && mkfontscale && mkfontdir && fc-cache -fv && fc-list :lang=zh && \
    rm -rf /var/lib/yum/yumdb/* && \
    echo "Asia/Shanghai" > /etc/timezone && \
    pip install --upgrade pip && pip install tensorflow flask -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com && rm -rf /var/lib/yum/yumdb/*
