# Quick-start guide

This guide will help you quickly configure and launch an instance of the Salon
Booking Guru with the intent of developing on the codebase.

---

## Database Set-up

To allow the API to have persistent data storage, you must first have a
PostgreSQL database running, we have provided a script in to start a PostgreSQL
instance using docker:

```bash
postgres/run.sh
```

### Database logs

```
docker logs -f postgres
```

**Note:** This database configuration should only ever be used for the purpose
of local development, since all credentials are stored in this public
repository.

## API Set-up

Now you have a PostgreSQL database up and running on your local machine, you
will be able to run the API. In the future we will achieve this by using Docker,
however for now, you must build the API yourself using golang's compiler.

First ensure you have [Golang](https://golang.org/doc/install) installed and
configured, then change into the
`api/` directory:

```bash
cd api/
```

Then use the `run.sh` script to install all dependencies, build a binary, and
run the binary:

```bash
./run.sh
```

For this to work, it is likely that this repository will need to exist in the
correct location of your GOPATH, you directory structure should look something
like this on a UNIX environment:

```
~/go/src/github.com/KarmaComputing/salon-booking-guru/
```

For more information on how the GOPATH works,
[check out this article by Digital Ocean](https://www.digitalocean.com/community/tutorials/understanding-the-gopath).

**Note:** The environment variables set in `api/run.sh` match those of the
postgres docker container started by `postgres/run.sh`.

## Front-end Set-up

Similarly to the API, we will eventually use docker to start an instance of the
front-end, however in the meantime you must use `yarn`.

First ensure you have [Yarn](https://yarnpkg.com/getting-started/install) installed.
Then change into the `front-end/` directory:

```bash
cd front-end/
```

Then install all dependencies using :

```bash
yarn
```

You will then be able to start the front-end server using:

```bash
yarn serve
```
