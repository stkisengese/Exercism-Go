// Package census simulates a system used to collect census data.
package census

import (
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name      string
		residents []*Resident
		want      int
	}{
		{name: "all data collected", residents: []*Resident{{Name: "Matthew Sanabria", Age: 29, Address: map[string]string{"street": "Main St."}}}, want: 1},
		{name: "no data collected", residents: []*Resident{{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.residents); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResident_HasRequiredInfo(t *testing.T) {
	tests := []struct {
		name     string
		resident *Resident
		want     bool
	}{
		{name: "no data collected", resident: &Resident{}, want: false},
		{name: "all data collected", resident: &Resident{Name: "Matthew Sanabria", Age: 29, Address: map[string]string{"street": "Main St."}}, want: true},
		{name: "missing name", resident: &Resident{Name: "", Age: 29, Address: map[string]string{"street": "Main St."}}, want: false},
		{name: "nil map as address", resident: &Resident{Name: "Rob Pike", Age: 0, Address: nil}, want: false},
		{name: "empty map as address", resident: &Resident{Name: "Rob Pike", Age: 0, Address: map[string]string{}}, want: false},
		{name: "missing street", resident: &Resident{Name: "Hossein", Age: 30, Address: map[string]string{"street": ""}}, want: false},
		{name: "age is optional", resident: &Resident{Name: "Rob Pike", Age: 0, Address: map[string]string{"street": "Main St."}}, want: true},
		{name: "unknown key with value that is not empty", resident: &Resident{Name: "Rob Pike", Address: map[string]string{"unknown key": "with value"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// r := &Resident{
			// 	Name:    tt.resident.Name,
			// 	Age:     tt.resident.Age,
			// 	Address: tt.resident.Address,
			//}
			if got := tt.resident.HasRequiredInfo(); got != tt.want {
				t.Errorf("%#v.HasRequiredInfo() = %t, want %t", tt.resident, got, tt.want)
			}
		})
	}
}

func TestNewResident(t *testing.T) {
	tests := []struct {
		name     string
		resident *Resident
		want     *Resident
	}{
		{
			name:     "all data collected",
			resident: &Resident{Name: "Matthew Sanabria", Age: 29, Address: map[string]string{"street": "Main St."}},
			want:     &Resident{Name: "Matthew Sanabria", Age: 29, Address: map[string]string{"street": "Main St."}},
		},
		{
			name:     "no data collected",
			resident: &Resident{},
			want:     &Resident{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResident(tt.resident.Name, tt.resident.Age, tt.resident.Address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResident() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResident_Delete(t *testing.T) {
	test := []struct {
		name     string
		resident *Resident
		want     *Resident
	}{
		{name: "no data collected", resident: &Resident{}, want: &Resident{}},
		{name: "all data collected", resident: &Resident{Name: "Matthew Sanabria", Age: 29, Address: map[string]string{"street": "Main St."}}, want: &Resident{}},
		{name: "some data collected", resident: &Resident{Name: "Rob Pike", Age: 0, Address: map[string]string{}}, want: &Resident{}},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			tc.resident.Delete()
			if tc.resident.Name != "" || tc.resident.Age != 0 || len(tc.resident.Address) != 0 {
				t.Errorf("Test %s failed: resident.Delete() = %#v, want %#v", tc.name, tc.resident, tc.want)
			}
		})
	}
}
