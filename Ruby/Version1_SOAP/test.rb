require 'socket'

ip = "52.7.155.169"

socket = TCPSocket.new(ip, 443)

puts "Conectado!"
socket.close