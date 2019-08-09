FROM golang 
COPY . .
EXPOSE 8080
CMD [ "go", "run", "main.go" ]

####Build a new docker IMAGE called mywebsite from the local Dockerfile
# docker build -t mywebsite .

####run (d = detached)
####run (p =  port mappings host:container)
####give it a name (RUNNINGSITE1)
####create from source images MYWEBSITE
# docker run -d -p 80:8080 --name runningsite1 mywebsite
