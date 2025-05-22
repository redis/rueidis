package rueidiscompat

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StructMap", func() {
	Context("newStructSpec", func() {
		It("should parse struct tags correctly", func() {
			type TestStruct struct {
				FieldA string `redis:"field_a"`
				FieldB int    `redis:"field_b,omitempty"`
				FieldC bool   `redis:"-"`
				FieldD float64
				FieldE *string `redis:"field_e"`
			}
			spec := newStructSpec(reflect.TypeOf(TestStruct{}), "redis")
			Expect(spec.m).To(HaveLen(3)) // FieldC and FieldD should be ignored
			Expect(spec.m).To(HaveKey("field_a"))
			Expect(spec.m["field_a"].index).To(Equal(0))
			Expect(spec.m).To(HaveKey("field_b"))
			Expect(spec.m["field_b"].index).To(Equal(1))
			Expect(spec.m).To(HaveKey("field_e"))
			Expect(spec.m["field_e"].index).To(Equal(4)) // Index of FieldE
		})

		It("should handle empty tags or missing tags", func() {
			type TestStruct struct {
				FieldA string `redis:""`
				FieldB int
				FieldC bool `redis:","` // Empty tag value after comma
			}
			spec := newStructSpec(reflect.TypeOf(TestStruct{}), "redis")
			Expect(spec.m).To(BeEmpty())
		})
	})

	Context("StructValue.Scan", func() {
		type SimpleStruct struct {
			StringField  string    `redis:"string_field"`
			IntField     int       `redis:"int_field"`
			BoolField    bool      `redis:"bool_field"`
			FloatField   float64   `redis:"float_field"`
			IntPtrField  *int      `redis:"int_ptr_field"`
			BoolPtrField *bool     `redis:"bool_ptr_field"`
			TimeField    time.Time `redis:"time_field"`
		}

		var s SimpleStruct
		var sv StructValue

		BeforeEach(func() {
			// Reset struct before each test
			s = SimpleStruct{}
			spec := newStructSpec(reflect.TypeOf(SimpleStruct{}), "redis")
			sv = StructValue{
				spec:  spec,
				value: reflect.ValueOf(&s).Elem(),
			}
		})

		It("should scan string fields", func() {
			err := sv.Scan("string_field", "hello world")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.StringField).To(Equal("hello world"))
		})

		It("should scan int fields", func() {
			err := sv.Scan("int_field", "12345")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.IntField).To(Equal(12345))
		})

		It("should return error for invalid int fields", func() {
			err := sv.Scan("int_field", "not-an-int")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("cannot scan redis.result not-an-int into struct field SimpleStruct.IntField of type int"))
		})
		
		It("should scan bool fields (true)", func() {
			err := sv.Scan("bool_field", "1")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.BoolField).To(BeTrue())
		})

		It("should scan bool fields (t)", func() {
			err := sv.Scan("bool_field", "t")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.BoolField).To(BeTrue())
		})

		It("should scan bool fields (false)", func() {
			err := sv.Scan("bool_field", "0")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.BoolField).To(BeFalse())
		})

		It("should scan bool fields (f)", func() {
			err := sv.Scan("bool_field", "f")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.BoolField).To(BeFalse())
		})

		It("should return error for invalid bool fields", func() {
			err := sv.Scan("bool_field", "not-a-bool")
			Expect(err).To(HaveOccurred())
		})

		It("should scan float fields", func() {
			err := sv.Scan("float_field", "123.45")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.FloatField).To(Equal(123.45))
		})

		It("should return error for invalid float fields", func() {
			err := sv.Scan("float_field", "not-a-float")
			Expect(err).To(HaveOccurred())
		})
		
		It("should scan int pointer fields", func() {
			err := sv.Scan("int_ptr_field", "67890")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.IntPtrField).NotTo(BeNil())
			Expect(*s.IntPtrField).To(Equal(67890))
		})

		It("should scan bool pointer fields (true)", func() {
			err := sv.Scan("bool_ptr_field", "true")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.BoolPtrField).NotTo(BeNil())
			Expect(*s.BoolPtrField).To(BeTrue())
		})
		
		It("should scan time.Time fields with RFC3339Nano format", func() {
			now := time.Now()
			timeStr := now.Format(time.RFC3339Nano)
			err := sv.Scan("time_field", timeStr)
			Expect(err).NotTo(HaveOccurred())
			Expect(s.TimeField.UnixNano()).To(Equal(now.UnixNano()))
		})

		It("should return error for invalid time.Time fields", func() {
			err := sv.Scan("time_field", "not-a-time")
			Expect(err).To(HaveOccurred())
		})

		It("should ignore fields not in struct spec", func() {
			err := sv.Scan("non_existent_field", "some_value")
			Expect(err).NotTo(HaveOccurred())
			// No changes expected to 's'
			Expect(s).To(Equal(SimpleStruct{}))
		})

		type StructWithScanner struct {
			CustomField CustomScannerType `redis:"custom_field"`
		}
		
		It("should use Scanner interface if available", func() {
			var s ScanStruct
			spec := newStructSpec(reflect.TypeOf(ScanStruct{}), "redis")
			svScanner := StructValue{
				spec: spec,
				value: reflect.ValueOf(&s).Elem(),
			}
			err := svScanner.Scan("field", "scanned_value")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Field.Scanned).To(BeTrue())
			Expect(s.Field.Value).To(Equal("scanned_value"))
		})
		
		It("should use TextUnmarshaler interface if available and Scanner is not", func() {
			var s UnmarshalStruct
			spec := newStructSpec(reflect.TypeOf(UnmarshalStruct{}), "redis")
			svUnmarshaler := StructValue{
				spec: spec,
				value: reflect.ValueOf(&s).Elem(),
			}
			err := svUnmarshaler.Scan("field", "unmarshaled_value")
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Field.Unmarshaled).To(BeTrue())
			Expect(s.Field.Value).To(Equal("unmarshaled_value"))
		})

	})
})

// Helper types for Scanner and TextUnmarshaler tests
type CustomScannerType struct {
	Value   string
	Scanned bool
}

func (cst *CustomScannerType) ScanRedis(s string) error {
	cst.Value = s
	cst.Scanned = true
	return nil
}

type ScanStruct struct {
	Field CustomScannerType `redis:"field"`
}


type CustomTextUnmarshalType struct {
	Value       string
	Unmarshaled bool
}

func (ctut *CustomTextUnmarshalType) UnmarshalText(text []byte) error {
	ctut.Value = string(text)
	ctut.Unmarshaled = true
	return nil
}

type UnmarshalStruct struct {
	Field CustomTextUnmarshalType `redis:"field"`
}

// Ginkgo test suite setup
func TestStructMap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StructMap Suite")
}

// Helper to convert string to *string
func strPtr(s string) *string {
	return &s
}

// Helper to convert int to *int
func intPtr(i int) *int {
	return &i
}

// Helper to convert bool to *bool
func boolPtr(b bool) *bool {
	return &b
}

// Helper to convert float64 to *float64
func float64Ptr(f float64) *float64 {
	return &f
}
