# Microencrypt
## What:
An http "microservice" that will asymmetrically encrypt whatever you hand it.
## Why:
I have written this code way too many times for scrapers. This endpoint is super
useful when you need to encrypt user data (username/password) for scrapers, but 
you don't want to store it server side. That way the server can encrypt it for
the client, and on each request you can send the encrypted data to the server
which can then use it to get the data from the site that you are scraping
## Contribution:
Please submit a pull request if you notice any possible issues with this
repo and write a test to accompany the issue so that it is easily reproducible,
and to ensure that the service is always working.