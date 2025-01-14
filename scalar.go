package graphql

// Note: These custom types are meant to be used in queries for now.
// But the plan is to switch to using native Go types (string, int, bool, time.Time, etc.).
// See https://github.com/shurcooL/githubv4/issues/9 for details.
//
// These custom types currently provide documentation, and their use
// is required for sending outbound queries. However, native Go types
// can be used for unmarshaling. Once https://github.com/shurcooL/githubv4/issues/9
// is resolved, native Go types can completely replace these.

type (
	// Boolean represents true or false values.
	Boolean bool

	// Float represents signed double-precision fractional values as
	// specified by IEEE 754.
	Float float64

	// ID represents a unique identifier that is Base64 obfuscated. It
	// is often used to refetch an object or as key for a cache. The ID
	// type appears in a JSON response as a String; however, it is not
	// intended to be human-readable. When expected as an input type,
	// any string (such as "VXNlci0xMA==") or integer (such as 4) input
	// value will be accepted as an ID.
	ID interface{}

	// Int represents non-fractional signed whole numeric values.
	// Int can represent values between -(2^31) and 2^31 - 1.
	Int int32

	// Int64Bit represents non-fractional signed whole numeric values.
	// Int64Bit can represent values between -(2^63) and 2^63 - 1.
	Int64Bit int64

	// String represents textual data as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	String string

	// Datetime represents textual datetime as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	Datetime string

	// Text represents textual data as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	Text string

	// GeoPoint represents textual data as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	GeoPoint string

	// EmailAddress represents textual data as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	EmailAddress string

	// InventoryitemableType represents textual data as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	InventoryitemableType string

	// InventoryItemDeviceStatus represents textual data as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	InventoryItemDeviceStatus string

	// PreseemStatus represents textual data as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	PreseemStatus string

	// Subdivision represents textual data as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	Subdivision string
)

// NewBoolean is a helper to make a new *Boolean.
func NewBoolean(v Boolean) *Boolean { return &v }

// NewFloat is a helper to make a new *Float.
func NewFloat(v Float) *Float { return &v }

// NewID is a helper to make a new *ID.
func NewID(v ID) *ID { return &v }

// NewInt is a helper to make a new *Int.
func NewInt(v Int) *Int { return &v }

// NewInt64Bit is a helper to make a new *Int64Bit.
func NewInt64Bit(v Int64Bit) *Int64Bit { return &v }

// NewString is a helper to make a new *String.
func NewString(v String) *String { return &v }

// NewDatetime is a helper to make a new *Datetime.
func NewDatetime(v Datetime) *Datetime { return &v }

// NewText is a helper to make a new *Text.
func NewText(v Text) *Text { return &v }

// NewGeoPoint is a helper to make a new *GeoPoint.
func NewGeoPoint(v GeoPoint) *GeoPoint { return &v }

// NewGeoPoint is a helper to make a new *GeoPoint.
func NewEmailAddress(v EmailAddress) *EmailAddress { return &v }

// NewInventoryitemableType is a helper to make a new *InventoryitemableType.
func NewInventoryitemableType(v InventoryitemableType) *InventoryitemableType { return &v }

// NewInventoryItemDeviceStatus is a helper to make a new *InventoryItemDeviceStatus.
func NewInventoryItemDeviceStatus(v InventoryItemDeviceStatus) *InventoryItemDeviceStatus { return &v }

// NewPreseemStatus is a helper to make a new *PreseemStatus.
func NewPreseemStatus(v PreseemStatus) *PreseemStatus { return &v }

// NewSubdivision is a helper to make a new *Subdivision.
func NewSubdivision(v Subdivision) *Subdivision { return &v }
