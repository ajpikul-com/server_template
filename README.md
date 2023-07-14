# ajpikul.com

This is a template for starting a webserver. It demonstrates logging, auth, and muxing. It expects SSL.

You will have to run `go mod init` and give it a name to build it.

```
-main.go defines a static fileserver, creates a URL path multiplexer, and attaches all paths to it.
-string.go` defines a string server, and provides a function to attach it to a multiplexer.
-systemInit.go handles some logging setup and flags
```

The `Makefile` handles some basic building and installing. There's a commented-out line in `make install` which you need to give your server access to port 80. You need to install `libcap2-bin` on debian, at least.\

I use letsencrypt and acme certbot to install my certificate. It's actually kind of a nightmare, but I use [auditmatic](github.com/ajpikul-com/auditmatic) to manage all my servers so there's useful stuff in there.

My user is part of the webweb group and the certs it needs AND the directories it serves are all owned by the group.
You can install `sudo apt install acl` and then use

```
setfactl -d -m g::rwx DIRECTORY
setfactl -d -m o::rx DIRECTORY
```

so that all files in that directory are given the proper permissions. That command gives all users read-access, so not suitable for your certs/keys.

## Depencies

It uses my [ilog](github.com/ajpikul-com/ilog), [simpauth](github.com/ajpikul-com/simpauth), [server_utils](github.com/ajpikul-com/server_utils).

*Remember to change your `git remote` after you clone it to your own fork!

## Analysis Tools
The `Makefile` contains some analysis tools but those need to be pulled out to seperate directory.
