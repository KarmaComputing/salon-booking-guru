# Deployment

Automatic deployment using Dokku and GitHub Actions

## Set Up GitHub Actions

- Create a yaml file within the path `.github/workflows/`

### In the file:

- Specify the GitHub event that triggers the workflow
    - Use the key ```on:``` to specify the event
        - From here you can add keys to specify the branch, we want main

### Now its time to create a job for the workflow run:

- Define a job ```jobs:```
    - Give the job a unique name by using a key
    - Specify the runner environment, we are using ubuntu-20.04

### Now create the steps for the job:

- Use the ```actions/checkout@v2``` to fetch the code from the main branch that we specified
    - You should specify ```fetch-depth: 0``` to fetch all history
    - Tell it which dockerfile you want used by specifiying a path, in this case `api/`

- Now make the first job! You must give the job an id and optionally a name
- We are using ```idoberko2/dokku-deploy-github-action@v1``` for this job
    - Within the ```with:``` key we specify all the sensitive details for the dokku server, store them in the repo secrets
        - For this deployment we specify ```ssh-private-key```, ```app-name``` and ```dokku-host```
        - Then specify the ```remote-branch```, as written before we want main
