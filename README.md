# socketandserve

Serves sockets in a folder as subdomains of `.dev`.

Say you have a folder, `/var/sockets`, and serve some webapps on sockets in that
folder, like

    /var/sockets
      - myapp.sock
      - otherthing.sock

you can then

    $ curl myapp.dev
    ...
    $ curl otherthing.dev
    ...


## Instructions

It will run on port 8080 by default, forwarding requests to `http://xxx.dev` to
the socket `xxx.sock` in the directory specified by `--socket-dir`. The only
thing left to do is to setup forwarding port 80 to port 8080.

On Ubuntu, at least, this is not _too_ hard.

1. `$ sudo apt-get install dnsmasq resolvconf`

2. Edit `/etc/dnsmasq.conf`, add

  ```
  address=/dev/127.0.0.1
  interface=lo
  no-dhcp-interface=lo
  ```

3. Reboot.

4. `$ sudo iptables -t nat -I OUTPUT -p tcp -d 127.0.0.1 --dport 80 -j REDIRECT --to-ports 8080`

5. It should work!

> Stolen from
> <https://mkrmr.wordpress.com/2011/07/15/using-dnsmasq-to-run-your-own-tld/>
> and
> <http://serverfault.com/questions/112795/how-can-i-run-a-server-on-linux-on-port-80-as-a-normal-user>.
