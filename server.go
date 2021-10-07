package citizenfx

import "syscall/js"

/**
 * Adds a rectangular blip for the specified coordinates/area.
 * It is recommended to use [SET_BLIP_ROTATION](#\_0xF87683CDF73C3F6E) and [SET_BLIP_COLOUR](#\_0x03D7FB09E75D6B7E) to make the blip not rotate along with the camera.
 * By default, the blip will show as a *regular* blip with the specified color/sprite if it is outside of the minimap view.
 * Example image:
 * ![minimap](https://w.wew.wtf/pdcjig.png)
 * ![big map](https://w.wew.wtf/zgcjcm.png)
 * (Native name is *likely* to actually be ADD_BLIP_FOR_AREA, but due to the usual reasons this can't be confirmed)
 * @param x The X coordinate of the center of the blip.
 * @param y The Y coordinate of the center of the blip.
 * @param z The Z coordinate of the center of the blip.
 * @param width The width of the blip.
 * @param height The height of the blip.
 * @return A handle to the blip.
 */
func (_server) AddBlipForArea(x float64, y float64, z float64, width float64, height float64) float64 {
	return 0.0
}

/**
 * Creates a blip for the specified coordinates. You can use `SET_BLIP_` natives to change the blip.
 * @param x The X coordinate to create the blip on.
 * @param y The Y coordinate.
 * @param z The Z coordinate.
 * @return A blip handle.
 */
func (_server) AddBlipForCoord(x float64, y float64, z float64) float64 {
	return 0.0
}

/**
 * Returns red ( default ) blip attached to entity.
 * Example:
 * Blip blip; //Put this outside your case or option
 * blip = HUD::ADD_BLIP_FOR_ENTITY(YourPedOrBodyguardName);
 * HUD::SET_BLIP_AS_FRIENDLY(blip, true);
 */
func (_server) AddBlipForEntity(entity float64) float64 {
	return 0.0
}

func (_server) AddBlipForRadius(posX float64, posY float64, posZ float64, radius float64) float64 {
	return 0.0
}

/**
 * Applies an Item from a PedDecorationCollection to a ped. These include tattoos and shirt decals.
 * collection - PedDecorationCollection filename hash
 * overlay - Item name hash
 * Example:
 * Entry inside "mpbeach_overlays.xml" -
 * <Item>
 * <uvPos x="0.500000" y="0.500000" />
 * <scale x="0.600000" y="0.500000" />
 * <rotation value="0.000000" />
 * <nameHash>FM_Hair_Fuzz</nameHash>
 * <txdHash>mp_hair_fuzz</txdHash>
 * <txtHash>mp_hair_fuzz</txtHash>
 * <zone>ZONE_HEAD</zone>
 * <type>TYPE_TATTOO</type>
 * <faction>FM</faction>
 * <garment>All</garment>
 * <gender>GENDER_DONTCARE</gender>
 * <award />
 * <awardLevel />
 * </Item>
 * Code:
 * PED::_0x5F5D1665E352A839(PLAYER::PLAYER_PED_ID(), MISC::GET_HASH_KEY("mpbeach_overlays"), MISC::GET_HASH_KEY("fm_hair_fuzz"))
 */
func (_server) AddPedDecorationFromHashes(ped float64, collection interface{}, overlay interface{}) {}

/**
 * Adds a handler for changes to a state bag.
 * The function called expects to match the following signature:
 * ```ts
 * function StateBagChangeHandler(bagName string, key string, value  interface{} {
	 return 0.0 }, reserved float64, replicated bool);
 * ```
 * *   **bagName**: The internal bag ID for the state bag which changed. This is usually `player:Source`, `entity:NetID`
 * or `localEntity:Handle`.
 * *   **key**: The changed key.
 * *   **value**: The new value stored at key. The old value is still stored in the state bag at the time this callback executes.
 * *   **reserved**: Currently unused.
 * *   **replicated**: Whether the set is meant to be replicated.
 * At this time, the change handler can't opt to reject changes.
 * @param keyFilter The key to check for, or null.
 * @param bagFilter The bag ID to check for, or null.
 * @param handler The handler function.
 * @return A cookie to remove the change handler.
*/
func (_server) AddStateBagChangeHandler(keyFilter string, bagFilter string, handler js.Func) float64 {
	return 0.0
}

/**
 * Applies a force to the specified entity.
 * **List of force types (p1)**:
 * ```
 * public enum ForceType
 * {
 * MinForce = 0,
 * MaxForceRot = 1,
 * MinForce2 = 2,
 * MaxForceRot2 = 3,
 * ForceNoRot = 4,
 * ForceRotPlusForce = 5
 * }
 * ```
 * Research/documentation on the gtaforums can be found [here](https://gtaforums.com/topic/885669-precisely-define-object-physics/) and [here](https://gtaforums.com/topic/887362-apply-forces-and-momentums-to-entityobject/).
 * @param entity The entity you want to apply a force on
 * @param forceType See native description above for a list of commonly used values
 * @param x Force amount (X)
 * @param y Force amount (Y)
 * @param z Force amount (Z)
 * @param offX Rotation/offset force (X)
 * @param offY Rotation/offset force (Y)
 * @param offZ Rotation/offset force (Z)
 * @param boneIndex (Often 0) Entity bone index
 * @param isDirectionRel (Usually false) Vector defined in local (body-fixed) coordinate frame
 * @param ignoreUpVec (Usually true)
 * @param isForceRel (Usually true) When true, force gets multiplied with the objects mass and different objects will have the same acceleration
 * @param p12 (Usually false)
 * @param p13 (Usually true)
 */
func (_server) ApplyForceToEntity(entity float64, forceType float64, x float64, y float64, z float64, offX float64, offY float64, offZ float64, boneIndex float64, isDirectionRel bool, ignoreUpVec bool, isForceRel bool, p12 bool, p13 bool) {
}

/**
 * Returns whether or not the specified player has enough information to start a commerce session for.
 * @param playerSrc The player handle
 * @return True or false.
 */
func (_server) CanPlayerStartCommerceSession(playerSrc string) bool

/**
 * Cancels the currently executing event.
 */
func (_server) CancelEvent() {}

/**
 * List of component/props ID
 * gtaxscripting.blogspot.com/2016/04/gta-v-peds-component-and-props.html
 */
func (_server) ClearPedProp(ped float64, propId float64) {}

func (_server) ClearPedSecondaryTask(ped float64) {}

func (_server) ClearPedTasks(ped float64) {}

/**
 * Immediately stops the pedestrian from whatever it's doing. They stop fighting, animations, etc. they forget what they were doing.
 */
func (_server) ClearPedTasksImmediately(ped float64) {}

/**
 * This executes at the same as speed as PLAYER::SET_PLAYER_WANTED_LEVEL(player, 0, false);
 * PLAYER::GET_PLAYER_WANTED_LEVEL(player); executes in less than half the time. Which means that it's worth first checking if the wanted level needs to be cleared before clearing. However, this is mostly about good code practice and can important in other situations. The difference in time in this example is negligible.
 */
func (_server) ClearPlayerWantedLevel(player string) {}

/**
 * Creates an object (prop) with the specified model at the specified position, offset on the Z axis by the radius of the object's model.
 * This object will initially be owned by the creating script as a mission entity, and the model should be loaded already (e.g. using REQUEST_MODEL).
 * @param modelHash The model to spawn.
 * @param x Spawn coordinate X component.
 * @param y Spawn coordinate Y component.
 * @param z Spawn coordinate Z component, 'ground level'.
 * @param isNetwork Whether to create a network object for the object. If false, the object exists only locally.
 * @param netMissionEntity Whether to register the object as pinned to the script host in the R\* network model.
 * @param doorFlag False to create a door archetype (archetype flag bit 26 set) as a door. Required to be set to true to create door models in network mode.
 * @return A script handle (fwScriptGuid index) for the object, or `0` if the object failed to be created.
 */
func (_server) CreateObject(modelHash interface{}, x float64, y float64, z float64, isNetwork bool, netMissionEntity bool, doorFlag bool) float64 {
	return 0.0
}

/**
 * Creates an object (prop) with the specified model centered at the specified position.
 * This object will initially be owned by the creating script as a mission entity, and the model should be loaded already (e.g. using REQUEST_MODEL).
 * @param modelHash The model to spawn.
 * @param x Spawn coordinate X component.
 * @param y Spawn coordinate Y component.
 * @param z Spawn coordinate Z component.
 * @param isNetwork Whether to create a network object for the object. If false, the object exists only locally.
 * @param netMissionEntity Whether to register the object as pinned to the script host in the R\* network model.
 * @param doorFlag False to create a door archetype (archetype flag bit 26 set) as a door. Required to be set to true to create door models in network mode.
 * @return A script handle (fwScriptGuid index) for the object, or `0` if the object failed to be created.
 */
func (_server) CreateObjectNoOffset(modelHash interface{}, x float64, y float64, z float64, isNetwork bool, netMissionEntity bool, doorFlag bool) float64 {
	return 0.0
}

/**
 * Creates a ped (biped character, pedestrian, actor) with the specified model at the specified position and heading.
 * This ped will initially be owned by the creating script as a mission entity, and the model should be loaded already
 * (e.g. using REQUEST_MODEL).
 * @param pedType Unused. Peds get set to CIVMALE/CIVFEMALE/etc. no matter the value specified.
 * @param modelHash The model of ped to spawn.
 * @param x Spawn coordinate X component.
 * @param y Spawn coordinate Y component.
 * @param z Spawn coordinate Z component.
 * @param heading Heading to face towards, in degrees.
 * @param isNetwork Whether to create a network object for the ped. If false, the ped exists only locally.
 * @param bScriptHostPed Whether to register the ped as pinned to the script host in the R\* network model.
 * @return A script handle (fwScriptGuid index) for the ped, or `0` if the ped failed to be created.
 */
func (_server) CreatePed(pedType float64, modelHash interface{}, x float64, y float64, z float64, heading float64, isNetwork bool, bScriptHostPed bool) float64 {
	return 0.0
}

func (_server) CreatePedInsideVehicle(vehicle float64, pedType float64, modelHash interface{}, seat float64, isNetwork bool, bScriptHostPed bool) float64 {
	return 0.0
}

/**
 * Creates a vehicle with the specified model at the specified position. This vehicle will initially be owned by the creating
 * script as a mission entity, and the model should be loaded already (e.g. using REQUEST_MODEL).
 * ```
 * NativeDB Added Parameter 8: BOOL p7
 * ```
 * @param modelHash The model of vehicle to spawn.
 * @param x Spawn coordinate X component.
 * @param y Spawn coordinate Y component.
 * @param z Spawn coordinate Z component.
 * @param heading Heading to face towards, in degrees.
 * @param isNetwork Whether to create a network object for the vehicle. If false, the vehicle exists only locally.
 * @param netMissionEntity Whether to register the vehicle as pinned to the script host in the R\* network model.
 * @return A script handle (fwScriptGuid index) for the vehicle, or `0` if the vehicle failed to be created.
 */
func (_server) CreateVehicle(modelHash interface{}, x float64, y float64, z float64, heading float64, isNetwork bool, netMissionEntity bool) float64 {
	return 0.0
}

/**
 * Deletes the specified entity.
 * @param entity The entity to delete.
 */
func (_server) DeleteEntity(entity float64) {}

func (_server) DeleteFunctionReference(referenceIdentity string) {}

func (_server) DeleteResourceKvp(key string) {}

/**
 * Nonsynchronous [DELETE_RESOURCE_KVP](#\_0x7389B5DF) operation; see [FLUSH_RESOURCE_KVP](#\_0x5240DA5A).
 * @param key The key to delete
 */
func (_server) DeleteResourceKvpNoSync(key string) {}

func (_server) DoesEntityExist(entity float64) bool

/**
 * Requests whether or not the player owns the specified SKU.
 * @param playerSrc The player handle
 * @param skuId The ID of the SKU.
 * @return A boolean.
 */
func (_server) DoesPlayerOwnSku(playerSrc string, skuId float64) bool

/**
 * Requests whether or not the player owns the specified package.
 * @param playerSrc The player handle
 * @param skuId The package ID on Tebex.
 * @return A boolean.
 */
func (_server) DoesPlayerOwnSkuExt(playerSrc string, skuId float64) bool

func (_server) DropPlayer(playerSrc string, reason string) {}

func (_server) DuplicateFunctionReference(referenceIdentity string) string

func (_server) EnableEnhancedHostSupport(enabled bool) {}

func (_server) EndFindKvp(handle float64) {}

