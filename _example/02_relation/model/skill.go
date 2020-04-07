// Code generated by eevee. DO NOT EDIT!

package model

import (
	"context"
	"encoding/json"
	"relation/dao"
	"relation/entity"
	"sort"
	"strconv"

	"golang.org/x/xerrors"
)

type SkillFinder interface {
	FindAll(context.Context) (*Skills, error)
	FindByID(context.Context, uint64) (*Skill, error)
	FindByIDs(context.Context, []uint64) (*Skills, error)
}

type Skill struct {
	*entity.Skill
	skillDAO         dao.Skill
	isAlreadyCreated bool
	savedValue       entity.Skill
	conv             ModelConverter
}

type Skills struct {
	values []*Skill
}

type SkillsCollection []*Skills

func NewSkill(value *entity.Skill, skillDAO dao.Skill) *Skill {
	return &Skill{
		Skill:    value,
		skillDAO: skillDAO,
	}
}

func NewSkills(entities entity.Skills) *Skills {
	return &Skills{values: make([]*Skill, 0, len(entities))}
}

func (m *Skills) newSkills(values []*Skill) *Skills {
	return &Skills{values: values}
}

func (m *Skills) Each(iter func(*Skill)) {
	if m == nil {
		return
	}
	for _, value := range m.values {
		iter(value)
	}
}

func (m *Skills) EachIndex(iter func(int, *Skill)) {
	if m == nil {
		return
	}
	for idx, value := range m.values {
		iter(idx, value)
	}
}

func (m *Skills) EachWithError(iter func(*Skill) error) error {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if err := iter(value); err != nil {
			return xerrors.Errorf("failed to iteration: %w", err)
		}
	}
	return nil
}

func (m *Skills) EachIndexWithError(iter func(int, *Skill) error) error {
	if m == nil {
		return nil
	}
	for idx, value := range m.values {
		if err := iter(idx, value); err != nil {
			return xerrors.Errorf("failed to iteration: %w", err)
		}
	}
	return nil
}

func (m *Skills) Map(mapFunc func(*Skill) *Skill) *Skills {
	if m == nil {
		return nil
	}
	mappedValues := []*Skill{}
	for _, value := range m.values {
		mappedValue := mapFunc(value)
		if mappedValue != nil {
			mappedValues = append(mappedValues, mappedValue)
		}
	}
	return m.newSkills(mappedValues)
}

func (m *Skills) Any(cond func(*Skill) bool) bool {
	if m == nil {
		return false
	}
	for _, value := range m.values {
		if cond(value) {
			return true
		}
	}
	return false
}

func (m *Skills) Some(cond func(*Skill) bool) bool {
	return m.Any(cond)
}

func (m *Skills) IsIncluded(cond func(*Skill) bool) bool {
	return m.Any(cond)
}

func (m *Skills) All(cond func(*Skill) bool) bool {
	if m == nil {
		return false
	}
	for _, value := range m.values {
		if !cond(value) {
			return false
		}
	}
	return true
}

func (m *Skills) Sort(compare func(*Skill, *Skill) bool) {
	if m == nil {
		return
	}
	sort.Slice(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *Skills) SortStable(compare func(*Skill, *Skill) bool) {
	if m == nil {
		return
	}
	sort.SliceStable(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *Skills) Find(cond func(*Skill) bool) *Skill {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if cond(value) {
			return value
		}
	}
	return nil
}

func (m *Skills) Filter(filter func(*Skill) bool) *Skills {
	if m == nil {
		return nil
	}
	filteredValues := []*Skill{}
	for _, value := range m.values {
		if filter(value) {
			filteredValues = append(filteredValues, value)
		}
	}
	return m.newSkills(filteredValues)
}

func (m *Skills) IsEmpty() bool {
	if m == nil {
		return true
	}
	if len(m.values) == 0 {
		return true
	}
	return false
}

func (m *Skills) At(idx int) *Skill {
	if m == nil {
		return nil
	}
	if idx < 0 {
		return nil
	}
	if len(m.values) > idx {
		return m.values[idx]
	}
	return nil
}

func (m *Skills) First() *Skill {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[0]
	}
	return nil
}

func (m *Skills) Last() *Skill {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[len(m.values)-1]
	}
	return nil
}

func (m *Skills) Compact() *Skills {
	if m == nil {
		return nil
	}
	compactedValues := []*Skill{}
	for _, value := range m.values {
		if value == nil {
			continue
		}
		compactedValues = append(compactedValues, value)
	}
	return m.newSkills(compactedValues)
}

