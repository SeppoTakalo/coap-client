# Simple CoAP client

This is very simple CoAP client, mainly develop as my excercise to learn Go language.

## Usage

```
Usage: coap-client [options] <URL>:
  -x string
    	Request method: GET PUT POST (default "GET")

```

## Example

```
❯ coap-client -x GET coap.me/test
welcome to the ETSI plugtest! last change: 2023-11-15 21:31:17 UTC
```
