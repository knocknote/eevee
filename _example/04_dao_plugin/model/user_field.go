// Code generated by eevee. DO NOT EDIT!

package model

import (
	"context"
	"daoplugin/dao"
	"daoplugin/entity"
	"encoding/json"
	"sort"
	"strconv"
	"time"

	"golang.org/x/xerrors"
)

type UserFieldFinder interface {
	FindAll(context.Context) (*UserFields, error)
	FindByFieldID(context.Context, uint64) (*UserField, error)
	FindByID(context.Context, uint64) (*UserField, error)
	FindByIDs(context.Context, []uint64) (*UserFields, error)
	FindByUserIDs(context.Context, []uint64) (*UserFields, error)
}

type UserField struct {
	*entity.UserField
	userFieldDAO     dao.UserField
	Field            func(context.Context) (*Field, error)
	isAlreadyCreated bool
	savedValue       entity.UserField
	conv             ModelConverter
}

type UserFields struct {
	values   []*UserField
	fieldIDs []uint64
	fields   *Fields
}

type UserFieldsCollection []*UserFields

func NewUserField(value *entity.UserField, userFieldDAO dao.UserField) *UserField {
	return &UserField{
		UserField:    value,
		userFieldDAO: userFieldDAO,
	}
}

func NewUserFields(entities entity.UserFields) *UserFields {
	return &UserFields{
		fieldIDs: entities.FieldIDs(),
		values:   make([]*UserField, 0, len(entities)),
	}
}

func (m *UserFields) newUserFields(values []*UserField) *UserFields {
	return &UserFields{
		fieldIDs: m.fieldIDs,
		fields:   m.fields,
		values:   values,
	}
}

func (m *UserFields) Each(iter func(*UserField)) {
	if m == nil {
		return
	}
	for _, value := range m.values {
		iter(value)
	}
}

func (m *UserFields) EachIndex(iter func(int, *UserField)) {
	if m == nil {
		return
	}
	for idx, value := range m.values {
		iter(idx, value)
	}
}