func (m *Skills) Add(args ...*Skill) *Skills {
	if m == nil {
		return nil
	}
	for _, value := range args {
		m.values = append(m.values, value)
	}
	return m
}

func (m *Skills) Merge(args ...*Skills) *Skills {
	if m == nil {
		return nil
	}
	for _, arg := range args {
		for _, value := range arg.values {
			m.values = append(m.values, value)
		}
	}
	return m
}

func (m *Skills) Len() int {
	if m == nil {
		return 0
	}
	return len(m.values)
}

func (m *SkillsCollection) Merge() *Skills {
	if m == nil {
		return nil
	}
	if len(*m) == 0 {
		return nil
	}
	if len(*m) == 1 {
		return (*m)[0]
	}
	values := []*Skill{}
	for _, collection := range *m {
		for _, value := range collection.values {
			values = append(values, value)
		}
	}
	return (*m)[0].newSkills(values)
}

func (m *Skill) ToJSON(ctx context.Context) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	buf = append(buf, '{')
	buf = append(buf, "\"id\":"...)
	buf = strconv.AppendUint(buf, m.ID, 10)
	buf = append(buf, ',')
	buf = append(buf, "\"skillEffect\":"...)
	buf = append(buf, strconv.Quote(m.SkillEffect)...)
	buf = append(buf, '}')
	return buf, nil
}

func (m *Skill) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	isWritten := false
	buf = append(buf, '{')
	if option.Exists("id") {
		buf = append(buf, "\"id\":"...)
		buf = strconv.AppendUint(buf, m.ID, 10)
		isWritten = true
	}
	if option.Exists("skill_effect") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"skillEffect\":"...)
		buf = append(buf, strconv.Quote(m.SkillEffect)...)
		isWritten = true
	}
	buf = append(buf, '}')
	return buf, nil
}

func (m *Skills) ToJSON(ctx context.Context) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	buf = append(buf, '[')
	for idx, value := range m.values {
		if idx != 0 {
			buf = append(buf, ',')
		}
		bytes, err := value.ToJSON(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to JSON: %w", err)
		}
		buf = append(buf, bytes...)
	}
	buf = append(buf, ']')
	return buf, nil
}

func (m *Skills) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	buf = append(buf, '[')
	for idx, value := range m.values {
		if idx != 0 {
			buf = append(buf, ',')
		}
		bytes, err := value.ToJSONWithOption(ctx, option)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to JSON: %w", err)
		}
		buf = append(buf, bytes...)
	}
	buf = append(buf, ']')
	return buf, nil
}

func (m *Skill) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Skill) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Skills) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Skills) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Skill) UnmarshalJSON(bytes []byte) error {
	var value struct {
		*entity.Skill
	}
	if err := json.Unmarshal(bytes, &value); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.Skill = value.Skill
	return nil
}

func (m *Skills) UnmarshalJSON(bytes []byte) error {
	var values []*Skill
	if err := json.Unmarshal(bytes, &values); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.values = values
	return nil
}

func (m *Skill) ToMap(ctx context.Context) (map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := map[string]interface{}{}
	value["id"] = m.ID
	value["skillEffect"] = m.SkillEffect
	return value, nil
}

func (m *Skill) ToMapWithOption(ctx context.Context, option *RenderOption) (map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := map[string]interface{}{}
	if option.Exists("id") {
		value["id"] = m.ID
	}
	if option.Exists("skill_effect") {
		value["skillEffect"] = m.SkillEffect
	}
	return value, nil
}

func (m *Skills) ToMap(ctx context.Context) ([]map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := []map[string]interface{}{}
	for _, v := range m.values {
		mapValue, err := v.ToMap(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to map: %w", err)
		}
		value = append(value, mapValue)
	}
	return value, nil
}

func (m *Skills) ToMapWithOption(ctx context.Context, option *RenderOption) ([]map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := []map[string]interface{}{}
	for _, v := range m.values {
		mapValue, err := v.ToMapWithOption(ctx, option)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to map: %w", err)
		}
		value = append(value, mapValue)
	}
	return value, nil
}

func (m *Skill) SetConverter(conv ModelConverter) {
	m.conv = conv
}

