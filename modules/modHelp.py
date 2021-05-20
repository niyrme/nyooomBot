# modHelp.py
from modules.mod import mod

class module(mod):
	def __init__(self):
		super().__init__()
		self.description = "Show help"
		self.how = "`?help`"

	async def run(self, *args):
		del args
		return "Helpful message!"
