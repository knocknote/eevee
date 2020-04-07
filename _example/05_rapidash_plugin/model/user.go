// Code generated by eevee. DO NOT EDIT!

package model

import (
	"context"
	"encoding/json"
	"rapidashplugin/dao"
	"rapidashplugin/entity"
	"sort"
	"strconv"
	"time"

	"golang.org/x/xerrors"
)

type UserFinder interface {
	FindAll(context.Context) (*Users, error)
	FindByGroupID(context.Context, uint64) (*Users, error)
	FindByGroupIDs(context.Context, []uint64) (*Users, error)
	FindByID(context.Context, uint64) (*User, error)
	FindByIDs(context.Context, []uint64) (*Users, error)
	FindByName(context.Context, string) (*User, error)
	FindByNames(context.Context, []string) (*Users, error)
	FindBySkillID(context.Context, uint64) (*Users, error)
	FindBySkillIDAndSkillRank(context.Context, uint64, int) (*User, error)
	FindBySkillIDs(context.Context, []uint64) (*Users, error)
	FindByWorldID(context.Context, uint64) (*Users, error)
	FindByWorldIDAndFieldID(context.Context, uint64, uint64) (*Users, error)
	FindByWorldIDs(context.Context, []uint64) (*Users, error)
}

type User struct {
	*entity.User
	userDAO          dao.User
	isAlreadyCreated bool
	savedValue       entity.User
	conv             ModelConverter
}

type Users struct {
	values []*User
}

type UsersCollection []*Users

func NewUser(value *entity.User, userDAO dao.User) *User {
	return &User{
		User:    value,
		userDAO: userDAO,
	}
}

func NewUsers(entities entity.Users) *Users {
	return &Users{values: make([]*User, 0, len(entities))}
}

func (m *Users) newUsers(values []*User) *Users {
	return &Users{values: values}
}

func (m *Users) Each(iter func(*User)) {
	if m == nil {
		return
	}
	for _, value := range m.values {
		iter(value)
	}
}

func (m *Users) EachIndex(iter func(int, *User)) {
	if m == nil {
		return
	}
	for idx, value := range m.values {
		iter(idx, value)
	}
}

