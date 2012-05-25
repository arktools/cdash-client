#Introduction

This is a client configuration for cdash. It sets up a machine to talk to CDash to allow scheduled builds.

#Setup
``` bash
sudo su
add user cdash
su cdash
git clone git://github.com/arktools/cdash-client.git ~/client
crontab -e
```
Example crontab
``` bash
* * * * * /home/cdash/client/go
```
