# va march 6 2021
# Robot Framework and Chrome, running headless in ubuntu
# based on  https://hub.docker.com/r/selenium/standalone-chrome

# https://robotframework.org/SeleniumLibrary/

FROM selenium/standalone-chrome:85.0


MAINTAINER va
LABEL description Robot Framework in Docker.


ENV DEBIAN_FRONTEND=noninteractive 

USER root

# locale
RUN apt-get update && apt-get install -y --no-install-recommends \
    locales \
    locales-all
    
# timezone
RUN apt-get update && apt-get install -y tzdata 
ENV TZ=America/Chicago
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


RUN apt-get update
RUN apt-get install -y python3 python3-pip  gcc wget curl unzip libdbus-glib-1-dev libgtk2.0-0

# python alias python3/pip3
RUN ln -sfn /usr/bin/python3 /usr/bin/python 
RUN ln -sfn /usr/bin/pip3 /usr/bin/pip


RUN pip install --upgrade pip
RUN pip install robotframework
RUN pip install robotframework-seleniumlibrary
RUN pip install robotframework-faker
RUN pip install robotframework-pabot
RUN pip install robotframework-csvlib
RUN pip install requests
RUN pip install robotframework-requests
RUN pip install webdrivermanager
RUN pip install PyYAML 
RUN pip install PyMySQL
RUN pip install mailer
RUN pip install prometheus_client

USER root

RUN apt-get update && apt-get install -y  mysql-client apt-utils curl git dnsutils iputils-ping \
    nginx wget traceroute nmap busybox jq vim net-tools tcpdump netcat iperf openssh-server


RUN mkdir /var/run/sshd

RUN locale-gen en_US.UTF-8
RUN dpkg-reconfigure locales
RUN echo LANG=en_US.UTF-8 >>  /etc/default/locale


# Add Tini
ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

# add robotuser
RUN useradd -d /home/robotuser -G users -m -s /bin/bash robotuser
USER robotuser


RUN /usr/bin/ssh-keygen -A 


# ssh
EXPOSE 22
# webserver
EXPOSE 8000


# tini as init to control child process:
ENTRYPOINT ["/tini", "--", "/docker-entrypoint.sh"]
