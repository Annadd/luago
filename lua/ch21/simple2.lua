co = coroutine.create(function ()
	print(coroutine.status(co)) --> running
	coroutine.resume(coroutine.create(function()
		print(coroutine.status(co)) --> normal
	end))
end)
print(coroutine.status(co)) --> suspended
coroutine.resume(co)
print(coroutine.status(co)) --> dead 
