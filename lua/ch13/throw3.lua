function lock2seconds(lock)
	lock:lock()
	local ok, msg = pcall(sleep, 2000)
	lock:unlock()

	if ok then
		print("ok")
	else
		print("error: " .. msg)
	end
end