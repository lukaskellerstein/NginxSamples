
If you want, you can generate your own certicate. Move yourself into ssl directory and run this command 

`sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./nginx.key -out ./nginx.crt`


Build

`sudo docker build -t testnginx ./nginx`


Run

`sudo docker run -d -p 80:80 -p 443:443 testnginx`

Open 

https://localhost