# mod.py
import logging
from abc import abstractmethod

class mod:
	def __init__(self, lgr: logging.Logger):
		self.logger: logging.Logger = lgr

	@abstractmethod
	async def run(self, *args):
		pass
