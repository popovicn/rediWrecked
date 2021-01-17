rediWrecked
-----------
Check URLs for response status code and redirect location.

### Arguments
```
  -i string
    	Input file path (required)
  -o string
    	Output file path (default "output.txt")
  -p int
    	Parallelism (default 50)
  -s	Simple CLI (tab separated)
```

### Example usage
```
> go run rediWrecked.go -i input.txt -o output.txt 

    ____ ____ ___  _ _ _ _ ____ ____ ____ _  _ ____ ___  
    |__/ |___ |  \ | | | | |__/ |___ |    |_/  |___ |  \ 
    |  \ |___ |__/ | |_|_| |  \ |___ |___ | \_ |___ |__/ 


200 https://www.qq.com/
200 https://www.wikipedia.org/
200 https://www.youtube.com/
200 https://www.facebook.com/
200 https://www.yahoo.com/
200 https://www.amazon.com/
301 https://www.twitter.com/ → https://twitter.com/
302 https://www.taobao.com/ → https://world.taobao.com
200 https://www.baidu.com/
```
Generated output file format: `[url, statusCode, redirect_url]`, tab separated:

```
https://www.qq.com/	200	
https://www.wikipedia.org/	200	
https://www.youtube.com/	200	
https://www.facebook.com/	200	
https://www.yahoo.com/	200	
https://www.amazon.com/	200	
https://www.twitter.com/	301	https://twitter.com/
https://www.taobao.com/	302	https://world.taobao.com
https://www.baidu.com/	200	
```
