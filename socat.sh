TTYS110="/dev/ttyS110"
TTYS111="/dev/ttyS111"
echo "You need sudo rights for this utility to work"
echo "This script creates two fake serial terminals,"
echo "$TTYS110 and $TTYS111 and links them together"
echo "Data written to one is echoed to the other and vice versa"
echo "Terminating this process ends the link"
sudo socat PTY,link=$TTYS110,user=$USER PTY,link=$TTYS111,user=$USER
