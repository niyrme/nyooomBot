# modDiscord.py
import logging
import os
import sys

import discord
from dotenv import load_dotenv

from modules import mod
from modules import modDice
from modules import modHelp
from modules import modPing

hdlr: logging.StreamHandler = logging.StreamHandler(sys.stdout)
hdlr.setFormatter(logging.Formatter("[%(name)s] | %(levelname)s | %(message)s"))

logger: logging.Logger = logging.getLogger("BOT-DISCORD")
logger.setLevel(logging.INFO)
logger.addHandler(hdlr)

class Client(discord.Client):
	# Login
	async def on_ready(self):
		logger.info("Logged in.")
		self.command: str = "?"
		self.cmds: dict[str, mod.mod] = {
			"dice": modDice.module(),
			"help": modHelp.module(),
			"ping": modPing.module(),
		}

	# Message-event
	async def on_message(self, msg: discord.Message):
		# Do not reply to self
		if msg.author == cl.user:
			return

		# Check if msg is a command
		if not str(msg.content).startswith(self.command) and len(str(msg.content)) <= 3:
			return

		response: str = ""

		logger.info(f"New Command: '{msg.content}'")

		command, *commandArgs = str(msg.content).split(' ')
		command = str(command[1:]).lower()
		c: str

		if command == "how":
			if len(commandArgs) == 0:
				response = "Not enough arguments given"
			else:
				if commandArgs[0].startswith(self.command):
					c = commandArgs[0][1:].strip()
				else:
					c = commandArgs[0].strip()

				if c == "how":
					response = "`?how {command}`"
				elif c == "desc":
					response = "`?desc {command}`"
				elif c in self.cmds:
					response = self.cmds[c].how
		elif command == "desc":
			if len(commandArgs) == 0:
				response = "Not enough arguments given"
			else:
				if commandArgs[0].startswith(self.command):
					c = commandArgs[0][1:].strip()
				else:
					c = commandArgs[0].strip()

				if c == "how":
					response = "Shows how to do a command."
				elif c == "desc":
					response = "Shows description of a command."
				elif c in self.cmds:
					response = self.cmds[c].description
		elif command in self.cmds:
			response = await self.cmds[command].run(commandArgs)


		if response:
			await msg.channel.send(response)
			return
		else:
			await msg.channel.send(f"Unknown command: `{msg.content}`")


if __name__ == '__main__':
	load_dotenv()
	cl: Client = Client()
	cl.run(os.getenv("DISCORD_TOKEN"))
