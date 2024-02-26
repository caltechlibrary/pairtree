
# Installation

*pairtree* is a command line program run from a shell like Bash. It allows you to organize encode
"Name as Text" documents to provide metadata for a directory.

Quick install with curl
-----------------------

If you are running macOS or Linux you can install released versions of newt
with the following curl command.

~~~
curl https://caltechlibrary.github.io/pairtree/installer.sh
~~~

## Compiled version

Compiled versions are available for Mac OS X (amd64 processor, macosx-amd64), Linux (amd64 process, linux-amd64), 
Windows (amd64 processor, windows-amd64) and Rapsberry Pi (arm7 processor, raspberry_pi_os-arm7)

VERSION_NUMBER is a [symantic version number](http://semver.org/) (e.g. v0.1.2)


For all the released version go to the project page on Github and click latest release

>    https://github.com/caltechlibrary/pairtree/releases/latest


| Platform    | Zip Filename                             |
|-------------|------------------------------------------|
| Windows     | pairtree-VERSION_NUMBER-windows-amd64.zip |
| Mac OS (Intel) | pairtree-VERSION_NUMBER-macos-amd64.zip  |
| Mac OS (M1) | pairtree-VERSION_NUMBER-macos-amd64.zip  |
| Linux/Intel | pairtree-VERSION_NUMBER-linux-amd64.zip   |
| Raspbery Pi | pairtree-VERSION_NUMBER-raspberry_pi_os-arm7.zip |


## The basic recipe

+ Find the Zip file listed matching the architecture you're running and download it
    + (e.g. if you're on a Windows 10 laptop/Surface with a amd64 style CPU you'd choose the Zip file with "windows-amd64" in the name).
+ Download the zip file and unzip the file.
+ Copy the contents of the folder named "bin" to a folder that is in your path 
    + (e.g. "$HOME/bin" is common).
+ Adjust your PATH if needed
    + (e.g. `export PATH="$HOME/bin:$PATH"`)
+ Test


### Mac OS X

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Make sure the new location in in our path
5. Test

Here's an example of the commands run in the Terminal App after downloading the 
zip file.

```shell
    cd Downloads/
    unzip pairtree-*-macosx-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    pairtree -version
```

### Windows

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell on Windows 10 after
downloading the zip file.

```shell
    cd Downloads/
    unzip pairtree-*-windows-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    pairtree -version
```


### Linux 

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip pairtree-*-linux-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    pairtree -version
```


### Raspberry Pi

Released version is for a Raspberry Pi 2 or later use (i.e. requires ARM 7 support).

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip pairtree-*-raspberry_pi_os-arm7.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    pairtree -version
```


## Compiling from source

_pairtree_ is "go gettable".  Use the "go get" command to download the dependant packages
as well as _pairtree_'s source code. 


```shell
    go get -u github.com/caltechlibrary/pairtree/...
```

Or clone the repstory and then compile

```shell
    cd
    git clone https://github.com/caltechlibrary/pairtree src/github.com/caltechlibrary/pairtree
    cd src/github.com/caltechlibrary/pairtree
    make
    make test
    make install
```

Compilation assumes [go](https://github.com/golang/go) v1.10 and [bleve](https://blevesearh.com) v0.6.0.

