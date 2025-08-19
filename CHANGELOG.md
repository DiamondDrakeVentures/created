<!-- markdownlint-disable MD024 -->

# Changelog

All notable changes in Create(D) will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).
This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
However, each version is only backwards compatible within the same Game Version.
For example, v1.2.0 is backwards compatible with v1.0.0 as long as they are both installed on the
same Minecraft version (e.g. 1.21.1).

## [UNRELEASED]

### Added

- Added Sepia Fungus Colony (Create: Garnished).
- Added Hypertube Accelerator (Create: Hypertubes).
- Added `polished` and `bricks` `StoneType`s (Every Compat (Stone Zone)).
- Added Forward Error Correction (FEC) to Opus encoder/decoder (Simple Voice Chat).
- Added Classical Chinese translation (Sophisticated Backpacks).
- Added HUD component showing the target blocks and entities (Jade).
- Added configuration menu (Ender Dragon Fight Remastered: Create(D) Edition).
- Added training dummy (MmmMmmMmmMmm).
- Added mods:
  - Configurable
  - Create Waystones Recipes
  - Ender Dragon Fight Remastered: Create(D) Edition
  - Global Packs
  - Jade
  - Jade Addons
  - MmmMmmMmmMmm
- Added resource packs:
  - Create Waystones Texture + GUI

### Changed

- Adjusted Limestone crushing/milling recipes (Create: Garnished).
- Adjusted Dragon's Breath fluids to be inclusive of the one from Create: Dragons Plus (Create: Garnished).
- Improved emoji compatibility (Chat Heads).
- Improved hypertube travel speed handling (Create: Hypertubes).
- Create Train Parts blocks are now minable (Create Train Parts).
- Renamed Mob Essence items to be number tiered (Create: Peaceful).
- Statue poses of copper golem can now be changed by right clicking (Friends&Foes).
- Changed copper golem sounds (Friends&Foes).
- Improved audio quality in bad networking conditions (Simple Voice Chat).
- Ticking entity actions are now only sent to players with the appropriate permissions to perform such actions (Neruina).
- Ads now die after the Dragon dies (Ender Dragon Fight Remastered: Create(D) Edition).
- Dragon is now invulnerable when the miniboss is active (Ender Dragon Fight Remastered: Create(D) Edition).
- Dragon egg now respawns at the end of every Dragon fight.
- Waystones, Portstones, and Sharestones can now only be crafted with Mechanical Crafters (Create Waystones Recipes).
- Waystones, Portstones, and Sharestones block texture and UI are now in-line with the Create aesthetics (Create Waystones Texture + GUI).
- Lodestones now cost an Iron Ingot rather than a Netherite Ingot to craft.
- Advancement for using a Lodestone now reside under the Adventure category.
- Updated mods:
  - Balm 21.0.48 => 21.0.49
  - Chat Heads 0.13.18 => 0.13.20
  - Create: Garnished 2.1.6.2 => 2.1.6.5
  - Create: Hypertubes 0.2.4 => 0.2.5
  - Create More: Parallel Pipes 1.1.0 => 1.1.1
  - Create: Peaceful 2.0.0 => 2.0.1
  - Create Train Parts 0.1.3 => 0.2.0
  - Every Compat (Stone Zone) 1.21-2.10.6 => 1.21-2.10.8
  - Friends&Foes 4.0.7 => 4.0.8
  - Fzzy Config 0.7.1+1.21+neoforge => 0.7.2+1.21+neoforge
  - Just Blahaj 2.0.3 => 2.0.5
  - KleeSlabs 21.1.6 => 21.1.7
  - KotlinLangForge 2.10.1-k2.2.0-3.0 => 2.10.3-k2.2.0-3.0
  - Neruina 2.3.1-beta.1 => 3.0.0
  - Puzzles Lib v21.1.36 => v21.1.38
  - Simple Voice Chat 2.5.35 => 2.5.36
  - Sophisiticated Backpacks 3.24.19.1292 => 3.24.19.1301
  - Sophisticated Core 1.3.59.1062 => 1.3.63.1086
  - Sophisticated Storage 1.4.46.1216 => 1.4.47.1223
  - Waystones 21.1.20 => 21.1.22
  - Wavey Capes 1.6.1.1 => 1.6.2

### Deprecated

### Removed

- Removed mods:
  - Ender Dragon Fight Remastered
  - Paxi

### Fixed

