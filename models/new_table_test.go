// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testNewTables(t *testing.T) {
	t.Parallel()

	query := NewTables(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testNewTablesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = newTable.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNewTablesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = NewTables(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNewTablesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := NewTableSlice{newTable}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testNewTablesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := NewTableExists(tx, newTable.ID)
	if err != nil {
		t.Errorf("Unable to check if NewTable exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NewTableExistsG to return true, but got false.")
	}
}
func testNewTablesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	newTableFound, err := FindNewTable(tx, newTable.ID)
	if err != nil {
		t.Error(err)
	}

	if newTableFound == nil {
		t.Error("want a record, got nil")
	}
}
func testNewTablesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = NewTables(tx).Bind(newTable); err != nil {
		t.Error(err)
	}
}

func testNewTablesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := NewTables(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNewTablesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTableOne := &NewTable{}
	newTableTwo := &NewTable{}
	if err = randomize.Struct(seed, newTableOne, newTableDBTypes, false, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}
	if err = randomize.Struct(seed, newTableTwo, newTableDBTypes, false, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTableOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = newTableTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := NewTables(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNewTablesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	newTableOne := &NewTable{}
	newTableTwo := &NewTable{}
	if err = randomize.Struct(seed, newTableOne, newTableDBTypes, false, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}
	if err = randomize.Struct(seed, newTableTwo, newTableDBTypes, false, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTableOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = newTableTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func newTableBeforeInsertHook(e boil.Executor, o *NewTable) error {
	*o = NewTable{}
	return nil
}

func newTableAfterInsertHook(e boil.Executor, o *NewTable) error {
	*o = NewTable{}
	return nil
}

func newTableAfterSelectHook(e boil.Executor, o *NewTable) error {
	*o = NewTable{}
	return nil
}

func newTableBeforeUpdateHook(e boil.Executor, o *NewTable) error {
	*o = NewTable{}
	return nil
}

func newTableAfterUpdateHook(e boil.Executor, o *NewTable) error {
	*o = NewTable{}
	return nil
}

func newTableBeforeDeleteHook(e boil.Executor, o *NewTable) error {
	*o = NewTable{}
	return nil
}

func newTableAfterDeleteHook(e boil.Executor, o *NewTable) error {
	*o = NewTable{}
	return nil
}

func newTableBeforeUpsertHook(e boil.Executor, o *NewTable) error {
	*o = NewTable{}
	return nil
}

func newTableAfterUpsertHook(e boil.Executor, o *NewTable) error {
	*o = NewTable{}
	return nil
}

func testNewTablesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &NewTable{}
	o := &NewTable{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, newTableDBTypes, false); err != nil {
		t.Errorf("Unable to randomize NewTable object: %s", err)
	}

	AddNewTableHook(boil.BeforeInsertHook, newTableBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	newTableBeforeInsertHooks = []NewTableHook{}

	AddNewTableHook(boil.AfterInsertHook, newTableAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	newTableAfterInsertHooks = []NewTableHook{}

	AddNewTableHook(boil.AfterSelectHook, newTableAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	newTableAfterSelectHooks = []NewTableHook{}

	AddNewTableHook(boil.BeforeUpdateHook, newTableBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	newTableBeforeUpdateHooks = []NewTableHook{}

	AddNewTableHook(boil.AfterUpdateHook, newTableAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	newTableAfterUpdateHooks = []NewTableHook{}

	AddNewTableHook(boil.BeforeDeleteHook, newTableBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	newTableBeforeDeleteHooks = []NewTableHook{}

	AddNewTableHook(boil.AfterDeleteHook, newTableAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	newTableAfterDeleteHooks = []NewTableHook{}

	AddNewTableHook(boil.BeforeUpsertHook, newTableBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	newTableBeforeUpsertHooks = []NewTableHook{}

	AddNewTableHook(boil.AfterUpsertHook, newTableAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	newTableAfterUpsertHooks = []NewTableHook{}
}
func testNewTablesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNewTablesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx, newTableColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNewTablesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = newTable.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testNewTablesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := NewTableSlice{newTable}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testNewTablesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := NewTables(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	newTableDBTypes = map[string]string{`Exchangeid`: `numeric`, `Filltype`: `character varying`, `Globaltradeid`: `numeric`, `ID`: `integer`, `Ordertype`: `character varying`, `Price`: `numeric`, `Quantity`: `numeric`, `Timestamping`: `timestamp without time zone`, `Total`: `numeric`, `Tradeid`: `numeric`}
	_               = bytes.MinRead
)

func testNewTablesUpdate(t *testing.T) {
	t.Parallel()

	if len(newTableColumns) == len(newTablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	if err = newTable.Update(tx); err != nil {
		t.Error(err)
	}
}

func testNewTablesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(newTableColumns) == len(newTablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	newTable := &NewTable{}
	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTableColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, newTable, newTableDBTypes, true, newTablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(newTableColumns, newTablePrimaryKeyColumns) {
		fields = newTableColumns
	} else {
		fields = strmangle.SetComplement(
			newTableColumns,
			newTablePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(newTable))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := NewTableSlice{newTable}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testNewTablesUpsert(t *testing.T) {
	t.Parallel()

	if len(newTableColumns) == len(newTablePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	newTable := NewTable{}
	if err = randomize.Struct(seed, &newTable, newTableDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = newTable.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert NewTable: %s", err)
	}

	count, err := NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &newTable, newTableDBTypes, false, newTablePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NewTable struct: %s", err)
	}

	if err = newTable.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert NewTable: %s", err)
	}

	count, err = NewTables(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
