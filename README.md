# kvdb
A key - value DB implementation

Build the program

Linux
```
make build
```

Mac
```
make build_mac
```

Tests
```
make test
```

Run the program
```
./kvdb
```

For now it supports the following commands.
```
SET key1 value1
GET key1
DEL key1
```
```aidl
SET counter 0
INCR counter
GET counter
INCRBY counter 10
```
The first argument is case-sensitive. Pass the argument as required by the command, there is no check in number of 
arguments and may raise a panic.

Use `exit` to terminate the program.

TODO:
- Currently only one client can connect. It should move to a centralized server which needs to have thread safe write/update operations.