func (m *Users) EachWithError(iter func(*User) error) error {
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

func (m *Users) EachIndexWithError(iter func(int, *User) error) error {
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

func (m *Users) Map(mapFunc func(*User) *User) *Users {
	if m == nil {
		return nil
	}
	mappedValues := []*User{}
	for _, value := range m.values {
		mappedValue := mapFunc(value)
		if mappedValue != nil {
			mappedValues = append(mappedValues, mappedValue)
		}
	}
	return m.newUsers(mappedValues)
}

func (m *Users) Any(cond func(*User) bool) bool {
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

func (m *Users) Some(cond func(*User) bool) bool {
	return m.Any(cond)
}

func (m *Users) IsIncluded(cond func(*User) bool) bool {
	return m.Any(cond)
}

func (m *Users) All(cond func(*User) bool) bool {
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

func (m *Users) Sort(compare func(*User, *User) bool) {
	if m == nil {
		return
	}
	sort.Slice(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *Users) SortStable(compare func(*User, *User) bool) {
	if m == nil {
		return
	}
	sort.SliceStable(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *Users) Find(cond func(*User) bool) *User {
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

func (m *Users) Filter(filter func(*User) bool) *Users {
	if m == nil {
		return nil
	}
	filteredValues := []*User{}
	for _, value := range m.values {
		if filter(value) {
			filteredValues = append(filteredValues, value)
		}
	}
	return m.newUsers(filteredValues)
}

func (m *Users) IsEmpty() bool {
	if m == nil {
		return true
	}
	if len(m.values) == 0 {
		return true
	}
	return false
}

func (m *Users) At(idx int) *User {
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

func (m *Users) First() *User {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[0]
	}
	return nil
}

func (m *Users) Last() *User {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[len(m.values)-1]
	}
	return nil
}

func (m *Users) Compact() *Users {
	if m == nil {
		return nil
	}
	compactedValues := []*User{}
	for _, value := range m.values {
		if value == nil {
			continue
		}
		compactedValues = append(compactedValues, value)
	}
	return m.newUsers(compactedValues)
}

func (m *Users) Add(args ...*User) *Users {
	if m == nil {
		return nil
	}
	for _, value := range args {
		m.values = append(m.values, value)
	}
	return m
}

func (m *Users) Merge(args ...*Users) *Users {
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

func (m *Users) Len() int {
	if m == nil {
		return 0
	}
	return len(m.values)
}

func (m *UsersCollection) Merge() *Users {
	if m == nil {
		return nil
	}
	if len(*m) == 0 {
		return nil
	}
	if len(*m) == 1 {
		return (*m)[0]
	}
	values := []*User{}
	for _, collection := range *m {
		for _, value := range collection.values {
			values = append(values, value)
		}
	}
	return (*m)[0].newUsers(values)
}

func (m *User) ToJSON(ctx context.Context) ([]byte, error) {
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
	buf = append(buf, "\"name\":"...)
	buf = append(buf, strconv.Quote(m.Name)...)
	buf = append(buf, ',')
	buf = append(buf, "\"sex\":"...)
	buf = append(buf, strconv.Quote(m.Sex)...)
	buf = append(buf, ',')
	buf = append(buf, "\"age\":"...)
	buf = strconv.AppendInt(buf, int64(m.Age), 10)
	buf = append(buf, ',')
	buf = append(buf, "\"skillID\":"...)
	buf = strconv.AppendUint(buf, m.SkillID, 10)
	buf = append(buf, ',')
	buf = append(buf, "\"skillRank\":"...)
	buf = strconv.AppendInt(buf, int64(m.SkillRank), 10)
	buf = append(buf, ',')
	buf = append(buf, "\"groupID\":"...)
	buf = strconv.AppendUint(buf, m.GroupID, 10)
	buf = append(buf, ',')
	buf = append(buf, "\"worldID\":"...)
	buf = strconv.AppendUint(buf, m.WorldID, 10)
	buf = append(buf, ',')
	buf = append(buf, "\"fieldID\":"...)
	buf = strconv.AppendUint(buf, m.FieldID, 10)
	buf = append(buf, ',')
	buf = append(buf, "\"createdAt\":"...)
	buf = strconv.AppendUint(buf, uint64(m.CreatedAt.Unix()), 10)
	buf = append(buf, ',')
	buf = append(buf, "\"updatedAt\":"...)
	buf = strconv.AppendUint(buf, uint64(m.UpdatedAt.Unix()), 10)
	buf = append(buf, '}')
	return buf, nil
}

func (m *User) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
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
	if option.Exists("name") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"name\":"...)
		buf = append(buf, strconv.Quote(m.Name)...)
		isWritten = true
	}
	if option.Exists("sex") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"sex\":"...)
		buf = append(buf, strconv.Quote(m.Sex)...)
		isWritten = true
	}
	if option.Exists("age") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"age\":"...)
		buf = strconv.AppendInt(buf, int64(m.Age), 10)
		isWritten = true
	}
	if option.Exists("skill_id") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"skillID\":"...)
		buf = strconv.AppendUint(buf, m.SkillID, 10)
		isWritten = true
	}
	if option.Exists("skill_rank") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"skillRank\":"...)
		buf = strconv.AppendInt(buf, int64(m.SkillRank), 10)
		isWritten = true
	}
	if option.Exists("group_id") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"groupID\":"...)
		buf = strconv.AppendUint(buf, m.GroupID, 10)
		isWritten = true
	}
	if option.Exists("world_id") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"worldID\":"...)
		buf = strconv.AppendUint(buf, m.WorldID, 10)
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
	buf = append(buf, '}')
	return buf, nil
}

func (m *Users) ToJSON(ctx context.Context) ([]byte, error) {
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

func (m *Users) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
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

func (m *User) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *User) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Users) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Users) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *User) UnmarshalJSON(bytes []byte) error {
	var value struct {
		*entity.User
	}
	if err := json.Unmarshal(bytes, &value); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.User = value.User
	return nil
}

func (m *Users) UnmarshalJSON(bytes []byte) error {
	var values []*User
	if err := json.Unmarshal(bytes, &values); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.values = values
	return nil
}

func (m *User) ToMap(ctx context.Context) (map[string]interface{}, error) {
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
	value["name"] = m.Name
	value["sex"] = m.Sex
	value["age"] = m.Age
	value["skillID"] = m.SkillID
	value["skillRank"] = m.SkillRank
	value["groupID"] = m.GroupID
	value["worldID"] = m.WorldID
	value["fieldID"] = m.FieldID
	value["createdAt"] = m.CreatedAt
	value["updatedAt"] = m.UpdatedAt
	return value, nil
}

