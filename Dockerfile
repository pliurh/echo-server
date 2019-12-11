FROM centos/tools
# RUN yum install epel-release -y
# RUN yum install golang -y && go build
COPY echo-server /usr/bin
CMD [ "/usr/bin/echo-server" ]
