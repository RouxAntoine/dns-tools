# dns-tools

dns relative tools like lookup or wireformat converter

## build

```
make wireformat
make http-resolver 
```

## run

With get resolution

`echo -n 'q80BAAABAAAAAAAAA3d3dwdleGFtcGxlA2NvbQAAAQAB' | base64 -d | curl -sH 'content-type: application/dns-message' --data-binary @- https://cloudflare-dns.com/dns-query -o - | ./bin/wireformat -f -`

With post resolution

`curl -s -H accept: application/dns-message https://cloudflare-dns.com/dns-query?dns=q80BAAABAAAAAAAAA3d3dwdleGFtcGxlA2NvbQAAAQAB | ./bin/wireformat -f -`

Wireformat encode and decode

`./bin/wireformat -s 'www.example.com' | ./bin/wireformat -f -`

`echo 'www.example.com' | ./bin/wireformat | ./bin/wireformat -f -`


Dns query and wireformat encode/decode

`echo 'www.google.com' | ./bin/wireformat -s - | ./bin/http-resolver | ./bin/wireformat -f -`