func (m *Skill) Create(ctx context.Context) error {
	if m.skillDAO == nil {
		// for testing
		return nil
	}
	if m.isAlreadyCreated {
		return xerrors.New("this instance has already created")
	}
	if err := m.skillDAO.Create(ctx, m.Skill); err != nil {
		return xerrors.Errorf("failed to Create: %w", err)
	}
	m.savedValue = *m.Skill
	m.isAlreadyCreated = true
	return nil
}

func (m *Skill) Update(ctx context.Context) error {
	if m.skillDAO == nil {
		// for testing
		return nil
	}
	isRequiredUpdate := false
	if m.savedValue.ID != m.ID {
		isRequiredUpdate = true
	}
	if m.savedValue.SkillEffect != m.SkillEffect {
		isRequiredUpdate = true
	}
	if !isRequiredUpdate {
		return nil
	}
	if err := m.skillDAO.Update(ctx, m.Skill); err != nil {
		return xerrors.Errorf("failed to Update: %w", err)
	}
	m.savedValue = *m.Skill
	return nil
}

func (m *Skill) Delete(ctx context.Context) error {
	if m.skillDAO == nil {
		// for testing
		return nil
	}
	if err := m.skillDAO.DeleteByID(ctx, m.ID); err != nil {
		return xerrors.Errorf("failed to Delete: %w", err)
	}
	return nil
}

func (m *Skill) SetAlreadyCreated(isAlreadyCreated bool) {
	m.isAlreadyCreated = isAlreadyCreated
}

func (m *Skill) SetSavedValue(savedValue *entity.Skill) {
	m.savedValue = *savedValue
}

func (m *Skill) Save(ctx context.Context) error {
	if m.isAlreadyCreated {
		if err := m.Update(ctx); err != nil {
			return xerrors.Errorf("failed to Update: %w", err)
		}
		return nil
	}
	if err := m.Create(ctx); err != nil {
		return xerrors.Errorf("failed to Create: %w", err)
	}
	return nil
}

func (m *Skills) Create(ctx context.Context) error {
	if err := m.EachWithError(func(v *Skill) error {
		if err := v.Create(ctx); err != nil {
			return xerrors.Errorf("failed to Create: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Skills: %w", err)
	}
	return nil
}

func (m *Skills) Update(ctx context.Context) error {
	if err := m.EachWithError(func(v *Skill) error {
		if err := v.Update(ctx); err != nil {
			return xerrors.Errorf("failed to Update: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Skills: %w", err)
	}
	return nil
}

func (m *Skills) Save(ctx context.Context) error {
	if err := m.EachWithError(func(v *Skill) error {
		if err := v.Save(ctx); err != nil {
			return xerrors.Errorf("failed to Save: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Skills: %w", err)
	}
	return nil
}

func (m *Skills) UniqueID() *Skills {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *Skill) bool {
		if _, exists := filterMap[value.ID]; exists {
			return false
		}
		filterMap[value.ID] = struct{}{}
		return true
	})
}

func (m *Skills) GroupByID() map[uint64]*Skills {
	if m == nil {
		return nil
	}
	values := map[uint64]*Skills{}
	for _, value := range m.values {
		if _, exists := values[value.ID]; !exists {
			values[value.ID] = &Skills{}
		}
		values[value.ID].Add(value)
	}
	return values
}

func (m *Skills) IDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.ID)
	}
	return values
}

func (m *Skills) UniqueSkillEffect() *Skills {
	if m == nil {
		return nil
	}
	filterMap := map[string]struct{}{}
	return m.Filter(func(value *Skill) bool {
		if _, exists := filterMap[value.SkillEffect]; exists {
			return false
		}
		filterMap[value.SkillEffect] = struct{}{}
		return true
	})
}

func (m *Skills) GroupBySkillEffect() map[string]*Skills {
	if m == nil {
		return nil
	}
	values := map[string]*Skills{}
	for _, value := range m.values {
		if _, exists := values[value.SkillEffect]; !exists {
			values[value.SkillEffect] = &Skills{}
		}
		values[value.SkillEffect].Add(value)
	}
	return values
}

func (m *Skills) SkillEffects() []string {
	if m == nil {
		return nil
	}
	values := []string{}
	for _, value := range m.values {
		values = append(values, value.SkillEffect)
	}
	return values
}

func (m *Skills) FirstByID(a0 uint64) *Skill {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.ID != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Skills) FilterByID(a0 uint64) *Skills {
	if m == nil {
		return nil
	}
	values := []*Skill{}
	for _, value := range m.values {
		if value.ID != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newSkills(values)
}
