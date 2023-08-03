# server_template

This is literally just my server with less services and with the config stripped.

This is a template for starting a webserver. It demonstrates logging, auth, and muxing. It expects SSL.

You will have to run `go mod init` and give it a name to build it.

The `Makefile` handles some basic building and installing. There's a commented-out line in `make install` which you need to give your server access to port 80. You need to install `libcap2-bin` on debian, at least.

I use letsencrypt and acme certbot to install my certificate. Installing can be tough, but the command is `certbot certonly --manual --preferred-challenges dns -d domain1,*.domnain1,domain2,*.domain2`

My user is part of the `webweb` group and the certs it needs AND the directories it serves are all owned by the group.
You can install `sudo apt install acl` and then use

```
chgrp webweb DIRECTORY
chmod g+ws DIRECTORY
setfacl -d -m g::rwx DIRECTORY
setfacl -d -m o::rx DIRECTORY
```

so that all files in that directory are given the proper permissions. That command gives all users read-access, so not suitable for your certs/keys.

## Dependencies

It uses my [ilog](github.com/ajpikul-com/ilog), [uwho](github.com/ajpikul-com/uwho), [sutils](github.com/ajpikul-com/sutils), fundamentally.

*Remember to change your `git remote` after you clone it to your own fork!
