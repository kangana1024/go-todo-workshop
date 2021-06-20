# Install Golang

## Windows

- Download Git [click](https://golang.org/doc/install) And Install

- ตั้งค่า Variables

Step 1 :

![Search System Properties](https://github.com/kangana1024/go-todo-workshop/blob/main/assets/s1.png?raw=true "Search System Properties")

Step 2 :

เปิด `gitbash` พิมพ์คำสั่ง mkdir `%USERPROFILE%\go`

สร้าง folder `pkg,bin,src` ใว้ใน c:\USERNAME\go หรือ เข้าที่ CMD ด้วยคำสั่ง `cd %USERPROFILE%\go` ก่อนแล้ว `mkdir` สร้างทีละอัน เราจะเก็บ code workshop ไว้ที่ `src`

Step 3 :

Varliable name : `GOPATH`

Varliable value : `%USERPROFILE%\go`
![Edit GOPATH](https://github.com/kangana1024/go-todo-workshop/blob/main/assets/s2.png?raw=true "Edit GOPATH")

## Mac OR Linux

- เข้าไปที่ [Golang Download](https://golang.org/dl/) โหลด pkg

### MAC

```bash
cd /tmp
curl https://golang.org/dl/go1.16.5.darwin-amd64.pkg
sudo open ./go1.16.5.darwin-amd64.pkg
```

### Linux

สร้าง path เก็บ project ไว้ที่ `~/go` ด้วยคำสั่ง `mkdir $HOME/go`

```bash
sudo apt update
sudo apt upgrade

cd /tmp
curl -O -L https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
tar -xf go1.16.5.linux-amd64.tar.gz
sudo chown -R root:root ./go
sudo mv -v go /usr/local
vim ~/.bash_profile

```

มาบรรทัดล่างสุด กดปุ่ม o 1 ที

```bash
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

กด `:wq` ออกมาเพื่อ save

จากนั้นพิมพ์คำสั่ง `source ~/.bash_profile` เพื่อให้ระบบรู้จัก go
