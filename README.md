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
200 https://www.apple.com/
302 https://www.microsoft.com/ → https://www.microsoft.com/sr-latn-rs/
302 https://www.paypal.com/ → https://www.paypal.com/rs/home
301 https://www.live.com/ → https://outlook.live.com/owa/
301 https://www.adcash.com/ → https://adcash.com/
301 https://www.kickass.so/ → http://www.kickass.so/
200 https://www.wikipedia.org/
200 https://www.amazon.de/
301 https://www.vk.com/ → https://vk.com/
403 https://www.163.com/
200 https://www.youtube.com/

```
Generated output file format: `[url, statusCode, redirect_url]`, tab separated:

```
https://www.qq.com/	200	
https://www.apple.com/	200	
https://www.microsoft.com/	302	https://www.microsoft.com/sr-latn-rs/
https://www.amazon.de/	200	
https://www.kickass.so/	301	http://www.kickass.so/
https://www.live.com/	301	https://outlook.live.com/owa/
https://www.adcash.com/	301	https://adcash.com/
https://www.wikipedia.org/	200	
https://www.google.co.in/	200	
https://www.google.co.uk/	200	
https://www.google.de/	200	
https://www.google.ru/	200	
https://www.google.it/	200	
https://www.google.com.br/	200	
https://www.sina.com.cn/	200	
https://www.google.es/	200	
https://www.youtube.com/	200	
```
