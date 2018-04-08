# sudo docker build -t go-anywhere .

# sudo docker run --rm --name="crowd-sell" --net="host" -v `pwd`/server:/go/src/main -v `pwd`/server-bin:/go/bin -u `id -u`:`cut -d: -f3 < <(getent group $(whoami))` -it go-anywhere bash
sudo docker run --rm --name="crowd-sell" --net="host" -v `pwd`/server:/go/src/main -v `pwd`/server-bin:/go/bin -it go-anywhere bash
