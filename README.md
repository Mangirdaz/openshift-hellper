# Openmshift Hellper
Small utility to help organise CI&CD flow of openshift objects

### Build
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

### Run
    ./openshift-hellper <flags>
```
NAME:
   Openshift Hellper - A new cli application

USAGE:
   openshift-hellper [global options] command [command options] [arguments...]
   
VERSION:
   0.0.1
   
AUTHOR(S):
   Mangirdas Judeikis <info@judeikis.lt> 
   
COMMANDS:
     imagestream, is  modify ImageStream with new version of images
     help, h          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```
### ImageStreams flows step by step

    (Openshift) Build Docker image 
    (Openshift) Test it
    (Jenkins) Tag with todays date 
    (Jenkins) clone GIT project (jenkins) with IS deffinition  
    (Openshift Hellper) Update IS file with new image 
    (Jenkins) Create new Breanch, and push file
    (Jenkins) Create Pull request
    (Git) Pull request approved
    (Ansible) Distribute IS

```
#Example for Openshift hellper
/openshift-hellper is edit --file=example.json --name=python27 --latest=true
```