func (m *User) ToMapWithOption(ctx context.Context, option *RenderOption) (map[string]interface{}, error) {
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
	if option.Exists("name") {
		value["name"] = m.Name
	}
	if option.Exists("sex") {
		value["sex"] = m.Sex
	}
	if option.Exists("age") {
		value["age"] = m.Age
	}
	if option.Exists("skill_id") {
		value["skillID"] = m.SkillID
	}
	if option.Exists("skill_rank") {
		value["skillRank"] = m.SkillRank
	}
	if option.Exists("group_id") {
		value["groupID"] = m.GroupID
	}
	if option.Exists("world_id") {
		value["worldID"] = m.WorldID
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
	return value, nil
}

func (m *Users) ToMap(ctx context.Context) ([]map[string]interface{}, error) {
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

func (m *Users) ToMapWithOption(ctx context.Context, option *RenderOption) ([]map[string]interface{}, error) {
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

func (m *User) SetConverter(conv ModelConverter) {
	m.conv = conv
}

func (m *User) Create(ctx context.Context) error {
	if m.userDAO == nil {
		// for testing
		return nil
	}
	if m.isAlreadyCreated {
		return xerrors.New("this instance has already created")
	}
	if err := m.userDAO.Create(ctx, m.User); err != nil {
		return xerrors.Errorf("failed to Create: %w", err)
	}
	m.savedValue = *m.User
	m.isAlreadyCreated = true
	return nil
}

func (m *User) Update(ctx context.Context) error {
	if m.userDAO == nil {
		// for testing
		return nil
	}
	isRequiredUpdate := false
	if m.savedValue.ID != m.ID {
		isRequiredUpdate = true
	}
	if m.savedValue.Name != m.Name {
		isRequiredUpdate = true
	}
	if m.savedValue.Sex != m.Sex {
		isRequiredUpdate = true
	}
	if m.savedValue.Age != m.Age {
		isRequiredUpdate = true
	}
	if m.savedValue.SkillID != m.SkillID {
		isRequiredUpdate = true
	}
	if m.savedValue.SkillRank != m.SkillRank {
		isRequiredUpdate = true
	}
	if m.savedValue.GroupID != m.GroupID {
		isRequiredUpdate = true
	}
	if m.savedValue.WorldID != m.WorldID {
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
	if err := m.userDAO.Update(ctx, m.User); err != nil {
		return xerrors.Errorf("failed to Update: %w", err)
	}
	m.savedValue = *m.User
	return nil
}

func (m *User) Delete(ctx context.Context) error {
	if m.userDAO == nil {
		// for testing
		return nil
	}
	if err := m.userDAO.DeleteByID(ctx, m.ID); err != nil {
		return xerrors.Errorf("failed to Delete: %w", err)
	}
	return nil
}

func (m *User) SetAlreadyCreated(isAlreadyCreated bool) {
	m.isAlreadyCreated = isAlreadyCreated
}

func (m *User) SetSavedValue(savedValue *entity.User) {
	m.savedValue = *savedValue
}

func (m *User) Save(ctx context.Context) error {
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

func (m *Users) Create(ctx context.Context) error {
	if err := m.EachWithError(func(v *User) error {
		if err := v.Create(ctx); err != nil {
			return xerrors.Errorf("failed to Create: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Users: %w", err)
	}
	return nil
}

func (m *Users) Update(ctx context.Context) error {
	if err := m.EachWithError(func(v *User) error {
		if err := v.Update(ctx); err != nil {
			return xerrors.Errorf("failed to Update: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Users: %w", err)
	}
	return nil
}

func (m *Users) Save(ctx context.Context) error {
	if err := m.EachWithError(func(v *User) error {
		if err := v.Save(ctx); err != nil {
			return xerrors.Errorf("failed to Save: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Users: %w", err)
	}
	return nil
}

func (m *Users) UniqueID() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.ID]; exists {
			return false
		}
		filterMap[value.ID] = struct{}{}
		return true
	})
}

func (m *Users) GroupByID() map[uint64]*Users {
	if m == nil {
		return nil
	}
	values := map[uint64]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.ID]; !exists {
			values[value.ID] = &Users{}
		}
		values[value.ID].Add(value)
	}
	return values
}

func (m *Users) IDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.ID)
	}
	return values
}

func (m *Users) UniqueName() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[string]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.Name]; exists {
			return false
		}
		filterMap[value.Name] = struct{}{}
		return true
	})
}

