FROM centos:centos6
MAINTAINER Ewerton Assis <earaujoassis@gmail.com>

LABEL Description="This image is used to start the space microservice" Version="0.1"
RUN rpm -Uvh http://download.fedoraproject.org/pub/epel/6/i386/epel-release-6-8.noarch.rpm
RUN yum install -y golang postgresql-devel openssl-devel
