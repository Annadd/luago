co = coroutine.create(function(a, b, c)
	print(a, b, c)
	while true do
		print(coroutine.yield())
	end
end)

coroutine.resume(co, 1, 2, 3) --> 1 2 3
coroutine.resume(co, 4, 5, 6) --> 4 5 6
coroutine.resume(co, 7, 8, 9) --> 7 8 9