- Fixed potential dedicated server crash with Waystones (Balm).
- Fixed inventory button refusing on target-bound warp conditions before a target is chosen (Waystones).
- Fixed feature cycle order crash when `wildWaystoneStyle` was set to `MOSSY` (Waystones).
- Fixed Farmer's Delight cutting compatibility recipes loading (Create: Garnished).
- Fixed tool validity for Farmer's Delight cutting compatibility recipes (Create: Garnished).
- Fixed `create:packages` tag loading (Create: Garnished).
- Fixed compatibility recipe(s) with Create: Framed (Create: Garnished).
- Fixed detection system to detect similar blocks (Every Compat (Stone Zone)).
- Fixed Create blocks drops (Every Compat (Stone Zone)).
- Fixed es_mx translation (Friends&Foes).
- Fixed "no items match result" message duplication when the window gets resized (Sophisticated Core).
- Fixed calculation of total player XP in XP Pump Upgrade (Sophisticated Core).
- Fixed issues with Tank Upgrade I/O when handling fluids with NBT components (Sophisticated Core).
- Fixed an issue where chests and shulker boxes are see through when there are hidden tier or upgrade slots displayed when the player has storage tool or relevant upgrade in hand (Sophisticated Storage).

### Security

## [0.2.1] - 2025-08-13

### Fixed

- Fixed an issue where Advanced Backups is not installed on the client.
  While it is a server sided mod, Advanced Backups does show progress toasts if installed on the client and the player has the appropriate permissions.

## [0.2.0] - 2025-08-13

### Added

- Added backup solution (Advanced Backups).
  By default, backups are taken incrementally every 15 minutes, with a full backup taken about every six hours.
- Added shark plushies (Just Blahaj).
- Added Nightvision enchantment, which can be applied to Spyglass (Nightvision Spyglass).
  Spyglasses enchanted with Nightvision grant the nightvision effect while looking through them.
- Added player availability and streaming status on the tablist (Status).
  Status changes are done manually by the player.