/**
 * Internal function for ensuring an entity has a state bag.
 */
func (_server) EnsureEntityStateBag(entity float64) {}

func (_server) ExecuteCommand(commandString string) {}

func (_server) FindKvp(handle float64) string

func (_server) FlagServerAsPrivate(private_ bool) {}

/**
 * Nonsynchronous operations will not wait for a disk/filesystem flush before returning from a write or delete call. They will be much faster than their synchronous counterparts (e.g., bulk operations), however, a system crash may lose the data to some recent operations.
 * This native ensures all `_NO_SYNC` operations are synchronized with the disk/filesystem.
 */
func (_server) FlushResourceKvp() {}

/**
 * Freezes or unfreezes an entity preventing its coordinates to change by the player if set to `true`. You can still change the entity position using SET_ENTITY_COORDS.
 * @param entity The entity to freeze/unfreeze.
 * @param toggle Freeze or unfreeze entity.
 */
func (_server) FreezeEntityPosition(entity float64, toggle bool) {}

func (_server) GetAirDragMultiplierForPlayersVehicle(playerSrc string) float64 {
	return 0.0
}

/**
 * Returns all object handles known to the server.
 * The data returned adheres to the following layout:
 * ```
 * [127, 42, 13, 37]
 * ```
 * @return An object containing a list of object handles.
 */
func (_server) GetAllObjects() interface{} {
	return 0.0
}

/**
 * Returns all peds handles known to the server.
 * The data returned adheres to the following layout:
 * ```
 * [127, 42, 13, 37]
 * ```
 * @return An object containing a list of peds handles.
 */
func (_server) GetAllPeds() interface{} {
	return 0.0
}

/**
 * Returns all vehicle handles known to the server.
 * The data returned adheres to the following layout:
 * ```
 * [127, 42, 13, 37]
 * ```
 * @return An object containing a list of vehicle handles.
 */
func (_server) GetAllVehicles() interface{} {
	return 0.0
}

func (_server) GetBlipSprite(self float64) float64 {
	return 0.0
}

/**
 * Returns the current console output buffer.
 * @return The most recent game console output, as a string.
 */
func (_server) GetConsoleBuffer() string

func (_server) GetConvar(varName string, default_ string) string

func (_server) GetConvarInt(varName string, default_ float64) float64 {
	return 0.0
}

/**
 * Returns the name of the currently executing resource.
 * @return The name of the resource.
 */
func (_server) GetCurrentResourceName() string

/**
 * Gets the current coordinates for a specified entity. This native is used server side when using OneSync.
 * See [GET_ENTITY_COORDS](#\_0x3FEF770D40960D5A) for client side.
 * @param entity The entity to get the coordinates from.
 * @return The current entity coordinates.
 */
func (_server) GetEntityCoords(entity float64) []float64 { return []float64{0.0} }

func (_server) GetEntityHeading(entity float64) float64 {
	return 0.0
}

/**
 * Currently it only works with peds.
 */
func (_server) GetEntityHealth(entity float64) float64 {
	return 0.0
}

/**
 * Currently it only works with peds.
 */
func (_server) GetEntityMaxHealth(entity float64) float64 {
	return 0.0
}

func (_server) GetEntityModel(entity float64) float64 {
	return 0.0
}

/**
 * This native gets an entity's population type.
 * @param entity the entity to obtain the population type from
 * @return Returns the population type ID, defined by the below enumeration:```cpp
enum ePopulationType
{
	POPTYPE_UNKNOWN = 0,
	POPTYPE_RANDOM_PERMANENT,
	POPTYPE_RANDOM_PARKED,
	POPTYPE_RANDOM_PATROL,
	POPTYPE_RANDOM_SCENARIO,
	POPTYPE_RANDOM_AMBIENT,
	POPTYPE_PERMANENT,
	POPTYPE_MISSION,
	POPTYPE_REPLAY,
	POPTYPE_CACHE,
	POPTYPE_TOOL
};
```
*/
func (_server) GetEntityPopulationType(entity float64) float64 {
	return 0.0
}

func (_server) GetEntityRotation(entity float64) []float64 { return []float64{0.0} }

func (_server) GetEntityRotationVelocity(entity float64) []float64 { return []float64{0.0} }

/**
 * Gets the routing bucket for the specified entity.
 * Routing buckets are also known as 'dimensions' or 'virtual worlds' in past echoes, however they are population-aware.
 * @param entity The entity to get the routing bucket for.
 * @return The routing bucket ID.
 */
func (_server) GetEntityRoutingBucket(entity float64) float64 {
	return 0.0
}

func (_server) GetEntityScript(entity float64) string

/**
 * Gets the current speed of the entity in meters per second.
 * ```
 * To convert to MPH: speed * 2.236936
 * To convert to KPH: speed * 3.6
 * ```
 * @param entity The entity to get the speed of
 * @return The speed of the entity in meters per second
 */
func (_server) GetEntitySpeed(entity float64) float64 {
	return 0.0
}

func (_server) GetEntityType(entity float64) float64 {
	return 0.0
}

func (_server) GetEntityVelocity(entity float64) []float64 { return []float64{0.0} }

/**
 * Returns the internal build number of the current game being executed.
 * Possible values:
 * *   FiveM
 * *   1604
 * *   2060
 * *   2189
 * *   2372
 * *   RedM
 * *   1311
 * *   1355
 * *   1436
 * *   LibertyM
 * *   43
 * *   FXServer
 * *   0
 * @return The build number, or **0** if no build number is known.
 */
func (_server) GetGameBuildNumber() float64 {
	return 0.0
}

/**
 * Returns the current game being executed.
 * Possible values:
 * | Return value | Meaning                        |
 * | ------------ | ------------------------------ |
 * | `fxserver`   | Server-side code ('Duplicity') |
 * | `fivem`      | FiveM for GTA V                |
 * | `libertym`   | LibertyM for GTA IV            |
 * | `redm`       | RedM for Red Dead Redemption 2 |
 * @return The game the script environment is running in.
 */
func (_server) GetGameName() string

/**
 * Gets the current game timer in milliseconds.
 * @return The game time.
 */
func (_server) GetGameTimer() float64 {
	return 0.0
}

/**
 * This native converts the passed string to a hash.
 */
func (_server) GetHashKey(model string) float64 {
	return 0.0
}

func (_server) GetHostId() string

func (_server) GetInstanceId() float64 {
	return 0.0
}

func (_server) GetInvokingResource() string

func (_server) GetIsVehicleEngineRunning(vehicle float64) bool

func (_server) GetIsVehiclePrimaryColourCustom(vehicle float64) bool

func (_server) GetIsVehicleSecondaryColourCustom(vehicle float64) bool

/**
 * See the client-side [GET_LANDING_GEAR_STATE](#\_0x9B0F3DCA3DB0F4CD) native for a description of landing gear states.
 * @param vehicle The vehicle to check.
 * @return The current state of the vehicles landing gear.
 */
func (_server) GetLandingGearState(vehicle float64) float64 {
	return 0.0
}

/**
 * Seat indexes:
 * *   1 = Driver
 * *   2 = Front Right Passenger
 * *   3 = Back Left Passenger
 * *   4 = Back Right Passenger
 * *   5 = Further Back Left Passenger (vehicles > 4 seats)
 * *   6 = Further Back Right Passenger (vehicles > 4 seats)
 * *   etc.
 * @param vehicle The target vehicle.
 * @param index The seat index.
 * @return The last ped in the specified seat of the passed vehicle. Returns 0 if the specified seat was never occupied or the last ped does not exist anymore.
 */
func (_server) GetLastPedInVehicleSeat(vehicle float64, index float64) float64 {
	return 0.0
}

func (_server) GetNumPlayerIdentifiers(playerSrc string) float64 {
	return 0.0
}

func (_server) GetNumPlayerIndices() float64 {
	return 0.0
}

func (_server) GetNumPlayerTokens(playerSrc string) float64 {
	return 0.0
}

/**
 * Gets the amount of metadata values with the specified key existing in the specified resource's manifest.
 * See also: [Resource manifest](https://docs.fivem.net/resources/manifest/)
 * @param resourceName The resource name.
 * @param metadataKey The key to look up in the resource manifest.
 */
func (_server) GetNumResourceMetadata(resourceName string, metadataKey string) float64 {
	return 0.0
}

func (_server) GetNumResources() float64 {
	return 0.0
}

func (_server) GetPasswordHash(password string) string

func (_server) GetPedArmour(ped float64) float64 {
	return 0.0
}

func (_server) GetPedCauseOfDeath(ped float64) float64 {
	return 0.0
}

func (_server) GetPedDesiredHeading(ped float64) float64 {
	return 0.0
}

/**
 * Seat indexes:
 * *   1 = Driver
 * *   2 = Front Right Passenger
 * *   3 = Back Left Passenger
 * *   4 = Back Right Passenger
 * *   5 = Further Back Left Passenger (vehicles > 4 seats)
 * *   6 = Further Back Right Passenger (vehicles > 4 seats)
 * *   etc.
 * @param vehicle The target vehicle.
 * @param index The seat index.
 * @return The ped in the specified seat of the passed vehicle. Returns 0 if the specified seat is not occupied.
 */
func (_server) GetPedInVehicleSeat(vehicle float64, index float64) float64 {
	return 0.0
}

func (_server) GetPedMaxHealth(ped float64) float64 {
	return 0.0
}

/**
 * Gets the script task command currently assigned to the ped.
 * @param ped The target ped.
 * @return The script task command currently assigned to the ped. A value of 0x811E343C denotes no script task is assigned.
 */
func (_server) GetPedScriptTaskCommand(ped float64) float64 {
	return 0.0
}

/**
 * Gets the stage of the peds scripted task.
 * @param ped The target ped.
 * @return The stage of the ped's scripted task. A value of 3 denotes no script task is assigned.
 */
func (_server) GetPedScriptTaskStage(ped float64) float64 {
	return 0.0
}

/**
 * Get the last entity that damaged the ped. This native is used server side when using OneSync.
 * @param ped The target ped
 * @return The entity id. Returns 0 if the ped has not been damaged recently.
 */
func (_server) GetPedSourceOfDamage(ped float64) float64 {
	return 0.0
}

/**
 * Get the entity that killed the ped. This native is used server side when using OneSync.
 * @param ped The target ped
 * @return The entity id. Returns 0 if the ped has no killer.
 */
func (_server) GetPedSourceOfDeath(ped float64) float64 {
	return 0.0
}

/**
 * Gets the type of a ped's specific task given an index of the CPedTaskSpecificDataNode nodes.
 * A ped will typically have a task at index 0, if a ped has multiple tasks at once they will be in the order 0, 1, 2, etc.
 * @param ped The target ped.
 * @param index A zero-based index with a maximum value of 7.
 * @return The type of the specific task.
1604: A value of 530 denotes no script task is assigned or an invalid input was given.
2060+: A value of 531 denotes no script task is assigned or an invalid input was given.
*/
func (_server) GetPedSpecificTaskType(ped float64, index float64) float64 {
	return 0.0
}

/**
 * Gets the current camera rotation for a specified player. This native is used server side when using OneSync.
 * @param playerSrc The player handle.
 * @return The player's camera rotation. Values are in radians.
 */
func (_server) GetPlayerCameraRotation(playerSrc string) []float64 { return []float64{0.0} }

func (_server) GetPlayerEndpoint(playerSrc string) string

/**
 * Gets the current fake wanted level for a specified player. This native is used server side when using OneSync.
 * @param playerSrc The target player
 * @return The fake wanted level
 */
func (_server) GetPlayerFakeWantedLevel(playerSrc string) float64 {
	return 0.0
}

func (_server) GetPlayerFromIndex(index float64) string

func (_server) GetPlayerGuid(playerSrc string) string

func (_server) GetPlayerIdentifier(playerSrc string, identifier float64) string

func (_server) GetPlayerInvincible(playerSrc string) bool

func (_server) GetPlayerLastMsg(playerSrc string) float64 {
	return 0.0
}

func (_server) GetPlayerMaxArmour(playerSrc string) float64 {
	return 0.0
}

func (_server) GetPlayerMaxHealth(playerSrc string) float64 {
	return 0.0
}

/**
 * A getter for [SET_PLAYER_MELEE_WEAPON_DAMAGE_MODIFIER](#\_0x4A3DC7ECCC321032).
 * @param playerId The player index.
 * @return Returns player melee weapon damage modifier value.
 */
func (_server) GetPlayerMeleeWeaponDamageModifier(playerId string) float64 {
	return 0.0
}

func (_server) GetPlayerName(playerSrc string) string

