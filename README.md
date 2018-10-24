<!---
    Copyright 2018 storyicon@foxmail.com
 
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at
 
        http://www.apache.org/licenses/LICENSE-2.0
 
    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
-->

**We are looking for contributors**! Please check the [ROADMAP](https://github.com/storyicon/graphquery/blob/master/ROADMAP.md) to see how you can help ❤️

---

# GraphQuery-HTTP

![GraphQuery](https://raw.githubusercontent.com/storyicon/graphquery/master/docs/screenshot/graphquery.png)   

[GraphQuery](https://github.com/storyicon/graphquery) is a query language and execution engine tied to any backend service.GraphQuery-PlayGround is a web application for practicing, learning and testing GraphQuery.

GraphQuery-http is a service that implements `cross-language` calls to GraphQuery via the `HTTP` protocol.

## How to use

### 1. Use binary distribution

Go to the [Release](https://github.com/storyicon/graphquery-http/releases) page, download and unzip the corresponding binary package according to your system type, run the service, and query GraphQuery by requesting `127.0.0.1:8558`.

#### Request format:             

> Request Address: `127.0.0.1:8558`  

> Request Method: `POST`   

> Request Parameters:          

| Name | Must | Type   | Default | Intro                   |
|------|------|--------|--------|--------------------------|
| document | ture | string |     | document to be parsed   |
| expression | true | int    |     | GraphQuery expression |
                  
> Response          
                 
| Name | Type | Default | Intro                   |
|------|------|--------|--------|--------------------------|
| data | string |     | analyzed result   |
| error |  string   |     | errors that occur during parsing |
| timecost | int    |  0   | time spent |

> Example: Curl

```bash
curl 127.0.0.1:8559 -X POST -d "document=<title>hellow world</title>" -d "expression=title `css(\"title\")`"
```
Response:
```
{"data":"hellow world","error":"","timecost":0}
```

> Example: Python

```python
import requests
response = requests.post("http://127.0.0.1:8559", data={
    "document": '<title>hellow world</title>',
    "expression": 'title `css("title")`',
})
print(response.text)
```
Response:
```json
{"data":"hellow world","error":"","timecost":0}
```

You can adjust the behavior of GraphQuery-http by parameters when starting the binary.       
Support parameter:      
1. `-debug`: You can turn off the debug mode when you start the service by adding `-debug=false` after the command line.
2. `-port`: You can add `-port=` to the command line to specify the listening port when you start the service.
3. `-h` or `-help`: You can add `-h` or `-help` to the command line to view all the custom parameters when you start the service.       


### 2. Compile

If you don't want to use the [release version](https://github.com/storyicon/graphquery-http/releases) and want to modify the source code, you can read the following steps to compile.

```
go get github.com/storyicon/graphquery-http
```

Find the downloaded `storyicon/graphquery-http` in GOPATH, 

```
go build service.go
```
Execute the obtained binary package `service`, the following steps are the same as [above](#Request-format:).

## Tips         
Communication between services via HTTP will inevitably increase the consumption of time. After testing, the delay is about 10ms. If you use native Go to call directly, the time consumption will be at the microsecond level. In fact, the 10ms delay generated by GraphQuery-http's http communication is negligible for the network IO generated by the crawler, but if you want to to be more efficient, you can try to use `RPC` for communication (such as [grpc](https://github.com/grpc/grpc)).      

