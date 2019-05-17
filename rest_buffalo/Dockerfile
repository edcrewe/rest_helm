# Dont use https://hub.docker.com/r/gobuffalo/buffalo ... since its monolithic with postgres in same docker
# So instead use standard golang image and install buffalo from binary
FROM golang:latest

ENV GOPATH=/go
ENV BP=$GOPATH/src/github.com/gobuffalo/buffalo
RUN go version

RUN curl -sL https://deb.nodesource.com/setup_8.x | bash \
&& apt-get update && apt-get install -y -q build-essential nodejs yarn sqlite3 libsqlite3-dev

# RUN wget https://github.com/gobuffalo/buffalo/releases/download/v0.14.3/buffalo_0.14.3_linux_amd64.tar.gz
# RUN tar -xvzf buffalo_0.14.3_linux_amd64.tar.gz
# RUN mv buffalo /usr/local/bin/buffalo

WORKDIR $GOPATH

RUN go get -u github.com/golang/dep/cmd/dep \
&& go get -tags sqlite -v -u github.com/gobuffalo/pop \
&& go get -tags sqlite -v -u github.com/gobuffalo/buffalo-pop \
&& go get -v -u github.com/gobuffalo/packr/packr \
&& go get -v -u github.com/gobuffalo/packr/v2/packr2 \
&& go get -v -u github.com/markbates/filetest \
&& go get -v -u github.com/markbates/grift \
&& go get -v -u github.com/markbates/refresh \
&& go get -u github.com/alecthomas/gometalinter \
&& gometalinter --install

RUN mkdir -p $GOPATH/src/rest_helm
WORKDIR $GOPATH/src/rest_helm

RUN go get -v -u github.com/gobuffalo/buffalo/buffalo \
&& go get -v -u github.com/gobuffalo/mw-csrf \
&& go get -v -u github.com/gobuffalo/mw-forcessl \
&& go get -v -u github.com/gobuffalo/mw-i18n \
&& go get -v -u github.com/gobuffalo/mw-paramlogger
&& go get -v -u github.com/swaggo/swag/cmd/swag

# this will cache the npm install step, unless package.json changes
ADD package.json .

# RUN cd $GOPATH/src/rest_helm && buffalo setup || :

ADD . .

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 3000

# Uncomment to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app

CMD buffalo dev