func (m *UserFields) EachWithError(iter func(*UserField) error) error {
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

func (m *UserFields) EachIndexWithError(iter func(int, *UserField) error) error {
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

func (m *UserFields) Map(mapFunc func(*UserField) *UserField) *UserFields {
	if m == nil {
		return nil
	}
	mappedValues := []*UserField{}
	for _, value := range m.values {
		mappedValue := mapFunc(value)
		if mappedValue != nil {
			mappedValues = append(mappedValues, mappedValue)
		}
	}
	return m.newUserFields(mappedValues)
}

func (m *UserFields) Any(cond func(*UserField) bool) bool {
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

func (m *UserFields) Some(cond func(*UserField) bool) bool {
	return m.Any(cond)
}

func (m *UserFields) IsIncluded(cond func(*UserField) bool) bool {
	return m.Any(cond)
}

func (m *UserFields) All(cond func(*UserField) bool) bool {
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

func (m *UserFields) Sort(compare func(*UserField, *UserField) bool) {
	if m == nil {
		return
	}
	sort.Slice(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *UserFields) SortStable(compare func(*UserField, *UserField) bool) {
	if m == nil {
		return
	}
	sort.SliceStable(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *UserFields) Find(cond func(*UserField) bool) *UserField {
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

func (m *UserFields) Filter(filter func(*UserField) bool) *UserFields {
	if m == nil {
		return nil
	}
	filteredValues := []*UserField{}
	for _, value := range m.values {
		if filter(value) {
			filteredValues = append(filteredValues, value)
		}
	}
	return m.newUserFields(filteredValues)
}

func (m *UserFields) IsEmpty() bool {
	if m == nil {
		return true
	}
	if len(m.values) == 0 {
		return true
	}
	return false
}

func (m *UserFields) At(idx int) *UserField {
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

func (m *UserFields) First() *UserField {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[0]
	}
	return nil
}

func (m *UserFields) Last() *UserField {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[len(m.values)-1]
	}
	return nil
}

func (m *UserFields) Compact() *UserFields {
	if m == nil {
		return nil
	}
	compactedValues := []*UserField{}
	for _, value := range m.values {
		if value == nil {
			continue
		}
		compactedValues = append(compactedValues, value)
	}
	return m.newUserFields(compactedValues)
}

func (m *UserFields) Add(args ...*UserField) *UserFields {
	if m == nil {
		return nil
	}
	for _, value := range args {
		m.values = append(m.values, value)
	}
	return m
}

func (m *UserFields) Merge(args ...*UserFields) *UserFields {
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

func (m *UserFields) Len() int {
	if m == nil {
		return 0
	}
	return len(m.values)
}

func (m *UserFieldsCollection) Merge() *UserFields {
	if m == nil {
		return nil
	}
	if len(*m) == 0 {
		return nil
	}
	if len(*m) == 1 {
		return (*m)[0]
	}
	values := []*UserField{}
	for _, collection := range *m {
		for _, value := range collection.values {
			values = append(values, value)
		}
	}
	return (*m)[0].newUserFields(values)
}

func (m *UserField) ToJSON(ctx context.Context) ([]byte, error) {
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
	buf = append(buf, "\"userID\":"...)
	buf = strconv.AppendUint(buf, m.UserID, 10)
	buf = append(buf, ',')
	buf = append(buf, "\"fieldID\":"...)
	buf = strconv.AppendUint(buf, m.FieldID, 10)
	buf = append(buf, ',')
	buf = append(buf, "\"createdAt\":"...)
	buf = strconv.AppendUint(buf, uint64(m.CreatedAt.Unix()), 10)
	buf = append(buf, ',')
	buf = append(buf, "\"updatedAt\":"...)
	buf = strconv.AppendUint(buf, uint64(m.UpdatedAt.Unix()), 10)
	buf = append(buf, ',')
	field, err := m.Field(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get Field: %w", err)
	}
	fieldBytes, err := field.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	buf = append(buf, "\"field\":"...)
	buf = append(buf, fieldBytes...)
	buf = append(buf, '}')
	return buf, nil
}

func (m *UserField) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
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
	if option.Exists("user_id") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"userID\":"...)
		buf = strconv.AppendUint(buf, m.UserID, 10)
		isWritten = true
	}
	if option.Exists("field_id") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"fieldID\":"...)
		buf = strconv.AppendUint(buf, m.FieldID, 10)
		isWritten = true
	}
	if option.Exists("created_at") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"createdAt\":"...)
		buf = strconv.AppendUint(buf, uint64(m.CreatedAt.Unix()), 10)
		isWritten = true
	}
	if option.Exists("updated_at") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"updatedAt\":"...)
		buf = strconv.AppendUint(buf, uint64(m.UpdatedAt.Unix()), 10)
		isWritten = true
	}
	if option.IsIncludeAll {
		if isWritten {
			buf = append(buf, ',')
		}
		field, err := m.Field(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot get Field: %w", err)
		}
		fieldBytes, err := field.ToJSON(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to JSON: %w", err)
		}
		buf = append(buf, "\"field\":"...)
		buf = append(buf, fieldBytes...)
		isWritten = true
	} else if opt := option.IncludeOption("field"); opt != nil {
		if isWritten {
			buf = append(buf, ',')
		}
		field, err := m.Field(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot get Field: %w", err)
		}
		fieldBytes, err := field.ToJSONWithOption(ctx, opt)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to JSON: %w", err)
		}
		buf = append(buf, "\"field\":"...)
		buf = append(buf, fieldBytes...)
		isWritten = true
	}
	buf = append(buf, '}')
	return buf, nil
}

func (m *UserFields) ToJSON(ctx context.Context) ([]byte, error) {
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

func (m *UserFields) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
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

func (m *UserField) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *UserField) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *UserFields) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *UserFields) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *UserField) UnmarshalJSON(bytes []byte) error {
	var value struct {
		*entity.UserField
	}
	if err := json.Unmarshal(bytes, &value); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.UserField = value.UserField
	return nil
}

func (m *UserFields) UnmarshalJSON(bytes []byte) error {
	var values []*UserField
	if err := json.Unmarshal(bytes, &values); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.values = values
	return nil
}

func (m *UserField) ToMap(ctx context.Context) (map[string]interface{}, error) {
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
	value["userID"] = m.UserID
	value["fieldID"] = m.FieldID
	value["createdAt"] = m.CreatedAt
	value["updatedAt"] = m.UpdatedAt
	field, err := m.Field(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get Field: %w", err)
	}
	fieldMap, err := field.ToMap(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to map: %w", err)
	}
	value["field"] = fieldMap
	return value, nil
}

func (m *UserField) ToMapWithOption(ctx context.Context, option *RenderOption) (map[string]interface{}, error) {
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
	if option.Exists("user_id") {
		value["userID"] = m.UserID
	}
	if option.Exists("field_id") {
		value["fieldID"] = m.FieldID
	}
	if option.Exists("created_at") {
		value["createdAt"] = m.CreatedAt
	}
	if option.Exists("updated_at") {
		value["updatedAt"] = m.UpdatedAt
	}
	if option.IsIncludeAll {
		field, err := m.Field(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot get Field: %w", err)
		}
		fieldMap, err := field.ToMap(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to map: %w", err)
		}
		value["field"] = fieldMap
	} else if opt := option.IncludeOption("field"); opt != nil {
		field, err := m.Field(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot get Field: %w", err)
		}
		fieldMap, err := field.ToMapWithOption(ctx, opt)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to map: %w", err)
		}
		value["field"] = fieldMap
	}
	return value, nil
}

