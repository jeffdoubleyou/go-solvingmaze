package models

type Item struct {
	Sku               string        `json:"sku,omitempty"`
	Name              string        `json:"name"`
	Qty               int           `json:"qty,integer"`
	Shelf             string        `json:"shelf,models.Shelf,omitempty"`
	UnitPrice         float64       `json:"unitPrice,float64,omitempty"`
	Value             int           `json:"value,int,omitempty"`
	Weight            float64       `json:"weight,float64,omitempty"`
	WeightUnit        string        `json:"weightUnit"`
	Dimensions        []*Dimension  `json:"dimensions,[]models.Dimension"`
	DimensionsUnit    string        `json:"dimensionsUnit,omitempty"`
	Rotatable         bool          `json:"rotatable,omitempty"`
	Packable          bool          `json:"packable,omitempty"`
	Strappable        bool          `json:"strappable,omitempty"`
	VoidFiller        bool          `json:"voidFiller,omitempty"`
	VoidDimensions    []*Dimension  `json:"voidDimensions,[]models.Dimension,omitempty"`
	PrepackageMinQty  int           `json:"prepackMinQty,int,omitempty"`
	Prepackages       []*Prepackage `json:"prepackages,[]models.Prepackage,omitempty"`
	Stack             *Stack        `json:"stack,models.Stack,omitempty"`
	Packaging         *Packaging    `json:"packaging,models.Packaging,omitempty"`
	PreferContainers  []string      `json:"preferContainers,[]string,omitempty"`
	ExcludeContainers []string      `json:"excludeContainers,[]string,omitempty"`
	ExcludeServices   []string      `json:"excludeServices,[]string,omitempty"`
	Group             int           `json:"group,int,omitempty"`
}

type PackedItem struct {
	Sku        string      `json:"sku,omitempty"`
	Name       string      `json:"name"`
	Qty        int         `json:"qty,integer"`
	VoidSku    string      `json:"voidSku,omitempty"`
	VoidIndex  int         `json:"voidIndex,integer,omitempty"`
	Prepackage *Prepackage `json:"prepackage,models.Prepackage,omitempty"`
}
