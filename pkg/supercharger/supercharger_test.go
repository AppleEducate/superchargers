package supercharger

import (
	"encoding/json"
	"testing"

	"github.com/dewski/spatial"
	"github.com/stretchr/testify/assert"
)

func TestSuperchargerEquality(t *testing.T) {
	a := Supercharger{
		Address: "1234",
	}

	b := Supercharger{
		Address: "1234",
	}

	assert.True(t, a.Equal(b))
}

func TestSuperchargerImportBaiduGeo(t *testing.T) {
	body := `{
		"baidu_lat": "1.1",
		"baidu_lng": "1.1",
		"latitude": "0.0",
		"longitude": "0.0"
	}`
	var a Supercharger
	err := json.Unmarshal([]byte(body), &a)
	assert.NoError(t, err)

	point := spatial.Point{
		Lat: 1.1,
		Lng: 1.1,
	}

	assert.Equal(t, point, a.Geo)
}

func TestSuperchargerImportBaiduGeoWithMissingData(t *testing.T) {
	body := `{
		"baidu_lat": "1.1",
		"latitude": "1.2",
		"longitude": "1.2"
	}`
	var a Supercharger
	err := json.Unmarshal([]byte(body), &a)
	assert.NoError(t, err)

	point := spatial.Point{
		Lat: 1.2,
		Lng: 1.2,
	}

	assert.Equal(t, point, a.Geo)
}

func TestSuperchargerEqualityAddress(t *testing.T) {
	a := Supercharger{
		Address: "1234",
	}

	b := Supercharger{
		Address: "1235",
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityAddressLine1(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		AddressLine1: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		AddressLine1: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityAddressLine2(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		AddressLine2: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		AddressLine2: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityAddressNotes(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		AddressNotes: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		AddressNotes: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityAmenities(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		Amenities: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		Amenities: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityChargers(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		Chargers: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		Chargers: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityCity(t *testing.T) {
	a := Supercharger{
		City: "1234",
	}

	b := Supercharger{
		City: "1235",
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityCommonName(t *testing.T) {
	a := Supercharger{
		CommonName: "1234",
	}

	b := Supercharger{
		CommonName: "1235",
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityCountry(t *testing.T) {
	a := Supercharger{
		Country: "1234",
	}

	b := Supercharger{
		Country: "1235",
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityDestinationChargerLogo(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		DestinationChargerLogo: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		DestinationChargerLogo: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityDestinationWebsite(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		DestinationWebsite: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		DestinationWebsite: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityDirectionsLink(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		DirectionsLink: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		DirectionsLink: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityEmails(t *testing.T) {
	emails := EmailList{
		Email{Email: "1234", Label: "4567"},
		Email{Email: "8910", Label: "1234"},
	}
	a := Supercharger{
		Emails: emails,
	}

	emails = EmailList{
		Email{Email: "01234", Label: "4567"},
		Email{Email: "8910", Label: "1234"},
	}
	b := Supercharger{
		Emails: emails,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityGeocode(t *testing.T) {
	a := Supercharger{
		Geocode: "1234",
	}

	b := Supercharger{
		Geocode: "1235",
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityHours(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		Hours: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		Hours: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityIsGallery(t *testing.T) {
	a := Supercharger{
		IsGallery: true,
	}

	b := Supercharger{
		IsGallery: false,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityKioskPinX(t *testing.T) {
	pointer := int64(1234)
	a := Supercharger{
		KioskPinX: &pointer,
	}

	pointerb := int64(1235)
	b := Supercharger{
		KioskPinX: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityKioskPinY(t *testing.T) {
	pointer := int64(1234)
	a := Supercharger{
		KioskPinY: &pointer,
	}

	pointerb := int64(1235)
	b := Supercharger{
		KioskPinY: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityKioskZoomPinX(t *testing.T) {
	pointer := int64(1234)
	a := Supercharger{
		KioskZoomPinX: &pointer,
	}

	pointerb := int64(1235)
	b := Supercharger{
		KioskZoomPinX: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityKioskZoomPinY(t *testing.T) {
	pointer := int64(1234)
	a := Supercharger{
		KioskZoomPinY: &pointer,
	}

	pointerb := int64(1235)
	b := Supercharger{
		KioskZoomPinY: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityGeo(t *testing.T) {
	a := Supercharger{
		Geo: spatial.Point{
			Lat: 0.0,
			Lng: 0.0,
		},
	}

	b := Supercharger{
		Geo: spatial.Point{
			Lat: 0.1,
			Lng: 0.0,
		},
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityLocationID(t *testing.T) {
	a := Supercharger{
		LocationID: "1234",
	}

	b := Supercharger{
		LocationID: "1235",
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityLocationType(t *testing.T) {
	list := LocationList{"a", "b"}
	a := Supercharger{
		LocationType: list,
	}

	list = LocationList{"a", "c"}
	b := Supercharger{
		LocationType: list,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityNid(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			assert.Fail(t, "Did not panic")
		}
	}()

	a := Supercharger{
		Nid: 1234,
	}

	b := Supercharger{
		Nid: 1235,
	}

	a.Equal(b)
}

func TestSuperchargerEqualityOpenSoon(t *testing.T) {
	a := Supercharger{
		OpenSoon: true,
	}

	b := Supercharger{
		OpenSoon: false,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityPath(t *testing.T) {
	a := Supercharger{
		Path: "1234",
	}

	b := Supercharger{
		Path: "1235",
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityPostalCode(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		PostalCode: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		PostalCode: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityProvinceState(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		ProvinceState: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		ProvinceState: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityRegion(t *testing.T) {
	a := Supercharger{
		Region: "1234",
	}

	b := Supercharger{
		Region: "1235",
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualitySalesPhone(t *testing.T) {
	list := PhoneList{
		Phone{Number: "1234", Label: "5678"},
		Phone{Number: "9101", Label: "2345"},
	}
	a := Supercharger{
		SalesPhone: list,
	}

	list = PhoneList{
		Phone{Number: "01234", Label: "5678"},
		Phone{Number: "9101", Label: "2345"},
	}
	b := Supercharger{
		SalesPhone: list,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualitySalesRepresentative(t *testing.T) {
	a := Supercharger{
		SalesRepresentative: true,
	}

	b := Supercharger{
		SalesRepresentative: false,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualitySubRegion(t *testing.T) {
	pointer := "1234"
	a := Supercharger{
		SubRegion: &pointer,
	}

	pointerb := "1235"
	b := Supercharger{
		SubRegion: &pointerb,
	}

	assert.False(t, a.Equal(b))
}

func TestSuperchargerEqualityTitle(t *testing.T) {
	a := Supercharger{
		Title: "1234",
	}

	b := Supercharger{
		Title: "1235",
	}

	assert.False(t, a.Equal(b))
}
