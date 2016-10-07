### decorator class
class logit(object):
	def __init__(self, logfile='out.log'):
		self.logfile  = logfile
	def __call__(self, func):
		log_string = func.__name__ + "was called"
		print log_string
		with open(self.logfile, 'a') as f:
			f.write(log_string + "\n")
		#
		self.notify()
	def notify(self):
		pass

@logit
def myfunc1():
	pass
