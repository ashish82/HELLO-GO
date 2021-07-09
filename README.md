# HELLO-GO
Go Rest API skelton
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
# How to install Go
1- download the go software from below link - https://golang.org/
# Go version 1.16.5 
https://golang.org/dl/go1.16.5.darwin-amd64.pkg


3- I am providing you the steps to install GO and setup on local machine
 
 3.1- After install this package
  go1.16.5.darwin-amd64.pkg 

 3.2- open the terminal and put the below command to check the go version

 ~ go version

 below command will give you list of enviornment variable, we need get the `GOPATH` value only 

 ~ go env  

 2.3 go to that folder (GOPATH="/Users/`username`/go")

 cd /Users/`username`/go
 
 create below three directory in the same directory

 ~ mkdir src

 ~ mkdir pkg

 ~ mkdir bin

 3- Clone/Copy the project inside the `src` directory 

git clone git@github.com:ashish82/HELLO-GO.git

4- go to HELLO-GO directory and build the project

~ go build

5- run the program
~ ./HELLO-GO

6- below is the api 
`Health-API`

curl --location --request GET 'http://localhost:9000/hello-go/health'

`COMMENT-API`

curl --location --request GET 'http://localhost:9000/hello-go/allComments'


6- Download the <b> visual studio</b> code from below link as per os- https://code.visualstudio.com/

# Install plugin in visual studio for go lang
Installing 10 tools at 
/Users/systemuser/go/bin in module mode.

  gopkgs <br>
  go-outline <br>
  gotests <br>
  gomodifytags <br>
  impl <br>
  goplay <br>
  dlv <br>
  dlv-dap <br>
  staticcheck <br>
  gopls <br>
