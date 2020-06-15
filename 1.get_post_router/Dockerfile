FROM alpine

WORKDIR $GOPATH/src/Snai.Gin/1.get_post_router
COPY get_post_router /
WORKDIR /

RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && echo 'Asia/Shanghai' >/etc/timezone

EXPOSE 8080

ENTRYPOINT ["./get_post_router"]

# docker 常用命令
# docker build -t get_post_router:v1.0 .
# docker run -d -p 8080:8080 --name get_post_router get_post_router:v1.0
# docker ps -a
# docker stop
# docker rm
# docker rmi
# docker exec -it CONTAINER_ID sh
# ifconfig eth0
# date
