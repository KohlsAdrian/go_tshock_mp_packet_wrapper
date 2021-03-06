## Tshock Multiplayer Packet Wrapper


# What does it do?

It wraps all packets from the game server into Go struct data, so you can read the game data from JSON easily

# How to use it

# 1 - Install the [Multiplayer Packet Wrapper Plugin](https://github.com/KohlsAdrian/tshock_multiplayer_packet_wrapper_plugin) on your tShock Terraria Server

# 2 - Create 2 endpoints

2.1 - Packet streaming as `localhost:10000/` (POST)

2.2. - Read the last packet read as `localhost:10000/lastPacket` (GET)

# 3 - Done!
