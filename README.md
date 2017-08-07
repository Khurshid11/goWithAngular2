# Go with Angular2 #

## Prerequisites ##

* **Client**: Node, npm, angular-cli
* **Server**: Go

## Getting started: ##
* git clone git@github.com:Khurshid11/goWithAngular2.git

* Install latest version of node by following command (If your node version is old):
step 1: sudo apt-get update
step 2: sudo apt-get install build-essential libssl-dev
step 3: curl https://raw.githubusercontent.com/creationix/nvm/v0.33.2/install.sh | bash

Close and reopen your terminal to start using nvm 
	OR
source ~/.profile


step 4: nvm --version

step 5: nvm install node


* make sure you change database to your own or create the same as in the code
* Compile and run project (Angular with Golang) `$ cd client
* $ ng build --env=prod (if ng command is not found, make sure you install angular-cli)  
* $ cd ..
* $ go run main.go

Open `http://localhost:3000/` 😊
