
## 環境

* AmazonLinux2023
* Nginx 1.24.0
* Golang 1.23.2


## ローカル

```bash
$ docker-compose up
```

http://localhost:8080/ にアクセス



## サーバー

### build

```bash
$ cd /path/to/helloworld-prod-go/app
$ go build main.go
```

### watch

```bash
$ go install github.com/air-verse/air@latest
$ cd /path/to/helloworld-prod-go/app
$ air -c .air.toml
```


### Supervisord

#### ログ

```bash
$ cd /var/log
$ sudo mkdir supervisor
```


#### Supervisord 起動

```bash
# systemd で管理している場合
$ sudo systemctl start supervisor
$ sudo systemctl enable supervisor

# systemd で管理していない場合
$ sudo supervisord -c /etc/supervisord.conf
```


#### Supervisord 停止

```bash
# systemd で管理している場合
$ sudo systemctl stop supervisor

# systemd で管理していない場合
$ sudo supervisorctl shutdown
```



#### プロダクト起動

```bash
$ sudo supervisorctl start helloworld-prod-go
```


#### プロダクト停止

```bash
$ sudo supervisorctl stop helloworld-prod-go
```

