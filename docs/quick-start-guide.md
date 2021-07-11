# Quick-Start Guide

This guide will help you quickly configure and launch an instance of the Salon
Booking Guru web API.

---

## Installation

### Docker

To get started with a docker container, you must first build an image with the
following command:
```bash
docker build . -t karmacomputing/salon-booking-guru
```

Once this is complete you will be able to run a container:
```bash
docker run -it karmacomputing/salon-booking-guru
```

### Golang

Alternatively, you may wish to build the binary directly using the golang
compiler, for this, you can use the helper script we have provided:
```bash
./run.sh
```
