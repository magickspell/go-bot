package product

import (
	"reflect"
	"testing"
)

// тест
func TestAllProducts(t *testing.T) {
	expected := allProducts

	if !reflect.DeepEqual(allProducts, expected) {
		t.Errorf("Expected allProducts to be %v, but got %v", expected, allProducts)
	}
}