func (m *Users) GroupByName() map[string]*Users {
	if m == nil {
		return nil
	}
	values := map[string]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.Name]; !exists {
			values[value.Name] = &Users{}
		}
		values[value.Name].Add(value)
	}
	return values
}

func (m *Users) Names() []string {
	if m == nil {
		return nil
	}
	values := []string{}
	for _, value := range m.values {
		values = append(values, value.Name)
	}
	return values
}

func (m *Users) UniqueSex() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[string]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.Sex]; exists {
			return false
		}
		filterMap[value.Sex] = struct{}{}
		return true
	})
}

func (m *Users) GroupBySex() map[string]*Users {
	if m == nil {
		return nil
	}
	values := map[string]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.Sex]; !exists {
			values[value.Sex] = &Users{}
		}
		values[value.Sex].Add(value)
	}
	return values
}

func (m *Users) Sexes() []string {
	if m == nil {
		return nil
	}
	values := []string{}
	for _, value := range m.values {
		values = append(values, value.Sex)
	}
	return values
}

func (m *Users) UniqueAge() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[int]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.Age]; exists {
			return false
		}
		filterMap[value.Age] = struct{}{}
		return true
	})
}

func (m *Users) GroupByAge() map[int]*Users {
	if m == nil {
		return nil
	}
	values := map[int]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.Age]; !exists {
			values[value.Age] = &Users{}
		}
		values[value.Age].Add(value)
	}
	return values
}

func (m *Users) Ages() []int {
	if m == nil {
		return nil
	}
	values := []int{}
	for _, value := range m.values {
		values = append(values, value.Age)
	}
	return values
}

func (m *Users) UniqueSkillID() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.SkillID]; exists {
			return false
		}
		filterMap[value.SkillID] = struct{}{}
		return true
	})
}

func (m *Users) GroupBySkillID() map[uint64]*Users {
	if m == nil {
		return nil
	}
	values := map[uint64]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.SkillID]; !exists {
			values[value.SkillID] = &Users{}
		}
		values[value.SkillID].Add(value)
	}
	return values
}

func (m *Users) SkillIDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.SkillID)
	}
	return values
}

func (m *Users) UniqueSkillRank() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[int]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.SkillRank]; exists {
			return false
		}
		filterMap[value.SkillRank] = struct{}{}
		return true
	})
}

func (m *Users) GroupBySkillRank() map[int]*Users {
	if m == nil {
		return nil
	}
	values := map[int]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.SkillRank]; !exists {
			values[value.SkillRank] = &Users{}
		}
		values[value.SkillRank].Add(value)
	}
	return values
}

func (m *Users) SkillRanks() []int {
	if m == nil {
		return nil
	}
	values := []int{}
	for _, value := range m.values {
		values = append(values, value.SkillRank)
	}
	return values
}

func (m *Users) UniqueGroupID() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.GroupID]; exists {
			return false
		}
		filterMap[value.GroupID] = struct{}{}
		return true
	})
}

func (m *Users) GroupByGroupID() map[uint64]*Users {
	if m == nil {
		return nil
	}
	values := map[uint64]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.GroupID]; !exists {
			values[value.GroupID] = &Users{}
		}
		values[value.GroupID].Add(value)
	}
	return values
}

func (m *Users) GroupIDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.GroupID)
	}
	return values
}

func (m *Users) UniqueWorldID() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.WorldID]; exists {
			return false
		}
		filterMap[value.WorldID] = struct{}{}
		return true
	})
}

func (m *Users) GroupByWorldID() map[uint64]*Users {
	if m == nil {
		return nil
	}
	values := map[uint64]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.WorldID]; !exists {
			values[value.WorldID] = &Users{}
		}
		values[value.WorldID].Add(value)
	}
	return values
}

func (m *Users) WorldIDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.WorldID)
	}
	return values
}

func (m *Users) UniqueFieldID() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.FieldID]; exists {
			return false
		}
		filterMap[value.FieldID] = struct{}{}
		return true
	})
}

