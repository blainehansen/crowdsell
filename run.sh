# sudo docker build -t crowd-sell .

sudo docker run --rm --name="crowd-sell" --net="host" -v `pwd`/src:/go/src/crowd-sell -v `pwd`/bin:/go/bin -v `pwd`/pkg:/go/pkg -it crowd-sell bash
