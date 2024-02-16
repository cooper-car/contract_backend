FROM golang:alpine

WORKDIR /app

RUN apk update && apk add bash

# 複製所有代碼到容器中
COPY . .

# 設置文件權限
RUN chmod +x /app/wait-for-it.sh

# 下載並整理模塊
RUN go mod download && go mod tidy

# 構建 Go 應用程序
RUN go build -o . cmd/main.go

# 開始時運行 wait-for-it.sh 並啟動應用程序
CMD ["./wait-for-it.sh", "mysql:3306", "--", "./main"]