
Build

`sudo docker build -t testnginx ./nginx`


Run

`sudo docker run -d -p 80:80 -p 8080:8080 testnginx`

Open 

http://localhost:80
and
http://localhost:8080