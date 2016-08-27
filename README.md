#What this is 
A simple little golang:alpine docker image running a Gorilla mux webserver.  I use it for API work when I don't feel like VPNing in to the office or working with live API sources.  I pull the results from what ever API i'll be working with and drop them in the mapped data directory.  Then I configure my code stuff to call localhost:8080.  I normally keep web fetch routines separet from the work objects so when I ready to actually use or test it, I just call whatever working api call function vs the one calling localhost.

Currently the server is only configured to give proper responses to json and html documents but you can modify main.go to make it do your bidding as needed. 


#How to run me

I keep an alias for it in my .bash_rc like the below that maps DOCKER:8080 <-> LOCALHOST:8080 and maps a data directory from my local machine to a mount on the docker container where the webserver can access it. 

```
alias docker-go-ws='docker run -p8080:8080 -d -v /Users/whancock/Dockers/docker-go-ws/data:/go/src/app/data -t  goweb'
```


#How to use me

```
git clone https://github.com/gogosphere/_goweb.git

cd _goweb

docker build -t $tagyourimage .
```
_Results_

```
Sending build context to Docker daemon 14.05 MB
Step 1 : FROM golang:alpine
 ---> 52493611af1e
Step 2 : ADD . /go/src/app
 ---> 2876e29e8037
Removing intermediate container e3baa2385aa7
Step 3 : RUN apk add --update go git
 ---> Running in a33a5a8334a0
fetch http://dl-cdn.alpinelinux.org/alpine/v3.4/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.4/community/x86_64/APKINDEX.tar.gz
(1/6) Installing libssh2 (1.7.0-r0)
(2/6) Installing libcurl (7.50.1-r0)
(3/6) Installing expat (2.1.1-r1)
(4/6) Installing pcre (8.38-r1)
(5/6) Installing git (2.8.3-r0)
(6/6) Installing go (1.6.3-r0)
Executing busybox-1.24.2-r9.trigger
OK: 206 MiB in 18 packages
 ---> 64b6011ec520
Removing intermediate container a33a5a8334a0
Step 4 : RUN cd /go/src/app/ && go get -v
 ---> Running in fcdf123775e5
github.com/gorilla/mux (download)
github.com/gorilla/mux
app
 ---> d985e1c6de68
Removing intermediate container fcdf123775e5
Step 5 : RUN cd /go/src/app/ && go run ./main.go
 ---> Running in a7a1f86bc847
```

After "Step 5" completes, open a new terminal window and run the following.

```
Bills-iMac:docker-go-ws whancock$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
d96434fea9d1        goweb               "/bin/sh -c 'cd /go/s"   3 seconds ago       Up 2 seconds        0.0.0.0:8080->8080/tcp   awesome_fermat
```

 
Take the CONTAINER ID and then commit the container with a human tag which we can use to invoke the container to run in the future.

```
docker commit d96434fea9d1 $humanname
```


At this point you are ready to run the container whenever you wish for a containerized Gorilla mux webserver for faking APIs or just local content with out all the hastle  and configuration laying around.  If you're ocd about your computer you should appreciate this.
