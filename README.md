# Matrix FX

## Configuration for windows build:

- Run this command:
  - mkdir -p build/windows/amd64 build/windows/arm64
    GOOS=windows GOARCH=amd64 go build -o build/windows/amd64/MatrixFx.exe MatrixFX.go
    GOOS=windows GOARCH=arm64 go build -o build/windows/arm64/MatrixFx.exe MatrixFX.go
