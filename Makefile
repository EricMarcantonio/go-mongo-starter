filename = "starter"
windows:
	GOOS=windows GOARCH=386 go build -o bin/$(filename)_windows.exe ./src

linux:
	GOOS=linux GOARCH=386 go build -o bin/$(filename)_linux ./src

macos:
	GOOS=darwin GOARCH=386 go build -o bin/$(filename)_macos ./src
