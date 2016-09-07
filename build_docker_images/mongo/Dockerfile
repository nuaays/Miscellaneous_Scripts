FROM ubuntu:latest
MAINTAINER Yang Sen <nuaays@gmail.com>

RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10
RUN echo "deb http://repo.mongodb.org/apt/ubuntu "$(lsb_release -sc)"/mongodb-org/3.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-3.0.list

#RUN apt-get update && apt-get install -y mongodb-org
RUN apt-get update && \
    apt-get install -y mongodb-org=3.0.1 \
                       mongodb-org-server=3.0.1 \
                       mongodb-org-shell=3.0.1 \
                       mongodb-org-mongos=3.0.1 \
                       mongodb-org-tools=3.0.1 

RUN apt-get autoclean && apt-get -y autoremove
RUN mkdir -p /data/db

EXPOSE 27017

ENTRYPOINT ["/usr/bin/mongod"] 