- Added cosmetic hats (Cake's Cosmetics).
- Added a HUD overlay displaying coordinates, weather, and time (Immersive Overlay).
  The appropriate items must be present in the player inventory:
  - Compass for X and Z coordinate
  - Altimeter for Y coordinate
  - Clock for time and weather information
- Added aircrafts (Immersive Aircraft).
- Added ornithopter gliders (Create: Ornithopter Glider).
- Added teas and related things (Brewing Tea Tales).
- Added owls (Animal Garden - Owls).
- Added vanity armor slots (VanitySlots).
- Added accessory slots (Accessories).
- Added compatibility layer for Curios API (Curious Compat Layer for Accessories).
- Added global game rules (The Ruler).

### Changed

- Updated the Neoforge loader to 21.1.196.
- The game will abort the launch when the minimum memory requirement is not met (Load Support).
  The minimum requirement is 4 GB.
- Memory usage has been further optimized (Saturn).
- Singleplayer world can now be favorited (Cherised World).
  Favorited worlds are pinned to the top and cannot be deleted.
- Singleplayer world creation screen has been modernized (Modern World Creation).
- Singleplayer world thumbnails are now updated on every load (Updating World Icon).
- Flowers can be bonemealed, which will grow more of that specific flower (Flourish).
  This is a port of the Bedrock Edition mechanic.
- Statistics screen is now searchable (SearchStats).
- Viewing through a Spyglass can be done with a hotkey.
  The Spyglass does not need to be equipped but it must still be present in the inventory.
- Superflat world no longer spawns slimes (Superflat World No Slimes).
- WorldEdit now uses dedicated items as wands (WorldEdit Items).
- Most enchantments now show descriptions of their functionality (Enchantment Descriptions).
- Most potions and food items now show descriptions of their functionality (Effect Insights).
- Chains can now connect to fenceposts and walls (Connectible Chains).
- Crosshair now changes based on what the target (Dynamic Crosshair).
  It is also hidden when there is no target.
- Compass now always points north.
- Coordinates are no longer shown in the F3 debug screen (game rule `reducedDebugInfo`).
- Invite-only server button now shows `Create(D) SMP`.
- Waystones no longer generate in the wild (Waystones).
  This does not stop Towers of the Wild from generating the Waystone towers.
- Title screen background image is now from an earlier time in the Minecraft day.
- Centered the title contraption in the menu screen background image.
- Items and recipes list no longer are no longer displayed on the sides of the inventory screen by default (Distraction Free Recipes).
  Search an item or recipe to display them again.
- Replaced inventory items and recipes list with EMI (EMI).

### Removed

- Removed the ability to pick up entities and block entities (Carry On).
  There are compatibility issues with Carry On.
  The specifics of the issues are difficult to triage.
- Removed Elytra Trims.
  Conflicts with Customizable Elytra and many other things.
- Removed Macaw's Bridges.
- Removed Macaw's Roofs.
- Removed Macaw's Stairs.
- Removed zoom hotkey without Spyglass (Zume).
  The zoom hotkey now requires a Spyglass to be present in the inventory.
- Removed Curious API.
  Accessories + Curios Compat Layer for Accessories provides the functionality of Curios API.
- Removed JEI add-ons:
  - Just Enough Archaeology
  - Just Enough Effect Descriptions
  - Just Enough Professions
  - Just Enough Resources

### Fixed

- Fixed an issue where the title screen footer can sometimes overlap with the options button.
- Fixed a critical issue where Create Clipboard can be modified and causes items to be duplicated (Create Bugfix: Clipboard Patch).
- Fixed a critical issue where Create Schematic can be modified and causes items to be duplicated (Create Bugfix: Schematic Patch).
- Packet inconsistencies no longer disconnects the player (Network Protocol Disconnect).
  This is done at the risk of desynced state.
  Do not connect to untrusted servers as there is no longer a protection against malicious packets.

## [0.1.0] - 2025-08-03

### Added

- Added performance mods:
  - Alternate Current
  - bad packets
  - BadOptimizations
  - Chunky
  - Clumps
  - Dynamic FPS
  - Entity Culling
  - Fast IP Ping
  - Fast Item Frames
  - Fast Paintings
  - FastQuit
  - FerriteCore
  - GpuTape
  - ImmediatelyFast
  - Lithium
  - ModernFix
  - Noisium
  - Observable
  - Packet Fixer
  - Sodium
  - Sodium Extra
  - spark
  - Structure Layout Optimizer
- Added QoL mods:
  - AppleSkin
  - Automatic Tool Swap
  - Bad Wither No Coookie - Reloaded
  - Better Advancements
  - Better dispenser crafting recipe
  - Better Safe Bed
  - Better Third Person
  - Better Tridents
  - BetterF3
  - Boat Item View
  - Carry On
  - Chat Heads
  - Chunks fade in
  - Cobweb
  - Comfortable Campfires
  - Comforts
  - Companion
  - Controlling
  - Crash Assistant
  - Cubes Without Borders
  - Custom Nether Portal
  - Customizable Elytra
  - Death Backup
  - Default Options
  - Despawning Eggs Hatch
  - Ding
  - Disable End Portals
  - Do a Barrel Roll
  - Double Doors
  - Elytra Slot
  - Elytra Trims
  - Emojiful
  - Fated Inventory
  - FindMe
  - Glass Breaker
  - Hand Over Your Items
  - Horse Expert
  - Infinite Trading
  - InvMove
  - Just Enough Archaeology
  - Just Enough Beacons Reforged
  - Just Enough Breeding
  - Just Enough Effect Descriptions
  - Just Enough Items
  - Just Enough Professions
  - Just Enough Resources
  - KleeSlabs
  - Longer Chat History
  - Lootr
  - Maintenance Mode
  - Max Health Fix
  - Measurements
  - MixinTrace Reforged
  - Model Gap Fix
  - Mouse Tweaks
  - Neruina - Ticking Entity Fixer
  - NetherPortalFix
  - No Chat Reports
  - No Telemetry
  - OAuth Account Manager
  - Ping Wheel
  - Reese's Sodium Options
  - RightClickHarvest
  - Scaffolding Drops Nearby
  - Scholar
  - Screenshot to Clipboard
  - Screenshot Viewer
  - Server Chat Sync
  - Sign Copy
  - Sit
  - Smithing Template Viewer
  - SophisticatedFix
  - SwingThrough
  - Too Fast
  - TrashSlot
  - VoidTotem
  - Wall-Jump TXF
  - Weaker Spiderwebs
  - WorldEdit
  - WorldEdit CUI
  - Yeetus Experimentus
  - Zume
- Added Create ecosystem mods:
  - Create
  - Create Compressed
  - Create Encased
  - Create Extra Casing
  - Create Jetpack
  - Create Mechanical Extruder
  - Create Mechanical Spawner
  - Create More: Parallel Pipes
  - Create Ore Excavation
  - Create Railways Navigator
  - Create Slice & Dice
  - Create Smart Bounds
  - Create Train Parts
  - Create Train Tank Mod
  - Create Train Utilities
  - Create: Aquatic Ambitions
  - Create: Bells & Whistles
  - Create: Central Kitchen
  - Create: Connected
  - Create: Copycats+
  - Create: Curious Jetpack & Backtank
  - Create: Delivery Director
  - Create: Design n' Decor
  - Create: Diesel Genrators
  - Create: Dragons Plus
  - Create: Dreams & Desires
  - Create: Drinks
  - Create: Easy Stone Generators
  - Create: Easy Structures
  - Create: Enchantment Industry
  - Create: Escalated
  - Create: Extended Wrenches
  - Create: Extra Gauges
  - Create: Framed
  - Create: Garnished
  - Create: Hypertubes
  - Create: Integrated Farming
  - Create: Lazy Engines
  - Create: Liquid Fuel
  - Create: Misc and Things
  - Create: Molten Vents
  - Create: New Beginnings
  - Create: Numismatics
  - Create: Oxidized
  - Create: Peaceful
  - Create: Portable Packaging
  - Create: Power Loader
  - Create: Quality of Life
  - Create: Simple Ore Doubling
  - Create: Sound of Steam
  - Create: Synthetic Pressure
  - Create: Trading floor
  - Create: Trimmed
  - Create: Ultimate Factory
  - Create: Vibrant Vaults
  - Create: Winery
- Added mapping and navigation mods:
  - AA4 Atlas
  - Antique Atlas
  - Antique Trains
  - Chalk
  - Surveyor Map Framework
  - Towers of the Wild Modded
  - Waystones
- Added compatibility mods:
  - Architectury API
  - Better Beds
  - Connector
  - Every Compat (Stone Zone)
  - Every Compat (Wood Good)
  - Forgified Fabric API
  - InvMoveCompats
  - Iris & Oculus Flywheel Compat
  - Kotlin for Forge
  - KotlinLangForge
  - Macaw's Biomes O' Plenty
  - Rechiseled: Create
  - Sophisticated Backpacks Create Integration
  - Sophisticated Storate Create Integration
  - Sophisticated Storage in Motion
- Added mods:
  - Additional Lanterns
  - AddonsLib
  - Advancement Plaques
  - Almanac
  - Amendments
  - Aquatic Additions
  - Armor Statues
  - Balm
  - Bedspreads
  - Berry Good
  - Better Archeology
  - Better Combat
  - Biomes O' Plenty
  - Block Runner
  - Bookshelf
  - Caelus API
  - CERBON's API
  - Chef's Delight
  - Clayworks
  - Cloth Config API
  - Coastal Waves
  - Collective
  - Continuity
  - Corn Delight
  - CraterLib
  - Critter Armory
  - Curios API
  - Custom Credits
  - Death Counter
  - Death Finder
  - DragonLib
  - Drippy Loading Screen
  - Easy Magic
  - Ecologics
  - End's Delight
  - Ender Dragon Fight Remastered
  - Entity Model Features
  - Entity Texture Features
  - Expanded Delight
  - Extended Bone Meal
  - FancyMenu
  - Farmer's Delight
  - Forgeshot
  - Friends&Foes - Beekeeper Hut
  - Friends&Foes - Flowery Moonblooms
  - Friends&Foes
  - Fright's Delight
  - Fusion (Connected Textures)
  - Fzzy Config
  - GlitchCore
  - Iceberg
  - iChunUtil
  - Iris Shaders
  - Item Borders
  - Item Highlighter
  - Item Obliterator
  - JamLib
  - Kiwi
  - Konkrete
  - Macaw's Bridges
  - Macaw's Doors
  - Macaw's Fences and Walls
  - Macaw's Furniture
  - Macaw's Holidays
  - Macaw's Lights and Lamps
  - Macaw's Paintings
  - Macaw's Paths and Pavings
  - Macaw's Roofs
  - Macaw's Stairs
  - Macaw's Trapdoors
  - Macaw's Windows
  - Mechanicals Lib
  - Melody
  - Moonlight Lib
  - MultiBeds
  - My Nether's Delight
  - Mysterious Mountain Lib
  - Necronomicon API
  - Nether Weather
  - Not Enough Animations
  - Ocean's Delight
  - Panda's Falling Tree's
  - PandaLib
  - Paxi
  - Pehkui
  - playerAnimator
  - Prickle
  - Prism
  - Puzzles Lib
  - Random Shulker Colours
  - Realistic Bees
  - Rechiseled
  - Resourceful Config
  - Resourceful Lib
  - Responsive Shields
  - Searchables
  - SetiphianCore
  - Simple Discord RPC
  - Simple Voice Chat
  - Simply Swords
  - Size Shifting Potions
  - Snow Under Trees
  - Snow! Real Magic!
  - Sophisticated Backpack/Storage Emerald Upgrade
  - Sophisticated Backpacks
  - Sophisticated Core
  - Sophisticated Storage
  - Soul Fire'd
  - Storage Labels
  - SuperMartijn642's Config Lib
  - SuperMartijn642's Core Lib
  - Supplementaries
  - TerraBlender
  - Tiny Item Animations
  - Tips
  - Too Many Paintings!
  - U Team Core
  - Wavey Capes
  - Woodworks
  - YetAnotherConfigLib (YACL)
  - YUNG's API
  - YUNG's Better Desert Temples
  - YUNG's Better Dungeons
  - YUNG's Better End Island
  - YUNG's Better Jungle Temples
  - YUNG's Better Mineshafts
  - YUNG's Better Nether Fortresses
  - YUNG's Better Ocean Monuments
  - YUNG's Better Strongholds
  - YUNG's Better Witch Huts
  - YUNG's Bridges
- Added shaders:
  - BSL
  - Complementary Shaders - Reimagined
  - Complementary Shaders - Unbound
  - Photon Shaders
  - Rethinking Voxels
