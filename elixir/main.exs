IO.puts "Hello world from Elixir"
:apple
IO.puts "apple" == :apple
IO.puts :apple
IO.puts is_binary(:apple)


add = fn a,b -> IO.puts a + b end
add.(1, 3)

array = [1, 2, 3]
IO.puts hd(array)