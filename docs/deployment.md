# Deployment

Automatic deployment using Dokku and GitHub Actions

## Create VPS

- Create a VPS using any preferred service
    - Choose Ubuntu 20.04 LTS
    - Choose a hostname with the domain name that you want to use for your dokku apps

## Setup Dokku

We are going to install Dokku v0.24.10

### On the VPS

- Install Dokku by using the following commands `wget https://raw.githubusercontent.com/dokku/dokku/v0.24.10/bootstrap.sh;
sudo DOKKU_TAG=v0.24.10 bash bootstrap.sh`

- Change to the Dokku user `su dokku`
    - If dokku was able to resolve the hostname for the VPS then a global domain will have already been added
        - To check this use the command `dokku domains:report --global` and `dokku domains:add-global <domain>`

### Set Up SSH Keys and Virtualhost Settings

On a web browser:
- Navigate to the ip or domain name of the dokku server
    - Add your desired SSH public keys for remote access
    - Tick "Use virtualhost naming for apps" as we may have multiple dokku apps

## Quick Security

- Now that you have added your SSH public key you may disable password access
    - Open `/etc/ssh/sshd_config in a text editor and change the line `#PasswordAuthentication yes` to `PasswordAuthentication no`

## Deploy Dokku

### On the VPS

- Create an app using `dokku apps:create salon-booking-guru`

Now we want to set the deploy branch on Dokku.
The default deploy branch is `master` but our repository uses `main`.
- Change it by using the command `dokku git:set salon-booking-guru deploy-branch main`

## Dockerfile configuration

For this version of dokku (v0.24.10) Dockerfile deployment is only recognised when there is a dockerfile in the root directory of the repository.
- The api Dockerfile in this repository is found under `api/`
    - A current workaround means that there is an empty Dockerfile in the root

## Dokku Dockerfile Path

Now that dokku recognises that we want to deploy using Dockerimage we want to now give it the correct path to the api Dockerfile

### On the VPS

- Run this command to add the correct Dockerfile path for the build phase of deployment `dokku docker-options:add salon-booking-guru build --file=/home/dokku/salon-booking-guru/api/Dockerfile`

- We must also manually add the Dockerfile (outside of version control) to the bare repo within dokku, so that it can build it.
    - Use `mkdir -p /home/dokku/salon-booking-guru/api/ && cd /home/dokku/salon-booking-guru/api/ && wget https://raw.githubusercontent.com/KarmaComputing/salon-booking-guru/main/api/Dockerfile`


### Quick Check

**Note:** Still yet to figure out how to set build context in dokku so there will be Go module errors

To make sure everything is in order, you can manually push to Dokku.
- `git clone` the repository on your local machine
    - Add the dokku remote: `git remote add dokku dokku@<ip/domain>:salon-booking-guru`
    - Now `git push dokku main`

## Set Up GitHub Actions

- Create a yaml file within the path `.github/workflows/api-deploy.yml`

### In the file:

- Specify the GitHub event that triggers the workflow
    - Use the key ```on:``` to specify the event, we want push
        - From here you can add keys to specify the branch, we want main

**Note:** `workflow_dispatch` will allow you to rerun GitHub actions on demand.


It will look like this:
```
name: Deploy API
on:
  workflow_dispatch:
  push:
    branches:
    - main
```

### Now its time to create a job for the workflow run:

- Define a job ```jobs:```
    - Give the job a unique name by using a key
    - Specify the runner environment, we are using ubuntu-20.04

### Now create the steps for the job:

- Use the ```actions/checkout@v2``` to fetch the code from the main branch that we specified
    - You should specify ```fetch-depth: 0``` to fetch all history

- Now make the first job! You must give the job an id and optionally a name
- We are using ```idoberko2/dokku-deploy-github-action@v1``` for this job
    - Within the ```with:``` key we specify all the sensitive details for the dokku server, store them in the repo secrets
        - For this deployment we specify ```SSH-PRIVATE-KEY```, ```APP-NAME``` and ```DOKKU-HOST```
        - Then specify the ```REMOTE-BRANCH```, as written before we want main

This will look like:
```
jobs:
  deploy:
    runs-on: ubuntu-20.04
    steps:
    - uses: actions/checkout@v2
      with:
        fetch-depth: 0
    - id: deploy
      name: Deploy to Dokku
      uses: idoberko2/dokku-deploy-github-action@v1
      with:
        ssh-private-key: ${{ secrets.SSH_PRIVATE_KEY }}
        dokku-host: ${{ secrets.DOKKU_HOST }}
        app-name: ${{ secrets.DOKKU_APP_NAME }}
        remote-branch: main
```