func (m *UserFields) ToMap(ctx context.Context) ([]map[string]interface{}, error) {
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

func (m *UserFields) ToMapWithOption(ctx context.Context, option *RenderOption) ([]map[string]interface{}, error) {
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

func (m *UserField) SetConverter(conv ModelConverter) {
	m.conv = conv
}

func (m *UserField) Create(ctx context.Context) error {
	if m.userFieldDAO == nil {
		// for testing
		return nil
	}
	if m.isAlreadyCreated {
		return xerrors.New("this instance has already created")
	}
	if err := m.userFieldDAO.Create(ctx, m.UserField); err != nil {
		return xerrors.Errorf("failed to Create: %w", err)
	}
	m.savedValue = *m.UserField
	m.isAlreadyCreated = true
	return nil
}

func (m *UserField) Update(ctx context.Context) error {
	if m.userFieldDAO == nil {
		// for testing
		return nil
	}
	isRequiredUpdate := false
	if m.savedValue.ID != m.ID {
		isRequiredUpdate = true
	}
	if m.savedValue.UserID != m.UserID {
		isRequiredUpdate = true
	}
	if m.savedValue.FieldID != m.FieldID {
		isRequiredUpdate = true
	}
	if m.savedValue.CreatedAt != m.CreatedAt {
		isRequiredUpdate = true
	}
	if m.savedValue.UpdatedAt != m.UpdatedAt {
		isRequiredUpdate = true
	}
	if !isRequiredUpdate {
		return nil
	}
	if err := m.userFieldDAO.Update(ctx, m.UserField); err != nil {
		return xerrors.Errorf("failed to Update: %w", err)
	}
	m.savedValue = *m.UserField
	return nil
}

func (m *UserField) Delete(ctx context.Context) error {
	if m.userFieldDAO == nil {
		// for testing
		return nil
	}
	if err := m.userFieldDAO.DeleteByID(ctx, m.ID); err != nil {
		return xerrors.Errorf("failed to Delete: %w", err)
	}
	return nil
}

func (m *UserField) SetAlreadyCreated(isAlreadyCreated bool) {
	m.isAlreadyCreated = isAlreadyCreated
}

func (m *UserField) SetSavedValue(savedValue *entity.UserField) {
	m.savedValue = *savedValue
}

func (m *UserField) Save(ctx context.Context) error {
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

func (m *UserFields) Create(ctx context.Context) error {
	if err := m.EachWithError(func(v *UserField) error {
		if err := v.Create(ctx); err != nil {
			return xerrors.Errorf("failed to Create: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for UserFields: %w", err)
	}
	return nil
}

func (m *UserFields) Update(ctx context.Context) error {
	if err := m.EachWithError(func(v *UserField) error {
		if err := v.Update(ctx); err != nil {
			return xerrors.Errorf("failed to Update: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for UserFields: %w", err)
	}
	return nil
}

func (m *UserFields) Save(ctx context.Context) error {
	if err := m.EachWithError(func(v *UserField) error {
		if err := v.Save(ctx); err != nil {
			return xerrors.Errorf("failed to Save: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for UserFields: %w", err)
	}
	return nil
}

func (m *UserFields) UniqueID() *UserFields {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *UserField) bool {
		if _, exists := filterMap[value.ID]; exists {
			return false
		}
		filterMap[value.ID] = struct{}{}
		return true
	})
}

func (m *UserFields) GroupByID() map[uint64]*UserFields {
	if m == nil {
		return nil
	}
	values := map[uint64]*UserFields{}
	for _, value := range m.values {
		if _, exists := values[value.ID]; !exists {
			values[value.ID] = &UserFields{}
		}
		values[value.ID].Add(value)
	}
	return values
}

func (m *UserFields) IDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.ID)
	}
	return values
}

func (m *UserFields) UniqueUserID() *UserFields {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *UserField) bool {
		if _, exists := filterMap[value.UserID]; exists {
			return false
		}
		filterMap[value.UserID] = struct{}{}
		return true
	})
}

func (m *UserFields) GroupByUserID() map[uint64]*UserFields {
	if m == nil {
		return nil
	}
	values := map[uint64]*UserFields{}
	for _, value := range m.values {
		if _, exists := values[value.UserID]; !exists {
			values[value.UserID] = &UserFields{}
		}
		values[value.UserID].Add(value)
	}
	return values
}

func (m *UserFields) UserIDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.UserID)
	}
	return values
}

func (m *UserFields) UniqueFieldID() *UserFields {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *UserField) bool {
		if _, exists := filterMap[value.FieldID]; exists {
			return false
		}
		filterMap[value.FieldID] = struct{}{}
		return true
	})
}

