###Commit a Project to Git
 
 Git is a distributed version control system.Git is used from the configuration management.
 In this project commit to git, the go file is created in the local repository and commit to the remote repository.  
 
 ##Requirements
 
 Linux centos Operating system 
 Linux Centos Operating system is used as the local computer with the working directory.
 Github.com
 Created a remote repository in the github account(ND0709/Network-Interfaces)
 Go language Programming
 Go file are created in the local repository to list the information about network interfaces.
 
 ##Initial configurations
 How to setup
 Linux sytem
 
 * Check whether we have the updated versions of the following softwares using the terminal
 $ git version
 git version 2.27.0
 $ go version
 go version go1.15.14 linux/amd64
 * check the present working directory whether it is in the user directory or not.
 $pwd
 /home/deepika
 * If not change it to user directory 
 $cd /home/user
  
 Github
 * A new repository for this project with the public privacy settings and README file is created.
 * This repository is a remote repository for this project.
 * clone to the local directory with

back to terminal
* cloning the repository in the local repository
$ git clone https://github.com/ND0709/Network-Interfaces.git
enter login id and password to get the accessibility into remote respository
* create the go file for the project in the repository directory
$ vim Ip.go
* Insert the go program to list the informationabout the network interfaces
* check whether the program is executable or not. If not make adequate changes. 
$ go run IP.go
* Adding the go file to the repository
$git add IP.go
* commiting the go file to the repository
$git commit -m IP.go
 the commit command will allow commit the change in the local repository.
 to commit changes to the remote repository, push command is used.
 $git push
 
##Conculsion
The go file is create and committed into the github.
