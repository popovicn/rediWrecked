rediWrecked
-----------
Check URLs for response status code and redirect location.

#### Arguments
```
  -i string
    	Input file path (required)
  -o string
    	Output file path (default "output.txt")
  -p int
    	Parallelism (default 40)
```

#### Example usage
```
> go run rediWrecked.go -i input.txt -o output.txt 
http://trickest.com	301	https://trickest.com/
https://trickest.com	200 
```