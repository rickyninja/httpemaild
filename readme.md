# httpemaild

This is an http service that runs on a host where I have already configured sending mail.  It may use
[email to SMS gateways](https://en.wikipedia.org/wiki/SMS_API) to send a notice to your phone.

# example usage

```
curl -vvv -X POST --data '{ "recipient": "1234567890@txt.att.net", "message": "foo"}' http://localhost:8080/minecraft/notify
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> POST /minecraft/notify HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.52.1
> Accept: */*
> Content-Length: 72
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 72 out of 72 bytes
< HTTP/1.1 200 OK
< Date: Tue, 11 Feb 2020 07:18:35 GMT
< Content-Length: 0
<
* Curl_http_done: called premature == 0
* Connection #0 to host localhost left intact
```
