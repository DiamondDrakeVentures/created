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

- Added missing translations to tags.
- Added an overlay showing what other players are doing (inventory screen, pause menu, etc.) over their player models (What Are They Up To).
- Added a title overlay showing the current dimension and Waystone name when switching dimensions or navigating through Waystones (Traveler's Titles).
- Added mods:
  - Chorus Fruit Drops Nearby
  - CoroUtil
  - Legendary Tooltips
  - Recipe Fixer
  - Traveler's Titles
  - What Are They Up To

### Changed

- Updated NeoForge loader 21.1.203 => 21.1.205.
- Chorus Fruit now drops at the base of the Chorus Tree (Chorus Fruit Drops Nearby).
- Changed item tooltips for more immersion (Legendary Tooltips).
- Pump Upgrade now works with fluids in the world even when worn by player or in player inventory (Sophisticated Core).
- When poaching fails, the entity may transform into another entity now (My Nether's Delight).
- Striders now stay warm when they have the Pungency effect (My Nether's Delight).
- Tools now wear out faster if they have the Curse of Poaching (My Nether's Delight).
- Bullet Pepper drop rates are now lower (My Nether's Delight).
- Trophies now push entities when powered or moved by Pistons (My Nether's Delight).
- Recipes are adjusted to include Pepper Powder (My Nether's Delight).
- Liquids that boost the Leteos process now use their own tag (My Nether's Delight).
- Updated mods:
  - EMI Addon: Extra Mod Integrations 1.0.2 => 1.0.3
  - Ender Dragon Fight Remastered: Create(D) Edition 5.0.0-2cbec9c => 5.0.0-8ba9aa2
  - Kiwi 15.6.2 => 15.7.0
  - Moonlight Lib 2.22.2 => 2.22.6
  - My Nether's Delight 1.8 => 1.9
  - Sophisticated Core 1.3.64.1090 => 1.3.65.1096
  - Supplementaries 3.4.9 => 3.4.12

### Deprecated

### Removed

- Removed mods:
  - Animal Garden - Owl

### Fixed

- Fixed an issue where entities (including players) cannot move when in any depth of water ([#5](https://github.com/DiamondDrakeVentures/created/issues/5)).
- Fixed missing tags (Supplementaries).

### Security

## [0.3.0] - 2025-08-27

### Added

- Added Sepia Fungus Colony (Create: Garnished).
- Added Hypertube Accelerator (Create: Hypertubes).
- Added `polished` and `bricks` `StoneType`s (Every Compat (Stone Zone)).
- Added Forward Error Correction (FEC) to Opus encoder/decoder (Simple Voice Chat).
- Added Classical Chinese translation (Sophisticated Backpacks).
- Added HUD component showing the target blocks and entities (Jade).
- Added configuration menu (Ender Dragon Fight Remastered: Create(D) Edition).
- Added training dummy (MmmMmmMmmMmm).
- Added a recipe to demagnetize Lodestone Compass.
- Added Crushing and Milling recipes for White Sand.
- Added Crushing and Milling recipes for Orange Sand.
- Added Milling recipe for Black Sand.
- Added Ending recipe for Bucket of Dragon's Breath.
- Added Hydraulic Compacting recipe for processing Graphite.
- Added Hydraulic Compacting recipe for Diamond from Diamond Shards.
- Added a new ponder to factory gauge to explain the expanded recipes feature (Create: Extra Gauges).
- Added Integer Selector (Create: Extra Gauges).
- Added the String Gauge (Create: Extra Gauges).
- Added the Display Collector (Create: Extra Gauges).
- Added cache system configurable in client side config (Create: Extra Gauges).
- Added Passive Gauge (Create: Extra Gauges).
- Added recipe for Display Collector (Create: Extra Gauges).
- Added string connection to all panels (Create: Extra Gauges).
- Added an option to lock the body rotation for custom bow animation (Not Enough Animations).
- Added an overlay showing picked up items (Pick Up Notifier).
- Added a UI to Name Tag (Easy Anvil).
- Added `amendments:cauldron_crafting` recipe type (Amendments).
- Added Handheld Saw (Create: Dreams & Desires).
- Added Handheld Drill (Create: Dreams & Desires).
- Added keybinds for Handheld Drill (Create: Dreams & Desires).
- Added Golden Mixer (Create: Dreams & Desires).
- Added alternative recipe for Straw (Create Crafts & Additions).
- Added Amethyst Magnum Torch (Magnum Torch).
- Added Diamond Magnum Torch (Magnum Torch).
- Added Emerald Magnum Torch (Magnum Torch).
- Added mods:
  - Configurable
  - Create Waystones Recipes
  - Diagonal Fences
  - Diagonal Windows
  - Easy Anvil
  - Ender Dragon Fight Remastered: Create(D) Edition
  - Global Packs
  - Jade
  - Jade Addons
  - Leaves Be Gone
  - Magnum Torch
  - MmmMmmMmmMmm
  - Pick Up Notifier
  - Simple Custom Early Loading
- Added resource packs:
  - Create Waystones Texture + GUI

### Changed

- Updated NeoForge loader 21.1.196 => 21.1.203.
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
- Tweaked Milling and Crushing recipes for Red Sand.
- Tweaked Crushing recipes with Blackstone as ingredient.
- Tweaked Ending recipe for End Stone.
- Tweaked Compacting recipe for Graphite.
- Tweaked Crushing recipe for recyling Impure Diamond.
- Supplementaries Rope can now be crafted with Flax and Straws.
- Safety Net can now be crafted (uncrafted) into Supplementaries Ropes instead of Farmer's Delight Ropes.
- Farmer's Delight Rope and Supplementaries Rope can be converted to and from each other.
- Ornithopter Glider now requires Supplementaries Ropes.
- Farmer's Delight Rope is no longer a Catenary item.
- All Chains are now Catenary compatible.
- Mobs with Sophisticated Backpacks no longer plays jukebox music.
- Rebalanced armor (Aquatic Additions).
- Bulk Fermenters and Distillation Towers now only require one heat source (Create: Diesel Generators).
- Bulk Fermenters now can only hold up to 1,000 mB per tank per fluid type (Create: Diesel Generators).
- Modified some recipes (Create: Diesel Generators).
- Factory gagues can now craft larger recipes as long as there are less than 10 different items in the recipe (Create: Extra Gauges).
- Small recipe can match larger auto crafters if configured in the GUI (Create: Extra Gauges).
- Panels now update instantly with a configurable max update per tick (Create: Extra Gauges).
- Levers and analog levers now directly connect to gauges without redstone links, though redstone updates are not instantaneous (Create: Extra Gauges).
- Changed textures of Counter and Comparator Gauges (Create: Extra Gauges).
- Integer Gauge can now change the stock amount of a factory gauge (Create: Extra Gauges).
- Rich Soil Farmland can now be hydrated by any fluid type capable of hydration (Farmer's Delight).
- Wild Cabbage and Sea Beet can now generate on any biome tagged with `minecraft:is_beach` (Farmer's Delight).
- Cooking Pot can no longer be accessed in Spectator mode.
- Fences now connect diagonally (Diagonal Fences).
- Glass Panes now connect diagonally (Diagonal Windows).
- Iron Bars now connect diagonally (Diagonal Windows).
- Items now stay inside the Anvil (Easy Anvil).
- Anvil cost scale a lot less aggressively (Easy Anvil).
- Anvils are repairable with an Iron Block (Easy Anvil).
- Name Tags can be renamed at no cost (Easy Anvil).
- Name field now accepts formatting codes (Easy Anvil).
- Buntings can now be placed on walls and ceilings (Supplementaries).
- Lunch Basket and Sack can now use unique GUI textures (Supplementaries).
- Lunch Basket can now be dyed all colors similar to Leather items (Supplementaries).
- Bomb and Slingshot will now emit equally spaced particles (Supplementaries).
- Candle Holder recipes now use one Ingot instead of three (Supplementaries).
- Lunch Basket will now be set on a cooldown when eating an item with cooldown (Supplementaries).
- All liquids can now bubble (Amendments).
- All recipes now work both in-world and with hand interaction (Amendments).
- Improved Sign screen renderer (Amendments).
- Tweaked recipes (Create: Dreams & Desires).
- Adjusted Jetpack recipes to utilize Immersive Aircraft sail instead of Elytra.
- Updated mods:
  - Amendments 1.2.24 => 2.0.5
  - Aquatic Additions 4.0 => 4.1
  - Balm 21.0.48 => 21.0.49
  - Bookshelf 21.1.67 => 21.1.68
  - Chat Heads 0.13.18 => 0.13.20
  - Configurable 3.2.1 => 3.2.5
  - Create Crafts & Additions 1.5.1 => 1.5.2
  - Create: Diesel Generators 1.3.4 => 1.3.5
  - Create: Dreams & Desires 2.2a-BETA => 2.2c-BETA
  - Create: Extra Gauges 1.1.1 => 2.0.5
  - Create: Garnished 2.1.6.2 => 2.1.6.5
  - Create: Hypertubes 0.2.4 => 0.2.5
  - Create More: Parallel Pipes 1.1.0 => 1.1.1
  - Create: Peaceful 2.0.0 => 2.0.1
  - Create Train Parts 0.1.3 => 0.2.0
  - Enchantment Descriptions 21.1.7 => 21.1.8
  - Entity Model Features 2.4.1 => 3.0.1
  - Entity Texture Features 6.2.9 => 7.0.1
  - Every Compat (Stone Zone) 1.21-2.10.6 => 1.21-2.10.8
  - Geckolib 4.7.6 => 4.7.7
  - FancyMenu 3.6.4 => 3.7.0
  - Farmer's Delight 1.2.8 => 1.2.9
  - Fusion 1.2.10 => 1.2.11a
  - Friends&Foes 4.0.7 => 4.0.8
  - Fzzy Config 0.7.1+1.21+neoforge => 0.7.2+1.21+neoforge
  - Just Blahaj 2.0.3 => 2.0.5
  - Kiwi 15.6.1 => 15.6.2
  - KleeSlabs 21.1.6 => 21.1.7
  - KotlinLangForge 2.10.1-k2.2.0-3.0 => 2.10.5-k2.2.10-3.0
  - Moonlight Lib 2.20.7 => 2.22.2
  - Mysterious Mountain Lib 1.1.8 => 1.2.14
  - Neruina 2.3.1-beta.1 => 3.0.3
  - Not Enough Animations 1.10.1 => 1.10.2
  - Packet Fixer 3.1.4 => 3.2.0
  - Puzzles Lib v21.1.36 => v21.1.38
  - Simple Voice Chat 2.5.35 => 2.5.36
  - Sophisiticated Backpacks 3.24.19.1292 => 3.24.21.1314
  - Sophisticated Core 1.3.59.1062 => 1.3.64.1090
  - Sophisticated Storage 1.4.46.1216 => 1.4.50.1235
  - Supplementaries 3.3.5 => 3.4.9
  - Waystones 21.1.20 => 21.1.22
  - Wavey Capes 1.6.1.1 => 1.6.2

### Deprecated

- Disabled Dragon's Breath Bucket from Create: Garnished.
- Disabled Create: Ornithopter Glider Rope.

### Removed

- Removed duplicate Ending recipe for Chorus Fruit.
- Removed Ending recipe for Crying Obsidian.
- Removed duplicate Ending recipe for End Stone.
- Removed Ending recipe for Magma Cream.
- Removed Packing recipe for Diamond from Coal Blocks.
- Removed Shaped Crafting recipe for Diamond from Diamond Shards.
- Removed Compacting recipe for processing Graphite.
- Removed duplicate Trident recipe.
- Removed recipe for Create: Ornithopter Glider Rope.
- Removed recipe for Rope and Nail using Strings.
- Removed recipe for Farmer's Delight Rope using Straws.
- Removed mods:
  - Ender Dragon Fight Remastered
  - Multibeds
  - ShetiPhianCore
  - Paxi
  - Panda's Falling Trees
  - PandaLib

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
- Fixed Dragon's Breath Ending recipe.
- Fixed Bulk Fermenters voiding items (Create: Diesel Generators).
- Fixed modular engine crash on servers (Create: Diesel Generators).
- Fixed train engine sounds disappearing after reloading the world (Create: Diesel Generators).
- Fixed panels rendering/connections color issues (Create: Extra Gauges).
- Fixed issues with Factory Gauges (Create: Extra Gauges).
- Fixed an issue with large autocrafting when passing by a Re-packager (Create: Extra Gauges).
- Fixed Cooking Pot allowing players to cook and store up to 99 servings (Farmer's Delight).
- Fixed Create integration milling recipes using invalid field, causing them to never finish processing (Farmer's Delight).
- Fixed Rabbit Stew not returning bowls after being eating (Farmer's Delight).
- Fixed a potential datapack error with cooking recipes in KubeJS (Farmer's Delight).
- Fixed an issue with the eating animation (Not Enough Animations).
- Fixed an issue with animations when the left hand is the main hand (Not Enough Animations).
- Fixed Bunting recipes (Supplementaries).
- Fixed Wall Sign model being one pixel too tall (Amendments).
- Fixed an issue where Stew recipes gave less than intended (Amendments).
- Fixed an issue where Smart Hoppper void items (Create: Dreams & Desires).
- Fixed updating issue with Smart Hopper (Create: Dreams & Desires).
- Fixed Redstone Relay (Create Crafts & Additions).
- Fixed broken compatibility recipes (Create Crafts & Additions).

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
