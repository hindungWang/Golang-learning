docker build -t goserver .

docker run -d -v //d//tmp:/tmp --name mygo -p 9090:9090 -p 80:80 -p 53:53 -p 8800:8800 goserver

docker ps

sleep 3

curl 192.168.99.100:9090/do

docker logs mygo