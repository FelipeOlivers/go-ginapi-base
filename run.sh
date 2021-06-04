sudo docker build . -t go-ginapi
sudo docker run -i -t -d -p 8000:8000 go-ginapi