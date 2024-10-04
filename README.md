# CLI Todo App
Todo app for console.

### Dependencies:

* Go 1.23.1
* <a href="https://github.com/alexeyco/simpletable">simpletable</a> 1.0.0

### Installation:

1. Clone the project.

2. Open console and execute command:
```commandline
go mod tidy
```
3. Run command below to check if it works correctly:
```commandline
make todo/run
```

4. Run command below to compile app:
```commandline
make todo/build
```

After this you should get `mytodo` file.
```commandline
./mytodo -list
```
You can set binary name:
```commandline
make todo/build BINARY_NAME=custom_todo
```
5. You can add `todo` to your system's `bin` directory:
```commandline
make todo/install
```
Note: I use Fedora, so you `bin` directory may be different. Default path is `/usr/bin`, but you can set it (and binary name):
```commandline
make todo/install BINARY_NAME=custom_todo INSTALL_DIR=/custom/path
```

After that
```commandline
mytodo -list #custom_todo -list
```
6. Delete local binary file (`./mytodo`):
```commandline
make todo/clean
```
Set `BINARY_NAME=custom_todo` if you set custom name.

7. Uninstall from `bin` directory:
```commandline
make todo/uninstall
```
OR
```commandline
make uninstall BINARY_NAME=custom_todo INSTALL_DIR=/custom/path
```
You can ommit `BINARY_NAME` or/and `INSTALL_DIR` if you don`t need.

Note: if you have set `BINARY_NAME` or/and `INSTALL_DIR` you have to write them in next commands you will use.
So, if you have created file `make todo/build BINARY_NAME=other_todo`, and you want to install it globally
You have to write `make todo/install BINARY_NAME=custom_todo`
### Features:
I will use default name `mytodo` for examples.
* Add todo:
```commandline
mytodo -add task for example
```
*List todos:
```commandline
mytodo -list
```
* Complete todo. You need to set id:
```commandline
mytodo -complete=1
```
*Delete todo.You need to set id:
```commandline
mytodo -del=1
```
#### Storage:
Todo list is stored to the default root directory that use for user-specific configuration data (For me it is $HOME/.config).
Relative path is `todo/todos.json`.
