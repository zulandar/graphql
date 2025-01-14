package graphql_test

import (
	"testing"

	"github.com/zulandar/graphql"
)

func TestNewScalars(t *testing.T) {
	if got := graphql.NewBoolean(false); got == nil {
		t.Error("NewBoolean returned nil")
	}
	if got := graphql.NewFloat(0.0); got == nil {
		t.Error("NewFloat returned nil")
	}
	// ID with underlying type string.
	if got := graphql.NewID(""); got == nil {
		t.Error("NewID returned nil")
	}
	// ID with underlying type int.
	if got := graphql.NewID(0); got == nil {
		t.Error("NewID returned nil")
	}
	if got := graphql.NewInt(0); got == nil {
		t.Error("NewInt returned nil")
	}
	if got := graphql.NewInt64Bit(0); got == nil {
		t.Error("NewInt64Bit returned nil")
	}
	if got := graphql.NewString(""); got == nil {
		t.Error("NewString returned nil")
	}
	if got := graphql.NewDatetime(""); got == nil {
		t.Error("NewDatetime returned nil")
	}
	if got := graphql.NewText(""); got == nil {
		t.Error("NewText returned nil")
	}
	if got := graphql.NewGeoPoint(""); got == nil {
		t.Error("NewGeoPoint returned nil")
	}
}
