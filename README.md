# pi [pí]
[![Travis Build Status](https://travis-ci.org/a-know/pi.svg?branch=master)][travis]

[travis]: https://travis-ci.org/a-know/pi


## Description

`pi` (`[pí]`) is a CLI tool for Pixela ([pixe.la](https://pixe.la/))


![](https://user-images.githubusercontent.com/1097533/53243207-84d04680-36ea-11e9-8465-f73d62b4b502.png)


## Installation

    % go get github.com/a-know/pi/cmd/pi

## Synopsis

    % pi users create --username a-know --token thisissecret --agree-terms-of-service yes --not-minor yes
    % export PIXELA_USER_TOKEN=thisissecret
    % export PIXELA_USER_NAME=a-know
    % pi graphs create -g my-first-graph -n "My first graph" -i commits -t int -c shibafu -z "Asia/Tokyo" -s none
    % pi pixel post -g my-first-graph -d 20190101 -q 5 -o "{\"key\":\"value\"}"
    % pi graphs svg -g my-first-graph | xargs open

## Available commands

```sh
  graphs    operate Graphs
  pixel     operate Pixel in Graph
  users     operate Users
  version   display version
  webhooks  operate Webhooks
```

### Subcommands
#### `users`
```
  create  create User
  delete  delete User
  update  update User Token
```


#### `graphs`
```
  create  create Graph
  delete  delete Graph
  detail  get Graph detail URL
  get     get Graph Definitions
  pixels  get Graph Pixels
  svg     get SVG Graph URL
  update  update Graph Definition
```


#### `pixel`
```
  decrement  decrement a Pixel
  delete     delete a Pixel
  get        get a Pixel
  increment  increment a Pixel
  post       post a Pixel
  update     update a Pixel
```

#### `webhooks`
```
  create  create a Webhook
  delete  delete a Webhook
  get     get registered Webhooks
  invoke  invoke Webhook
```


## Options
Please see the running result each subcommands with `-h`.


## CI running count

[![CI running count](https://pixe.la/v1/users/pi/graphs/ci-count)][ci-count]

[ci-count]: https://pixe.la/v1/users/pi/graphs/ci-count.html



## Author

[a-know](https://github.com/a-know)
