* Clone repo:
```
cd ~/go/src && git clone git@github.com:duypx-ltt/sample-go.git sample
```

* Run:
```
cp .env.example .env
cp config/app.env.example config/app.env
cp config/proxy.env.example config/proxy.env
cd ~/go/src/sample
docker-compose up -d
```

* URL: http://localhost
* Test: 
    - http://localhost/hello
    - http://localhost/terms-of-use
