# mod.py
from abc import abstractmethod

class mod:
	description: str
	how: str
	def __init__(self):
		pass

	@abstractmethod
	async def run(self, *args):
		pass
