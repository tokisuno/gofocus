# gofocus
This is a port of [stream-pomodoro](https://github.com/tokisuno/stream-pomodoro) written in Go.

After hearing about the TypeScript interpreter being ported over to Go, it made me remember my time with go and how much I enjoyed making things in it. So I ported over some of what I had in C over to Go and I think this is what I want to continue using. Go is simple and efficient. I don't really care about memory management for the time being because most of what I do doesn't require it. C is still great though. It's just not what I need for the time being.

# Requirements
- go
- raylib-go

# How to use
## Linux
1. `git clone git@github.com:tokisuno/gofocus.git` into your desired directory
2. `go build`
3. run the executable!

## Windows
1. TBD

# Mappings
- ``<C-y>`` :: Next screen/scene
- ``<C-a>`` :: Increment session counter +1
- ``<C-x>`` :: Decrement session counter -1

# TODO
- [ ] Figure out how to make a GitHub release
- [ ] Make GitHub releases for both Linux and Windows

# credits
- [font](https://fonts.google.com/specimen/JetBrains+Mono)
- [sfx](https://rpg.hamsterrepublic.com/ohrrpgce/Main_Page)
