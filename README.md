#Introduction

This is a client configuration for cdash. It sets up a machine to talk to CDash to allow scheduled builds.

#Setup

1. Setup cdash user.
``` bash
sudo su
add user cdash
su cdash
```

1. Clone the client repository. 
``` bash
git clone git://github.com/arktools/cdash-client.git ~/client
```

1. Set the properties in machine.cdash.xml.
1. Set the properties in machine.ctest if different.
1. Configure crontab to automatically handle jobs from cdash.
``` bash
crontab -e
```
Example crontab
``` bash
* * * * * /home/cdash/client/go
```