func (_server) GetPlayerPed(playerSrc string) float64 {
	return 0.0
}

func (_server) GetPlayerPing(playerSrc string) float64 {
	return 0.0
}

/**
 * Gets the routing bucket for the specified player.
 * Routing buckets are also known as 'dimensions' or 'virtual worlds' in past echoes, however they are population-aware.
 * @param playerSrc The player to get the routing bucket for.
 * @return The routing bucket ID.
 */
func (_server) GetPlayerRoutingBucket(playerSrc string) float64 {
	return 0.0
}

func (_server) GetPlayerTeam(playerSrc string) float64 {
	return 0.0
}

/**
 * Gets the amount of time player has spent evading the cops.
 * Counter starts and increments only when cops are chasing the player.
 * If the player is evading, the timer will pause.
 * @param playerSrc The target player
 * @param lastPursuit False = CurrentPursuit, True = LastPursuit
 * @return Returns -1, if the player is not wanted or wasn't in pursuit before, depending on the lastPursuit parameter
Returns 0, if lastPursuit == False and the player has a wanted level, but the pursuit has not started yet
Otherwise, will return the milliseconds of the pursuit.
*/
func (_server) GetPlayerTimeInPursuit(playerSrc string, lastPursuit bool) float64 {
	return 0.0
}

/**
 * Gets a player's token. Tokens can be used to enhance banning logic, however are specific to a server.
 * @param playerSrc A player.
 * @param index Index between 0 and GET_NUM_PLAYER_TOKENS.
 * @return A token value.
 */
func (_server) GetPlayerToken(playerSrc string, index float64) string

/**
 * Gets the current known coordinates for the specified player from cops perspective. This native is used server side when using OneSync.
 * @param playerSrc The target player
 * @return The player's position known by police. Vector zero if the player has no wanted level.
 */
func (_server) GetPlayerWantedCentrePosition(playerSrc string) []float64 { return []float64{0.0} }

/**
 * Returns given players wanted level server-side.
 * @param playerSrc The target player
 * @return The wanted level
 */
func (_server) GetPlayerWantedLevel(playerSrc string) float64 {
	return 0.0
}

/**
 * A getter for [SET_PLAYER_WEAPON_DAMAGE_MODIFIER](#\_0xCE07B9F7817AADA3).
 * @param playerId The player index.
 * @return The value of player weapon damage modifier.
 */
func (_server) GetPlayerWeaponDamageModifier(playerId string) float64 {
	return 0.0
}

/**
 * A getter for [SET_PLAYER_WEAPON_DEFENSE_MODIFIER](#\_0x2D83BC011CA14A3C).
 * @param playerId The player index.
 * @return The value of player weapon defense modifier.
 */
func (_server) GetPlayerWeaponDefenseModifier(playerId string) float64 {
	return 0.0
}

/**
 * A getter for [\_SET_PLAYER_WEAPON_DEFENSE_MODIFIER\_2](#\_0xBCFDE9EDE4CF27DC).
 * @param playerId The player index.
 * @return The value of player weapon defense modifier 2.
 */
func (_server) GetPlayerWeaponDefenseModifier_2(playerId string) float64 {
	return 0.0
}

/**
 * Returns all commands that are registered in the command system.
 * The data returned adheres to the following layout:
 * ```
 * [
 * {
 * "name": "cmdlist"
 * },
 * {
 * "name": "command1"
 * }
 * ]
 * ```
 * @return An object containing registered commands.
 */
func (_server) GetRegisteredCommands() interface{} {
	return 0.0
}

func (_server) GetResourceByFindIndex(findIndex float64) string

/**
 * A getter for [SET_RESOURCE_KVP_FLOAT](#\_0x9ADD2938).
 * @param key The key to fetch
 * @return A float that contains the value stored in the Kvp or nil/null if none.
 */
func (_server) GetResourceKvpFloat(key string) float64 {
	return 0.0
}

/**
 * A getter for [SET_RESOURCE_KVP_INT](#\_0x6A2B1E8).
 * @param key The key to fetch
 * @return A int that contains the value stored in the Kvp or nil/null if none.
 */
func (_server) GetResourceKvpInt(key string) float64 {
	return 0.0
}

/**
 * A getter for [SET_RESOURCE_KVP](#\_0x21C7A35B).
 * @param key The key to fetch
 * @return A string that contains the value stored in the Kvp or nil/null if none.
 */
func (_server) GetResourceKvpString(key string) string

/**
 * Gets the metadata value at a specified key/index from a resource's manifest.
 * See also: [Resource manifest](https://docs.fivem.net/resources/manifest/)
 * @param resourceName The resource name.
 * @param metadataKey The key in the resource manifest.
 * @param index The value index, in a range from \[0..GET_NUM_RESOURCE_METDATA-1].
 */
func (_server) GetResourceMetadata(resourceName string, metadataKey string, index float64) string

/**
 * Returns the physical on-disk path of the specified resource.
 * @param resourceName The name of the resource.
 * @return The resource directory name, possibly without trailing slash.
 */
func (_server) GetResourcePath(resourceName string) string

/**
 * Returns the current state of the specified resource.
 * @param resourceName The name of the resource.
 * @return The resource state. One of `"missing", "started", "starting", "stopped", "stopping", "uninitialized" or "unknown"`.
 */
func (_server) GetResourceState(resourceName string) string

/**
 * Returns a hash of selected ped weapon.
 * @param ped The target ped.
 * @return The weapon hash.
 */
func (_server) GetSelectedPedWeapon(ped float64) float64 {
	return 0.0
}

/**
 * Returns the value of a state bag key.
 * @return Value.
 */
func (_server) GetStateBagValue(bagName string, key string) interface{} {
	return 0.0
}

func (_server) GetTrainCarriageEngine(train float64) float64 {
	return 0.0
}

func (_server) GetTrainCarriageIndex(train float64) float64 {
	return 0.0
}

