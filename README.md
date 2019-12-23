# Setting and Running

## First

[SettingGOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Windows

default go path `%USERPROFILE%\go` ext. `C:\Users\username\go`

Your workspace can be located wherever you like,
but we'll use `C:\go-work` in this example.

**NOTE:** `GOPATH` must not be the same path as your Go installation.

- Create folder at `C:\go-work`.
- Right click on "Start" and click on "Control Panel". Select "System and Security", then click on "System".
- From the menu on the left, select the "Advanced systems settings".
- Click the "Environment Variables" button at the bottom.
- Click "New" from the "User variables" section.
- Type `GOPATH` into the "Variable name" field.
- Type `C:\go-work` into the "Variable value" field.
- Click OK.

## Second

Type

```bash
go install main
```

And then

```bash
./bin/main
```

Or type

```bash
go run main
```
