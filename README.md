# cpath.go
cpath.go is very simple app written by Go. It use to change value of environment variable.

## Installation
Make sure you have the a working Go environment (go 1.1 is *required*). [See the install instructions](http://golang.org/doc/install.html). And cpath.go depends on cli.go. So before install cpath.go, you must install cli.go.

To install cli.go, simply run:
    
    $ go get github.com/codegangsta/cli

And install cpath.go, simply run:

    $ go get github.com/htoooth/cpath
    
## Getting Started
Currently cpath.go support **windows7 only** because I'm working at windows7 and use some windows commands, such as SETX. Make sure your PATH includes to the `$GOPATH/bin` directory so your commands can be easily used:

Here are some help.
> $ cpath

    NAME:
       cpath - Change current user's environment variable PATH.
    
    USAGE:
       cpath [global options] command [command options] [arguments...]
    
    VERSION:
       1.0.0
    
    COMMANDS:
       list, l   List current PATH
       append, a Append new path in PATH
       search, s Search key word in PATH
       delete, d Delete count number
       create, c Create name value
       help, h   Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
       --versionprint the version
       --help, -h   show help



> $ cpath list

    0.  c:\mygo\bin
    1.  C:\Ruby200\bin
    2.  C:\Program Files\Haskell\bin
    3.  C:\Program Files\Haskell Platform\2012.4.0.0\lib\extralibs\bin
    4.  C:\Program Files\Haskell Platform\2012.4.0.0\bin
    5.  C:\Python27\Lib\site-packages\PyQt4
    6.  C:\Program Files\Common Files\Microsoft Shared\Windows Live
    7.  C:\Windows\system32
    8.  C:\Windows
    9.  C:\Windows\System32\Wbem
    10. C:\Windows\System32\WindowsPowerShell\v1.0\

> $ cpath append your_path

> $ cpath search key_word

> $ cpath delete list_number

**Note**: There are some commands like *append, search, delete* make effect in your PATH variable.

> $ cpath create variable\_name variable\_value

**Note**: The *create* command create **new variable** in your system environment.

## About
cpath.go is written by htoo and thanks to cli.go by [Code Gangsta](http://codegangsta.io).

