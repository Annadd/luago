function lock2seconds(lock)
	if not lock:tryLock() then
		error("Unable to acquire the lock!")
	end
	pcall(function()
		sleep(2000)
	end)
	lock:unlock()
end