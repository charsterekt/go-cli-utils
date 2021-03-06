# Go CLI Utils

A collection of short, sweet, and simple CLI utils written in Golang, primarily for learning purposes. However, they do have their practical uses. The currently available tools are detailed below:

- - - - 

### Network Info CLI

A simple CLI that takes a host URL as an argument and is capable of looking up info about the provided host such as: <br>

* CNAME records
* MX records
* IP addresses
* Nameservers

Example:

`./network-lookup ns --host google.com`

- - - -

### File Downloader CLI

A simple CLI that takes a URL containing a file as the argument and downloads and saves the file at the location into the current working directory. If no file is directly linked in the URL, it will download the entire webpage. It may be customized to bulk download files.

Example:

`./downloader dl --link https://images.firstpost.com/fpimages/1200x800/fixed/jpg/2021/05/Below-Zero-Out-Now_opt.jpg`

- - - -

More tools may be added in the future.
