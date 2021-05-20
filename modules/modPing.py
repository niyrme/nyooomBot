# modPing.py
from modules.mod import mod

class module(mod):
	def __init__(self):
		super().__init__()
		self.description = "Reply `Pong!`"
		self.how = "`?ping`"

	async def run(self, *args):
		del args
		return "Pong!"
