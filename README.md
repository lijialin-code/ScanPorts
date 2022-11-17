# ScanPorts



BuildCommand:  CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o ScanPort  main.go

RunCommand:  ./ScanPort -ip="127.0.0.1" -ports="22,81"