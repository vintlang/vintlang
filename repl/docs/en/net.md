# HTTP with vint

You can access the internet via http protocol using the `net` module.

## Importing

Import the module with:
```
import net
```

## Methods

### get()

Use this as GET method. It can either accept one positional argument which will be the URL:

```
import net

net.get("http://google.com")
```

Or you can use keyword arguments to pass in parameters and headers as shown below. Note that headers and parameters must be a dictionary:

```
import net

url = "http://mysite.com"
headers = {"Authentication": "Bearer XXXX"}

net.get(url=url, headers=headers, body=params)
```

### post()

Use this as POST method. Use keyword arguments to pass in parameters and headers as shown below. Note that headers and parameters must be a dictionary:

```
import net

url = "http://mysite.com"
headers = {"Authentication": "Bearer XXXX"}
params = {"key": "Value"}

net.post(url=url, headers=headers, body=params)
```
