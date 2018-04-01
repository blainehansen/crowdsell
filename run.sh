# sudo docker build -t crowd-sell .

sudo docker run --rm --name="crowd-sell" --net="host" -v `pwd`/server:/go/src/main -v `pwd`/server-bin:/go/bin -u `id -u`:`cut -d: -f3 < <(getent group $(whoami))` -it crowd-sell bash
