# Pokemon Download
This is a simple golang program that downloads pokemon episodes in the background based on the start and end episode numbers provided. This program downloads rpisodes from [Pokemon360](https://pokemon360.me).

This program can be run on a raspberry pi which has a stable wifi connection. In this way you would not have to manually download all the episodes.

The program will download one episode at a time.

P.s. there are a total of 1087 episodes thats why the program does no download all of them at once😅.

## Project setup
This app can be run on the latest version of go.

### To run the app:

> ```go run main.go <start episode number> <end episode number>```

Example
> `go run main.go 1 5` 

this downloads first 5 episodes

_Gotta catch em all_