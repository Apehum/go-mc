// Code generated by gen_soundid.go. DO NOT EDIT.

package soundid

// SoundID represents a sound ID used in the minecraft protocol.
type SoundID int32

// SoundNames - map of ids to names for sounds.
var SoundNames = map[SoundID]string{ 
	0: "ambient.cave",
	1: "ambient.basalt_deltas.additions",
	2: "ambient.basalt_deltas.loop",
	3: "ambient.basalt_deltas.mood",
	4: "ambient.crimson_forest.additions",
	5: "ambient.crimson_forest.loop",
	6: "ambient.crimson_forest.mood",
	7: "ambient.nether_wastes.additions",
	8: "ambient.nether_wastes.loop",
	9: "ambient.nether_wastes.mood",
	10: "ambient.soul_sand_valley.additions",
	11: "ambient.soul_sand_valley.loop",
	12: "ambient.soul_sand_valley.mood",
	13: "ambient.warped_forest.additions",
	14: "ambient.warped_forest.loop",
	15: "ambient.warped_forest.mood",
	16: "ambient.underwater.enter",
	17: "ambient.underwater.exit",
	18: "ambient.underwater.loop",
	19: "ambient.underwater.loop.additions",
	20: "ambient.underwater.loop.additions.rare",
	21: "ambient.underwater.loop.additions.ultra_rare",
	22: "block.amethyst_block.break",
	23: "block.amethyst_block.chime",
	24: "block.amethyst_block.fall",
	25: "block.amethyst_block.hit",
	26: "block.amethyst_block.place",
	27: "block.amethyst_block.step",
	28: "block.amethyst_cluster.break",
	29: "block.amethyst_cluster.fall",
	30: "block.amethyst_cluster.hit",
	31: "block.amethyst_cluster.place",
	32: "block.amethyst_cluster.step",
	33: "block.ancient_debris.break",
	34: "block.ancient_debris.step",
	35: "block.ancient_debris.place",
	36: "block.ancient_debris.hit",
	37: "block.ancient_debris.fall",
	38: "block.anvil.break",
	39: "block.anvil.destroy",
	40: "block.anvil.fall",
	41: "block.anvil.hit",
	42: "block.anvil.land",
	43: "block.anvil.place",
	44: "block.anvil.step",
	45: "block.anvil.use",
	46: "item.armor.equip_chain",
	47: "item.armor.equip_diamond",
	48: "item.armor.equip_elytra",
	49: "item.armor.equip_generic",
	50: "item.armor.equip_gold",
	51: "item.armor.equip_iron",
	52: "item.armor.equip_leather",
	53: "item.armor.equip_netherite",
	54: "item.armor.equip_turtle",
	55: "entity.armor_stand.break",
	56: "entity.armor_stand.fall",
	57: "entity.armor_stand.hit",
	58: "entity.armor_stand.place",
	59: "entity.arrow.hit",
	60: "entity.arrow.hit_player",
	61: "entity.arrow.shoot",
	62: "item.axe.strip",
	63: "item.axe.scrape",
	64: "item.axe.wax_off",
	65: "entity.axolotl.attack",
	66: "entity.axolotl.death",
	67: "entity.axolotl.hurt",
	68: "entity.axolotl.idle_air",
	69: "entity.axolotl.idle_water",
	70: "entity.axolotl.splash",
	71: "entity.axolotl.swim",
	72: "block.azalea.break",
	73: "block.azalea.fall",
	74: "block.azalea.hit",
	75: "block.azalea.place",
	76: "block.azalea.step",
	77: "block.azalea_leaves.break",
	78: "block.azalea_leaves.fall",
	79: "block.azalea_leaves.hit",
	80: "block.azalea_leaves.place",
	81: "block.azalea_leaves.step",
	82: "block.bamboo.break",
	83: "block.bamboo.fall",
	84: "block.bamboo.hit",
	85: "block.bamboo.place",
	86: "block.bamboo.step",
	87: "block.bamboo_sapling.break",
	88: "block.bamboo_sapling.hit",
	89: "block.bamboo_sapling.place",
	90: "block.barrel.close",
	91: "block.barrel.open",
	92: "block.basalt.break",
	93: "block.basalt.step",
	94: "block.basalt.place",
	95: "block.basalt.hit",
	96: "block.basalt.fall",
	97: "entity.bat.ambient",
	98: "entity.bat.death",
	99: "entity.bat.hurt",
	100: "entity.bat.loop",
	101: "entity.bat.takeoff",
	102: "block.beacon.activate",
	103: "block.beacon.ambient",
	104: "block.beacon.deactivate",
	105: "block.beacon.power_select",
	106: "entity.bee.death",
	107: "entity.bee.hurt",
	108: "entity.bee.loop_aggressive",
	109: "entity.bee.loop",
	110: "entity.bee.sting",
	111: "entity.bee.pollinate",
	112: "block.beehive.drip",
	113: "block.beehive.enter",
	114: "block.beehive.exit",
	115: "block.beehive.shear",
	116: "block.beehive.work",
	117: "block.bell.use",
	118: "block.bell.resonate",
	119: "block.big_dripleaf.break",
	120: "block.big_dripleaf.fall",
	121: "block.big_dripleaf.hit",
	122: "block.big_dripleaf.place",
	123: "block.big_dripleaf.step",
	124: "entity.blaze.ambient",
	125: "entity.blaze.burn",
	126: "entity.blaze.death",
	127: "entity.blaze.hurt",
	128: "entity.blaze.shoot",
	129: "entity.boat.paddle_land",
	130: "entity.boat.paddle_water",
	131: "block.bone_block.break",
	132: "block.bone_block.fall",
	133: "block.bone_block.hit",
	134: "block.bone_block.place",
	135: "block.bone_block.step",
	136: "item.bone_meal.use",
	137: "item.book.page_turn",
	138: "item.book.put",
	139: "block.blastfurnace.fire_crackle",
	140: "item.bottle.empty",
	141: "item.bottle.fill",
	142: "item.bottle.fill_dragonbreath",
	143: "block.brewing_stand.brew",
	144: "block.bubble_column.bubble_pop",
	145: "block.bubble_column.upwards_ambient",
	146: "block.bubble_column.upwards_inside",
	147: "block.bubble_column.whirlpool_ambient",
	148: "block.bubble_column.whirlpool_inside",
	149: "item.bucket.empty",
	150: "item.bucket.empty_axolotl",
	151: "item.bucket.empty_fish",
	152: "item.bucket.empty_lava",
	153: "item.bucket.empty_powder_snow",
	154: "item.bucket.fill",
	155: "item.bucket.fill_axolotl",
	156: "item.bucket.fill_fish",
	157: "item.bucket.fill_lava",
	158: "item.bucket.fill_powder_snow",
	159: "block.cake.add_candle",
	160: "block.calcite.break",
	161: "block.calcite.step",
	162: "block.calcite.place",
	163: "block.calcite.hit",
	164: "block.calcite.fall",
	165: "block.campfire.crackle",
	166: "block.candle.ambient",
	167: "block.candle.break",
	168: "block.candle.extinguish",
	169: "block.candle.fall",
	170: "block.candle.hit",
	171: "block.candle.place",
	172: "block.candle.step",
	173: "entity.cat.ambient",
	174: "entity.cat.stray_ambient",
	175: "entity.cat.death",
	176: "entity.cat.eat",
	177: "entity.cat.hiss",
	178: "entity.cat.beg_for_food",
	179: "entity.cat.hurt",
	180: "entity.cat.purr",
	181: "entity.cat.purreow",
	182: "block.cave_vines.break",
	183: "block.cave_vines.fall",
	184: "block.cave_vines.hit",
	185: "block.cave_vines.place",
	186: "block.cave_vines.step",
	187: "block.cave_vines.pick_berries",
	188: "block.chain.break",
	189: "block.chain.fall",
	190: "block.chain.hit",
	191: "block.chain.place",
	192: "block.chain.step",
	193: "block.chest.close",
	194: "block.chest.locked",
	195: "block.chest.open",
	196: "entity.chicken.ambient",
	197: "entity.chicken.death",
	198: "entity.chicken.egg",
	199: "entity.chicken.hurt",
	200: "entity.chicken.step",
	201: "block.chorus_flower.death",
	202: "block.chorus_flower.grow",
	203: "item.chorus_fruit.teleport",
	204: "entity.cod.ambient",
	205: "entity.cod.death",
	206: "entity.cod.flop",
	207: "entity.cod.hurt",
	208: "block.comparator.click",
	209: "block.composter.empty",
	210: "block.composter.fill",
	211: "block.composter.fill_success",
	212: "block.composter.ready",
	213: "block.conduit.activate",
	214: "block.conduit.ambient",
	215: "block.conduit.ambient.short",
	216: "block.conduit.attack.target",
	217: "block.conduit.deactivate",
	218: "block.copper.break",
	219: "block.copper.step",
	220: "block.copper.place",
	221: "block.copper.hit",
	222: "block.copper.fall",
	223: "block.coral_block.break",
	224: "block.coral_block.fall",
	225: "block.coral_block.hit",
	226: "block.coral_block.place",
	227: "block.coral_block.step",
	228: "entity.cow.ambient",
	229: "entity.cow.death",
	230: "entity.cow.hurt",
	231: "entity.cow.milk",
	232: "entity.cow.step",
	233: "entity.creeper.death",
	234: "entity.creeper.hurt",
	235: "entity.creeper.primed",
	236: "block.crop.break",
	237: "item.crop.plant",
	238: "item.crossbow.hit",
	239: "item.crossbow.loading_end",
	240: "item.crossbow.loading_middle",
	241: "item.crossbow.loading_start",
	242: "item.crossbow.quick_charge_1",
	243: "item.crossbow.quick_charge_2",
	244: "item.crossbow.quick_charge_3",
	245: "item.crossbow.shoot",
	246: "block.deepslate_bricks.break",
	247: "block.deepslate_bricks.fall",
	248: "block.deepslate_bricks.hit",
	249: "block.deepslate_bricks.place",
	250: "block.deepslate_bricks.step",
	251: "block.deepslate.break",
	252: "block.deepslate.fall",
	253: "block.deepslate.hit",
	254: "block.deepslate.place",
	255: "block.deepslate.step",
	256: "block.deepslate_tiles.break",
	257: "block.deepslate_tiles.fall",
	258: "block.deepslate_tiles.hit",
	259: "block.deepslate_tiles.place",
	260: "block.deepslate_tiles.step",
	261: "block.dispenser.dispense",
	262: "block.dispenser.fail",
	263: "block.dispenser.launch",
	264: "entity.dolphin.ambient",
	265: "entity.dolphin.ambient_water",
	266: "entity.dolphin.attack",
	267: "entity.dolphin.death",
	268: "entity.dolphin.eat",
	269: "entity.dolphin.hurt",
	270: "entity.dolphin.jump",
	271: "entity.dolphin.play",
	272: "entity.dolphin.splash",
	273: "entity.dolphin.swim",
	274: "entity.donkey.ambient",
	275: "entity.donkey.angry",
	276: "entity.donkey.chest",
	277: "entity.donkey.death",
	278: "entity.donkey.eat",
	279: "entity.donkey.hurt",
	280: "block.dripstone_block.break",
	281: "block.dripstone_block.step",
	282: "block.dripstone_block.place",
	283: "block.dripstone_block.hit",
	284: "block.dripstone_block.fall",
	285: "block.pointed_dripstone.break",
	286: "block.pointed_dripstone.step",
	287: "block.pointed_dripstone.place",
	288: "block.pointed_dripstone.hit",
	289: "block.pointed_dripstone.fall",
	290: "block.pointed_dripstone.land",
	291: "block.pointed_dripstone.drip_lava",
	292: "block.pointed_dripstone.drip_water",
	293: "block.pointed_dripstone.drip_lava_into_cauldron",
	294: "block.pointed_dripstone.drip_water_into_cauldron",
	295: "block.big_dripleaf.tilt_down",
	296: "block.big_dripleaf.tilt_up",
	297: "entity.drowned.ambient",
	298: "entity.drowned.ambient_water",
	299: "entity.drowned.death",
	300: "entity.drowned.death_water",
	301: "entity.drowned.hurt",
	302: "entity.drowned.hurt_water",
	303: "entity.drowned.shoot",
	304: "entity.drowned.step",
	305: "entity.drowned.swim",
	306: "item.dye.use",
	307: "entity.egg.throw",
	308: "entity.elder_guardian.ambient",
	309: "entity.elder_guardian.ambient_land",
	310: "entity.elder_guardian.curse",
	311: "entity.elder_guardian.death",
	312: "entity.elder_guardian.death_land",
	313: "entity.elder_guardian.flop",
	314: "entity.elder_guardian.hurt",
	315: "entity.elder_guardian.hurt_land",
	316: "item.elytra.flying",
	317: "block.enchantment_table.use",
	318: "block.ender_chest.close",
	319: "block.ender_chest.open",
	320: "entity.ender_dragon.ambient",
	321: "entity.ender_dragon.death",
	322: "entity.dragon_fireball.explode",
	323: "entity.ender_dragon.flap",
	324: "entity.ender_dragon.growl",
	325: "entity.ender_dragon.hurt",
	326: "entity.ender_dragon.shoot",
	327: "entity.ender_eye.death",
	328: "entity.ender_eye.launch",
	329: "entity.enderman.ambient",
	330: "entity.enderman.death",
	331: "entity.enderman.hurt",
	332: "entity.enderman.scream",
	333: "entity.enderman.stare",
	334: "entity.enderman.teleport",
	335: "entity.endermite.ambient",
	336: "entity.endermite.death",
	337: "entity.endermite.hurt",
	338: "entity.endermite.step",
	339: "entity.ender_pearl.throw",
	340: "block.end_gateway.spawn",
	341: "block.end_portal_frame.fill",
	342: "block.end_portal.spawn",
	343: "entity.evoker.ambient",
	344: "entity.evoker.cast_spell",
	345: "entity.evoker.celebrate",
	346: "entity.evoker.death",
	347: "entity.evoker_fangs.attack",
	348: "entity.evoker.hurt",
	349: "entity.evoker.prepare_attack",
	350: "entity.evoker.prepare_summon",
	351: "entity.evoker.prepare_wololo",
	352: "entity.experience_bottle.throw",
	353: "entity.experience_orb.pickup",
	354: "block.fence_gate.close",
	355: "block.fence_gate.open",
	356: "item.firecharge.use",
	357: "entity.firework_rocket.blast",
	358: "entity.firework_rocket.blast_far",
	359: "entity.firework_rocket.large_blast",
	360: "entity.firework_rocket.large_blast_far",
	361: "entity.firework_rocket.launch",
	362: "entity.firework_rocket.shoot",
	363: "entity.firework_rocket.twinkle",
	364: "entity.firework_rocket.twinkle_far",
	365: "block.fire.ambient",
	366: "block.fire.extinguish",
	367: "entity.fish.swim",
	368: "entity.fishing_bobber.retrieve",
	369: "entity.fishing_bobber.splash",
	370: "entity.fishing_bobber.throw",
	371: "item.flintandsteel.use",
	372: "block.flowering_azalea.break",
	373: "block.flowering_azalea.fall",
	374: "block.flowering_azalea.hit",
	375: "block.flowering_azalea.place",
	376: "block.flowering_azalea.step",
	377: "entity.fox.aggro",
	378: "entity.fox.ambient",
	379: "entity.fox.bite",
	380: "entity.fox.death",
	381: "entity.fox.eat",
	382: "entity.fox.hurt",
	383: "entity.fox.screech",
	384: "entity.fox.sleep",
	385: "entity.fox.sniff",
	386: "entity.fox.spit",
	387: "entity.fox.teleport",
	388: "block.roots.break",
	389: "block.roots.step",
	390: "block.roots.place",
	391: "block.roots.hit",
	392: "block.roots.fall",
	393: "block.furnace.fire_crackle",
	394: "entity.generic.big_fall",
	395: "entity.generic.burn",
	396: "entity.generic.death",
	397: "entity.generic.drink",
	398: "entity.generic.eat",
	399: "entity.generic.explode",
	400: "entity.generic.extinguish_fire",
	401: "entity.generic.hurt",
	402: "entity.generic.small_fall",
	403: "entity.generic.splash",
	404: "entity.generic.swim",
	405: "entity.ghast.ambient",
	406: "entity.ghast.death",
	407: "entity.ghast.hurt",
	408: "entity.ghast.scream",
	409: "entity.ghast.shoot",
	410: "entity.ghast.warn",
	411: "block.gilded_blackstone.break",
	412: "block.gilded_blackstone.fall",
	413: "block.gilded_blackstone.hit",
	414: "block.gilded_blackstone.place",
	415: "block.gilded_blackstone.step",
	416: "block.glass.break",
	417: "block.glass.fall",
	418: "block.glass.hit",
	419: "block.glass.place",
	420: "block.glass.step",
	421: "item.glow_ink_sac.use",
	422: "entity.glow_item_frame.add_item",
	423: "entity.glow_item_frame.break",
	424: "entity.glow_item_frame.place",
	425: "entity.glow_item_frame.remove_item",
	426: "entity.glow_item_frame.rotate_item",
	427: "entity.glow_squid.ambient",
	428: "entity.glow_squid.death",
	429: "entity.glow_squid.hurt",
	430: "entity.glow_squid.squirt",
	431: "entity.goat.ambient",
	432: "entity.goat.death",
	433: "entity.goat.eat",
	434: "entity.goat.hurt",
	435: "entity.goat.long_jump",
	436: "entity.goat.milk",
	437: "entity.goat.prepare_ram",
	438: "entity.goat.ram_impact",
	439: "entity.goat.screaming.ambient",
	440: "entity.goat.screaming.death",
	441: "entity.goat.screaming.eat",
	442: "entity.goat.screaming.hurt",
	443: "entity.goat.screaming.long_jump",
	444: "entity.goat.screaming.milk",
	445: "entity.goat.screaming.prepare_ram",
	446: "entity.goat.screaming.ram_impact",
	447: "entity.goat.step",
	448: "block.grass.break",
	449: "block.grass.fall",
	450: "block.grass.hit",
	451: "block.grass.place",
	452: "block.grass.step",
	453: "block.gravel.break",
	454: "block.gravel.fall",
	455: "block.gravel.hit",
	456: "block.gravel.place",
	457: "block.gravel.step",
	458: "block.grindstone.use",
	459: "entity.guardian.ambient",
	460: "entity.guardian.ambient_land",
	461: "entity.guardian.attack",
	462: "entity.guardian.death",
	463: "entity.guardian.death_land",
	464: "entity.guardian.flop",
	465: "entity.guardian.hurt",
	466: "entity.guardian.hurt_land",
	467: "block.hanging_roots.break",
	468: "block.hanging_roots.fall",
	469: "block.hanging_roots.hit",
	470: "block.hanging_roots.place",
	471: "block.hanging_roots.step",
	472: "item.hoe.till",
	473: "entity.hoglin.ambient",
	474: "entity.hoglin.angry",
	475: "entity.hoglin.attack",
	476: "entity.hoglin.converted_to_zombified",
	477: "entity.hoglin.death",
	478: "entity.hoglin.hurt",
	479: "entity.hoglin.retreat",
	480: "entity.hoglin.step",
	481: "block.honey_block.break",
	482: "block.honey_block.fall",
	483: "block.honey_block.hit",
	484: "block.honey_block.place",
	485: "block.honey_block.slide",
	486: "block.honey_block.step",
	487: "item.honeycomb.wax_on",
	488: "item.honey_bottle.drink",
	489: "entity.horse.ambient",
	490: "entity.horse.angry",
	491: "entity.horse.armor",
	492: "entity.horse.breathe",
	493: "entity.horse.death",
	494: "entity.horse.eat",
	495: "entity.horse.gallop",
	496: "entity.horse.hurt",
	497: "entity.horse.jump",
	498: "entity.horse.land",
	499: "entity.horse.saddle",
	500: "entity.horse.step",
	501: "entity.horse.step_wood",
	502: "entity.hostile.big_fall",
	503: "entity.hostile.death",
	504: "entity.hostile.hurt",
	505: "entity.hostile.small_fall",
	506: "entity.hostile.splash",
	507: "entity.hostile.swim",
	508: "entity.husk.ambient",
	509: "entity.husk.converted_to_zombie",
	510: "entity.husk.death",
	511: "entity.husk.hurt",
	512: "entity.husk.step",
	513: "entity.illusioner.ambient",
	514: "entity.illusioner.cast_spell",
	515: "entity.illusioner.death",
	516: "entity.illusioner.hurt",
	517: "entity.illusioner.mirror_move",
	518: "entity.illusioner.prepare_blindness",
	519: "entity.illusioner.prepare_mirror",
	520: "item.ink_sac.use",
	521: "block.iron_door.close",
	522: "block.iron_door.open",
	523: "entity.iron_golem.attack",
	524: "entity.iron_golem.damage",
	525: "entity.iron_golem.death",
	526: "entity.iron_golem.hurt",
	527: "entity.iron_golem.repair",
	528: "entity.iron_golem.step",
	529: "block.iron_trapdoor.close",
	530: "block.iron_trapdoor.open",
	531: "entity.item_frame.add_item",
	532: "entity.item_frame.break",
	533: "entity.item_frame.place",
	534: "entity.item_frame.remove_item",
	535: "entity.item_frame.rotate_item",
	536: "entity.item.break",
	537: "entity.item.pickup",
	538: "block.ladder.break",
	539: "block.ladder.fall",
	540: "block.ladder.hit",
	541: "block.ladder.place",
	542: "block.ladder.step",
	543: "block.lantern.break",
	544: "block.lantern.fall",
	545: "block.lantern.hit",
	546: "block.lantern.place",
	547: "block.lantern.step",
	548: "block.large_amethyst_bud.break",
	549: "block.large_amethyst_bud.place",
	550: "block.lava.ambient",
	551: "block.lava.extinguish",
	552: "block.lava.pop",
	553: "entity.leash_knot.break",
	554: "entity.leash_knot.place",
	555: "block.lever.click",
	556: "entity.lightning_bolt.impact",
	557: "entity.lightning_bolt.thunder",
	558: "entity.lingering_potion.throw",
	559: "entity.llama.ambient",
	560: "entity.llama.angry",
	561: "entity.llama.chest",
	562: "entity.llama.death",
	563: "entity.llama.eat",
	564: "entity.llama.hurt",
	565: "entity.llama.spit",
	566: "entity.llama.step",
	567: "entity.llama.swag",
	568: "entity.magma_cube.death_small",
	569: "block.lodestone.break",
	570: "block.lodestone.step",
	571: "block.lodestone.place",
	572: "block.lodestone.hit",
	573: "block.lodestone.fall",
	574: "item.lodestone_compass.lock",
	575: "entity.magma_cube.death",
	576: "entity.magma_cube.hurt",
	577: "entity.magma_cube.hurt_small",
	578: "entity.magma_cube.jump",
	579: "entity.magma_cube.squish",
	580: "entity.magma_cube.squish_small",
	581: "block.medium_amethyst_bud.break",
	582: "block.medium_amethyst_bud.place",
	583: "block.metal.break",
	584: "block.metal.fall",
	585: "block.metal.hit",
	586: "block.metal.place",
	587: "block.metal_pressure_plate.click_off",
	588: "block.metal_pressure_plate.click_on",
	589: "block.metal.step",
	590: "entity.minecart.inside.underwater",
	591: "entity.minecart.inside",
	592: "entity.minecart.riding",
	593: "entity.mooshroom.convert",
	594: "entity.mooshroom.eat",
	595: "entity.mooshroom.milk",
	596: "entity.mooshroom.suspicious_milk",
	597: "entity.mooshroom.shear",
	598: "block.moss_carpet.break",
	599: "block.moss_carpet.fall",
	600: "block.moss_carpet.hit",
	601: "block.moss_carpet.place",
	602: "block.moss_carpet.step",
	603: "block.moss.break",
	604: "block.moss.fall",
	605: "block.moss.hit",
	606: "block.moss.place",
	607: "block.moss.step",
	608: "entity.mule.ambient",
	609: "entity.mule.angry",
	610: "entity.mule.chest",
	611: "entity.mule.death",
	612: "entity.mule.eat",
	613: "entity.mule.hurt",
	614: "music.creative",
	615: "music.credits",
	616: "music_disc.11",
	617: "music_disc.13",
	618: "music_disc.blocks",
	619: "music_disc.cat",
	620: "music_disc.chirp",
	621: "music_disc.far",
	622: "music_disc.mall",
	623: "music_disc.mellohi",
	624: "music_disc.pigstep",
	625: "music_disc.stal",
	626: "music_disc.strad",
	627: "music_disc.wait",
	628: "music_disc.ward",
	629: "music.dragon",
	630: "music.end",
	631: "music.game",
	632: "music.menu",
	633: "music.nether.basalt_deltas",
	634: "music.nether.nether_wastes",
	635: "music.nether.soul_sand_valley",
	636: "music.nether.crimson_forest",
	637: "music.nether.warped_forest",
	638: "music.under_water",
	639: "block.nether_bricks.break",
	640: "block.nether_bricks.step",
	641: "block.nether_bricks.place",
	642: "block.nether_bricks.hit",
	643: "block.nether_bricks.fall",
	644: "block.nether_wart.break",
	645: "item.nether_wart.plant",
	646: "block.stem.break",
	647: "block.stem.step",
	648: "block.stem.place",
	649: "block.stem.hit",
	650: "block.stem.fall",
	651: "block.nylium.break",
	652: "block.nylium.step",
	653: "block.nylium.place",
	654: "block.nylium.hit",
	655: "block.nylium.fall",
	656: "block.nether_sprouts.break",
	657: "block.nether_sprouts.step",
	658: "block.nether_sprouts.place",
	659: "block.nether_sprouts.hit",
	660: "block.nether_sprouts.fall",
	661: "block.fungus.break",
	662: "block.fungus.step",
	663: "block.fungus.place",
	664: "block.fungus.hit",
	665: "block.fungus.fall",
	666: "block.weeping_vines.break",
	667: "block.weeping_vines.step",
	668: "block.weeping_vines.place",
	669: "block.weeping_vines.hit",
	670: "block.weeping_vines.fall",
	671: "block.wart_block.break",
	672: "block.wart_block.step",
	673: "block.wart_block.place",
	674: "block.wart_block.hit",
	675: "block.wart_block.fall",
	676: "block.netherite_block.break",
	677: "block.netherite_block.step",
	678: "block.netherite_block.place",
	679: "block.netherite_block.hit",
	680: "block.netherite_block.fall",
	681: "block.netherrack.break",
	682: "block.netherrack.step",
	683: "block.netherrack.place",
	684: "block.netherrack.hit",
	685: "block.netherrack.fall",
	686: "block.note_block.basedrum",
	687: "block.note_block.bass",
	688: "block.note_block.bell",
	689: "block.note_block.chime",
	690: "block.note_block.flute",
	691: "block.note_block.guitar",
	692: "block.note_block.harp",
	693: "block.note_block.hat",
	694: "block.note_block.pling",
	695: "block.note_block.snare",
	696: "block.note_block.xylophone",
	697: "block.note_block.iron_xylophone",
	698: "block.note_block.cow_bell",
	699: "block.note_block.didgeridoo",
	700: "block.note_block.bit",
	701: "block.note_block.banjo",
	702: "entity.ocelot.hurt",
	703: "entity.ocelot.ambient",
	704: "entity.ocelot.death",
	705: "entity.painting.break",
	706: "entity.painting.place",
	707: "entity.panda.pre_sneeze",
	708: "entity.panda.sneeze",
	709: "entity.panda.ambient",
	710: "entity.panda.death",
	711: "entity.panda.eat",
	712: "entity.panda.step",
	713: "entity.panda.cant_breed",
	714: "entity.panda.aggressive_ambient",
	715: "entity.panda.worried_ambient",
	716: "entity.panda.hurt",
	717: "entity.panda.bite",
	718: "entity.parrot.ambient",
	719: "entity.parrot.death",
	720: "entity.parrot.eat",
	721: "entity.parrot.fly",
	722: "entity.parrot.hurt",
	723: "entity.parrot.imitate.blaze",
	724: "entity.parrot.imitate.creeper",
	725: "entity.parrot.imitate.drowned",
	726: "entity.parrot.imitate.elder_guardian",
	727: "entity.parrot.imitate.ender_dragon",
	728: "entity.parrot.imitate.endermite",
	729: "entity.parrot.imitate.evoker",
	730: "entity.parrot.imitate.ghast",
	731: "entity.parrot.imitate.guardian",
	732: "entity.parrot.imitate.hoglin",
	733: "entity.parrot.imitate.husk",
	734: "entity.parrot.imitate.illusioner",
	735: "entity.parrot.imitate.magma_cube",
	736: "entity.parrot.imitate.phantom",
	737: "entity.parrot.imitate.piglin",
	738: "entity.parrot.imitate.piglin_brute",
	739: "entity.parrot.imitate.pillager",
	740: "entity.parrot.imitate.ravager",
	741: "entity.parrot.imitate.shulker",
	742: "entity.parrot.imitate.silverfish",
	743: "entity.parrot.imitate.skeleton",
	744: "entity.parrot.imitate.slime",
	745: "entity.parrot.imitate.spider",
	746: "entity.parrot.imitate.stray",
	747: "entity.parrot.imitate.vex",
	748: "entity.parrot.imitate.vindicator",
	749: "entity.parrot.imitate.witch",
	750: "entity.parrot.imitate.wither",
	751: "entity.parrot.imitate.wither_skeleton",
	752: "entity.parrot.imitate.zoglin",
	753: "entity.parrot.imitate.zombie",
	754: "entity.parrot.imitate.zombie_villager",
	755: "entity.parrot.step",
	756: "entity.phantom.ambient",
	757: "entity.phantom.bite",
	758: "entity.phantom.death",
	759: "entity.phantom.flap",
	760: "entity.phantom.hurt",
	761: "entity.phantom.swoop",
	762: "entity.pig.ambient",
	763: "entity.pig.death",
	764: "entity.pig.hurt",
	765: "entity.pig.saddle",
	766: "entity.pig.step",
	767: "entity.piglin.admiring_item",
	768: "entity.piglin.ambient",
	769: "entity.piglin.angry",
	770: "entity.piglin.celebrate",
	771: "entity.piglin.death",
	772: "entity.piglin.jealous",
	773: "entity.piglin.hurt",
	774: "entity.piglin.retreat",
	775: "entity.piglin.step",
	776: "entity.piglin.converted_to_zombified",
	777: "entity.piglin_brute.ambient",
	778: "entity.piglin_brute.angry",
	779: "entity.piglin_brute.death",
	780: "entity.piglin_brute.hurt",
	781: "entity.piglin_brute.step",
	782: "entity.piglin_brute.converted_to_zombified",
	783: "entity.pillager.ambient",
	784: "entity.pillager.celebrate",
	785: "entity.pillager.death",
	786: "entity.pillager.hurt",
	787: "block.piston.contract",
	788: "block.piston.extend",
	789: "entity.player.attack.crit",
	790: "entity.player.attack.knockback",
	791: "entity.player.attack.nodamage",
	792: "entity.player.attack.strong",
	793: "entity.player.attack.sweep",
	794: "entity.player.attack.weak",
	795: "entity.player.big_fall",
	796: "entity.player.breath",
	797: "entity.player.burp",
	798: "entity.player.death",
	799: "entity.player.hurt",
	800: "entity.player.hurt_drown",
	801: "entity.player.hurt_freeze",
	802: "entity.player.hurt_on_fire",
	803: "entity.player.hurt_sweet_berry_bush",
	804: "entity.player.levelup",
	805: "entity.player.small_fall",
	806: "entity.player.splash",
	807: "entity.player.splash.high_speed",
	808: "entity.player.swim",
	809: "entity.polar_bear.ambient",
	810: "entity.polar_bear.ambient_baby",
	811: "entity.polar_bear.death",
	812: "entity.polar_bear.hurt",
	813: "entity.polar_bear.step",
	814: "entity.polar_bear.warning",
	815: "block.polished_deepslate.break",
	816: "block.polished_deepslate.fall",
	817: "block.polished_deepslate.hit",
	818: "block.polished_deepslate.place",
	819: "block.polished_deepslate.step",
	820: "block.portal.ambient",
	821: "block.portal.travel",
	822: "block.portal.trigger",
	823: "block.powder_snow.break",
	824: "block.powder_snow.fall",
	825: "block.powder_snow.hit",
	826: "block.powder_snow.place",
	827: "block.powder_snow.step",
	828: "entity.puffer_fish.ambient",
	829: "entity.puffer_fish.blow_out",
	830: "entity.puffer_fish.blow_up",
	831: "entity.puffer_fish.death",
	832: "entity.puffer_fish.flop",
	833: "entity.puffer_fish.hurt",
	834: "entity.puffer_fish.sting",
	835: "block.pumpkin.carve",
	836: "entity.rabbit.ambient",
	837: "entity.rabbit.attack",
	838: "entity.rabbit.death",
	839: "entity.rabbit.hurt",
	840: "entity.rabbit.jump",
	841: "event.raid.horn",
	842: "entity.ravager.ambient",
	843: "entity.ravager.attack",
	844: "entity.ravager.celebrate",
	845: "entity.ravager.death",
	846: "entity.ravager.hurt",
	847: "entity.ravager.step",
	848: "entity.ravager.stunned",
	849: "entity.ravager.roar",
	850: "block.nether_gold_ore.break",
	851: "block.nether_gold_ore.fall",
	852: "block.nether_gold_ore.hit",
	853: "block.nether_gold_ore.place",
	854: "block.nether_gold_ore.step",
	855: "block.nether_ore.break",
	856: "block.nether_ore.fall",
	857: "block.nether_ore.hit",
	858: "block.nether_ore.place",
	859: "block.nether_ore.step",
	860: "block.redstone_torch.burnout",
	861: "block.respawn_anchor.ambient",
	862: "block.respawn_anchor.charge",
	863: "block.respawn_anchor.deplete",
	864: "block.respawn_anchor.set_spawn",
	865: "block.rooted_dirt.break",
	866: "block.rooted_dirt.fall",
	867: "block.rooted_dirt.hit",
	868: "block.rooted_dirt.place",
	869: "block.rooted_dirt.step",
	870: "entity.salmon.ambient",
	871: "entity.salmon.death",
	872: "entity.salmon.flop",
	873: "entity.salmon.hurt",
	874: "block.sand.break",
	875: "block.sand.fall",
	876: "block.sand.hit",
	877: "block.sand.place",
	878: "block.sand.step",
	879: "block.scaffolding.break",
	880: "block.scaffolding.fall",
	881: "block.scaffolding.hit",
	882: "block.scaffolding.place",
	883: "block.scaffolding.step",
	884: "block.sculk_sensor.clicking",
	885: "block.sculk_sensor.clicking_stop",
	886: "block.sculk_sensor.break",
	887: "block.sculk_sensor.fall",
	888: "block.sculk_sensor.hit",
	889: "block.sculk_sensor.place",
	890: "block.sculk_sensor.step",
	891: "entity.sheep.ambient",
	892: "entity.sheep.death",
	893: "entity.sheep.hurt",
	894: "entity.sheep.shear",
	895: "entity.sheep.step",
	896: "item.shield.block",
	897: "item.shield.break",
	898: "block.shroomlight.break",
	899: "block.shroomlight.step",
	900: "block.shroomlight.place",
	901: "block.shroomlight.hit",
	902: "block.shroomlight.fall",
	903: "item.shovel.flatten",
	904: "entity.shulker.ambient",
	905: "block.shulker_box.close",
	906: "block.shulker_box.open",
	907: "entity.shulker_bullet.hit",
	908: "entity.shulker_bullet.hurt",
	909: "entity.shulker.close",
	910: "entity.shulker.death",
	911: "entity.shulker.hurt",
	912: "entity.shulker.hurt_closed",
	913: "entity.shulker.open",
	914: "entity.shulker.shoot",
	915: "entity.shulker.teleport",
	916: "entity.silverfish.ambient",
	917: "entity.silverfish.death",
	918: "entity.silverfish.hurt",
	919: "entity.silverfish.step",
	920: "entity.skeleton.ambient",
	921: "entity.skeleton.converted_to_stray",
	922: "entity.skeleton.death",
	923: "entity.skeleton_horse.ambient",
	924: "entity.skeleton_horse.death",
	925: "entity.skeleton_horse.hurt",
	926: "entity.skeleton_horse.swim",
	927: "entity.skeleton_horse.ambient_water",
	928: "entity.skeleton_horse.gallop_water",
	929: "entity.skeleton_horse.jump_water",
	930: "entity.skeleton_horse.step_water",
	931: "entity.skeleton.hurt",
	932: "entity.skeleton.shoot",
	933: "entity.skeleton.step",
	934: "entity.slime.attack",
	935: "entity.slime.death",
	936: "entity.slime.hurt",
	937: "entity.slime.jump",
	938: "entity.slime.squish",
	939: "block.slime_block.break",
	940: "block.slime_block.fall",
	941: "block.slime_block.hit",
	942: "block.slime_block.place",
	943: "block.slime_block.step",
	944: "block.small_amethyst_bud.break",
	945: "block.small_amethyst_bud.place",
	946: "block.small_dripleaf.break",
	947: "block.small_dripleaf.fall",
	948: "block.small_dripleaf.hit",
	949: "block.small_dripleaf.place",
	950: "block.small_dripleaf.step",
	951: "block.soul_sand.break",
	952: "block.soul_sand.step",
	953: "block.soul_sand.place",
	954: "block.soul_sand.hit",
	955: "block.soul_sand.fall",
	956: "block.soul_soil.break",
	957: "block.soul_soil.step",
	958: "block.soul_soil.place",
	959: "block.soul_soil.hit",
	960: "block.soul_soil.fall",
	961: "particle.soul_escape",
	962: "block.spore_blossom.break",
	963: "block.spore_blossom.fall",
	964: "block.spore_blossom.hit",
	965: "block.spore_blossom.place",
	966: "block.spore_blossom.step",
	967: "entity.strider.ambient",
	968: "entity.strider.happy",
	969: "entity.strider.retreat",
	970: "entity.strider.death",
	971: "entity.strider.hurt",
	972: "entity.strider.step",
	973: "entity.strider.step_lava",
	974: "entity.strider.eat",
	975: "entity.strider.saddle",
	976: "entity.slime.death_small",
	977: "entity.slime.hurt_small",
	978: "entity.slime.jump_small",
	979: "entity.slime.squish_small",
	980: "block.smithing_table.use",
	981: "block.smoker.smoke",
	982: "entity.snowball.throw",
	983: "block.snow.break",
	984: "block.snow.fall",
	985: "entity.snow_golem.ambient",
	986: "entity.snow_golem.death",
	987: "entity.snow_golem.hurt",
	988: "entity.snow_golem.shoot",
	989: "entity.snow_golem.shear",
	990: "block.snow.hit",
	991: "block.snow.place",
	992: "block.snow.step",
	993: "entity.spider.ambient",
	994: "entity.spider.death",
	995: "entity.spider.hurt",
	996: "entity.spider.step",
	997: "entity.splash_potion.break",
	998: "entity.splash_potion.throw",
	999: "item.spyglass.use",
	1000: "item.spyglass.stop_using",
	1001: "entity.squid.ambient",
	1002: "entity.squid.death",
	1003: "entity.squid.hurt",
	1004: "entity.squid.squirt",
	1005: "block.stone.break",
	1006: "block.stone_button.click_off",
	1007: "block.stone_button.click_on",
	1008: "block.stone.fall",
	1009: "block.stone.hit",
	1010: "block.stone.place",
	1011: "block.stone_pressure_plate.click_off",
	1012: "block.stone_pressure_plate.click_on",
	1013: "block.stone.step",
	1014: "entity.stray.ambient",
	1015: "entity.stray.death",
	1016: "entity.stray.hurt",
	1017: "entity.stray.step",
	1018: "block.sweet_berry_bush.break",
	1019: "block.sweet_berry_bush.place",
	1020: "block.sweet_berry_bush.pick_berries",
	1021: "enchant.thorns.hit",
	1022: "entity.tnt.primed",
	1023: "item.totem.use",
	1024: "item.trident.hit",
	1025: "item.trident.hit_ground",
	1026: "item.trident.return",
	1027: "item.trident.riptide_1",
	1028: "item.trident.riptide_2",
	1029: "item.trident.riptide_3",
	1030: "item.trident.throw",
	1031: "item.trident.thunder",
	1032: "block.tripwire.attach",
	1033: "block.tripwire.click_off",
	1034: "block.tripwire.click_on",
	1035: "block.tripwire.detach",
	1036: "entity.tropical_fish.ambient",
	1037: "entity.tropical_fish.death",
	1038: "entity.tropical_fish.flop",
	1039: "entity.tropical_fish.hurt",
	1040: "block.tuff.break",
	1041: "block.tuff.step",
	1042: "block.tuff.place",
	1043: "block.tuff.hit",
	1044: "block.tuff.fall",
	1045: "entity.turtle.ambient_land",
	1046: "entity.turtle.death",
	1047: "entity.turtle.death_baby",
	1048: "entity.turtle.egg_break",
	1049: "entity.turtle.egg_crack",
	1050: "entity.turtle.egg_hatch",
	1051: "entity.turtle.hurt",
	1052: "entity.turtle.hurt_baby",
	1053: "entity.turtle.lay_egg",
	1054: "entity.turtle.shamble",
	1055: "entity.turtle.shamble_baby",
	1056: "entity.turtle.swim",
	1057: "ui.button.click",
	1058: "ui.loom.select_pattern",
	1059: "ui.loom.take_result",
	1060: "ui.cartography_table.take_result",
	1061: "ui.stonecutter.take_result",
	1062: "ui.stonecutter.select_recipe",
	1063: "ui.toast.challenge_complete",
	1064: "ui.toast.in",
	1065: "ui.toast.out",
	1066: "entity.vex.ambient",
	1067: "entity.vex.charge",
	1068: "entity.vex.death",
	1069: "entity.vex.hurt",
	1070: "entity.villager.ambient",
	1071: "entity.villager.celebrate",
	1072: "entity.villager.death",
	1073: "entity.villager.hurt",
	1074: "entity.villager.no",
	1075: "entity.villager.trade",
	1076: "entity.villager.yes",
	1077: "entity.villager.work_armorer",
	1078: "entity.villager.work_butcher",
	1079: "entity.villager.work_cartographer",
	1080: "entity.villager.work_cleric",
	1081: "entity.villager.work_farmer",
	1082: "entity.villager.work_fisherman",
	1083: "entity.villager.work_fletcher",
	1084: "entity.villager.work_leatherworker",
	1085: "entity.villager.work_librarian",
	1086: "entity.villager.work_mason",
	1087: "entity.villager.work_shepherd",
	1088: "entity.villager.work_toolsmith",
	1089: "entity.villager.work_weaponsmith",
	1090: "entity.vindicator.ambient",
	1091: "entity.vindicator.celebrate",
	1092: "entity.vindicator.death",
	1093: "entity.vindicator.hurt",
	1094: "block.vine.break",
	1095: "block.vine.fall",
	1096: "block.vine.hit",
	1097: "block.vine.place",
	1098: "block.vine.step",
	1099: "block.lily_pad.place",
	1100: "entity.wandering_trader.ambient",
	1101: "entity.wandering_trader.death",
	1102: "entity.wandering_trader.disappeared",
	1103: "entity.wandering_trader.drink_milk",
	1104: "entity.wandering_trader.drink_potion",
	1105: "entity.wandering_trader.hurt",
	1106: "entity.wandering_trader.no",
	1107: "entity.wandering_trader.reappeared",
	1108: "entity.wandering_trader.trade",
	1109: "entity.wandering_trader.yes",
	1110: "block.water.ambient",
	1111: "weather.rain",
	1112: "weather.rain.above",
	1113: "block.wet_grass.break",
	1114: "block.wet_grass.fall",
	1115: "block.wet_grass.hit",
	1116: "block.wet_grass.place",
	1117: "block.wet_grass.step",
	1118: "entity.witch.ambient",
	1119: "entity.witch.celebrate",
	1120: "entity.witch.death",
	1121: "entity.witch.drink",
	1122: "entity.witch.hurt",
	1123: "entity.witch.throw",
	1124: "entity.wither.ambient",
	1125: "entity.wither.break_block",
	1126: "entity.wither.death",
	1127: "entity.wither.hurt",
	1128: "entity.wither.shoot",
	1129: "entity.wither_skeleton.ambient",
	1130: "entity.wither_skeleton.death",
	1131: "entity.wither_skeleton.hurt",
	1132: "entity.wither_skeleton.step",
	1133: "entity.wither.spawn",
	1134: "entity.wolf.ambient",
	1135: "entity.wolf.death",
	1136: "entity.wolf.growl",
	1137: "entity.wolf.howl",
	1138: "entity.wolf.hurt",
	1139: "entity.wolf.pant",
	1140: "entity.wolf.shake",
	1141: "entity.wolf.step",
	1142: "entity.wolf.whine",
	1143: "block.wooden_door.close",
	1144: "block.wooden_door.open",
	1145: "block.wooden_trapdoor.close",
	1146: "block.wooden_trapdoor.open",
	1147: "block.wood.break",
	1148: "block.wooden_button.click_off",
	1149: "block.wooden_button.click_on",
	1150: "block.wood.fall",
	1151: "block.wood.hit",
	1152: "block.wood.place",
	1153: "block.wooden_pressure_plate.click_off",
	1154: "block.wooden_pressure_plate.click_on",
	1155: "block.wood.step",
	1156: "block.wool.break",
	1157: "block.wool.fall",
	1158: "block.wool.hit",
	1159: "block.wool.place",
	1160: "block.wool.step",
	1161: "entity.zoglin.ambient",
	1162: "entity.zoglin.angry",
	1163: "entity.zoglin.attack",
	1164: "entity.zoglin.death",
	1165: "entity.zoglin.hurt",
	1166: "entity.zoglin.step",
	1167: "entity.zombie.ambient",
	1168: "entity.zombie.attack_wooden_door",
	1169: "entity.zombie.attack_iron_door",
	1170: "entity.zombie.break_wooden_door",
	1171: "entity.zombie.converted_to_drowned",
	1172: "entity.zombie.death",
	1173: "entity.zombie.destroy_egg",
	1174: "entity.zombie_horse.ambient",
	1175: "entity.zombie_horse.death",
	1176: "entity.zombie_horse.hurt",
	1177: "entity.zombie.hurt",
	1178: "entity.zombie.infect",
	1179: "entity.zombified_piglin.ambient",
	1180: "entity.zombified_piglin.angry",
	1181: "entity.zombified_piglin.death",
	1182: "entity.zombified_piglin.hurt",
	1183: "entity.zombie.step",
	1184: "entity.zombie_villager.ambient",
	1185: "entity.zombie_villager.converted",
	1186: "entity.zombie_villager.cure",
	1187: "entity.zombie_villager.death",
	1188: "entity.zombie_villager.hurt",
	1189: "entity.zombie_villager.step",
}

// GetSoundNameByID helper method
func GetSoundNameByID(id SoundID) (string, bool) {
	name, ok := SoundNames[id]
	return name, ok
}