func (m *Users) GroupByFieldID() map[uint64]*Users {
	if m == nil {
		return nil
	}
	values := map[uint64]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.FieldID]; !exists {
			values[value.FieldID] = &Users{}
		}
		values[value.FieldID].Add(value)
	}
	return values
}

func (m *Users) FieldIDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.FieldID)
	}
	return values
}

func (m *Users) UniqueCreatedAt() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[time.Time]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.CreatedAt]; exists {
			return false
		}
		filterMap[value.CreatedAt] = struct{}{}
		return true
	})
}

func (m *Users) GroupByCreatedAt() map[time.Time]*Users {
	if m == nil {
		return nil
	}
	values := map[time.Time]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.CreatedAt]; !exists {
			values[value.CreatedAt] = &Users{}
		}
		values[value.CreatedAt].Add(value)
	}
	return values
}

func (m *Users) CreatedAts() []time.Time {
	if m == nil {
		return nil
	}
	values := []time.Time{}
	for _, value := range m.values {
		values = append(values, value.CreatedAt)
	}
	return values
}

func (m *Users) UniqueUpdatedAt() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[time.Time]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.UpdatedAt]; exists {
			return false
		}
		filterMap[value.UpdatedAt] = struct{}{}
		return true
	})
}

func (m *Users) GroupByUpdatedAt() map[time.Time]*Users {
	if m == nil {
		return nil
	}
	values := map[time.Time]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.UpdatedAt]; !exists {
			values[value.UpdatedAt] = &Users{}
		}
		values[value.UpdatedAt].Add(value)
	}
	return values
}

func (m *Users) UpdatedAts() []time.Time {
	if m == nil {
		return nil
	}
	values := []time.Time{}
	for _, value := range m.values {
		values = append(values, value.UpdatedAt)
	}
	return values
}

func (m *Users) FirstByID(a0 uint64) *User {
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

func (m *Users) FilterByID(a0 uint64) *Users {
	if m == nil {
		return nil
	}
	values := []*User{}
	for _, value := range m.values {
		if value.ID != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newUsers(values)
}

func (m *Users) FirstByName(a0 string) *User {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.Name != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Users) FilterByName(a0 string) *Users {
	if m == nil {
		return nil
	}
	values := []*User{}
	for _, value := range m.values {
		if value.Name != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newUsers(values)
}

func (m *Users) FirstBySkillIDAndSkillRank(a0 uint64, a1 int) *User {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.SkillID != a0 {
			continue
		}
		if value.SkillRank != a1 {
			continue
		}
		return value
	}
	return nil
}

func (m *Users) FilterBySkillIDAndSkillRank(a0 uint64, a1 int) *Users {
	if m == nil {
		return nil
	}
	values := []*User{}
	for _, value := range m.values {
		if value.SkillID != a0 {
			continue
		}
		if value.SkillRank != a1 {
			continue
		}
		values = append(values, value)
	}
	return m.newUsers(values)
}

func (m *Users) FirstBySkillID(a0 uint64) *User {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.SkillID != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Users) FilterBySkillID(a0 uint64) *Users {
	if m == nil {
		return nil
	}
	values := []*User{}
	for _, value := range m.values {
		if value.SkillID != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newUsers(values)
}

func (m *Users) FirstByGroupID(a0 uint64) *User {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.GroupID != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Users) FilterByGroupID(a0 uint64) *Users {
	if m == nil {
		return nil
	}
	values := []*User{}
	for _, value := range m.values {
		if value.GroupID != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newUsers(values)
}

func (m *Users) FirstByWorldIDAndFieldID(a0 uint64, a1 uint64) *User {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.WorldID != a0 {
			continue
		}
		if value.FieldID != a1 {
			continue
		}
		return value
	}
	return nil
}

func (m *Users) FilterByWorldIDAndFieldID(a0 uint64, a1 uint64) *Users {
	if m == nil {
		return nil
	}
	values := []*User{}
	for _, value := range m.values {
		if value.WorldID != a0 {
			continue
		}
		if value.FieldID != a1 {
			continue
		}
		values = append(values, value)
	}
	return m.newUsers(values)
}

func (m *Users) FirstByWorldID(a0 uint64) *User {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.WorldID != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Users) FilterByWorldID(a0 uint64) *Users {
	if m == nil {
		return nil
	}
	values := []*User{}
	for _, value := range m.values {
		if value.WorldID != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newUsers(values)
}