func (_server) GetVehicleBodyHealth(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleColours(vehicle float64) (float64, float64) { return 0.0, 0.0 }

func (_server) GetVehicleCustomPrimaryColour(vehicle float64) (float64, float64, float64) {
	return 0.0, 0.0, 0.0
}

func (_server) GetVehicleCustomSecondaryColour(vehicle float64) (float64, float64, float64) {
	return 0.0, 0.0, 0.0
}

func (_server) GetVehicleDashboardColour(vehicle float64, color float64) float64 {
	return 0.0
}

func (_server) GetVehicleDirtLevel(vehicle float64) float64 {
	return 0.0
}

/**
 * enum VehicleLockStatus = {
 * None = 0,
 * Unlocked = 1,
 * Locked = 2,
 * LockedForPlayer = 3,
 * StickPlayerInside = 4, -- Doesn't allow players to exit the vehicle with the exit vehicle key.
 * CanBeBrokenInto = 7, -- Can be broken into the car. If the glass is broken, the value will be set to 1
 * CanBeBrokenIntoPersist = 8, -- Can be broken into persist
 * CannotBeTriedToEnter = 10, -- Cannot be tried to enter (Nothing happens when you press the vehicle enter key).
 * }
 */
func (_server) GetVehicleDoorLockStatus(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleDoorStatus(vehicle float64) float64 {
	return 0.0
}

/**
 * Currently it only works when set to "all players".
 */
func (_server) GetVehicleDoorsLockedForPlayer(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleEngineHealth(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleExtraColours(vehicle float64) (float64, float64) { return 0.0, 0.0 }

/**
 * Gets the flight nozzel position for the specified vehicle. See the client-side [\_GET_VEHICLE_FLIGHT_NOZZLE_POSITION](#\_0xDA62027C8BDB326E) native for usage examples.
 * @param vehicle The vehicle to check.
 * @return The flight nozzel position between 0.0 (flying normally) and 1.0 (VTOL mode)
 */
func (_server) GetVehicleFlightNozzlePosition(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleHandbrake(vehicle float64) bool

func (_server) GetVehicleHeadlightsColour(vehicle float64) float64 {
	return 0.0
}

/**
 * Gets the lock on state for the specified vehicle. See the client-side [GET_VEHICLE_HOMING_LOCKON_STATE](#\_0xE6B0E8CFC3633BF0) native for a description of lock on states.
 * @param vehicle The vehicle to check.
 * @return The lock on state.
 */
func (_server) GetVehicleHomingLockonState(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleInteriorColour(vehicle float64, color float64) float64 {
	return 0.0
}

func (_server) GetVehicleLightsState(vehicle float64) (bool, interface{}, interface{}) {
	return false, nil, nil
}

func (_server) GetVehicleLivery(vehicle float64) float64 {
	return 0.0
}

/**
 * Gets the vehicle that is locked on to for the specified vehicle.
 * @param vehicle The vehicle to check.
 * @return The vehicle that is locked on. 0 returned if no vehicle is locked on.
 */
func (_server) GetVehicleLockOnTarget(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleNumberPlateText(vehicle float64) string

func (_server) GetVehicleNumberPlateTextIndex(vehicle float64) float64 {
	return 0.0
}

/**
 * Gets the vehicle the specified Ped is/was in depending on bool value. This native is used server side when using OneSync.
 * @param ped The target ped
 * @param lastVehicle False = CurrentVehicle, True = LastVehicle
 * @return The vehicle id. Returns 0 if the ped is/was not in a vehicle.
 */
func (_server) GetVehiclePedIsIn(ped float64, lastVehicle bool) float64 {
	return 0.0
}

func (_server) GetVehiclePetrolTankHealth(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleRadioStationIndex(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleRoofLivery(vehicle float64) float64 {
	return 0.0
}

/**
 * Returns the type of the passed vehicle.
 * ### Vehicle types
 * *   automobile
 * *   bike
 * *   boat
 * *   heli
 * *   plane
 * *   submarine
 * *   trailer
 * *   train
 * @param vehicle The vehicle's entity handle.
 * @return If the entity is a vehicle, the vehicle type. If it is not a vehicle, the return value will be null.
 */
func (_server) GetVehicleType(vehicle float64) string

func (_server) GetVehicleTyreSmokeColor(vehicle float64) (float64, float64, float64) {
	return 0.0, 0.0, 0.0
}

func (_server) GetVehicleWheelType(vehicle float64) float64 {
	return 0.0
}

func (_server) GetVehicleWindowTint(vehicle float64) float64 {
	return 0.0
}

func (_server) GiveWeaponComponentToPed(ped float64, weaponHash interface{}, componentHash interface{}) {
}

func (_server) GiveWeaponToPed(ped float64, weaponHash interface{}, ammoCount float64, isHidden bool, bForceInHand bool) {
}

func (_server) HasEntityBeenMarkedAsNoLongerNeeded(vehicle float64) bool

func (_server) HasVehicleBeenOwnedByPlayer(vehicle float64) bool

func (_server) InvokeFunctionReference(referenceIdentity string, argsSerialized string, argsLength float64, retvalLength float64) (string, float64) {
	return "", 0.0
}

func (_server) IsAceAllowed(object string) bool

/**
 * Gets whether or not this is the CitizenFX server.
 * @return A boolean value.
 */
func (_server) IsDuplicityVersion() bool

/**
 * This native checks if the given entity is visible.
 * @return Returns `true` if the entity is visible, `false` otherwise.
 */
func (_server) IsEntityVisible(entity float64) bool

/**
 * This native checks if the given ped is a player.
 * @return Returns `true` if the ped is a player, `false` otherwise.
 */
func (_server) IsPedAPlayer(ped float64) bool

func (_server) IsPlayerAceAllowed(playerSrc string, object string) bool

/**
 * Requests whether or not the commerce data for the specified player has loaded.
 * @param playerSrc The player handle
 * @return A boolean.
 */
func (_server) IsPlayerCommerceInfoLoaded(playerSrc string) bool

/**
 * Requests whether or not the commerce data for the specified player has loaded from Tebex.
 * @param playerSrc The player handle
 * @return A boolean.
 */
func (_server) IsPlayerCommerceInfoLoadedExt(playerSrc string) bool

/**
 * This will return true if the player is evading wanted level, meaning that the wanted level stars are blink.
 * Otherwise will return false.
 * If the player is not wanted, it simply returns false.
 * @param playerSrc The target player
 * @return boolean value, depending if the player is evading his wanted level or not.
 */
func (_server) IsPlayerEvadingWantedLevel(playerSrc string) bool

func (_server) IsPlayerUsingSuperJump(playerSrc string) bool

func (_server) IsPrincipalAceAllowed(principal string, object string) bool

func (_server) IsVehicleEngineStarting(vehicle float64) bool

func (_server) IsVehicleExtraTurnedOn(vehicle float64, extraId float64) bool

func (_server) IsVehicleSirenOn(vehicle float64) bool

func (_server) IsVehicleTyreBurst(vehicle float64, wheelID float64, completely bool) bool

/**
 * Requests the commerce data for the specified player, including the owned SKUs. Use `IS_PLAYER_COMMERCE_INFO_LOADED` to check if it has loaded.
 * @param playerSrc The player handle
 */
func (_server) LoadPlayerCommerceData(playerSrc string) {}

/**
 * Requests the commerce data from Tebex for the specified player, including the owned SKUs. Use `IS_PLAYER_COMMERCE_INFO_LOADED` to check if it has loaded.
 * @param playerSrc The player handle
 */
func (_server) LoadPlayerCommerceDataExt(playerSrc string) {}

/**
 * Reads the contents of a text file in a specified resource.
 * If executed on the client, this file has to be included in `files` in the resource manifest.
 * Example: `local data = LoadResourceFile("devtools", "data.json")`
 * @param resourceName The resource name.
 * @param fileName The file in the resource.
 * @return The file contents
 */
func (_server) LoadResourceFile(resourceName string, fileName string) string

/**
 * Create a permanent voice channel.
 * @param id ID of the channel.
 */
func (_server) MumbleCreateChannel(id float64) {}

func (_server) NetworkGetEntityFromNetworkId(netId float64) float64 {
	return 0.0
}

/**
 * Returns the owner ID of the specified entity.
 * @param entity The entity to get the owner for.
 * @return On the server, the server ID of the entity owner. On the client, returns the player/slot ID of the entity owner.
 */
func (_server) NetworkGetEntityOwner(entity float64) float64 {
	return 0.0
}

/**
 * Returns the first owner ID of the specified entity.
 * @param entity The entity to get the first owner for.
 * @return The server ID of the first entity owner.
 */
func (_server) NetworkGetFirstEntityOwner(entity float64) float64 {
	return 0.0
}

func (_server) NetworkGetNetworkIdFromEntity(entity float64) float64 {
	return 0.0
}

func (_server) NetworkGetVoiceProximityOverride(playerSrc string) []float64 { return []float64{0.0} }

func (_server) PerformHttpRequestInternal(requestData string, requestDataLength float64) float64 {
	return 0.0
}

func (_server) PerformHttpRequestInternalEx(requestData interface{}) float64 {
	return 0.0
}

/**
 * Prints 'structured trace' data to the server `file descriptor 3` channel. This is not generally useful outside of
 * server monitoring utilities.
 * @param jsonString JSON data to submit as `payload` in the `script_structured_trace` event.
 */
func (_server) PrintStructuredTrace(jsonString string) {}

/**
 * Scope entry for profiler.
 * @param scopeName Scope name.
 */
func (_server) ProfilerEnterScope(scopeName string) {}

/**
 * Scope exit for profiler.
 */
func (_server) ProfilerExitScope() {}

/**
 * Returns true if the profiler is active.
 * @return True or false.
 */
func (_server) ProfilerIsRecording() bool

/**
 * Registered commands can be executed by entering them in the client console (this works for client side and server side registered commands). Or by entering them in the server console/through an RCON client (only works for server side registered commands). Or if you use a supported chat resource, like the default one provided in the cfx-server-data repository, then you can enter the command in chat by prefixing it with a `/`.
 * Commands registered using this function can also be executed by resources, using the [`ExecuteCommand` native](#\_0x561C060B).
 * The restricted bool is not used on the client side. Permissions can only be checked on the server side, so if you want to limit your command with an ace permission automatically, make it a server command (by registering it in a server script).
 * **Example result**:
 * ![](https://i.imgur.com/TaCnG09.png)
 * @param commandName The command you want to register.
 * @param handler A handler function that gets called whenever the command is executed.
 * @param restricted If this is a server command and you set this to true, then players will need the command.yourCommandName ace permission to execute this command.
 */
func (_server) RegisterCommand(commandName string, handler js.Func, restricted bool) {}

/**
 * Registers a listener for console output messages.
 * @param listener A function of `(channel string, message string) => void`. The message might contain `\n`.
 */
func (_server) RegisterConsoleListener(listener js.Func) {}

/**
 * An internal function which allows the current resource's HLL script runtimes to receive state for the specified event.
 * @param eventName An event name, or "\*" to disable HLL event filtering for this resource.
 */
func (_server) RegisterResourceAsEventHandler(eventName string) {}

/**
 * **Experimental**: This native may be altered or removed in future versions of CitizenFX without warning.
 * Registers a cached resource asset with the resource system, similar to the automatic scanning of the `stream/` folder.
 * @param resourceName The resource to add the asset to.
 * @param fileName A file name in the resource.
 * @return A cache string to pass to `REGISTER_STREAMING_FILE_FROM_CACHE` on the client.
 */
func (_server) RegisterResourceAsset(resourceName string, fileName string) string

/**
 * Registers a build task factory for resources.
 * The function should return an object (msgpack map) with the following fields:
 * ```
 * {
 * // returns whether the specific resource should be built
 * shouldBuild = func(resourceName string): bool,
 * // asynchronously start building the specific resource.
 * // call cb when completed
 * build = func(resourceName string, cb: func(success: bool, status string): void): void
 * }
 * ```
 * @param factoryId The identifier for the build task.
 * @param factoryFn The factory function.
 */
func (_server) RegisterResourceBuildTaskFactory(factoryId string, factoryFn js.Func) {}

/**
 * setting the last params to false it does that same so I would suggest its not a toggle
 */
func (_server) RemoveAllPedWeapons(ped float64, p1 bool) {}

/**
 * In the C++ SDK, this seems not to work-- the blip isn't removed immediately. I use it for saving cars.
 * E.g.:
 * Ped pped = PLAYER::PLAYER_PED_ID();
 * Vehicle v = PED::GET_VEHICLE_PED_IS_USING(pped);
 * Blip b = HUD::ADD_BLIP_FOR_ENTITY(v);
 * works fine.
 * But later attempting to delete it with:
 * Blip b = HUD::GET_BLIP_FROM_ENTITY(v);
 * if (HUD::DOES_BLIP_EXIST(b)) HUD::REMOVE_BLIP(&b);
 * doesn't work. And yes, doesn't work without the DOES_BLIP_EXIST check either. Also, if you attach multiple blips to the same thing (say, a vehicle), and that thing disappears, the blips randomly attach to other things (in my case, a vehicle).
 * Thus for me, HUD::REMOVE_BLIP(&b) only works if there's one blip, (in my case) the vehicle is marked as no longer needed, you drive away from it and it eventually despawns, AND there is only one blip attached to it. I never intentionally attach multiple blips but if the user saves the car, this adds a blip. Then if they delete it, it is supposed to remove the blip, but it doesn't. Then they can immediately save it again, causing another blip to re-appear.
 * -------------
 * Passing the address of the variable instead of the value works for me.
 * e.g.
 * int blip = HUD::ADD_BLIP_FOR_ENTITY(ped);
 * HUD::REMOVE_BLIP(&blip);
 * Remove blip will currently crash your game, just artificially remove the blip by setting the sprite to a id that is 'invisible'.
 * --
 * It crashes my game.
 */
func (_server) RemoveBlip(blip float64) {}

/**
 * **Experimental**: This native may be altered or removed in future versions of CitizenFX without warning.
 * Removes a handler for changes to a state bag.
 * @param cookie The cookie.
 */
func (_server) RemoveStateBagChangeHandler(cookie float64) {}

func (_server) RemoveWeaponComponentFromPed(ped float64, weaponHash interface{}, componentHash interface{}) {
}

/**
 * This native removes a specified weapon from your selected ped.
 * Weapon Hashes: pastebin.com/0wwDZgkF
 * Example:
 * C#:
 * Function.Call(Hash.REMOVE_WEAPON_FROM_PED, Game.Player.Character, 0x99B507EA);
 * C++:
 * WEAPON::REMOVE_WEAPON_FROM_PED(PLAYER::PLAYER_PED_ID(), 0x99B507EA);
 * The code above removes the knife from the player.
 */
func (_server) RemoveWeaponFromPed(ped float64, weaponHash interface{}) {}

/**
 * Requests the specified player to buy the passed SKU. This'll pop up a prompt on the client, which upon acceptance
 * will open the browser prompting further purchase details.
 * @param playerSrc The player handle
 * @param skuId The ID of the SKU.
 */
func (_server) RequestPlayerCommerceSession(playerSrc string, skuId float64) {}

/**
 * Writes the specified data to a file in the specified resource.
 * Using a length of `-1` will automatically detect the length assuming the data is a C string.
 * @param resourceName The name of the resource.
 * @param fileName The name of the file.
 * @param data The data to write to the file.
 * @param dataLength The length of the written data.
 * @return A value indicating if the write succeeded.
 */
func (_server) SaveResourceFile(resourceName string, fileName string, data string, dataLength float64) bool

/**
 * Schedules the specified resource to run a tick as soon as possible, bypassing the server's fixed tick rate.
 * @param resourceName The resource to tick.
 */
func (_server) ScheduleResourceTick(resourceName string) {}

/**
 * <!--
 * _loc1_.map((name, idx) => `| ${idx} | ${name} | ![${name}](https://runtime.fivem.net/blips/${name}.svg) |`).join('\n')
 * -->
 * Sets the displayed sprite for a specific blip.
 * There's a [list of sprites](https://docs.fivem.net/game-references/blips/) on the FiveM documentation site.
 * @param blip The blip to change.
 * @param spriteId The sprite ID to set.
 */
func (_server) SetBlipSprite(blip float64, spriteId float64) {}

func (_server) SetConvar(varName string, value string) {}

func (_server) SetConvarReplicated(varName string, value string) {}

func (_server) SetConvarServerInfo(varName string, value string) {}

func (_server) SetCurrentPedWeapon(ped float64, weaponHash interface{}, bForceInHand bool) {}

/**
 * Sets the coordinates (world position) for a specified entity, offset by the radius of the entity on the Z axis.
 * @param entity The entity to change coordinates for.
 * @param xPos The X coordinate.
 * @param yPos The Y coordinate.
 * @param zPos The Z coordinate, ground level.
 * @param alive Unused by the game, potentially used by debug builds of GTA in order to assert whether or not an entity was alive.
 * @param deadFlag Whether to disable physics for dead peds, too, and not just living peds.
 * @param ragdollFlag A special flag used for ragdolling peds.
 * @param clearArea Whether to clear any entities in the target area.
 */
func (_server) SetEntityCoords(entity float64, xPos float64, yPos float64, zPos float64, alive bool, deadFlag bool, ragdollFlag bool, clearArea bool) {
}

/**
 * It overrides the default distance culling radius of an entity. Set to `0.0` to reset.
 * If you want to interact with an entity outside of your players' scopes set the radius to a huge number.
 * @param entity The entity handle to override the distance culling radius.
 * @param radius The new distance culling radius.
 */
func (_server) SetEntityDistanceCullingRadius(entity float64, radius float64) {}

/**
 * Set the heading of an entity in degrees also known as "Yaw".
 * @param entity The entity to set the heading for.
 * @param heading The heading in degrees.
 */
func (_server) SetEntityHeading(entity float64, heading float64) {}

func (_server) SetEntityRotation(entity float64, pitch float64, roll float64, yaw float64, rotationOrder float64, p5 bool) {
}

/**
 * Sets the routing bucket for the specified entity.
 * Routing buckets are also known as 'dimensions' or 'virtual worlds' in past echoes, however they are population-aware.
 * @param entity The entity to set the routing bucket for.
 * @param bucket The bucket ID.
 */
func (_server) SetEntityRoutingBucket(entity float64, bucket float64) {}

/**
 * Note that the third parameter(denoted as z) is "up and down" with positive numbers encouraging upwards movement.
 */
func (_server) SetEntityVelocity(entity float64, x float64, y float64, z float64) {}

func (_server) SetGameType(gametypeName string) {}

func (_server) SetHttpHandler(handler js.Func) {}

func (_server) SetMapName(mapName string) {}

/**
 * NativeDB Added Parameter 4: BOOL p3
 */
func (_server) SetPedAmmo(ped float64, weaponHash interface{}, ammo float64) {}

/**
 * Sets the armor of the specified ped.
 * ped: The Ped to set the armor of.
 * amount: A value between 0 and 100 indicating the value to set the Ped's armor to.
 */
func (_server) SetPedArmour(ped float64, amount float64) {}

func (_server) SetPedCanRagdoll(ped float64, toggle bool) {}

/**
 * This native is used to set component variation on a ped. Components, drawables and textures IDs are related to the ped model.
 * ### MP Freemode list of components
 * **0**: Face\
 * **1**: Mask\
 * **2**: Hair\
 * **3**: Torso\
 * **4**: Leg\
 * **5**: Parachute / bag\
 * **6**: Shoes\
 * **7**: Accessory\
 * **8**: Undershirt\
 * **9**: Kevlar\
 * **10**: Badge\
 * **11**: Torso 2
 * ### Related and useful natives
 * [GET_NUMBER_OF_PED_DRAWABLE_VARIATIONS](#\_0x27561561732A7842)\
 * [GET_NUMBER_OF_PED_TEXTURE_VARIATIONS](#\_0x8F7156A3142A6BAD)
 * [List of component/props ID](gtaxscripting.blogspot.com/2016/04/gta-v-peds-component-and-props.html) of player_two with examples
 * @param ped The ped handle.
 * @param componentId The component that you want to set.
 * @param drawableId The drawable id that is going to be set.
 * @param textureId The texture id of the drawable.
 * @param paletteId 0 to 3.
 */
func (_server) SetPedComponentVariation(ped float64, componentId float64, drawableId float64, textureId float64, paletteId float64) {
}

/**
 * cpp
 * // Potential names and hash collisions included as comments
 * enum ePedConfigFlags {
 * _0x67D1A445 = 0,
 * _0xC63DE95E = 1,
 * CPED_CONFIG_FLAG_NoCriticalHits = 2,
 * CPED_CONFIG_FLAG_DrownsInWater = 3,
 * CPED_CONFIG_FLAG_DisableReticuleFixedLockon = 4,
 * _0x37D196F4 = 5,
 * _0xE2462399 = 6,
 * CPED_CONFIG_FLAG_UpperBodyDamageAnimsOnly = 7,
 * _0xEDDEB838 = 8,
 * _0xB398B6FD = 9,
 * _0xF6664E68 = 10,
 * _0xA05E7CA3 = 11,
 * _0xCE394045 = 12,
 * CPED_CONFIG_FLAG_NeverLeavesGroup = 13,
 * _0xCD8D1411 = 14,
 * _0xB031F1A9 = 15,
 * _0xFE65BEE3 = 16,
 * CPED_CONFIG_FLAG_BlockNonTemporaryEvents = 17,
 * _0x380165BD = 18,
 * _0x07C045C7 = 19,
 * _0x583B5E2D = 20,
 * _0x475EDA58 = 21,
 * _0x8629D05B = 22,
 * _0x1522968B = 23,
 * CPED_CONFIG_FLAG_IgnoreSeenMelee = 24,
 * _0x4CC09C4B = 25,
 * _0x034F3053 = 26,
 * _0xD91BA7CC = 27,
 * _0x5C8DC66E = 28,
 * _0x8902EAA0 = 29,
 * _0x6580B9D2 = 30,
 * _0x0EF7A297 = 31,
 * _0x6BF86E5B = 32,
 * CPED_CONFIG_FLAG_DieWhenRagdoll = 33,
 * CPED_CONFIG_FLAG_HasHelmet = 34,
 * CPED_CONFIG_FLAG_UseHelmet = 35,
 * _0xEEB3D630 = 36,
 * _0xB130D17B = 37,
 * _0x5F071200 = 38,
 * CPED_CONFIG_FLAG_DisableEvasiveDives = 39,
 * _0xC287AAFF = 40,
 * _0x203328CC = 41,
 * CPED_CONFIG_FLAG_DontInfluenceWantedLevel = 42,
 * CPED_CONFIG_FLAG_DisablePlayerLockon = 43,
 * CPED_CONFIG_FLAG_DisableLockonToRandomPeds = 44,
 * _0xEC4A8ACF = 45,
 * _0xDB115BFA = 46,
 * CPED_CONFIG_FLAG_PedBeingDeleted = 47,
 * CPED_CONFIG_FLAG_BlockWeaponSwitching = 48,
 * _0xF8E99565 = 49,
 * _0xDD17FEE6 = 50,
 * _0x7ED9B2C9 = 51,
 * _0x655E8618 = 52,
 * _0x5A6C1F6E = 53,
 * _0xD749FC41 = 54,
 * _0x357F63F3 = 55,
 * _0xC5E60961 = 56,
 * _0x29275C3E = 57,
 * CPED_CONFIG_FLAG_IsFiring = 58,
 * CPED_CONFIG_FLAG_WasFiring = 59,
 * CPED_CONFIG_FLAG_IsStanding = 60,
 * CPED_CONFIG_FLAG_WasStanding = 61,
 * CPED_CONFIG_FLAG_InVehicle = 62,
 * CPED_CONFIG_FLAG_OnMount = 63,
 * CPED_CONFIG_FLAG_AttachedToVehicle = 64,
 * CPED_CONFIG_FLAG_IsSwimming = 65,
 * CPED_CONFIG_FLAG_WasSwimming = 66,
 * CPED_CONFIG_FLAG_IsSkiing = 67,
 * CPED_CONFIG_FLAG_IsSitting = 68,
 * CPED_CONFIG_FLAG_KilledByStealth = 69,
 * CPED_CONFIG_FLAG_KilledByTakedown = 70,
 * CPED_CONFIG_FLAG_Knockedout = 71,
 * _0x3E3C4560 = 72,
 * _0x2994C7B7 = 73,
 * _0x6D59D275 = 74,
 * CPED_CONFIG_FLAG_UsingCoverPoint = 75,
 * CPED_CONFIG_FLAG_IsInTheAir = 76,
 * _0x2D493FB7 = 77,
 * CPED_CONFIG_FLAG_IsAimingGun = 78,
 * _0x14D69875 = 79,
 * _0x40B05311 = 80,
 * _0x8B230BC5 = 81,
 * _0xC74E5842 = 82,
 * _0x9EA86147 = 83,
 * _0x674C746C = 84,
 * _0x3E56A8C2 = 85,
 * _0xC144A1EF = 86,
 * _0x0548512D = 87,
 * _0x31C93909 = 88,
 * _0xA0269315 = 89,
 * _0xD4D59D4D = 90,
 * _0x411D4420 = 91,
 * _0xDF4AEF0D = 92,
 * CPED_CONFIG_FLAG_ForcePedLoadCover = 93,
 * _0x300E4CD3 = 94,
 * _0xF1C5BF04 = 95,
 * _0x89C2EF13 = 96,
 * CPED_CONFIG_FLAG_VaultFromCover = 97,
 * _0x02A852C8 = 98,
 * _0x3D9407F1 = 99,
 * _0x319B4558 = 100,
 * CPED_CONFIG_FLAG_ForcedAim = 101,
 * _0xB942D71A = 102,
 * _0xD26C55A8 = 103,
 * _0xB89E703B = 104,
 * CPED_CONFIG_FLAG_ForceReload = 105,
 * _0xD9E73DA2 = 106,
 * _0xFF71DC2C = 107,
 * _0x1E27E8D8 = 108,
 * _0xF2C53966 = 109,
 * _0xC4DBE247 = 110,
 * _0x83C0A4BF = 111,
 * _0x0E0FAF8C = 112,
 * _0x26616660 = 113,
 * _0x43B80B79 = 114,
 * _0x0D2A9309 = 115,
 * _0x12C1C983 = 116,
 * CPED_CONFIG_FLAG_BumpedByPlayer = 117,
 * _0xE586D504 = 118,
 * _0x52374204 = 119,
 * CPED_CONFIG_FLAG_IsHandCuffed = 120,
 * CPED_CONFIG_FLAG_IsAnkleCuffed = 121,
 * CPED_CONFIG_FLAG_DisableMelee = 122,
 * _0xFE714397 = 123,
 * _0xB3E660BD = 124,
 * _0x5FED6BFD = 125,
 * _0xC9D6F66F = 126,
 * _0x519BC986 = 127,
 * CPED_CONFIG_FLAG_CanBeAgitated = 128,
 * _0x9A4B617C = 129, // CPED_CONFIG_FLAG_FaceDirInsult
 * _0xDAB70E9F = 130,
 * _0xE569438A = 131,
 * _0xBBC77D6D = 132,
 * _0xCB59EF0F = 133,
 * _0x8C5EA971 = 134,
 * CPED_CONFIG_FLAG_IsScuba = 135,
 * CPED_CONFIG_FLAG_WillArrestRatherThanJack = 136,
 * _0xDCE59B58 = 137,
 * CPED_CONFIG_FLAG_RidingTrain = 138,
 * CPED_CONFIG_FLAG_ArrestResult = 139,
 * CPED_CONFIG_FLAG_CanAttackFriendly = 140,
 * _0x98A4BE43 = 141,
 * _0x6901E731 = 142,
 * _0x9EC9BF6C = 143,
 * _0x42841A8F = 144,
 * CPED_CONFIG_FLAG_ShootingAnimFlag = 145,
 * CPED_CONFIG_FLAG_DisableLadderClimbing = 146,
 * CPED_CONFIG_FLAG_StairsDetected = 147,
 * CPED_CONFIG_FLAG_SlopeDetected = 148,
 * _0x1A15670B = 149,
 * _0x61786EE5 = 150,
 * _0xCB9186BD = 151,
 * _0xF0710152 = 152,
 * _0x43DFE310 = 153,
 * _0xC43C624E = 154,
 * CPED_CONFIG_FLAG_CanPerformArrest = 155,
 * CPED_CONFIG_FLAG_CanPerformUncuff = 156,
 * CPED_CONFIG_FLAG_CanBeArrested = 157,
 * _0xF7960FF5 = 158,
 * _0x59564113 = 159,
 * _0x0C6C3099 = 160,
 * _0x645F927A = 161,
 * _0xA86549B9 = 162,
 * _0x8AAF337A = 163,
 * _0x13BAA6E7 = 164,
 * _0x5FB9D1F5 = 165,
 * CPED_CONFIG_FLAG_IsInjured = 166,
 * _0x6398A20B = 167,
 * _0xD8072639 = 168,
 * _0xA05B1845 = 169,
 * _0x83F6D220 = 170,
 * _0xD8430331 = 171,
 * _0x4B547520 = 172,
 * _0xE66E1406 = 173,
 * _0x1C4BFE0C = 174,
 * _0x90008BFA = 175,
 * _0x07C7A910 = 176,
 * _0xF15F8191 = 177,
 * _0xCE4E8BE2 = 178,
 * _0x1D46E4F2 = 179,
 * CPED_CONFIG_FLAG_IsInCustody = 180,
 * _0xE4FD9B3A = 181,
 * _0x67AE0812 = 182,
 * CPED_CONFIG_FLAG_IsAgitated = 183,
 * CPED_CONFIG_FLAG_PreventAutoShuffleToDriversSeat = 184,
 * _0x7B2D325E = 185,
 * CPED_CONFIG_FLAG_EnableWeaponBlocking = 186,
 * CPED_CONFIG_FLAG_HasHurtStarted = 187,
 * CPED_CONFIG_FLAG_DisableHurt = 188,
 * CPED_CONFIG_FLAG_PlayerIsWeird = 189,
 * _0x32FC208B = 190,
 * _0x0C296E5A = 191,
 * _0xE63B73EC = 192,
 * _0x04E9CC80 = 193,
 * CPED_CONFIG_FLAG_UsingScenario = 194,
 * CPED_CONFIG_FLAG_VisibleOnScreen = 195,
 * _0xD88C58A1 = 196,
 * _0x5A3DCF43 = 197, // CPED_CONFIG_FLAG_AvoidUnderSide
 * _0xEA02B420 = 198,
 * _0x3F559CFF = 199,
 * _0x8C55D029 = 200,
 * _0x5E6466F6 = 201,
 * _0xEB5AD706 = 202,
 * _0x0EDDDDE7 = 203,
 * _0xA64F7B1D = 204,
 * _0x48532CBA = 205,
 * _0xAA25A9E7 = 206,
 * _0x415B26B9 = 207,
 * CPED_CONFIG_FLAG_DisableExplosionReactions = 208,
 * CPED_CONFIG_FLAG_DodgedPlayer = 209,
 * _0x67405504 = 210,
 * _0x75DDD68C = 211,
 * _0x2AD879B4 = 212,
 * _0x51486F91 = 213,
 * _0x32F79E21 = 214,
 * _0xBF099213 = 215,
 * _0x054AC8E2 = 216,
 * _0x14E495CC = 217,
 * _0x3C7DF9DF = 218,
 * _0x848FFEF2 = 219,
 * CPED_CONFIG_FLAG_DontEnterLeadersVehicle = 220,
 * _0x2618E1CF = 221,
 * _0x84F722FA = 222,
 * _0xD1B87B1F = 223,
 * _0x728AA918 = 224,
 * CPED_CONFIG_FLAG_DisablePotentialToBeWalkedIntoResponse = 225,
 * CPED_CONFIG_FLAG_DisablePedAvoidance = 226,
 * _0x59E91185 = 227,
 * _0x1EA7225F = 228,
 * CPED_CONFIG_FLAG_DisablePanicInVehicle = 229,
 * _0x6DCA7D88 = 230,
 * _0xFC3E572D = 231,
 * _0x08E9F9CF = 232,
 * _0x2D3BA52D = 233,
 * _0xFD2F53EA = 234,
 * _0x31A1B03B = 235,
 * CPED_CONFIG_FLAG_IsHoldingProp = 236,
 * _0x82ED0A66 = 237, // CPED_CONFIG_FLAG_BlocksPathingWhenDead
 * _0xCE57C9A3 = 238,
 * _0x26149198 = 239,
 * _0x1B33B598 = 240,
 * _0x719B6E87 = 241,
 * _0x13E8E8E8 = 242,
 * _0xF29739AE = 243,
 * _0xABEA8A74 = 244,
 * _0xB60EA2BA = 245,
 * _0x536B0950 = 246,
 * _0x0C754ACA = 247,
 * _0x69D28F3E = 248,
 * _0x12659168 = 249,
 * _0x1BDF2F04 = 250,
 * _0x7728FAA3 = 251,
 * _0x6A807ED8 = 252,
 * CPED_CONFIG_FLAG_OnStairs = 253,
 * _0xE1A2F73F = 254,
 * _0x5B3697C8 = 255,
 * _0xF1EB20A9 = 256,
 * _0x8B7DF407 = 257,
 * _0x329DCF1A = 258,
 * _0x8D90DD1B = 259,
 * _0xB8A292B7 = 260,
 * _0x8374B087 = 261,
 * _0x2AF558F0 = 262,
 * _0x82251455 = 263,
 * _0x30CF498B = 264,
 * _0xE1CD50AF = 265,
 * _0x72E4AE48 = 266,
 * _0xC2657EA1 = 267,
 * _0x29FF6030 = 268,
 * _0x8248A5EC = 269,
 * CPED_CONFIG_FLAG_OnStairSlope = 270,
 * _0xA0897933 = 271,
 * CPED_CONFIG_FLAG_DontBlipCop = 272,
 * CPED_CONFIG_FLAG_ClimbedShiftedFence = 273,
 * _0xF7823618 = 274,
 * _0xDC305CCE = 275, // CPED_CONFIG_FLAG_KillWhenTrapped
 * CPED_CONFIG_FLAG_EdgeDetected = 276,
 * _0x92B67896 = 277,
 * _0xCAD677C9 = 278,
 * CPED_CONFIG_FLAG_AvoidTearGas = 279,
 * _0x5276AC7B = 280,
 * _0x1032692A = 281,
 * _0xDA23E7F1 = 282,
 * _0x9139724D = 283,
 * _0xA1457461 = 284,
 * _0x4186E095 = 285,
 * _0xAC68E2EB = 286,
 * CPED_CONFIG_FLAG_RagdollingOnBoat = 287,
 * CPED_CONFIG_FLAG_HasBrandishedWeapon = 288,
 * _0x1B9EE8A1 = 289,
 * _0xF3F5758C = 290,
 * _0x2A9307F1 = 291,
 * _0x7403D216 = 292,
 * _0xA06A3C6C = 293,
 * CPED_CONFIG_FLAG_DisableShockingEvents = 294,
 * _0xF8DA25A5 = 295,
 * _0x7EF55802 = 296,
 * _0xB31F1187 = 297,
 * _0x84315402 = 298,
 * _0x0FD69867 = 299,
 * _0xC7829B67 = 300,
 * CPED_CONFIG_FLAG_DisablePedConstraints = 301,
 * _0x6D23CF25 = 302,
 * _0x2ADA871B = 303,
 * _0x47BC8A58 = 304,
 * _0xEB692FA5 = 305,
 * _0x4A133C50 = 306,
 * _0xC58099C3 = 307,
 * _0xF3D76D41 = 308,
 * _0xB0EEE9F2 = 309,
 * CPED_CONFIG_FLAG_IsInCluster = 310,
 * _0x0FA153EF = 311,
 * _0xD73F5CD3 = 312,
 * _0xD4136C22 = 313,
 * _0xE404CA6B = 314,
 * _0xB9597446 = 315,
 * _0xD5C98277 = 316,
 * _0xD5060A9C = 317,
 * _0x3E5F1CBB = 318,
 * _0xD8BE1D54 = 319,
 * _0x0B1F191F = 320,
 * _0xC995167A = 321,
 * CPED_CONFIG_FLAG_HasHighHeels = 322,
 * _0x86B01E54 = 323,
 * _0x3A56FE15 = 324,
 * _0xC03B736C = 325, // CPED_CONFIG_FLAG_SpawnedAtScenario
 * _0xBBF47729 = 326,
 * _0x22B668A8 = 327,
 * _0x2624D4D4 = 328,
 * CPED_CONFIG_FLAG_DisableTalkTo = 329,
 * CPED_CONFIG_FLAG_DontBlip = 330,
 * CPED_CONFIG_FLAG_IsSwitchingWeapon = 331,
 * _0x630F55F3 = 332,
 * _0x150468FD = 333,
 * _0x914EBD6B = 334,
 * _0x79AF3B6D = 335,
 * _0x75C7A632 = 336,
 * _0x52D530E2 = 337,
 * _0xDB2A90E0 = 338,
 * _0x5922763D = 339,
 * _0x12ADB567 = 340,
 * _0x105C8518 = 341,
 * _0x106F703D = 342,
 * _0xED152C3E = 343,
 * _0xA0EFE6A8 = 344,
 * _0xBF348C82 = 345,
 * _0xCDDFE830 = 346,
 * _0x7B59BD9B = 347,
 * _0x0124C788 = 348,
 * CPED_CONFIG_FLAG_EquipJetpack = 349,
 * _0x08D361A5 = 350,
 * _0xE13D1F7C = 351,
 * _0x40E25FB9 = 352,
 * _0x930629D9 = 353,
 * _0xECCF0C7F = 354,
 * _0xB6E9613B = 355,
 * _0x490C0478 = 356,
 * _0xE8865BEA = 357,
 * _0xF3C34A29 = 358,
 * CPED_CONFIG_FLAG_IsDuckingInVehicle = 359,
 * _0xF660E115 = 360,
 * _0xAB0E6DED = 361,
 * CPED_CONFIG_FLAG_HasReserveParachute = 362,
 * CPED_CONFIG_FLAG_UseReserveParachute = 363,
 * _0x5C5D9CD3 = 364,
 * _0x8F7701F3 = 365,
 * _0xBC4436AD = 366,
 * _0xD7E07D37 = 367,
 * _0x03C4FD24 = 368,
 * _0x7675789A = 369,
 * _0xB7288A88 = 370,
 * _0xC06B6291 = 371,
 * _0x95A4A805 = 372,
 * _0xA8E9A042 = 373,
 * CPED_CONFIG_FLAG_NeverLeaveTrain = 374,
 * _0xBAC674B3 = 375,
 * _0x147F1FFB = 376,
 * _0x4376DD79 = 377,
 * _0xCD3DB518 = 378,
 * _0xFE4BA4B6 = 379,
 * _0x5DF03A55 = 380,
 * _0xBCD816CD = 381,
 * _0xCF02DD69 = 382,
 * _0xF73AFA2E = 383,
 * _0x80B9A9D0 = 384,
 * _0xF601F7EE = 385,
 * _0xA91350FC = 386,
 * _0x3AB23B96 = 387,
 * CPED_CONFIG_FLAG_IsClimbingLadder = 388,
 * CPED_CONFIG_FLAG_HasBareFeet = 389,
 * _0xB4B1CD4C = 390,
 * _0x5459AFB8 = 391,
 * _0x54F27667 = 392,
 * _0xC11D3E8F = 393,
 * _0x5419EB3E = 394,
 * _0x82D8DBB4 = 395,
 * _0x33B02D2F = 396,
 * _0xAE66176D = 397,
 * _0xA2692593 = 398,
 * _0x714C7E31 = 399,
 * _0xEC488AC7 = 400,
 * _0xAE398504 = 401,
 * _0xABC58D72 = 402,
 * _0x5E5B9591 = 403,
 * _0x6BA1091E = 404,
 * _0x77840177 = 405,
 * _0x1C7ACAC4 = 406,
 * _0x124420E9 = 407,
 * _0x75A65587 = 408,
 * _0xDFD2D55B = 409,
 * _0xBDD39919 = 410,
 * _0x43DEC267 = 411,
 * _0xE42B7797 = 412,
 * CPED_CONFIG_FLAG_IsHolsteringWeapon = 413,
 * _0x4F8149F5 = 414,
 * _0xDD9ECA7A = 415,
 * _0x9E7EF9D2 = 416,
 * _0x2C6ED942 = 417,
 * CPED_CONFIG_FLAG_IsSwitchingHelmetVisor = 418,
 * _0xA488727D = 419,
 * _0xCFF5F6DE = 420,
 * _0x6D614599 = 421,
 * CPED_CONFIG_FLAG_DisableVehicleCombat = 422,
 * _0xFE401D26 = 423,
 * CPED_CONFIG_FLAG_FallsLikeAircraft = 424,
 * _0x2B42AE82 = 425,
 * _0x7A95734F = 426,
 * _0xDF4D8617 = 427,
 * _0x578F1F14 = 428,
 * CPED_CONFIG_FLAG_DisableStartEngine = 429,
 * CPED_CONFIG_FLAG_IgnoreBeingOnFire = 430,
 * _0x153C9500 = 431,
 * _0xCB7A632E = 432,
 * _0xDE727981 = 433,
 * CPED_CONFIG_FLAG_DisableHomingMissileLockon = 434,
 * _0x12BBB935 = 435,
 * _0xAD0A1277 = 436,
 * _0xEA6AA46A = 437,
 * CPED_CONFIG_FLAG_DisableHelmetArmor = 438,
 * _0xCB7F3A1E = 439,
 * _0x50178878 = 440,
 * _0x051B4F0D = 441,
 * _0x2FC3DECC = 442,
 * _0xC0030B0B = 443,
 * _0xBBDAF1E9 = 444,
 * _0x944FE59C = 445,
 * _0x506FBA39 = 446,
 * _0xDD45FE84 = 447,
 * _0xE698AE75 = 448,
 * _0x199633F8 = 449,
 * CPED_CONFIG_FLAG_PedIsArresting = 450,
 * CPED_CONFIG_FLAG_IsDecoyPed = 451,
 * _0x3A251D83 = 452,
 * _0xA56F6986 = 453,
 * _0x1D19C622 = 454,
 * _0xB68D3EAB = 455,
 * CPED_CONFIG_FLAG_CanBeIncapacitated = 456,
 * _0x4BD5EBAD = 457,
 * }
 */
func (_server) SetPedConfigFlag(ped float64, flagId float64, value bool) {}

/**
 * Sets Ped Default Clothes
 */
func (_server) SetPedDefaultComponentVariation(ped float64) {}

/**
 * Used for freemode (online) characters.
 * Indices:
 * 1.  black
 * 2.  very light blue/green
 * 3.  dark blue
 * 4.  brown
 * 5.  darker brown
 * 6.  light brown
 * 7.  blue
 * 8.  light blue
 * 9.  pink
 * 10. yellow
 * 11. purple
 * 12. black
 * 13. dark green
 * 14. light brown
 * 15. yellow/black pattern
 * 16. light colored spiral pattern
 * 17. shiny red
 * 18. shiny half blue/half red
 * 19. half black/half light blue
 * 20. white/red perimter
 * 21. green snake
 * 22. red snake
 * 23. dark blue snake
 * 24. dark yellow
 * 25. bright yellow
 * 26. all black
 * 27. red small pupil
 * 28. devil blue/black
 * 29. white small pupil
 * 30. glossed over
 */
func (_server) SetPedEyeColor(ped float64, index float64) {}

/**
 * Sets the various freemode face features, e.g. nose length, chin shape. Scale ranges from -1.0 to 1.0.
 * Index can be 0 - 19
 * SET_PED_M*
 * Here is the list of names. It starts at 0 and runs in sequence
 * Face_Feature
 * Nose_Width
 * Nose_Peak_Hight
 * Nose_Peak_Lenght
 * Nose_Bone_High
 * Nose_Peak_Lowering
 * Nose_Bone_Twist
 * EyeBrown_High
 * EyeBrown_Forward
 * Cheeks_Bone_High
 * Cheeks_Bone_Width
 * Cheeks_Width
 * Eyes_Openning
 * Lips_Thickness
 * Jaw_Bone_Width: Bone size to sides
 * Jaw_Bone_Back_Lenght: Bone size to back
 * Chimp_Bone_Lowering: Go Down
 * Chimp_Bone_Lenght: Go forward
 * Chimp_Bone_Width
 * Chimp_Hole
 * Neck_Thikness
 */
func (_server) SetPedFaceFeature(ped float64, index float64, scale float64) {}

/**
 * Used for freemode (online) characters.
 */
func (_server) SetPedHairColor(ped float64, colorID float64, highlightColorID float64) {}

/**
 * The "shape" parameters control the shape of the ped's face. The "skin" parameters control the skin tone. ShapeMix and skinMix control how much the first and second IDs contribute,(typically mother and father.) ThirdMix overrides the others in favor of the third IDs. IsParent is set for "children" of the player character's grandparents during old-gen character creation. It has unknown effect otherwise.
 * The IDs start at zero and go Male Non-DLC, Female Non-DLC, Male DLC, and Female DLC.
 * !!!Can someone add working example for this???
 * try this:
 * headBlendData headData;
 * _GET_PED_HEAD_BLEND_DATA(PLAYER_PED_ID(), &headData);
 * SET_PED_HEAD_BLEND_DATA(PLAYER_PED_ID(), headData.shapeFirst, headData.shapeSecond, headData.shapeThird, headData.skinFirst, headData.skinSecond
 * , headData.skinThird, headData.shapeMix, headData.skinMix, headData.skinThird, 0);
 * For more info please refer to this topic.
 * gtaforums.com/topic/858970-all-gtao-face-ids-pedset-ped-head-blend-data-explained
 */
func (_server) SetPedHeadBlendData(ped float64, shapeFirstID float64, shapeSecondID float64, shapeThirdID float64, skinFirstID float64, skinSecondID float64, skinThirdID float64, shapeMix float64, skinMix float64, thirdMix float64, isParent bool) {
}

/**
 * OverlayID ranges from 0 to 12, index from 0 to _GET_NUM_OVERLAY_VALUES(overlayID)-1, and opacity from 0.0 to 1.0.
 * overlayID       Part                  Index, to disable
 * 0               Blemishes             0 - 23, 255
 * 1               Facial Hair           0 - 28, 255
 * 2               Eyebrows              0 - 33, 255
 * 3               Ageing                0 - 14, 255
 * 4               Makeup                0 - 74, 255
 * 5               Blush                 0 - 6, 255
 * 6               Complexion            0 - 11, 255
 * 7               Sun Damage            0 - 10, 255
 * 8               Lipstick              0 - 9, 255
 * 9               Moles/Freckles        0 - 17, 255
 * 10              Chest Hair            0 - 16, 255
 * 11              Body Blemishes        0 - 11, 255
 * 12              Add Body Blemishes    0 - 1, 255
 */
func (_server) SetPedHeadOverlay(ped float64, overlayID float64, index float64, opacity float64) {}

/**
 * Used for freemode (online) characters.
 * ColorType is 1 for eyebrows, beards, and chest hair; 2 for blush and lipstick; and 0 otherwise, though not called in those cases.
 * Called after SET_PED_HEAD_OVERLAY().
 */
func (_server) SetPedHeadOverlayColor(ped float64, overlayID float64, colorType float64, colorID float64, secondColorID float64) {
}

func (_server) SetPedIntoVehicle(ped float64, vehicle float64, seatIndex float64) {}

/**
 * This native is used to set prop variation on a ped. Components, drawables and textures IDs are related to the ped model.
 * ### MP Freemode list of props
 * **0**: Hat\
 * **1**: Glass\
 * **2**: Ear\
 * **6**: Watch\
 * **7**: Bracelet
 * ### Related and useful natives
 * [GET_NUMBER_OF_PED_PROP_DRAWABLE_VARIATIONS](#\_0x5FAF9754E789FB47)\
 * [GET_NUMBER_OF_PED_PROP_TEXTURE_VARIATIONS](#\_0xA6E7F1CEB523E171)
 * [List of component/props ID](gtaxscripting.blogspot.com/2016/04/gta-v-peds-component-and-props.html) of player_two with examples
 * @param ped The ped handle.
 * @param componentId The component that you want to set.
 * @param drawableId The drawable id that is going to be set.
 * @param textureId The texture id of the drawable.
 * @param attach Attached or not.
 */
func (_server) SetPedPropIndex(ped float64, componentId float64, drawableId float64, textureId float64, attach bool) {
}

/**
 * p1 is always 0 in R* scripts; and a quick disassembly seems to indicate that p1 is unused.
 * List of component/props ID:
 * gtaxscripting.blogspot.com/2016/04/gta-v-peds-component-and-props.html
 */
func (_server) SetPedRandomComponentVariation(ped float64, p1 float64) {}

/**
 * List of component/props ID
 * gtaxscripting.blogspot.com/2016/04/gta-v-peds-component-and-props.html
 */
func (_server) SetPedRandomProps(ped float64) {}

/**
 * PED::SET_PED_RESET_FLAG(PLAYER::PLAYER_PED_ID(), 240, 1);
 * Known values:
 */
func (_server) SetPedResetFlag(ped float64, flagId float64, doReset bool) {}

/**
 * p4/p5: Unusued in TU27
 * ### Ragdoll Types
 * **0**: CTaskNMRelax
 * **1**: CTaskNMScriptControl: Hardcoded not to work in networked environments.
 * **Else**: CTaskNMBalance
 * @param time1 Time(ms) Ped is in ragdoll mode; only applies to ragdoll types 0 and not 1.
 */
func (_server) SetPedToRagdoll(ped float64, time1 float64, time2 float64, ragdollType float64, p4 bool, p5 bool, p6 bool) {
}

/**
 * Return variable is never used in R*'s scripts.
 * Not sure what p2 does. It seems like it would be a time judging by it's usage in R*'s scripts, but didn't seem to affect anything in my testings.
 * x, y, and z are coordinates, most likely to where the ped will fall.
 * p7 is probably the force of the fall, but untested, so I left the variable name the same.
 * p8 to p13 are always 0f in R*'s scripts.
 * (Simplified) Example of the usage of the function from R*'s scripts:
 * ped::set_ped_to_ragdoll_with_fall(ped, 1500, 2000, 1, -entity::get_entity_forward_vector(ped), 1f, 0f, 0f, 0f, 0f, 0f, 0f);
 */
func (_server) SetPedToRagdollWithFall(ped float64, time float64, p2 float64, ragdollType float64, x float64, y float64, z float64, p7 float64, p8 float64, p9 float64, p10 float64, p11 float64, p12 float64, p13 float64) {
}

/**
 * Flags:
 * SPC_AMBIENT_SCRIPT = (1 << 1),
 * SPC_CLEAR_TASKS = (1 << 2),
 * SPC_REMOVE_FIRES = (1 << 3),
 * SPC_REMOVE_EXPLOSIONS = (1 << 4),
 * SPC_REMOVE_PROJECTILES = (1 << 5),
 * SPC_DEACTIVATE_GADGETS = (1 << 6),
 * SPC_REENABLE_CONTROL_ON_DEATH = (1 << 7),
 * SPC_LEAVE_CAMERA_CONTROL_ON = (1 << 8),
 * SPC_ALLOW_PLAYER_DAMAGE = (1 << 9),
 * SPC_DONT_STOP_OTHER_CARS_AROUND_PLAYER = (1 << 10),
 * SPC_PREVENT_EVERYBODY_BACKOFF = (1 << 11),
 * SPC_ALLOW_PAD_SHAKE = (1 << 12)
 * See: https://alloc8or.re/gta5/doc/enums/eSetPlayerControlFlag.txt
 */
func (_server) SetPlayerControl(player string, bHasControl bool, flags float64) {}

/**
 * Sets the culling radius for the specified player.
 * Set to `0.0` to reset.
 * @param playerSrc The player to set the culling radius for.
 * @param radius The radius.
 */
func (_server) SetPlayerCullingRadius(playerSrc string, radius float64) {}

/**
 * Simply sets you as invincible (Health will not deplete).
 * Use 0x733A643B5B0C53C1 instead if you want Ragdoll enabled, which is equal to:
 * *(DWORD *)(playerPedAddress + 0x188) |= (1 << 9);
 */
func (_server) SetPlayerInvincible(player string, toggle bool) {}

/**
 * Set the model for a specific Player. Be aware that this will destroy the current Ped for the Player and create a new one, any reference to the old ped should be reset
 * Make sure to request the model first and wait until it has loaded.
 */
func (_server) SetPlayerModel(player string, model interface{}) {}

/**
 * Sets the routing bucket for the specified player.
 * Routing buckets are also known as 'dimensions' or 'virtual worlds' in past echoes, however they are population-aware.
 * @param playerSrc The player to set the routing bucket for.
 * @param bucket The bucket ID.
 */
func (_server) SetPlayerRoutingBucket(playerSrc string, bucket float64) {}

/**
 * Call SET_PLAYER_WANTED_LEVEL_NOW for immediate effect
 * wantedLevel is an integer value representing 0 to 5 stars even though the game supports the 6th wanted level but no police will appear since no definitions are present for it in the game files
 * disableNoMission-  Disables When Off Mission- appears to always be false
 */
func (_server) SetPlayerWantedLevel(player string, wantedLevel float64, disableNoMission bool) {}

/**
 * A setter for [GET_RESOURCE_KVP_STRING](#\_0x5240DA5A).
 * @param key The key to set
 * @param value The value to write
 */
func (_server) SetResourceKvp(key string, value string) {}

/**
 * A setter for [GET_RESOURCE_KVP_FLOAT](#\_0x35BDCEEA).
 * @param key The key to set
 * @param value The value to write
 */
func (_server) SetResourceKvpFloat(key string, value float64) {}

/**
 * Nonsynchronous [SET_RESOURCE_KVP_FLOAT](#\_0x9ADD2938) operation; see [FLUSH_RESOURCE_KVP](#\_0x5240DA5A).
 * @param key The key to set
 * @param value The value to write
 */
func (_server) SetResourceKvpFloatNoSync(key string, value float64) {}

/**
 * A setter for [GET_RESOURCE_KVP_INT](#\_0x557B586A).
 * @param key The key to set
 * @param value The value to write
 */
func (_server) SetResourceKvpInt(key string, value float64) {}

/**
 * Nonsynchronous [SET_RESOURCE_KVP_INT](#\_0x6A2B1E8) operation; see [FLUSH_RESOURCE_KVP](#\_0x5240DA5A).
 * @param key The key to set
 * @param value The value to write
 */
func (_server) SetResourceKvpIntNoSync(key string, value float64) {}

/**
 * Nonsynchronous [SET_RESOURCE_KVP](#\_0x21C7A35B) operation; see [FLUSH_RESOURCE_KVP](#\_0x5240DA5A).
 * @param key The key to set
 * @param value The value to write
 */
func (_server) SetResourceKvpNoSync(key string, value string) {}

/**
 * Sets the entity lockdown mode for a specific routing bucket.
 * Lockdown modes are:
 * | Mode       | Meaning                                                    |
 * | ---------- | ---------------------------------------------------------- |
 * | `strict`   | No entities can be created by clients at all.              |
 * | `relaxed`  | Only script-owned entities created by clients are blocked. |
 * | `inactive` | Clients can create any entity they want.                   |
 * @param bucketId The routing bucket ID to adjust.
 * @param mode One of aforementioned modes.
 */
func (_server) SetRoutingBucketEntityLockdownMode(bucketId float64, mode string) {}

/**
 * Sets whether or not the specified routing bucket has automatically-created population enabled.
 * @param bucketId The routing bucket ID to adjust.
 * @param mode `true` to enable population, `false` to disable population.
 */
func (_server) SetRoutingBucketPopulationEnabled(bucketId float64, mode bool) {}

/**
 * Internal function for setting a state bag value.
 */
func (_server) SetStateBagValue(bagName string, keyName string, valueData string, valueLength float64, replicated bool) {
}

func (_server) SetVehicleAlarm(vehicle float64, state bool) {}

/**
 * p2 often set to 1000.0 in the decompiled scripts.
 */
func (_server) SetVehicleBodyHealth(vehicle float64, value float64) {}

/**
 * Sets the selected vehicle's colors to their default value (specific variant specified using the colorCombination parameter).
 * Range of possible values for colorCombination is currently unknown, I couldn't find where these values are stored either (Disquse's guess was vehicles.meta but I haven't seen it in there.)
 * @param vehicle The vehicle to modify.
 * @param colorCombination One of the default color values of the vehicle.
 */
func (_server) SetVehicleColourCombination(vehicle float64, colorCombination float64) {}

/**
 * colorPrimary & colorSecondary are the paint indexes for the vehicle.
 * For a list of valid paint indexes, view: pastebin.com/pwHci0xK
 */
func (_server) SetVehicleColours(vehicle float64, colorPrimary float64, colorSecondary float64) {}

/**
 * p1, p2, p3 are RGB values for color (255,0,0 for Red, ect)
 */
func (_server) SetVehicleCustomPrimaryColour(vehicle float64, r float64, g float64, b float64) {}

/**
 * p1, p2, p3 are RGB values for color (255,0,0 for Red, ect)
 */
func (_server) SetVehicleCustomSecondaryColour(vehicle float64, r float64, g float64, b float64) {}

/**
 * Sets the dirt level of the passed vehicle.
 * @param vehicle The vehicle to set.
 * @param dirtLevel A number between 0.0 and 15.0 representing the vehicles dirt level.
 */
func (_server) SetVehicleDirtLevel(vehicle float64, dirtLevel float64) {}

/**
 * See eDoorId declared in [`SET_VEHICLE_DOOR_SHUT`](#\_0x93D9BD300D7789E5)
 */
func (_server) SetVehicleDoorBroken(vehicle float64, doorIndex float64, deleteDoor bool) {}

/**
 * // Source GTA VC miss2 leak, matching constants for 0/2/4, testing
 * // They use 10 in am_mp_property_int, don't know what it does atm.
 * enum eCarLock {
 * CARLOCK_NONE = 0,
 * CARLOCK_UNLOCKED = 1,
 * CARLOCK_LOCKED = 2,
 * CARLOCK_LOCKOUT_PLAYER_ONLY = 3,
 * CARLOCK_LOCKED_PLAYER_INSIDE = 4,
 * CARLOCK_LOCKED_INITIALLY = 5,
 * CARLOCK_FORCE_SHUT_DOORS = 6,
 * CARLOCK_LOCKED_BUT_CAN_BE_DAMAGED = 7
 * };
 */
func (_server) SetVehicleDoorsLocked(vehicle float64, doorLockStatus float64) {}

func (_server) SetVehicleNumberPlateText(vehicle float64, plateText string) {}

func (_server) StartFindKvp(prefix string) float64 {
	return 0.0
}

func (_server) StartResource(resourceName string) bool

func (_server) StopResource(resourceName string) bool

/**
 * Makes the specified ped attack the target ped.
 * p2 should be 0
 * p3 should be 16
 */
func (_server) TaskCombatPed(ped float64, targetPed float64, p2 float64, p3 float64) {}

/**
 * Example:
 * TASK::TASK_DRIVE_BY(l_467[1 -- [[22]] ], PLAYER::PLAYER_PED_ID(), 0, 0.0, 0.0, 2.0, 300.0, 100, 0, ${firing_pattern_burst_fire_driveby});
 * Needs working example. Doesn't seem to do anything.
 * I marked p2 as targetVehicle as all these shooting related tasks seem to have that in common.
 * I marked p6 as distanceToShoot as if you think of GTA's Logic with the native SET_VEHICLE_SHOOT natives, it won't shoot till it gets within a certain distance of the target.
 * I marked p7 as pedAccuracy as it seems it's mostly 100 (Completely Accurate), 75, 90, etc. Although this could be the ammo count within the gun, but I highly doubt it. I will change this comment once I find out if it's ammo count or not.
 */
func (_server) TaskDriveBy(driverPed float64, targetPed float64, targetVehicle float64, targetX float64, targetY float64, targetZ float64, distanceToShoot float64, pedAccuracy float64, p8 bool, firingPattern interface{}) {
}

/**
 * speed 1.0 = walk, 2.0 = run
 * p5 1 = normal, 3 = teleport to vehicle, 16 = teleport directly into vehicle
 * p6 is always 0
 * @param seatIndex See eSeatPosition declared in [`IS_VEHICLE_SEAT_FREE`](#\_0x22AC59A870E6A669).
 */
func (_server) TaskEnterVehicle(ped float64, vehicle float64, timeout float64, seatIndex float64, speed float64, flag float64, p6 float64) {
}

func (_server) TaskEveryoneLeaveVehicle(vehicle float64) {}

func (_server) TaskGoStraightToCoord(ped float64, x float64, y float64, z float64, speed float64, timeout float64, targetHeading float64, distanceToSlide float64) {
}

/**
 * example from fm_mission_controller
 * TASK::TASK_GO_TO_COORD_ANY_MEANS(l_649, sub_f7e86(-1, 0), 1.0, 0, 0, 786603, 0xbf800000);
 */
func (_server) TaskGoToCoordAnyMeans(ped float64, x float64, y float64, z float64, speed float64, p5 float64, p6 bool, walkingStyle float64, p8 float64) {
}

/**
 * The entity will move towards the target until time is over (duration) or get in target's range (distance). p5 and p6 are unknown, but you could leave p5 = 1073741824 or 100 or even 0 (didn't see any difference but on the decompiled scripts, they use 1073741824 mostly) and p6 = 0
 * Note: I've only tested it on entity -> ped and target -> vehicle. It could work differently on other entities, didn't try it yet.
 * Example: TASK::TASK_GO_TO_ENTITY(pedHandle, vehicleHandle, 5000, 4.0, 100, 1073741824, 0)
 * Ped will run towards the vehicle for 5 seconds and stop when time is over or when he gets 4 meters(?) around the vehicle (with duration = -1, the task duration will be ignored).
 */
func (_server) TaskGoToEntity(entity float64, target float64, duration float64, distance float64, speed float64, p5 float64, p6 float64) {
}

/**
 * In the scripts, p3 was always -1.
 * p3 seems to be duration or timeout of turn animation.
 * Also facingPed can be 0 or -1 so ped will just raise hands up.
 */
func (_server) TaskHandsUp(ped float64, duration float64, facingPed float64, p3 float64, p4 bool) {}

/**
 * Flags are the same flags used in [`TASK_LEAVE_VEHICLE`](#\_0xD3DBCE61A490BE02)
 */
func (_server) TaskLeaveAnyVehicle(ped float64, p1 float64, flags float64) {}

/**
 * Flags from decompiled scripts:
 * 0 = normal exit and closes door.
 * 1 = normal exit and closes door.
 * 16 = teleports outside, door kept closed.
 * 64 = normal exit and closes door, maybe a bit slower animation than 0.
 * 256 = normal exit but does not close the door.
 * 4160 = ped is throwing himself out, even when the vehicle is still.
 * 262144 = ped moves to passenger seat first, then exits normally
 * Others to be tried out: 320, 512, 131072.
 */
func (_server) TaskLeaveVehicle(ped float64, vehicle float64, flags float64) {}

/**
 * [Animations list](https://alexguirre.github.io/animations-list/)
 * ```
 * float blendInSpeed > normal speed is 8.0f
 * ----------------------
 * float blendOutSpeed > normal speed is 8.0f
 * ----------------------
 * int duration: time in millisecond
 * ----------------------
 * -1 _ _ _ _ _ _ _> Default (see flag)
 * 0 _ _ _ _ _ _ _ > Not play at all
 * Small value _ _ > Slow down animation speed
 * Other _ _ _ _ _ > freeze player control until specific time (ms) has
 * _ _ _ _ _ _ _ _ _ passed. (No effect if flag is set to be
 * _ _ _ _ _ _ _ _ _ controllable.)
 * int flag:
 * ----------------------
 * enum eAnimationFlags
 * {
 * ANIM_FLAG_NORMAL = 0,
 * ANIM_FLAG_REPEAT = 1,
 * ANIM_FLAG_STOP_LAST_FRAME = 2,
 * ANIM_FLAG_UPPERBODY = 16,
 * ANIM_FLAG_ENABLE_PLAYER_CONTROL = 32,
 * ANIM_FLAG_CANCELABLE = 120,
 * };
 * Odd number : loop infinitely
 * Even number : Freeze at last frame
 * Multiple of 4: Freeze at last frame but controllable
 * 01 to 15 > Full body
 * 10 to 31 > Upper body
 * 32 to 47 > Full body > Controllable
 * 48 to 63 > Upper body > Controllable
 * ...
 * 001 to 255 > Normal
 * 256 to 511 > Garbled
 * ...
 * playbackRate:
 * values are between 0.0 and 1.0
 * lockX:
 * 0 in most cases 1 for rcmepsilonism8 and rcmpaparazzo_3
 * > 1 for mini@sprunk
 * lockY:
 * 0 in most cases
 * 1 for missfam5_yoga, missfra1mcs_2_crew_react
 * lockZ:
 * 0 for single player
 * Can be 1 but only for MP
 * ```
 */
func (_server) TaskPlayAnim(ped float64, animDictionary string, animationName string, blendInSpeed float64, blendOutSpeed float64, duration float64, flag float64, playbackRate float64, lockX bool, lockY bool, lockZ bool) {
}

/**
 * It's similar to the one above, except the first 6 floats let you specify the initial position and rotation of the task. (Ped gets teleported to the position).
 * [Animations list](https://alexguirre.github.io/animations-list/)
 * @param ped The target ped
 * @param animDict Name of the animation dictionary
 * @param animName Name of the animation
 * @param posX Initial X position of the task
 * @param posY Initial Y position of the task
 * @param posZ Initial Z position of the task
 * @param rotX Initial X rotation of the task, doesn't seem to have any effect
 * @param rotY Initial Y rotation of the task, doesn't seem to have any effect
 * @param rotZ Initial Z rotation of the task
 * @param animEnterSpeed Adjust character speed to fully enter animation
 * @param animExitSpeed Adjust character speed to fully exit animation (useless `ClearPedTasksImmediately()` is called)
 * @param duration Time in milliseconds
 * @param animTime Value between 0.0 and 1.0, lets you start an animation from the given point
 */
func (_server) TaskPlayAnimAdvanced(ped float64, animDict string, animName string, posX float64, posY float64, posZ float64, rotX float64, rotY float64, rotZ float64, animEnterSpeed float64, animExitSpeed float64, duration float64, flag float64, animTime float64, p14 float64, p15 float64) {
}

func (_server) TaskReactAndFleePed(ped float64, fleeTarget float64) {}

/**
 * Firing Pattern Hash Information: https://pastebin.com/Px036isB
 */
func (_server) TaskShootAtCoord(ped float64, x float64, y float64, z float64, duration float64, firingPattern interface{}) {
}

/**
 * //this part of the code is to determine at which entity the player is aiming, for example if you want to create a mod where you give orders to peds
 * Entity aimedentity;
 * Player player = PLAYER::PLAYER_ID();
 * PLAYER::_GET_AIMED_ENTITY(player, &aimedentity);
 * //bg is an array of peds
 * TASK::TASK_SHOOT_AT_ENTITY(bg[i], aimedentity, 5000, MISC::GET_HASH_KEY("FIRING_PATTERN_FULL_AUTO"));
 * in practical usage, getting the entity the player is aiming at and then task the peds to shoot at the entity, at a button press event would be better.
 * Firing Pattern Hash Information: https://pastebin.com/Px036isB
 */
func (_server) TaskShootAtEntity(entity float64, target float64, duration float64, firingPattern interface{}) {
}

func (_server) TaskWarpPedIntoVehicle(ped float64, vehicle float64, seatIndex float64) {}

func (_server) TempBanPlayer(playerSrc string, reason string) {}

/**
 * The backing function for TriggerClientEvent.
 */
func (_server) TriggerClientEventInternal(eventName string, eventTarget string, eventPayload string, payloadLength float64) {
}

/**
 * The backing function for TriggerEvent.
 */
func (_server) TriggerEventInternal(eventName string, eventPayload string, payloadLength float64) {}

/**
 * The backing function for TriggerLatentClientEvent.
 */
func (_server) TriggerLatentClientEventInternal(eventName string, eventTarget string, eventPayload string, payloadLength float64, bps float64) {
}

func (_server) VerifyPasswordHash(password string, hash string) bool

/**
 * Returns whether or not the currently executing event was canceled.
 * @return A boolean.
 */
func (_server) WasEventCanceled() bool
