## Basic Introduction
### 1.1
You can use this project to control the cron job with the cron expression.

The front_end project is [FrontEndProject](https://github.com/lsk569937453/rust_frontend)

The test   of the whole project  is  [testProject](http://lskyy.top/admin).

## Get started

### 2.1 server
```
git clone git@github.com:lsk569937453/go_backend.git

cd go_backend

go build
```

### 2.2 front
```
git clone https://github.com/lsk569937453/rust_frontend

cd rust_frontend

npm install

npm run build
```

### 2.3 copy the dist file to server
```
cd go_backend

mkdir resource
```

copy the dist file to the resource directory.

### 2.4 start the server
```
cd go_backend

./go_backend
```

Then you can could browse the url [http://localhost:9393/admin](http://localhost:9393/admin)

