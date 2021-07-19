# Salon Booking Guru

Online salon client booking and management system.

- Take deposits
- Online bookings
- Reminders to clients

# Documentation

[Read Documentation](https://karmacomputing.github.io/salon-booking-guru/)

## Contribte to the docs

To edit documentation locally
script:
```bash
./mkdocs.sh
```

Or alternatively, use the following docker command:
```bash
docker run --rm -it -p 8000:8000 -v ${PWD}:/docs squidfunk/mkdocs-material
