docker stop simpligo-pln && docker rm simpligo-pln && docker run -d --name simpligo-pln -p 80:8080 --add-host elasticsearch:172.17.0.1 --restart always simpligo-pln:$1