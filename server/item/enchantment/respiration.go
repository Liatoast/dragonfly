package enchantment

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// Respiration extends underwater breathing time by +15 seconds per enchantment level in addition to the default time of
// 15 seconds.
type Respiration struct{}

// Name ...
func (Respiration) Name() string {
	return "Respiration"
}

// MaxLevel ...
func (Respiration) MaxLevel() int {
	return 3
}

// Rarity ...
func (Respiration) Rarity() item.EnchantmentRarity {
	return item.EnchantmentRarityRare
}

// CompatibleWithEnchantment ...
func (Respiration) CompatibleWithEnchantment(item.EnchantmentType) bool {
	return true
}

// CompatibleWithItem ...
func (Respiration) CompatibleWithItem(i world.Item) bool {
	h, ok := i.(item.HelmetType)
	return ok && h.Helmet()
}
