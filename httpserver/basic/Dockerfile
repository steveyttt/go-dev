FROM golang 
COPY . .
# ENTRYPOINT ["./main"]
EXPOSE 8080
# CMD [ "bash" ]
CMD [ "go", "run", "main.go" ]

####Build a new docker image from this file called test1
# docker build -t test1 .

####run (i = interactive) (t = assign a tty ) (d = detached)
####give it a name (bashtest1)
####create from source images test1
# Docker run -itd --name bashtest2 test1

####
#docker exec -it bashtest2 bash

####docker run
# docker run -d -p 80:8080 --name lalala2 test22