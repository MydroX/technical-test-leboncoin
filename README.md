# Technical test Leboncoin

## Brief

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". 
The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,..."

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.



## Quick start (prod)

  ```bash 
  $ ./deploy.sh
  ```

## Run dev mode 
  
  ```bash
  $ echo GIN_MODE=debug > .env
  ```

  ```bash
  $ docker-compose -f deploy/docker-compose.yml -f deploy/dev/docker-compose.yml up --build
  ```

## Endpoint

Endpoint : http://localhost:3000

Body 
```
{
  "int1": [int],
  "int2": [int],
  "limit": [int],
  "str1": [string],
  "str2": [string]
}
```

## Curl request example

```bash
curl --location --request POST 'localhost:3000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "int1": 6,
    "int2": 2,
    "limit": 10,
    "str1": "pee",
    "str2": "wee"
}'
```