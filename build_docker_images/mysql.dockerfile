FROM centos:centos6
MAINTAINER Yang Sen <yangsen@zhongan.com>

#docker pull nuaays/mysqlserver:latest
#docker run -d -P nuaays/mysqlserver:latest

RUN yum install -y mysql-server mysql

RUN /etc/init.d/mysqld start && \
    mysql -e "grant all privileges on *.* to 'root'@'%' identified by 'test';"&& \
    mysql -e "grant all privileges on *.* to 'root'@'localhost' identified by 'test';"&& \
    mysql -u root -ptest -e "show databases;" 

EXPOSE 3306

CMD ["/usr/bin/mysqld_safe"]

