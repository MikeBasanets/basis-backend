IP_ADDR=$1

go build

ssh root@${IP_ADDR} "killall basis"
scp basis root@${IP_ADDR}:~/
scp -r ./static root@${IP_ADDR}:~/
ssh -n -f root@${IP_ADDR} "nohup ./basis > /dev/null 2>&1 &"

rm basis