func (m *UserFields) GroupByFieldID() map[uint64]*UserFields {
	if m == nil {
		return nil
	}
	values := map[uint64]*UserFields{}
	for _, value := range m.values {
		if _, exists := values[value.FieldID]; !exists {
			values[value.FieldID] = &UserFields{}
		}
		values[value.FieldID].Add(value)
	}
	return values
}

func (m *UserFields) FieldIDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.FieldID)
	}
	return values
}

func (m *UserFields) UniqueCreatedAt() *UserFields {
	if m == nil {
		return nil
	}
	filterMap := map[time.Time]struct{}{}
	return m.Filter(func(value *UserField) bool {
		if _, exists := filterMap[value.CreatedAt]; exists {
			return false
		}
		filterMap[value.CreatedAt] = struct{}{}
		return true
	})
}

func (m *UserFields) GroupByCreatedAt() map[time.Time]*UserFields {
	if m == nil {
		return nil
	}
	values := map[time.Time]*UserFields{}
	for _, value := range m.values {
		if _, exists := values[value.CreatedAt]; !exists {
			values[value.CreatedAt] = &UserFields{}
		}
		values[value.CreatedAt].Add(value)
	}
	return values
}

func (m *UserFields) CreatedAts() []time.Time {
	if m == nil {
		return nil
	}
	values := []time.Time{}
	for _, value := range m.values {
		values = append(values, value.CreatedAt)
	}
	return values
}

func (m *UserFields) UniqueUpdatedAt() *UserFields {
	if m == nil {
		return nil
	}
	filterMap := map[time.Time]struct{}{}
	return m.Filter(func(value *UserField) bool {
		if _, exists := filterMap[value.UpdatedAt]; exists {
			return false
		}
		filterMap[value.UpdatedAt] = struct{}{}
		return true
	})
}

func (m *UserFields) GroupByUpdatedAt() map[time.Time]*UserFields {
	if m == nil {
		return nil
	}
	values := map[time.Time]*UserFields{}
	for _, value := range m.values {
		if _, exists := values[value.UpdatedAt]; !exists {
			values[value.UpdatedAt] = &UserFields{}
		}
		values[value.UpdatedAt].Add(value)
	}
	return values
}

func (m *UserFields) UpdatedAts() []time.Time {
	if m == nil {
		return nil
	}
	values := []time.Time{}
	for _, value := range m.values {
		values = append(values, value.UpdatedAt)
	}
	return values
}

func (m *UserFields) FindField(ctx context.Context, fieldID uint64, finder FieldFinder) (*Field, error) {
	if m.fields != nil {
		return m.fields.FirstByID(fieldID), nil
	}
	fields, err := finder.FindByIDs(ctx, m.fieldIDs)
	if err != nil {
		return nil, xerrors.Errorf("failed to FindByIDs: %w", err)
	}
	if fields == nil {
		return nil, xerrors.New("cannot find record")
	}
	m.fields = fields
	return m.fields.FirstByID(fieldID), nil
}

func (m *UserFields) Fields(ctx context.Context) (*Fields, error) {
	if m == nil {
		return nil, nil
	}
	values := &Fields{}
	for _, value := range m.values {
		field, err := value.Field(ctx)
		if err != nil {
			return nil, xerrors.Errorf("failed to get Field: %w", err)
		}
		if field == nil {
			continue
		}
		values.Add(field)
	}
	return values, nil
}

func (m *UserFields) FirstByID(a0 uint64) *UserField {
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

func (m *UserFields) FilterByID(a0 uint64) *UserFields {
	if m == nil {
		return nil
	}
	values := []*UserField{}
	for _, value := range m.values {
		if value.ID != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newUserFields(values)
}

func (m *UserFields) FirstByUserIDAndFieldID(a0 uint64, a1 uint64) *UserField {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.UserID != a0 {
			continue
		}
		if value.FieldID != a1 {
			continue
		}
		return value
	}
	return nil
}

func (m *UserFields) FilterByUserIDAndFieldID(a0 uint64, a1 uint64) *UserFields {
	if m == nil {
		return nil
	}
	values := []*UserField{}
	for _, value := range m.values {
		if value.UserID != a0 {
			continue
		}
		if value.FieldID != a1 {
			continue
		}
		values = append(values, value)
	}
	return m.newUserFields(values)
}

func (m *UserFields) FirstByUserID(a0 uint64) *UserField {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.UserID != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *UserFields) FilterByUserID(a0 uint64) *UserFields {
	if m == nil {
		return nil
	}
	values := []*UserField{}
	for _, value := range m.values {
		if value.UserID != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newUserFields(values)
}
