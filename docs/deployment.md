# Deployment

Automatic deployment using Dokku and GitHub Actions

## Deploy a Dokku server

Requirements:

- Ubuntu 20.04 LTS
- [Dokku v0.24.10](https://dokku.com/)
- A domain name with a [DNS A record](https://en.wikipedia.org/wiki/List_of_DNS_record_types) pointing to your dokku server (optional)

## Dokku Installation

On the VPS install Dokku:

```bash
wget https://raw.githubusercontent.com/dokku/dokku/v0.24.10/bootstrap.sh;
sudo DOKKU_TAG=v0.24.10 bash bootstrap.sh
```

## Dokku configuration
Setup domain (optional)

> If dokku was able to resolve the hostname during install then the dokku domain will have already been added
  check with: `dokku domains:report --global`

If not, check your DNS A record is correctly pointing to your server, then configure dokku manually:

```bash
su dokku
dokku domains:add-global <your-domain>
```

#### Dokku SSH keys and Virtualhost Settings

Complete the Dokku installation from your web browser.

- Navigate to the ip or domain name of your server
- Add your desired SSH public keys for remote access
- Tick "Use virtualhost naming for apps" as we may have multiple dokku apps

#### Disable Password based authentication
- Verify your ssh access working with key based access (`ssh root@<your-ip>`) # login without a password
- Disable ssh password based authentication 
  - Open `/etc/ssh/sshd_config` 
  - Change the line `#PasswordAuthentication yes` to `PasswordAuthentication no`

## Deploy api & and front-end

### Create dokku apps front-end & api

```bash
dokku apps:create salon-booking-guru-front-end;
dokku apps:create salon-booking-guru-api;
dokku apps:create salon-booking-guru-database;
```
### Set deploy branch to `main`: 

```bash
dokku git:set salon-booking-guru-front-end deploy-branch main;
dokku git:set salon-booking-guru-api deploy-branch main;
dokku git:set salon-booking-guru-database deploy-branch main;
```

## Dockerfile configuration

For dokku v0.24.10 Dockerfile deployment is only recognised when there is a `Dockerfile` in the root directory of the repository. However this repo has components in subdirectories:

- [api/Dockerfile](https://github.com/KarmaComputing/salon-booking-guru/blob/main/api/Dockerfile)
- [front-end/Dockerfile](https://github.com/KarmaComputing/salon-booking-guru/blob/main/front-end/Dockerfile)
- [database/Dockerfile](https://github.com/KarmaComputing/salon-booking-guru/blob/main/database/Dockerfile)

A current workaround means that there is an empty Dockerfile in the root. This forces dokku to treat the apps as docker 
deployments (Dokku supports multiple deployment types).

### Set Dokku Dockerfile path

Configure api app `Dockerfile` location
```bash
dokku docker-options:add salon-booking-guru-api build --file=/home/dokku/salon-booking-guru-api/Dockerfile;
dokku docker-options:add salon-booking-guru-front-end build --file=/home/dokku/salon-booking-guru/front-end/Dockerfile;
dokku docker-options:add salon-booking-guru-database build --file=/home/dokku/salon-booking-guru-database/Dockerfile;
```

> If you make a mistake, you can check  the docker-options. `dokku docker-options:report salon-booking-guru-front-end`
  If you need to reset them:

>  1. take a copy of *all* existing settings
   `dokku docker-options:report salon-booking-guru-front-end`
  2. Clear the options `dokku docker-options:clear`
  3. Add the settings back you want for each stage
     e.g. `dokku docker-options:add salon-booking-guru-front-end deploy --restart=on-failure:10`


### Set dokku git pre-recieve hook to get Dockerfile
> This is an awful hack since when docku builds containers the build context is hardcoded to the 
root of the repo.

Add a pre-recieve hook on the dokku server to fetch the `Dockerfile` into the repo:

#### Front-end
```bash
cd /home/dokku/salon-booking-guru-front-end
vi ./hooks/pre-receive
```

Add to the file `hooks/pre-recieve`:
```bash
mkdir -p /home/dokku/salon-booking-guru/api/ && curl https://raw.githubusercontent.com/KarmaComputing/salon-booking-guru/main/front-end/Dockerfile > /home/dokku/salon-booking-guru/front-end/Dockerfile
```

#### API

```bash
cd /home/dokku/salon-booking-guru-api
vi ./hooks/pre-receive
```

Add to the file `hooks/pre-recieve`:
```bash 
#!/usr/bin/env bash
set -e
set -o pipefail

mkdir -p /home/dokku/salon-booking-guru-api/ && curl https://raw.githubusercontent.com/KarmaComputing/salon-booking-guru/main/api/Dockerfile > /home/dokku/salon-booking-guru-api/Dockerfile

cat | DOKKU_ROOT="/home/dokku" dokku git-hook salon-booking-guru-api
```

#### Database
```bash
cd /home/dokku/salon-booking-guru-database
vi ./hooks/pre-receive
```

Add to the file `hooks/pre-recieve`:
```bash
mkdir -p /home/dokku/salon-booking-guru-database && curl https://raw.githubusercontent.com/KarmaComputing/salon-booking-guru/main/database/Dockerfile > /home/dokku/salon-booking-guru-database/Dockerfile
```
### Quick Check

To make sure everything is in order, you can manually push to Dokku.
- `git clone` the repository on your local machine
- Add the dokku remote: `git remote add dokku dokku@<ip/domain>:salon-booking-guru-front-end`
- Now `git push dokku main`

If you get the following error: `remote: unable to prepare context: unable to evaluate symlinks in Dockerfile path: lstat /home/dokku/salon-booking-guru/api: no such file or directory`. This means that the Dockerfile is not already on the Dokku servers repository, make sure that its already there and check the pre-recieve hook so that it has the correct URL to the Dockerfile. It must start with `raw.githubusercontent.com`.

## Set Up GitHub Actions

### API Automated deployment

- Create a yaml file within the path `.github/workflows/api-deploy.yml`

- Specify the GitHub event that triggers the workflow
    - Use the key ```on:``` to specify the event, we want push
        - From here you can add keys to specify the branch, we want main

> `workflow_dispatch` will allow you to rerun GitHub actions on demand.
> We are using ```idoberko2/dokku-deploy-github-action@v1``` for this job

> Github secrets must be created in the repo settings:
  - `SSH-PRIVATE-KEY`
  - `DOKKU-HOST`

> Note: `REMOTE-BRANCH`, is set to `main` not `master`

#### Deploy api Github action

> This example may not be up to date. [View latest github actions.](https://github.com/KarmaComputing/salon-booking-guru/tree/main/.github/workflows)

```
---
# When a pull request is merged into the main branch, deploy the latest
# version of the API

name: Deploy API
on:
  workflow_dispatch:
  push:
    branches:
    - main
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
        app-name: salon-booking-guru-api
        remote-branch: main
...
```

## Automatic Let's Encrypt TLS Certificate

> Via [dokku-letsencrypt](https://github.com/dokku/dokku-letsencrypt)


### Install Let's Encrypt plugin
```
# as root on dokku
sudo dokku plugin:install https://github.com/dokku/dokku-letsencrypt.git
```
## Setup automatic certs for api and frontend

> Must use a valid email address
```
sudo -iu dokku
dokku config:set --global DOKKU_LETSENCRYPT_EMAIL=your@email.tld
```

### Certificate generation
```
# API
dokku letsencrypt:enable salon-booking-guru-front-end
# Front-end
dokku proxy:ports-add salon-booking-guru-api http:80:8085
dokku letsencrypt:enable salon-booking-guru-api
dokku proxy:ports-add salon-booking-guru-api https:443:8085

```
