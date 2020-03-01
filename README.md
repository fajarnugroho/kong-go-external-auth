# KONG Go External Authentication Plugin
## Introduction
A Kong External Authentication Plugin written in go

## Build
```bash
$ docker build -t reponame/imagename:tag .
```

## Config
```
Url: http://example.com
Method: GET
Headers: api-key,apiKey
```

`Url` : your authentication server URL
`Method` : Method will be used for request to your authentication server
`Headers` : header will be read from client request and will be send to
authentication server (comma separated)

## LICENSE

```
The MIT License (MIT)
=====================

Copyright (c) 2015 Panagis Tselentis

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

```
