/**
 * This hook transform property name from the backend to DM names for ease of
 * support for other APIs
 */
export function getPropertyMap(prop: keyof typeof IconMap) {
  let isPercent = true;
  if (["BASE", "FLAT", "ENERGY_REGEN_CONVERT"].some(e => prop.endsWith(e))) isPercent = false;

  return {
    dmValue: IconMap[prop],
    isPercent,
  };
}
export const IconMap = {
  HP_BASE: "IconMaxHP",
  HP_PERCENT: "IconMaxHP",
  HP_FLAT: "IconMaxHP",
  HP_CONVERT: "IconMaxHP",
  ATK_BASE: "IconAttack",
  ATK_PERCENT: "IconAttack",
  ATK_FLAT: "IconAttack",
  ATK_CONVERT: "IconAttack",
  DEF_BASE: "IconDefence",
  DEF_PERCENT: "IconDefence",
  DEF_FLAT: "IconDefence",
  DEF_CONVERT: "IconDefence",
  SPD_BASE: "IconSpeed",
  SPD_PERCENT: "IconSpeed",
  SPD_FLAT: "IconSpeed",
  SPD_CONVERT: "IconSpeed",
  CRIT_CHANCE: "IconCriticalDamage",
  CRIT_DMG: "IconCriticalDamage",
  ENERGY_REGEN: "IconEnergyRecovery",
  ENERGY_REGEN_CONVERT: "IconEnergyLimit",
  EFFECT_HIT_RATE: "IconStatusProbability",
  EFFECT_HIT_RATE_CONVERT: "IconStatusProbability",
  EFFECT_RES: "IconStatusResistance",
  EFFECT_RES_CONVERT: "IconStatusResistance",
  HEAL_BOOST: "IconHealRatio",
  HEAL_BOOST_CONVERT: "IconHealRatio",
  HEAL_TAKEN: "IconHealRatio",
  BREAK_EFFECT: "IconBreakUp",
  PHYSICAL_DMG_RES: "IconPhysicalResistanceDelta",
  FIRE_DMG_RES: "IconFireResistanceDelta",
  ICE_DMG_RES: "IconIceResistanceDelta",
  THUNDER_DMG_RES: "IconThunderResistanceDelta",
  QUANTUM_DMG_RES: "IconQuantumResistanceDelta",
  IMAGINARY_DMG_RES: "IconImaginaryResistanceDelta",
  WIND_DMG_RES: "IconWindResistanceDelta",
  FIRE_DMG_PERCENT: "IconFireAddedRatio",
  ICE_DMG_PERCENT: "IconIceAddedRatio",
  THUNDER_DMG_PERCENT: "IconThunderAddedRatio",
  QUANTUM_DMG_PERCENT: "IconQuantumAddedRatio",
  IMAGINARY_DMG_PERCENT: "IconImaginaryAddedRatio",
  WIND_DMG_PERCENT: "IconWindAddedRatio",
  PHYSICAL_DMG_PERCENT: "IconPhysicalAddedRatio",
} as const;
