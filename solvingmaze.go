package solvingmaze

import (
	"encoding/json"

	"github.com/jeffdoubleyou/go-solvingmaze/models"

	"github.com/mitchellh/mapstructure"
)

const (
	API_BASE = "https://api.solvingmaze.com/"
)

type SolvingMaze struct {
	client *client
}

type PackAndQuote struct {
	sm          *SolvingMaze
	items       []*models.Item
	destination *models.Address
}

type Destination struct{ address *models.Address }
type Item struct{ item *models.Item }
type Dimension struct{ dimension *models.Dimension }
type Stack struct{ stack *models.Stack }
type Packaging struct{ packaging *models.Packaging }
type Prepackage struct{ prepackage *models.Prepackage }

func New(ApiKey string, WarehouseId string, timeout int, debug bool) *SolvingMaze {
	client := NewClient(ApiKey, WarehouseId, timeout, debug)
	return &SolvingMaze{client}
}

func (sm *SolvingMaze) NewPackAndQuote() *PackAndQuote {
	var items []*models.Item
	var destination *models.Address
	return &PackAndQuote{sm, items, destination}
}

func (pq *PackAndQuote) Destination(props map[string]interface{}) (*Destination, error) {
	address := &models.Address{}
	err := mapstructure.Decode(props, &address)
	if err != nil {
		return nil, err
	}
	pq.destination = address
	return &Destination{address}, nil
}

func (pq *PackAndQuote) AddItem(props map[string]interface{}) (*Item, error) {

	item := &models.Item{
		Sku:            "",
		Name:           "",
		Qty:            1,
		WeightUnit:     "lb",
		Dimensions:     []*models.Dimension{},
		DimensionsUnit: "in",
		Rotatable:      true,
		Packable:       true,
		Strappable:     true,
	}

	if props["Dimensions"] != nil {
		for _, dim := range props["Dimensions"].([][]float64) {
			item.Dimensions = append(item.Dimensions, &models.Dimension{dim[0], dim[1], dim[2]})
		}
		delete(props, "Dimensions")
	}

	if props["VoidDimensions"] != nil {
		for _, dim := range props["VoidDimensions"].([][]float64) {
			item.VoidDimensions = append(item.VoidDimensions, &models.Dimension{dim[0], dim[1], dim[2]})
		}
		delete(props, "VoidDimensions")
	}

	if props["Prepackages"] != nil {
		/*for _, p := range props["Prepackages"].([][]float64) {
			item.Prepackages = append(item.Prepackages, &models.Prepackage{0, p[0], p[1], p[2].(string), p[3], p[4], p[5]})
		}*/
		delete(props, "Prepackages")
	}

	err := mapstructure.Decode(props, &item)

	if err != nil {
		return nil, err
	}

	i := &Item{item}
	pq.items = append(pq.items, i.item)
	return i, nil
}

func (pq *PackAndQuote) Calculate() (*models.PackResponse, error) {
	req := &models.PackRequest{}
	req.Items = pq.items
	req.Destination = pq.destination

	body, err := json.Marshal(req)
	resp, err := pq.sm.client.Post("calculate", body)

	if err != nil {
		return nil, err
	}

	res := &models.PackResponse{}
	err = json.Unmarshal(resp, res)
	return res, nil
}

func (item *Item) AddDimensions(props ...map[string]interface{}) *Dimension {
	dim := &models.Dimension{}
	if len(props) == 1 {
		mapstructure.Decode(props[0], dim)
	}
	item.item.Dimensions = append(item.item.Dimensions, dim)
	return &Dimension{dim}
}

func (dim *Dimension) Length(l float64) {
	dim.dimension.Length = l
}

func (dim *Dimension) Width(w float64) {
	dim.dimension.Width = w
}

func (dim *Dimension) Height(h float64) {
	dim.dimension.Height = h
}

func (item *Item) AddVoidDimensions(props ...map[string]interface{}) *Dimension {
	dim := &models.Dimension{}
	if len(props) == 1 {
		mapstructure.Decode(props[0], dim)
	}
	item.item.VoidDimensions = append(item.item.VoidDimensions, dim)
	return &Dimension{dim}
}

func (item *Item) AddPrepackage(props ...map[string]interface{}) *Prepackage {
	prepackage := &models.Prepackage{}
	if len(props) == 1 {
		mapstructure.Decode(props[0], prepackage)
	}
	item.item.Prepackages = append(item.item.Prepackages, prepackage)
	return &Prepackage{prepackage}
}

func (p *Prepackage) Length(l float64) {
	p.prepackage.Length = l
}

func (p *Prepackage) Width(w float64) {
	p.prepackage.Width = w
}

func (p *Prepackage) Height(h float64) {
	p.prepackage.Height = h
}

func (p *Prepackage) DimensionUnit(u string) {
	p.prepackage.DimensionUnit = u
}

func (p *Prepackage) Weight(w float64) {
	p.prepackage.Weight = w
}

func (p *Prepackage) Capacity(c int) {
	p.prepackage.Capacity = c
}

func (item *Item) AddStack(props ...map[string]interface{}) *Stack {
	stack := &models.Stack{}
	if len(props) == 1 {
		mapstructure.Decode(props[0], stack)
	}
	item.item.Stack = stack
	return &Stack{stack}
}

func (s *Stack) LengthIncrement(l float64) {
	s.stack.LengthIncrement = l
}

func (s *Stack) WidthIncrement(w float64) {
	s.stack.WidthIncrement = w
}

func (s *Stack) HeightIncrement(h float64) {
	s.stack.HeightIncrement = h
}

func (s *Stack) MaxQty(q int) {
	s.stack.MaxQty = q
}

func (item *Item) AddPackaging(props ...map[string]interface{}) *Packaging {
	packaging := &models.Packaging{}
	if len(props) == 1 {
		mapstructure.Decode(props[0], packaging)
	}
	item.item.Packaging = packaging
	return &Packaging{packaging}
}

func (p *Packaging) Length(l float64) {
	p.packaging.Length = l
}

func (p *Packaging) Width(w float64) {
	p.packaging.Width = w
}

func (p *Packaging) Height(h float64) {
	p.packaging.Height = h
}

func (p *Packaging) Weight(w float64) {
	p.packaging.Weight = w
}

func (p *Packaging) Irregular(i bool) {
	p.packaging.Irregular = i
}
